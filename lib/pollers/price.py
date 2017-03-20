from datetime import datetime
from lib.pollers.base import BasePoller


class PricePoller(BasePoller):
    __ES_INDEX_NAME = 'coinbase-price'

    def __init__(self):
        super(PricePoller, self).__init__()

    def poll_price(self):
        price_data = self.__get_bit_coin_price()
        self.__save_bit_coin_price(price_data.amount, price_data.currency)

    def __get_bit_coin_price(self):
        return self._coin_base_client.get_spot_price(currency_pair='BTC-USD')

    def __save_bit_coin_price(self, price, currency):
        body = {
            "price": float(price),
            "currency": currency,
            "timestamp": datetime.now()
        }
        self._es_client.index(index=PricePoller.__ES_INDEX_NAME, doc_type="price_data", body=body)

if __name__ == '__main__':
    PricePoller().poll_price()
