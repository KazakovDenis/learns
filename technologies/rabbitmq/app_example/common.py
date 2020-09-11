"""
Модуль с общей логикой работы с RabbitMQ
"""
from typing import Callable, Union

from pika import BlockingConnection, ConnectionParameters, BasicProperties
from pika.adapters.blocking_connection import BlockingChannel


class RabbitMQMixin:
    """Миксин для работы с очередью RabbitMQ"""

    rabbit: BlockingChannel

    @staticmethod
    def connect_to_rabbit():
        """Подключаемся к RabbitMQ и создаём очередь"""
        connection = BlockingConnection(ConnectionParameters())
        return connection.channel()

    def send_to_rabbit(self, message: Union[str, bytes], exchange: str = '', queue: str = ''):
        """Отправляем сообщение в очередь

        :param message: тело сообщения
        :param queue: ключ маршрутизации
        :param exchange: название обменника
        """
        self.rabbit.basic_publish(
            exchange=exchange,
            routing_key=queue,
            body=message,
            properties=BasicProperties(delivery_mode=2),
        )

    def consume_from_rabbit(self, callback: Callable, queue: str = ''):
        """Подписываемся на сообщения очереди

        :param callback: функция для обработки сообщения
        :param queue: название очереди
        """
        self.rabbit.basic_qos(prefetch_count=1)
        self.rabbit.basic_consume(
            queue=queue,
            on_message_callback=callback,
            auto_ack=True,
        )
