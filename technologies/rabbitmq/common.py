from pika import BlockingConnection, ConnectionParameters, BasicProperties


QUEUE = 'app_log'


class RabbitMQMixin:
    """Миксин для работы с очередью RabbitMQ"""

    def __init__(self):
        self.rabbit = self.connect_to_rabbit()

    @staticmethod
    def connect_to_rabbit(queue=QUEUE):
        """Подключаемся к RabbitMQ и создаём очередь"""
        connection = BlockingConnection(ConnectionParameters('localhost'))
        channel = connection.channel()
        channel.queue_declare(queue=queue, durable=True)
        return channel

    def send_to_rabbit(self, message, queue=QUEUE):
        """Отправляем сообщение в очередь

        :param message: тело сообщения
        :param queue: название очереди
        """
        self.rabbit.basic_publish(
            exchange='',
            routing_key=queue,
            body=message,
            properties=BasicProperties(delivery_mode=2)
        )

    def consume_from_rabbit(self, callback, queue=QUEUE):
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
