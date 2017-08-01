import logging
from datetime import datetime
from poloniex import Poloniex
from lib.services.base import BaseService
import lib.configs as configs
from lib.services.coinmarketcap_service import CoinMarketCapService
from lib.services.currency_convert_service import CurrencyConvertService


class PoloniexService(BaseService):
    EXCHANGE_NAME = 'poloninex'

    def __init__(self):
        super(PoloniexService, self).__init__()
        self.__client = Poloniex(configs.POLONIEX_API_KEY, configs.POLONIEX_API_SECRET)
        self.__coinmarketcap_service = CoinMarketCapService()
        self.__currency_converter = CurrencyConvertService()

    def get_balances(self):
        balance_data_by_type = self.__client.returnAvailableAccountBalances()
        coins_data = self.__coinmarketcap_service.get_all_coins_data_hash(limit=1000)
        balances_data = []
        now = datetime.utcnow()

        # there is a bug in poloniex client that when the response list is empty, the returned data
        # is an empty list instead of a dict
        if type(balance_data_by_type) is list:
            return balances_data

        for type, balance_by_currency in balance_data_by_type.iteritems():
            for symbol, native_balance in balance_by_currency.iteritems():
                native_balance = float(native_balance)
                usd_amount = coins_data[symbol]['price_usd'] * native_balance
                sgd_amount = self.__currency_converter.convert(usd_amount, 'USD', 'SGD')
                logging.info('Current Poloniex %s balance: SGD %s' % (symbol, sgd_amount))
                balances_data.append({
                    # 'timestamp': now,
                    'exchange': self.EXCHANGE_NAME,
                    'type': type,
                    'account': type,
                    'currency': symbol.upper(),
                    'balances': {
                        'native': native_balance,
                        'sgd': sgd_amount,
                        'usd': usd_amount
                    }
                })
        return balances_data