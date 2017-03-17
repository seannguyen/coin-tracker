from coinbase.wallet.client import Client
from datetime import datetime
from elasticsearch import Elasticsearch


coin_base_client = Client('no-api-key', 'no-api-secret')
bit_coin_price = coin_base_client.get_spot_price(currency_pair = 'BTC-USD')
print(bit_coin_price)

elasticsearch = Elasticsearch()
body = {
  "timestamp": datetime.now(), 
  "price": bit_coin_price.amount, 
  "currency": bit_coin_price.currency
}
elasticsearch.index(index="coinbase-price", doc_type="test-type", body=body)
