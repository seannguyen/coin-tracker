import abc
from lib.services import coin_base, elastic_search, sentry_service, bitfinex_service, currency_convert_service


class BasePoller(object):
    __ES_INDEX_NAME = 'coinbase-price'

    def __init__(self):
        super(BasePoller, self).__init__()
        self._coin_base_client = coin_base.client()
        self._coin_base_service = coin_base.CoinBaseService()
        self._bit_finex_service = bitfinex_service.BitFinexService()
        self._es_client = elastic_search.client()
        self._currency_converter = currency_convert_service.CurrencyConvertService()
        self.__sentry_client = sentry_service.client()


    def poll(self):
        try:
            self._execute()
        except Exception as e:
            self.__sentry_client.captureException()
            raise e

    @abc.abstractmethod
    def _execute(self):
        """Execute polling data"""
        return
