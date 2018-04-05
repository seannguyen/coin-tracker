from datetime import datetime
import logging
from lib.pollers.base import BasePoller
from lib.services.coinmarketcap_service import CoinMarketCapService


class InfoPoller(BasePoller):
    __ES_INDEX_NAME = BasePoller._get_es_index_name('coin-info')

    def __init__(self):
        super(InfoPoller, self).__init__()
        self._coinmarketcap_service = CoinMarketCapService()

    def _execute(self):
        coins_data = self._coinmarketcap_service.get_all_coins_data()
        for coin_data in coins_data:
            self.__save_coin_data(coin_data)

    def __save_coin_data(self, data):
        logging.info('Saving %s info to ElasticSearch' % data['name'])
        data['timestamp'] = datetime.utcnow()
        data['price_sgd'] = self._currency_converter.convert(data['price_usd'], 'USD', 'SGD')
        self._es_client.index(index=InfoPoller.__ES_INDEX_NAME, doc_type="coin_info_data", body=data)

if __name__ == '__main__':
    InfoPoller().poll()
