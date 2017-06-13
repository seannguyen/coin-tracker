from datetime import datetime, timedelta
import logging
import requests

import lib.configs as configs
from lib.pollers.info import InfoPoller
import lib.services.redis_service as redis_service


class PriceChangePoller(InfoPoller):
    __ES_INDEX_NAME = 'coin-price-change'
    __PRICE_CHANGE_THRESHOLD = 0.1
    __PRICE_CHANGE_TIME_RANGE_HOURS = 3
    __NOTIFICATION_SUSPENSION_KEY = 'NOTIFICATION_SUSPENSION:PRICE_CHANGE'
    __NOTIFICATION_SUSPENSION_KEY_EXPIRING_TIME = 60 * 60 * 3

    def __init__(self):
        super(PriceChangePoller, self).__init__()
        self._redis_client = redis_service.client()

    def _execute(self):
        coins_data = self._coinmarketcap_service.get_all_coins_data()
        for coin_data in coins_data:
            self.__poll_and_alert_price_change(coin_data)

    def __poll_and_alert_price_change(self, current_price_data):
        logging.info('Start polling %s price change' % current_price_data['name'])
        past_price_data = self.__get_past_price(current_price_data['id'])
        if not past_price_data:
            return

        logging.info('%s price %s hours ago: USD %s' % (current_price_data['name'],
                                                       self.__PRICE_CHANGE_TIME_RANGE_HOURS,
                                                       past_price_data['price_usd']))
        logging.info('%s price at the moment: USD %s' % (current_price_data['name'],
                                                        current_price_data['price_usd']))

        if not current_price_data['price_usd'] or not past_price_data['price_usd']:
            return
        change_ratio = float(current_price_data['price_usd']) / past_price_data['price_usd']
        logging.info('%s Price change ratio: %s' % (current_price_data['name'], change_ratio))
        self.__save(change_ratio, current_price_data['id'])
        self.__alert(change_ratio, current_price_data['name'])

    def __save(self, change_ratio, coin_id):
        logging.info('Saving %s change ratio to ElasticSearch' % coin_id)
        body = {
            "change_ratio": float(change_ratio),
            "timestamp": datetime.utcnow(),
            "type": coin_id
        }
        self._es_client.index(index=self.__ES_INDEX_NAME, doc_type="price_change_data", body=body)

    def __alert(self, change_ratio, name):
        redis_notification_suspension_key = "%s:%s" % (PriceChangePoller.__NOTIFICATION_SUSPENSION_KEY, name)
        if self._redis_client.exists(redis_notification_suspension_key):
            return
        lower_threshold = 1 - PriceChangePoller.__PRICE_CHANGE_THRESHOLD
        upper_threshold = 1 + PriceChangePoller.__PRICE_CHANGE_THRESHOLD
        if not(lower_threshold < change_ratio < upper_threshold):
            payload = {"text": "%s price is changing quickly %s%% compare to %s hours ago"
                               % (name, round(change_ratio * 100, 2), self.__PRICE_CHANGE_TIME_RANGE_HOURS)}
            requests.post(configs.SLACK_WEB_HOOK, data={'payload': str(payload)})
            self._redis_client.set(
                redis_notification_suspension_key,
                True,
                ex=PriceChangePoller.__NOTIFICATION_SUSPENSION_KEY_EXPIRING_TIME)

    def __get_past_price(self, coin_id):
        query_body = {
            'query': {
                'bool': {
                    'must': [
                        {
                            'match': {
                                "id": coin_id
                            }
                        },
                        {
                            'range': {
                                'timestamp': {
                                    'gte': (datetime.utcnow() - timedelta(hours=self.__PRICE_CHANGE_TIME_RANGE_HOURS)).isoformat(),
                                    'lte': (
                                        datetime.utcnow() - timedelta(hours=self.__PRICE_CHANGE_TIME_RANGE_HOURS - 1)).isoformat()
                                }
                            }
                        }
                    ]
                }

            }
        }
        res = self._es_client.search(index='coin-info', body=query_body)
        hits = res['hits']['hits']
        return hits[0]['_source'] if hits else None


if __name__ == '__main__':
    PriceChangePoller().poll()
