"""
Hello-world пример для работы с RabbitMQ.
Producer - отправитель сообщения.
"""
from pika import BlockingConnection, ConnectionParameters


# подключаемся к серверу
connection = BlockingConnection(ConnectionParameters('localhost'))
channel = connection.channel()

# проверяем, что очередь сущетсвует, или создаем новую
channel.queue_declare(queue='hello')

# отправляем сообщение
channel.basic_publish(
    exchange='',              # точка обмена
    routing_key='hello',      # имя очереди
    body='Hello World!'       # сообщение
)

connection.close()
