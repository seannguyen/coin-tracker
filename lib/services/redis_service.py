from redis import StrictRedis
from lib import configs


def client():
    return StrictRedis(host=configs.REDIS_HOST, port=configs.REDIS_PORT, db=0, password=configs.REDIS_PASSWORD)