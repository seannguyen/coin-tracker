from datetime import datetime
import logging

from lib.pollers.base import BasePoller


class GainPoller(BasePoller):
    __ES_INDEX_NAME = 'coinbase-gain'
    __COIN_BASE_COMMISSION = 0.015
    __BUY_ADJUST_AFTER_COMMISSION = 1 + __COIN_BASE_COMMISSION
    __SELL_ADJUST_AFTER_COMMISSION = 1 - __COIN_BASE_COMMISSION

    def __init__(self):
        super(GainPoller, self).__init__()

    def poll_gain(self):
        accounts = self.__get_accounts()
        for account in accounts:
            buy = 0
            sell = 0
            transactions = self.__get_transactions(account.id)
            for transaction in transactions:
                if transaction.status != 'completed':
                    continue
                amount = float(transaction.native_amount.amount)
                if transaction.type == 'buy':
                    buy += amount
                else:
                    sell += amount

            account_name = account.name
            balance = float(account.native_balance.amount)
            logging.info(balance)
            self.__save_gain(sell, buy, balance, account_name)
            self.__save_net_gain(sell, buy, balance, account_name)

    def __save_net_gain(self, sell, buy, balance, account_name):
        net_sell = sell * GainPoller.__SELL_ADJUST_AFTER_COMMISSION
        net_buy = buy * GainPoller.__BUY_ADJUST_AFTER_COMMISSION
        net_balance = balance * GainPoller.__SELL_ADJUST_AFTER_COMMISSION
        net_gain = net_sell - net_buy + net_balance
        self.__save(net_gain, 'net_gain', account_name)

    def __save_gain(self, sell, buy, balance, account_name):
        gain = sell - buy + balance
        self.__save(gain, 'gain', account_name)

    def __save(self, gain, type, account_name):
        body = {
            'timestamp': datetime.utcnow(),
            'account_name': account_name,
            'type': type,
            'gain': gain
        }
        logging.info(body)
        self._es_client.index(index=GainPoller.__ES_INDEX_NAME, doc_type="gain_data", body=body)

    def __get_transactions(self, account_id):
        res = self._coin_base_client.get_transactions(account_id)
        return res.data

    def __get_accounts(self):
        res = self._coin_base_client.get_accounts()
        return res.data

if __name__ == '__main__':
    GainPoller().poll_gain()
