import logging
from datetime import datetime
from itertools import takewhile
from lib.services.base import BaseService
from lib import configs
from bittrex.bittrex import Bittrex
from lib.services.currency_convert_service import CurrencyConvertService
from lib.services.coinmarketcap_service import CoinMarketCapService


class BittrexService(BaseService):
    EXCHANGE_NAME = 'bittrex'

    def __init__(self):
        self.__currency_converter = CurrencyConvertService()
        self.__coin_service = CoinMarketCapService()
        self.__client = Bittrex(configs.BITTREX_APY_KEY, configs.BITTREX_APY_SECRET)
        self.__currencies = None

    def get_balances(self):
        logging.info('Start query for Bittrex balance')
        result_data = []
        balances_data = self.__get_native_balances()

        now = datetime.utcnow()
        for data in balances_data:
            currency = data['Currency']
            logging.info('Processing %s' % currency)
            native_amount = float(data['Balance'])
            usd_amount = self.__get_price(currency) * native_amount
            sgd_amount = self.__currency_converter.convert(usd_amount, 'USD', 'SGD')
            logging.info('Current Bittrex balance: SGD %s' % sgd_amount)
            balance_data = {
                'timestamp': now,
                'exchange': self.EXCHANGE_NAME,
                'type': 'trading',
                'account': 'trading',
                'currency': currency.upper(),
                'balances': {
                    'native': native_amount,
                    'sgd': sgd_amount,
                    'usd': usd_amount
                }
            }
            result_data.append(balance_data)
        return result_data

    def __get_native_balances(self):
        return self.__client.get_balances()['result']

    def __get_price(self, symbol):
        currencies = self.__get_currencies()
        logging.info('Getting price for %s' % symbol)
        currency = next(item for item in currencies if item['Currency'] == symbol)
        return self.__coin_service.get_price(currency['CurrencyLong'])

    def __get_currencies(self):
        self.__currencies = self.__currencies or self.__client.get_currencies()['result']
        return self.__currencies

if __name__ == '__main__':
    print(BittrexService().get_balances())