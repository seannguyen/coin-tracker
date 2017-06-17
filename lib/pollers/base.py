import abc
from datetime import date
from lib.services import coin_base, elastic_search, sentry_service, bitfinex_service, currency_convert_service
from lib.services.poloniex_service import PoloniexService


class BasePoller(object):
    def __init__(self):
        super(BasePoller, self).__init__()
        self._coin_base_client = coin_base.client()
        self._coin_base_service = coin_base.CoinBaseService()
        self._bit_finex_service = bitfinex_service.BitFinexService()
        self._poloniex_service = PoloniexService()
        self._es_client = elastic_search.client()
        self._currency_converter = currency_convert_service.CurrencyConvertService()
        self.__sentry_client = sentry_service.client()


    def poll(self):
        try:
            self._execute()
        except Exception as e:
            self.__sentry_client.captureException()
            raise e

    @staticmethod
    def _get_es_index_name(base_name):
        return "%s-%s" % (base_name, ''.join(str(date.today()).split('-')))

    @abc.abstractmethod
    def _execute(self):
        """Execute polling data"""
        return
