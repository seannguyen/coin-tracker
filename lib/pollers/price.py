from datetime import datetime
import logging
import requests
from lib.pollers.base import BasePoller


class PricePoller(BasePoller):
    __ES_INDEX_NAME = 'coinbase-price'

    def __init__(self):
        super(PricePoller, self).__init__()

    def poll_price(self):
        price_data = self._get_bit_coin_price()
        self.__save_bit_coin_price(price_data['amount'], price_data['currency'])

    def _get_bit_coin_price(self):
        price_data = self.__get_price_by_url('https://api.coinbase.com/v2/prices/BTC-SGD/spot')
        logging.info('CoinBase BitCoin price: %s' % (price_data))
        return price_data

    def _get_ethereum_price(self):
        price_data = self.__get_price_by_url('https://api.coinbase.com/v2/prices/ETH-SGD/spot')
        logging.info('CoinBase Ethereum price: %s' % (price_data))
        return price_data

    def __get_price_by_url(self, url):
        # response = self._coin_base_client.get_spot_price(currency_pair='BTC-SGD')
        response = requests.get(url)
        if not response.ok:
            return False
        price_data = response.json()['data']
        return price_data

    def __save_bit_coin_price(self, price, currency):
        self.__save_price(price, currency, 'BTC')

    def __save_ethereum_price(self, price, currency):
        self.__save_price(price, currency, 'ETH')

    def __save_price(self, price, currency, type):
        logging.info('Saving CoinBase %s price to ElasticSearch' % type)
        body = {
            "price": float(price),
            "currency": currency,
            "type": type,
            "timestamp": datetime.utcnow()
        }
        self._es_client.index(index=PricePoller.__ES_INDEX_NAME, doc_type="price_data", body=body)

if __name__ == '__main__':
    PricePoller().poll_price()
