import logging
from datetime import datetime
import json
import base64
import hashlib
import requests
import hmac
import time
from lib.services.base import BaseService
from lib import configs
from lib.services.currency_convert_service import CurrencyConvertService


class BitFinexService(BaseService):
    EXCHANGE_NAME = 'bitfinex'

    def __init__(self):
        self.__currency_converter = CurrencyConvertService()

    def get_balances(self):
        logging.info('Start query for BitFinex balance')
        result_data = []
        balances_data = self.__get_native_balances()
        now = datetime.utcnow()
        for data in balances_data:
            usd_amount = self.__get_price(data['currency']) * float(data['amount'])
            sgd_amount = self.__currency_converter.convert(usd_amount, 'USD', 'SGD')
            logging.info('Current BitFinex %s balance: SGD %s' % (data['type'], sgd_amount))
            balance_data = {
                'timestamp': now,
                'exchange': self.EXCHANGE_NAME,
                'type': data['type'],
                'account': data['type'],
                'currency': data['currency'].upper(),
                'balances': {
                    'native': float(data['amount']),
                    'sgd': sgd_amount,
                    'usd': usd_amount
                }
            }
            result_data.append(balance_data)
        return result_data

    def __get_native_balances(self):
        url = '/v1/balances'
        body = {
            'request': url,
            'nonce': str(int(round(time.time() * 100000)))
        }
        body_json = json.dumps(body)
        payload = base64.b64encode(body_json)
        signature = hmac.new(bytes(configs.BITFINEX_API_SECRET), msg=payload, digestmod=hashlib.sha384).hexdigest()
        header = {
            'X-BFX-APIKEY': configs.BITFINEX_API_KEY,
            'X-BFX-PAYLOAD': payload,
            'X-BFX-SIGNATURE': signature
        }
        response = requests.post(configs.BITFINEX_BASE_URL + url, headers=header, data=body_json)
        response.raise_for_status()
        return response.json()

    def __get_price(self, currency, target_currency='usd'):
        response = requests.get(configs.BITFINEX_BASE_URL + '/v1/pubticker/%s' % currency + target_currency)
        response.raise_for_status()
        return float(response.json()['last_price'])


if __name__ == '__main__':
    print json.dumps(BitFinexService().get_balances(), indent=2)