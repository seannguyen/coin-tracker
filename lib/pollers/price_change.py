from datetime import datetime, timedelta
import logging
import requests
from lib.pollers.base import BasePoller


class PriceChangePoller(BasePoller):
    __ES_INDEX_NAME = 'coinbase-price'
    __PRICE_CHANGE_THRESHOLD = 0.05
    __PRICE_CHANGE_TIME_RANGE_HOURS = 10

    def __init__(self):
        super(PriceChangePoller, self).__init__()

    def poll_price_change(self):
        data1 = requests.get('https://api.coinbase.com/v2/prices/BTC-SGD/spot',
                     params={'date': '2017-03-21T01:24:53.169327'}).json()
        data2 = requests.get('https://api.coinbase.com/v2/prices/BTC-SGD/spot',
                             params={'date':'2017-03-21T02:24:53.169327'}).json()

        print(data1)
        print(data2)
        # current_time = datetime.utcnow()
        # past_time = current_time - timedelta(hours=PriceChangePoller.__PRICE_CHANGE_TIME_RANGE_HOURS)
        # past_price = self.__get_bit_coin_price(past_time)
        # current_price = self.__get_bit_coin_price(current_time)
        #
        # lower_threshold = 1 - PriceChangePoller.__PRICE_CHANGE_THRESHOLD
        # upper_threshold = 1 + PriceChangePoller.__PRICE_CHANGE_THRESHOLD
        # if lower_threshold < current_price / past_price < upper_threshold:
        #     pass

    def __get_bit_coin_price(self, time):
        print(time.isoformat())
        response = requests.get('https://api.coinbase.com/v2/prices/BTC-SGD/spot', params={'date': time.isoformat()})
        price_data = response.json()['data']
        price = float(price_data['amount'])
        print(price)
        return price

if __name__ == '__main__':
    PriceChangePoller().poll_price_change()
