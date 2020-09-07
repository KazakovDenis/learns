"""
Hello-world пример для работы с RabbitMQ.
Consumer - подписчик на сообщения.
"""
from logging import Logger, StreamHandler
from sys import stdout

from pika import BlockingConnection, ConnectionParameters

from .common import RabbitMQMixin


def receive_msg_example(queue: str):
    """Пример подписки на сообщения очереди

    :param queue: название очереди
    """

    # подключаемся к серверу
    connection = BlockingConnection(ConnectionParameters('localhost'))
    channel = connection.channel()

    # проверяем, что очередь сущетсвует, или создаем
    channel.queue_declare(
        queue=queue,
        durable=True,
    )

    # пишем callback для обработки сообщений
    def callback(ch, method, properties, body):
        """Обязательная фукнция для обработки сообщений. Должна принимать 4 параметра.

        :param ch: экземпляр BlockingChannel
        :param method: экземпляр Basic.Deliver
        :param properties: экземпляр BasicProperties
        :param body: непосредственно тело сообщения
        """
        print('Message: ', body)

        # отправка подтверждения, что сообщение обработано
        # не требуется, если auto_ack=True
        # если получение не подтверждается, сообщения не удаляются из очереди
        ch.basic_ack(delivery_tag=method.delivery_tag)

    # не давать сообщения подписчику, пока обработка предыдущего не подтверждена
    channel.basic_qos(prefetch_count=1)

    # подписываемся на очередь
    channel.basic_consume(
        queue=queue,                         # название очереди
        on_message_callback=callback,        # функция-обработчик
        # auto_ack=True,                     # подтверждать ли обработку сообщения автоматически
    )

    try:
        # запускаем мониторинг очереди
        channel.start_consuming()
    except KeyboardInterrupt:
        channel.stop_consuming()
        exit(0)


# Ниже пример на приложении:
class LogService(RabbitMQMixin):
    """Сервис, имитирующий подписчика на сообщения"""

    def __init__(self):
        self.log = Logger(self.__class__.__name__)
        handler = StreamHandler(stdout)
        handler.setLevel(10)
        self.log.addHandler(handler)

        RabbitMQMixin.__init__(self)
        self.consume_from_rabbit(self.on_message)

    def on_message(self, ch, method, properties, body):
        """Функция для обработки сообщения из очереди RabbitMQ"""
        msg = f'[LogService] [{method.routing_key}] GET request from {body.decode()} '
        self.log.info(msg)

    def run(self):
        try:
            print('[LogService] Service started')
            self.rabbit.start_consuming()
        finally:
            print('[LogService] Shutting down...')
            self.rabbit.stop_consuming()
            exit(0)


if __name__ == '__main__':
    LogService().run()
