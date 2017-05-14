from lib.services import coin_base, elastic_search


class BasePoller(object):
    __ES_INDEX_NAME = 'coinbase-price'

    def __init__(self):
        super(BasePoller, self).__init__()
        self._coin_base_client = coin_base.client()
        self._es_client = elastic_search.client()
