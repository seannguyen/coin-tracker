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
        # response = self._coin_base_client.get_spot_price(currency_pair='BTC-SGD')
        response = requests.get('https://api.coinbase.com/v2/prices/BTC-SGD/spot')
        if not response.ok:
            return False
        price_data = response.json()['data']
        logging.info('CoinBase BitCoin price: %s' % (price_data))
        return price_data

    def __save_bit_coin_price(self, price, currency):
        logging.info('Saving CoinBase balance to ElasticSearch')
        body = {
            "price": float(price),
            "currency": currency,
            "timestamp": datetime.now()
        }
        self._es_client.index(index=PricePoller.__ES_INDEX_NAME, doc_type="price_data", body=body)

if __name__ == '__main__':
    PricePoller().poll_price()
