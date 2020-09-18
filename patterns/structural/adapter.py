"""
Адаптер — структурный паттерн проектирования, который позволяет объектам с несовместимыми
интерфейсами работать вместе.

Применение:
- необходимо связать несовместимые части системы
- необходимо конвертировать данные от одной части системы в формат, подходящий для другой

Сущности: клиент, адаптер, сторонний сервис (иная часть системы)

Реализация: клиент взаимодействует с адаптером, в котором обёрнуты методы стороннего сервиса.
1. Адаптер объектов: принимает сторонний сервис в качестве атрибута, в своем методе вызывает методы стороннего сервиса
2. Адаптер классов: наследует классы Клиента и Сервиса.

Пример:
Клиент - сервис рассылки новостей, у которого реализован единый метод получения новостей.
Сторонний сервис - источник новостей, не имеющий единого метода.
Адаптер - класс, объединяющий методы стороннего сервиса в единый метод для клиента.
"""
from abc import ABC, abstractmethod
from random import randint


# Базовые сущности, описывающие схему работы паттерна
class ThirdPartyService(ABC):
    """Сторонний сервис, интерфейс которого не подходит клиентскому приложению"""


class Adapter(ABC):
    """Адаптер объекта, оборачивающий методы стороннего сервиса в пригодный для клиента интерфейс"""
    service: ThirdPartyService


class RequiredInterface(ABC):
    """Интерфейс, который должен реализовывать сторонний сервис"""
    @abstractmethod
    def get_news(self):
        pass


class Client(ABC):
    """Клиентский сервис, которому необходим адаптер для работы со сторонним сервисом"""
    source: RequiredInterface


# Непосредственно пример работы паттерна
class NewsService(ThirdPartyService):
    """Сервис - источник новостей"""

    @staticmethod
    def get_politics():
        return randint(0, 10)

    @staticmethod
    def get_economics():
        return randint(0, 10)


class NewsServiceAdapter(Adapter, RequiredInterface):
    """Адаптер новостного сервиса"""
    service: NewsService

    def __init__(self, service: NewsService):
        self.service = service

    def get_news(self):
        politics = self.service.get_politics()
        economics = self.service.get_economics()
        return politics + economics


class MailingService(Client):
    """Сервис рассылки"""

    def __init__(self, news_source: RequiredInterface):
        self.source = news_source

    def send_email(self, subscriber: str):
        news = self.source.get_news()
        print(f'News ({news}) has been sent to: {subscriber}')


if __name__ == '__main__':
    source = NewsService()
    adapter = NewsServiceAdapter(source)
    client = MailingService(adapter)
    client.send_email('Ivan Ivanov')
