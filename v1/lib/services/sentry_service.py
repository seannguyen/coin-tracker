from raven import Client

import lib.configs as configs


def client():
    return Client(configs.SENTRY_URL)
