"""
Модуль с приложением-подписчиком на сообщения очереди RabbitMQ.
"""
from logging import getLogger, Logger, StreamHandler
from sys import stdout
from threading import Thread, Event

from .common import RabbitMQMixin, EXCHANGE, EX_TYPE


class LogService(Thread, RabbitMQMixin):
    """Сервис, имитирующий подписчика на сообщения"""

    instances = 0
    xch_name = EXCHANGE
    xch_type = EX_TYPE

    def __init__(self, rabbit_key: str):
        Thread.__init__(self, daemon=True)
        LogService.instances += 1
        self.number = LogService.instances
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
        msg = f'[{self}] [{method.routing_key}] GET request "{path}" from {remote} '
        self.log.info(msg)

    def run(self):
        """Запускает 'прослушку' очереди"""
        print(f'[{self}] Service started')
        self.rabbit.start_consuming()

    def __str__(self):
        return f'LogService-{self.number}'


def run_many(*routing_keys):
    """Запускает несколько экземпляров LogService"""
    threads = [LogService(key) for key in routing_keys]
    event = Event()

    for thread in threads:
        thread.start()

    try:
        event.wait()
    except KeyboardInterrupt:
        event.set()


if __name__ == '__main__':
    run_many(
        'app.index',     # забирает логи только главной страницы
        'app.other',     # забирает логи остальных страниц
        'index',         # забирает логи всех страниц
    )
