"""
Hello-world пример для работы с RabbitMQ.
Consumer - подписчик на сообщения.
"""
from pika import BlockingConnection, ConnectionParameters


# подключаемся к серверу
connection = BlockingConnection(ConnectionParameters('localhost'))
channel = connection.channel()

# проверяем, что очередь сущетсвует, или создаем
channel.queue_declare(queue='hello')


# пишем callback для обработки сообщений
def callback(ch, method, properties, body):
    """Обязательная фукнция для обработки сообщений. Должна принимать 4 обязательных параметра.

    :param ch: экземпляр BlockingChannel
    :param method: экземпляр Basic.Deliver
    :param properties: экземпляр BasicProperties
    :param body: непосредственно тело сообщения
    """
    print('Message: ', body)


# подписываемся на очередь
channel.basic_consume(
    queue='hello',
    on_message_callback=callback,
)

# запускаем мониторинг очереди
channel.start_consuming()
