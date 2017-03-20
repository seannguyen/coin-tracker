from coinbase.wallet.client import Client

from lib import configs


def client():
    return Client(configs.API_KEY, configs.API_SECRET)
