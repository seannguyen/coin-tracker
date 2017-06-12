import logging
from datetime import datetime
from coinbase.wallet.client import Client
from lib.services.base import BaseService
from lib import configs
from lib.services.currency_convert_service import CurrencyConvertService


def client():
    return Client(configs.COINBASE_API_KEY, configs.COINBASE_API_SECRET)


class CoinBaseService(BaseService):
    EXCHANGE_NAME = 'coinbase'

    def __init__(self):
        self.__client = Client(configs.COINBASE_API_KEY, configs.COINBASE_API_SECRET)
        self.__currency_converter = CurrencyConvertService()

    def get_balances(self):
        logging.info('Start query for CoinBase balance')
        now = datetime.utcnow()
        balances_data = []
        accounts = self.__client.get_accounts().data
        for account in accounts:
            logging.info('Current CoinBase %s balance: %s %s' % (account['name'],
                                                        account['native_balance']['currency'],
                                                        account['native_balance']['amount']))
            sgd_amount = float(account['native_balance']['amount'])
            balance_data = {
                'timestamp': now,
                'exchange': self.EXCHANGE_NAME,
                'type': account['type'],
                'account': account['name'],
                'currency': account['currency'],
                'balances': {
                    'native': float(account['balance']['amount']),
                    'sgd': sgd_amount,
                    'usd': self.__currency_converter.convert(sgd_amount, 'SGD', 'USD')
                }
            }
            balances_data.append(balance_data)
        return balances_data
