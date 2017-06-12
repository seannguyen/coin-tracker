import requests


class InvalidCurrencyException(Exception):
    pass


class CurrencyConvertService(object):
    def __init__(self):
        self.__rates = {}

    def convert(self, amount, from_currency, to_currency):
        if amount is None:
            return None
        rate = self.__get_rate(from_currency.upper(), to_currency.upper())
        return rate * amount

    def __get_rate(self, from_currency, to_currency):
        if from_currency not in self.__rates:
            self.__rates[from_currency] = {}
        if to_currency in self.__rates[from_currency]:
            return self.__rates[from_currency][to_currency]

        params = { 'base': from_currency, 'symbols': to_currency }
        response = requests.get('http://api.fixer.io/latest', params=params)
        response.raise_for_status()
        response_json = response.json()
        if from_currency != response_json['base'] or response_json['rates'][to_currency] is None:
            raise InvalidCurrencyException
        self.__rates[from_currency][to_currency] = response_json['rates'][to_currency]
        return response_json['rates'][to_currency]
