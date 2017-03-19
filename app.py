from datetime import datetime
import coin_base_client
import es_client


class BitCoinPriceTracker:
    __ES_INDEX_NAME = 'coinbase-price'

    def __init__(self):
        self.__coin_base_client = coin_base_client.client()
        self.__es_client = es_client.client()

    def snapshot_price(self):
        price_data = self.__get_bit_coin_price()
        self.__save_bit_coin_price(price_data.amount, price_data.currency)

    def __get_bit_coin_price(self):
        return self.__coin_base_client.get_spot_price(urrency_pair = 'BTC-SGD')

    def __save_bit_coin_price(self, price, currency):
        self.__create_es_index()

        body = {
            "price": float(price),
            "currency": currency,
            "timestamp": datetime.now()
        }
        self.__es_client.index(index=BitCoinPriceTracker.__ES_INDEX_NAME, doc_type="price_data", body=body)

    def __create_es_index(self):
        indices = self.__es_client.indices
        if not indices.exists(index=BitCoinPriceTracker.__ES_INDEX_NAME):
            index_config = {
                "mappings": {
                    "price_data": {
                        "properties": {
                            "price": {
                                "type": "float",
                            }
                        }
                    }
                }
            }
            indices.create(index=BitCoinPriceTracker.__ES_INDEX_NAME, body=index_config)


if __name__ == '__main__':
    BitCoinPriceTracker().snapshot_price()
