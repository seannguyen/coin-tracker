from elasticsearch import Elasticsearch

es = Elasticsearch()

res = es.search(index="coinbase-price", body={"query": {"match_all": {}}})
print("Got %d Hits:" % res['hits']['total'])
for hit in res['hits']['hits']:
    print(hit['_source'])
    print("%(timestamp)s %(currency)s: %(price)s" % hit["_source"])
