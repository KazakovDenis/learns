"""
Одиночка — порождающий паттерн проектирования, который гарантирует, что у класса есть только один экземпляр,
и предоставляет к нему глобальную точку доступа.
"""


class Singleton:

    __instance = None

    def __new__(cls, *args, **kwargs):
        if not cls.__instance:
            cls.__instance = super().__new__(cls)
        return cls.__instance

    def business_logic(self):
        return id(self)


if __name__ == '__main__':
    s1, s2 = Singleton(), Singleton()
    print(
        'Does Singleton work correctly:',
        s1.business_logic() == s2.business_logic()
    )
