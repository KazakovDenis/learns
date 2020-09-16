"""
Модуль с общей логикой работы с RabbitMQ
"""
from typing import Union

from pika import BlockingConnection, ConnectionParameters, BasicProperties
from pika.adapters.blocking_connection import BlockingChannel


EXCHANGE = 'app'
EX_TYPE = 'direct'


class RabbitMQMixin:
    """Миксин для работы с очередью RabbitMQ"""

    rabbit: BlockingChannel

    @staticmethod
    def connect_to_rabbit(exchange_name: str = '', exchange_type: str = ''):
        """Подключаемся к RabbitMQ и создаём обменник

        :param exchange_name: название обменника
        :param exchange_type: тип обменника
        """
        connection = BlockingConnection(ConnectionParameters())
        channel = connection.channel()
        channel.exchange_declare(exchange_name, exchange_type)
        return channel

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

    def subscribe(self, exchange: str = '', *routing_keys: str):
        """Завязываем очереди на обменник и подписываемся

        :param exchange: название обменника
        :param routing_keys: ключи для привязки очередей
        """
        m = self.rabbit.queue_declare('', durable=True)
        queue = m.method.queue

        for key in routing_keys:
            self.rabbit.queue_bind(queue, exchange, key)

        self.rabbit.basic_consume(
            queue=queue,
            on_message_callback=self.on_message,
            auto_ack=True,
        )

    def on_message(self, *args):
        """Метод для обработки сообщения из очереди RabbitMQ"""
        pass
