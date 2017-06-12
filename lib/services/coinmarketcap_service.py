import requests
from lib.services.base import BaseService


class CoinMarketCapService(BaseService):
    def get_all_currencies_data(self):
        res = requests.get('https://api.coinmarketcap.com/v1/ticker')
        res.raise_for_status()
        return res.json()['data']
