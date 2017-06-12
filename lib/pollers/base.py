import abc
import logging
from raven.handlers.logging import SentryHandler
from raven.conf import setup_logging
from lib.services import coin_base, elastic_search, sentry_service


class BasePoller(object):
    __ES_INDEX_NAME = 'coinbase-price'

    def __init__(self):
        super(BasePoller, self).__init__()
        self._coin_base_client = coin_base.client()
        self._coin_base_service = coin_base.CoinBaseService()
        self._es_client = elastic_search.client()
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
