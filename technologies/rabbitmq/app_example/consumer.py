"""
Модуль с приложением-подписчиком на сообщения очереди RabbitMQ.
"""
from logging import getLogger, Logger, StreamHandler
from sys import stdout
from threading import Thread, Event

from .common import RabbitMQMixin


class LogService(Thread, RabbitMQMixin):
    """Сервис, имитирующий подписчика на сообщения"""

    instances = 0

    def __init__(self, rabbit_key: str):
        Thread.__init__(self, daemon=True)
        LogService.instances += 1

        self.number = LogService.instances
        self.log = self.get_logger()

        self.rabbit = self.connect_to_rabbit()
        self.consume_from_rabbit(self.on_message, rabbit_key)

    def get_logger(self) -> Logger:
        logger = getLogger(str(self))
        handler = StreamHandler(stdout)
        handler.setLevel(10)
        logger.addHandler(handler)
        return logger

    def declare(self, *routing_keys: str, exchange: str = ''):
        """Подключаемся к RabbitMQ и создаём очередь"""
        # if exchange:
        #     self.rabbit.exchange_declare(exchange, exchange_type='topic')

        for key in routing_keys:
            res = self.rabbit.queue_declare(key, durable=True)
            # if exchange:
            #     self.rabbit.queue_bind(res.method.queue, exchange, routing_key=key)

    def on_message(self, ch, method, properties, body):
        """Функция для обработки сообщения из очереди RabbitMQ"""
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

    for thread in threads:
        thread.start()

    try:
        Event().wait()
    except KeyboardInterrupt:
        raise SystemExit


if __name__ == '__main__':
    run_many(
        'app.index',     # забирает логи только главной страницы
        'app.other',     # забирает логи остальных страниц
        'app.*',         # забирает логи всех страниц
    )
