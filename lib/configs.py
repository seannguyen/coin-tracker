import os
import logging

# Logging
logging.basicConfig(filename='coinbase-tracker.log',
                    level=logging.INFO,
                    format='%(asctime)-15s %(message)s')

# CoinBase
API_KEY = 'qZZTKMz8DXuGajSl'
API_SECRET = os.getenv('API_SECRET', 'dummy_key')