from datetime import datetime, timedelta
import logging
import requests

import lib.configs as configs
from lib.pollers.price import PricePoller


class PriceChangePoller(PricePoller):
    __ES_INDEX_NAME = 'coinbase-price-change'
    __PRICE_CHANGE_THRESHOLD = 0.05
    __PRICE_CHANGE_TIME_RANGE_HOURS = 10

    def __init__(self):
        super(PriceChangePoller, self).__init__()

    def poll_price_change(self):
        self.__poll_and_alert_price_change('BTC')
        self.__poll_and_alert_price_change('ETH')
        self.__poll_and_alert_price_change('LTC')

    def __poll_and_alert_price_change(self, type):
        logging.info('Start polling %s price change' % type)
        past_price_data = self.__get_past_price(type)
        if not past_price_data:
            return

        logging.info('%s price %s hours ago: %s %s' % (type,
                                                       self.__PRICE_CHANGE_TIME_RANGE_HOURS,
                                                       past_price_data['currency'],
                                                       past_price_data['price']))
        current_price_data = self.__get_current_price(type)
        logging.info('%s price at the moment: %s %s' % (type,
                                                        current_price_data['currency'],
                                                        current_price_data['amount']))

        change_ratio = float(current_price_data['amount']) / past_price_data['price']
        logging.info('%s Price change ratio: %s' % (type, change_ratio))
        self.__save(change_ratio, type)
        self.__alert(change_ratio, type)

    def __save(self, change_ratio, type):
        logging.info('Saving %s change ratio to ElasticSearch' % type)
        body = {
            "change_ratio": float(change_ratio),
            "timestamp": datetime.utcnow(),
            "type": type
        }
        self._es_client.index(index=self.__ES_INDEX_NAME, doc_type="price_change_data", body=body)

    def __alert(self, change_ratio, type):
        lower_threshold = 1 - PriceChangePoller.__PRICE_CHANGE_THRESHOLD
        upper_threshold = 1 + PriceChangePoller.__PRICE_CHANGE_THRESHOLD
        if not(lower_threshold < change_ratio < upper_threshold):
            payload = {"text": "%s price is changing quickly %s%% compare to %s hours ago"
                               % (type, round(change_ratio * 100, 2), self.__PRICE_CHANGE_TIME_RANGE_HOURS)}
            requests.post(configs.SLACK_WEB_HOOK, data={'payload': str(payload)})

    def __get_past_price(self, type):
        query_body = {
            'query': {
                'bool': {
                    'must': [
                        {
                            'match': {
                                "type": type
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
        res = self._es_client.search(index='coinbase-price', body=query_body)
        hits = res['hits']['hits']
        return hits[0]['_source'] if hits else None

    def __get_current_price(self, type):
        if type == 'BTC':
            return self._get_bit_coin_price()
        elif type == 'ETH':
            return self._get_ethereum_price()
        elif type == 'LTC':
            return self._get_lite_coin_price()
        else:
            pass


if __name__ == '__main__':
    PriceChangePoller().poll_price_change()
