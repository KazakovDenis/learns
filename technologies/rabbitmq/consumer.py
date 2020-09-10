"""
'Hello world' пример для работы с RabbitMQ.
Consumer - подписчик на сообщения.
"""
from pika import BlockingConnection, ConnectionParameters


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


if __name__ == '__main__':
    receive_msg_example('test')
