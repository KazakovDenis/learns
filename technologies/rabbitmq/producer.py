"""
Hello-world пример для работы с RabbitMQ.
Producer - отправитель сообщения.
"""
from pika import BlockingConnection, BasicProperties, ConnectionParameters

from .app import *


# подключаемся к серверу
connection = BlockingConnection(ConnectionParameters('localhost'))
channel = connection.channel()

# проверяем, что очередь сущетсвует, или создаем новую
channel.queue_declare(
    queue=queue,                    # название
    durable=True,                   # объявить устойчивой
)

# отправляем сообщение
channel.basic_publish(
    exchange='',                    # точка обмена
    routing_key=queue,              # имя очереди
    body=msg,                       # сообщение
    properties=BasicProperties(
        delivery_mode=2,            # объявить устойчивым
    )
)

connection.close()
