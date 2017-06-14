from itertools import ifilter, imap
from functools import reduce
from datetime import datetime
import logging
from lib.pollers.base import BasePoller


class BalancePoller(BasePoller):
    __ES_INDEX_NAME = 'coin-balance'

    def __init__(self):
        super(BalancePoller, self).__init__()

    def _execute(self):
        # CoinBase
        coinbase_balances_data = self._coin_base_service.get_balances()
        for balance_data in coinbase_balances_data:
            self._es_client.index(index=BalancePoller.__ES_INDEX_NAME, doc_type="balance_data", body=balance_data)
        # BitFinex
        bitfinex_balances_data = self._bit_finex_service.get_balances()
        for balance_data in bitfinex_balances_data:
            self._es_client.index(index=BalancePoller.__ES_INDEX_NAME, doc_type="balance_data", body=balance_data)
        # Poloniex
        poloniex_balances_data = self._poloniex_service.get_balances()
        for balance_data in poloniex_balances_data:
            self._es_client.index(index=BalancePoller.__ES_INDEX_NAME, doc_type="balance_data", body=balance_data)

        # All
        global_balances_data = self.__get_global_balances_data(coinbase_balances_data +
                                                               bitfinex_balances_data +
                                                               poloniex_balances_data)
        global_sum = {'usd': 0, 'sgd': 0, 'native': 0}
        for balance_data in global_balances_data:
            self._es_client.index(index=BalancePoller.__ES_INDEX_NAME, doc_type="balance_data", body=balance_data)
            global_sum['usd'] += balance_data['balances']['usd']
            global_sum['sgd'] += balance_data['balances']['sgd']
        balance_data = {
            'timestamp': datetime.utcnow(),
            'exchange': 'global',
            'type': 'global',
            'account': 'global',
            'currency': 'global',
            'balances': global_sum
        }
        self._es_client.index(index=BalancePoller.__ES_INDEX_NAME, doc_type="balance_data", body=balance_data)

    def __get_global_balances_data(self, balances_data):
        currencies = set(imap(lambda balance_data: balance_data['currency'], balances_data))
        now = datetime.utcnow()
        global_balances_data = []
        for currency in currencies:
            currency_balances_data = ifilter(lambda data: data['currency'] == currency, balances_data)
            native_amount = 0
            usd_amount = 0
            sgd_amount = 0
            for currency_balance_data in currency_balances_data:
                native_amount += currency_balance_data['balances']['native']
                usd_amount += currency_balance_data['balances']['usd']
                sgd_amount += currency_balance_data['balances']['sgd']
            logging.info('Current global %s balance: SGD %s' % (currency, sgd_amount))
            balance_data = {
                'timestamp': now,
                'exchange': 'global',
                'type': 'global',
                'account': 'global',
                'currency': currency,
                'balances': {
                    'native': native_amount,
                    'sgd': sgd_amount,
                    'usd': usd_amount
                }
            }
            global_balances_data.append(balance_data)
        return global_balances_data


if __name__ == '__main__':
    BalancePoller().poll()
