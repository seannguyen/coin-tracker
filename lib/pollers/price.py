from datetime import datetime
import logging
import requests
from lib.pollers.base import BasePoller


class PricePoller(BasePoller):
    __ES_INDEX_NAME = 'coinbase-price'

    def __init__(self):
        super(PricePoller, self).__init__()

    def _execute(self):
        price_data = self._coin_base_service.get_price('BTC', 'SGD')
        self.__save_price(price_data['amount'], price_data['currency'], 'BTC')
        price_data = self._coin_base_service.get_price('ETH', 'SGD')
        self.__save_price(price_data['amount'], price_data['currency'], 'ETH')
        price_data = self._coin_base_service.get_price('LTC', 'SGD')
        self.__save_price(price_data['amount'], price_data['currency'], 'LTC')

    def __save_price(self, price, currency, type):
        logging.info('Saving %s price to ElasticSearch' % type)
        body = {
            "price": float(price),
            "currency": currency,
            "type": type,
            "timestamp": datetime.utcnow()
        }
        self._es_client.index(index=PricePoller.__ES_INDEX_NAME, doc_type="price_data", body=body)

if __name__ == '__main__':
    PricePoller().poll()
