"""
Фабричный метод — порождающий паттерн проектирования, который определяет общий интерфейс для создания объектов
в суперклассе, позволяя подклассам изменять тип создаваемых объектов.

Применение:
- заранее неизвестны типы и зависимости объектов, с которыми должен работать код
- необходимо предусмотреть возможность расширять части фреймворка или библиотеки

Сущности: создатель, продукты

Реализация: приводим однотипные продукты к единому интерфейсу, реализуем метод создания этих продуктов,
все места создания продуктов в коде заменяем вызовами фабричного метода

Пример:
Почтовое отделение - предоставляет клиенту разные способы доставки.
Почтовое отделение - создатель. Способы доставки - продукты.

Из головного почтового отделения можем создать 3 вида доставки: самолётом, поездом, автомобилем.
Но одному из наших сервисов нужна только локальная доставка, поэтому он определяет свои: автомобилем и курьером.
"""
from abc import ABC, abstractmethod


# Базовые сущности, описывающие схему работы паттерна
class Product(ABC):
    """Базовый класс для классов-продуктов, возвращаемых создателем. Должны иметь один интерфейс.
    В нашем случае продукт - способ доставки."""
    def __init__(self, address: str):
        self.address = address

    @abstractmethod
    def deliver(self):
        pass


class Creator(ABC):
    """Базовый класс для класса-создателя, использующего единый интерфейс создания продуктов"""
    @abstractmethod
    def create_delivery(self, address: str, mode: str) -> Product:
        """Фабричный метод"""


# Непосредственно пример работы паттерна
# Объекты нашего основного приложения
class AirDelivery(Product):
    """Класс-продукт Авиадоставка"""
    def deliver(self):
        print(f'Delivered by plane to "{self.address}"')


class AutoDelivery(Product):
    """Класс-продукт Автодоставка"""
    def deliver(self):
        print(f'Delivered by car to "{self.address}"')


class TrainDelivery(Product):
    """Класс-продукт Доставка поездом"""
    def deliver(self):
        print(f'Delivered by train to "{self.address}"')


class HeadPostOffice(Creator):
    """Класс-создатель. Головное почтовое отделение"""

    modes = {
        'express': AirDelivery,
        'standard': AutoDelivery,
        'economy': TrainDelivery,
    }

    def create_delivery(self, address: str, mode: str = 'standard') -> Product:
        """Фабричный метод, создающий Delivery, используя единый интерфейс продуктов"""
        delivery = self.modes.get(mode)
        if not delivery:
            raise NotImplementedError('No such delivery method yet')

        return delivery(address)


# Объекты расширения к нашему приложению
class CourierDelivery(Product):
    """Пользовательское расширение к нашему приложению - Курьерская доставка"""
    def deliver(self):
        print(f'Delivered by courier to "{self.address}"')


class LocalPostOffice(HeadPostOffice):
    """Пользовательское расширение к нашему приложению, использующее собственные типы доставки"""
    modes = {
        'standard': AutoDelivery,
        'local': CourierDelivery,
    }


if __name__ == '__main__':
    head_office = HeadPostOffice()
    to_mow = head_office.create_delivery('Moscow')
    to_mow.deliver()

    to_mrm = head_office.create_delivery('Murmansk', 'express')
    to_mrm.deliver()

    local_office = LocalPostOffice()
    to_nsk = local_office.create_delivery('Novosibirsk', 'local')
    to_nsk.deliver()
