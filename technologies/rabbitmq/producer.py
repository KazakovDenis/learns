"""
'Hello world' пример для работы с RabbitMQ.
Producer - отправитель сообщения.
"""
from pika import BlockingConnection, BasicProperties, ConnectionParameters


def send_msg_example(queue: str, message: str):
    """Пример отправки сообщения в очередь

    :param queue: название очереди
    :param message: тело сообщения
    """

    # подключаемся к серверу
    connection = BlockingConnection(ConnectionParameters('localhost'))
    channel = connection.channel()

    # проверяем, что очередь сущетсвует, или создаем новую
    # method = channel.queue_declare('') создаст временную очередь со случайным именем
    channel.queue_declare(
        queue=queue,                    # название
        durable=True,                   # объявить устойчивой
    )

    # необязательно: создаём обменник и связываем с очередью
    # channel.exchange_declare('logs', exchange_type='fanout')
    # channel.queue_bind(method.method.queue, 'logs')
    # с типом fanout при отправке сообщения routing_key можно не указывать

    # отправляем сообщение
    channel.basic_publish(
        exchange='',                    # точка обмена
        routing_key=queue,              # имя очереди
        body=message,                   # сообщение
        properties=BasicProperties(
            delivery_mode=2,            # объявить устойчивым
        )
    )
    connection.close()


if __name__ == '__main__':
    send_msg_example('test', 'Hello world!')
