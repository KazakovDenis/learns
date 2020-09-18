"""
Стратегия — поведенческий паттерн проектирования, который определяет семейство схожих алгоритмов и помещает каждый
из них в собственный класс, после чего алгоритмы можно взаимозаменять прямо во время исполнения программы.

Применение:
- нужно использовать разные вариации какого-то алгоритма внутри одного объекта
- есть множество похожих классов, отличающихся только некоторым поведением
- не следует обнажать детали реализации алгоритмов для других классов

Сущности: клиент, контекст, стратегии (1 и более).

Реализация: клиент взаимодействует с контекстом, контекст передаёт исполнение какой-то логики выбранной стратегии
через единый для всех стратегий интерфейс.

Пример:
Контекст - приложение-навигатор.
Стратегии: пеший маршрут, маршрут на авто, маршрут на метро.

Клиент сообщает пункты А и Б и выбирает стратегию, контекст получает от стратегии координаты маршрута, строит его и
возвращает клиенту результат.
"""
from abc import ABC, abstractmethod


# Базовые сущности, описывающие схему работы паттерна
class Strategy(ABC):
    """Базовый класс с единым интерфейсом стратегий. Интерфейс должен определяться контекстом"""

    @abstractmethod
    def get_coordinates(self, start, finish):
        pass


class Context(ABC):
    """Контекст не знает, какая ему будет передана стратегия, он использует только интерфейс"""
    _strategy: Strategy

    def set_strategy(self, strategy: Strategy):
        self._strategy = strategy


# Непосредственно пример работы паттерна
class WalkingStrategy(Strategy):
    def get_coordinates(self, start, finish):
        # здесь какая-то бизнес-логика
        longitude = start + finish
        return 'пешком: ' + str(longitude)


class SubwayStrategy(Strategy):
    def get_coordinates(self, start, finish):
        # здесь какая-то бизнес-логика
        longitude = (start + finish) * 2
        return 'на метро: ' + str(longitude)


class CarStrategy(Strategy):
    def get_coordinates(self, start, finish):
        # здесь какая-то бизнес-логика
        longitude = (start + finish) * 3
        return 'на автомобиле: ' + str(longitude)


class Navigator(Context):
    """Приложение, использующее паттерн"""

    def build_route(self, start, finish):
        """Метод, непосредственно использующий интерфейс стратегий"""
        coordinates = self._strategy.get_coordinates(start, finish)
        # здесь какая-то бизнес-логика
        return 'Маршрут ' + coordinates


if __name__ == '__main__':
    navigator = Navigator()
    strategies = WalkingStrategy(), SubwayStrategy(), CarStrategy()

    for s in strategies:
        navigator.set_strategy(s)
        route = navigator.build_route(1, 2)
        print(route)
