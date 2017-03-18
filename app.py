from coinbase.wallet.client import Client
from datetime import datetime
from elasticsearch import Elasticsearch

coin_base_client = Client('no-api-key', 'no-api-secret')
bit_coin_price = coin_base_client.get_spot_price()

index_name = 'coinbase-price'

client = Elasticsearch()

if not client.indices.exists(index=index_name):
  index_config = {
    "mappings": {
      "price_data": {
        "properties": {
          "price": {
            "type": "float",
          }
        }
      }
    }
  }
  client.indices.create(index=index_name, body=index_config)

body = {
  "price": float(bit_coin_price.amount),
  "currency": bit_coin_price.currency,
  "timestamp": datetime.now()
}

client.index(index="coinbase-price", doc_type="test-type", body=body)