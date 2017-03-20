from datetime import datetime
from services import coin_base, elastic_search


class BitCoinPriceTracker:
    __ES_INDEX_NAME = 'coinbase-price'

    def __init__(self):
        self.__coin_base_client = coin_base.client()
        self.__es_client = elastic_search.client()

    def snapshot_price(self):
        price_data = self.__get_bit_coin_price()
        self.__save_bit_coin_price(price_data.amount, price_data.currency)

    def __get_bit_coin_price(self):
        return self.__coin_base_client.get_spot_price(urrency_pair = 'BTC-SGD')

    def __save_bit_coin_price(self, price, currency):
        body = {
            "price": float(price),
            "currency": currency,
            "timestamp": datetime.now()
        }
        self.__es_client.index(index=BitCoinPriceTracker.__ES_INDEX_NAME, doc_type="price_data", body=body)

if __name__ == '__main__':
    BitCoinPriceTracker().snapshot_price()
