from datetime import datetime
from lib.services import coin_base, elastic_search
from lib.poller.base import BasePoller


class BalancePoller(BasePoller):
    __ES_INDEX_NAME = 'coinbase-balance'

    def __init__(self):
        super(BalancePoller, self).__init__()

    def poll_balance(self):
        accounts = self._coin_base_client.get_accounts().data
        for account in accounts:
            account['balance']['amount'] = float(account['balance']['amount'])
            account['native_balance']['amount'] = float(account['native_balance']['amount'])
            self._es_client.index(index=BalancePoller.__ES_INDEX_NAME, doc_type="balance_data", body=account)


if __name__ == '__main__':
    BalancePoller().poll_balance()
