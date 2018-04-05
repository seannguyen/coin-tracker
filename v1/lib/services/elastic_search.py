from elasticsearch import Elasticsearch
from lib import configs


def client():
    return Elasticsearch(hosts=configs.ELASTIC_SEARCH_HOST)