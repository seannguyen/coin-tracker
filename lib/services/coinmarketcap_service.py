import requests
from datetime import datetime
from itertools import imap
from lib.services.base import BaseService


class CoinMarketCapService(BaseService):
    def get_all_coins_data(self):
        res = requests.get('https://api.coinmarketcap.com/v1/ticker')
        res.raise_for_status()
        return imap(self.__preprocess_data, res.json())

    def __preprocess_data(self, data):
        number_fields = ['rank',
                         'price_usd',
                         'price_btc',
                         '24h_volume_usd',
                         'market_cap_usd',
                         'available_supply',
                         'total_supply',
                         'percent_change_1h',
                         'percent_change_24h',
                         'percent_change_7d',
                         'last_updated']
        for field in number_fields:
            if data[field] is not None:
                data[field] = float(data[field])
        if data['last_updated'] is not None:
            data['last_updated'] = datetime.fromtimestamp(data['last_updated'])
        return data