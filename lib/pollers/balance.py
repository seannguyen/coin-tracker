from datetime import datetime
import logging

from lib.pollers.base import BasePoller


class BalancePoller(BasePoller):
    __ES_INDEX_NAME = 'coin-balance'

    def __init__(self):
        super(BalancePoller, self).__init__()

    def _execute(self):
        #CoinBase
        balances_data = self._coin_base_service.get_balances()
        for balance_data in balances_data:
            self._es_client.index(index=BalancePoller.__ES_INDEX_NAME, doc_type="balance_data", body=balance_data)


if __name__ == '__main__':
    BalancePoller().poll()
