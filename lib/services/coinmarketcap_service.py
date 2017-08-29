import requests
from datetime import datetime
from itertools import imap, izip
from lib.services.base import BaseService
from pydash.collections import key_by


class CoinMarketCapService(BaseService):
    def get_all_coins_data_hash(self, limit=50):
        all_coins_data = self.get_all_coins_data(limit=limit)
        return key_by(all_coins_data, 'symbol')

    def get_all_coins_data(self, limit=50):
        res = requests.get('https://api.coinmarketcap.com/v1/ticker', params={'limit':limit})
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