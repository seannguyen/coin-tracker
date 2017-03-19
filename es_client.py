from elasticsearch import Elasticsearch


def client():
    return Elasticsearch()