from datetime import datetime
import logging

from lib.pollers.base import BasePoller


class BalancePoller(BasePoller):
    __ES_INDEX_NAME = 'coinbase-balance'

    def __init__(self):
        super(BalancePoller, self).__init__()

    def poll_balance(self):
        logging.info('Start query for CoinBase balance')
        accounts = self._coin_base_client.get_accounts().data
        for account in accounts:
            account['timestamp'] = datetime.utcnow()
            account['balance']['amount'] = float(account['balance']['amount'])
            account['native_balance']['amount'] = float(account['native_balance']['amount'])
            self._es_client.index(index=BalancePoller.__ES_INDEX_NAME, doc_type="balance_data", body=account)


if __name__ == '__main__':
    BalancePoller().poll_balance()
