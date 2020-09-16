"""
Модуль с приложением-подписчиком на сообщения очереди RabbitMQ.
"""
from logging import getLogger, Logger, StreamHandler
from sys import stdout
from threading import Thread, Event

from .common import RabbitMQMixin, EXCHANGE, EX_TYPE


class LogService(Thread, RabbitMQMixin):
    """Сервис, имитирующий подписчика на сообщения"""

    xch_name = EXCHANGE
    xch_type = EX_TYPE

    def __init__(self, rabbit_key: str):
        Thread.__init__(self, daemon=True)
        self.rabbit_key = rabbit_key
        self.log = self.get_logger()

        self.rabbit = self.connect_to_rabbit(self.xch_name, self.xch_type)
        self.subscribe(self.xch_name, rabbit_key)

    def get_logger(self) -> Logger:
        logger = getLogger(str(self))
        handler = StreamHandler(stdout)
        logger.addHandler(handler)
        logger.setLevel(10)
        handler.setLevel(10)
        return logger

    def on_message(self, ch, method, properties, body):
        """Метод для обработки сообщения из очереди RabbitMQ"""
        remote, path = body.decode().split()
        msg = f'{self} GET request from {remote} to "{path}"'
        self.log.info(msg)

    def run(self):
        """Запускает 'прослушку' очереди"""
        print(f'{self} Service started')
        self.rabbit.start_consuming()

    def __str__(self):
        return f'<LogService [{self.rabbit_key}]>'.ljust(25)


def run_many(*routing_keys):
    """Запускает несколько экземпляров LogService"""
    threads = [LogService(key) for key in routing_keys]

    for thread in threads:
        thread.start()

    event = Event()
    try:
        event.wait()
    except KeyboardInterrupt:
        event.set()


if __name__ == '__main__':
    run_many(
        'app.index',     # забирает логи только главной страницы
        'app.other',     # забирает логи остальных страниц
        'app.*',         # забирает логи всех страниц
    )
