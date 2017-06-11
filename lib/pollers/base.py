import abc
from lib.services import coin_base, elastic_search, sentry_service


class BasePoller(object):
    __ES_INDEX_NAME = 'coinbase-price'

    def __init__(self):
        super(BasePoller, self).__init__()
        self._coin_base_client = coin_base.client()
        self._es_client = elastic_search.client()
        self._sentry_client = sentry_service.client()

    def poll(self):
        try:
            self._execute()
        except Exception:
            self._sentry_client.captureException()

    @abc.abstractmethod
    def _execute(self):
        """Execute polling data"""
        return
