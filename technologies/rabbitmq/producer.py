"""
Hello-world пример для работы с RabbitMQ.
Producer - отправитель сообщения.
"""
from datetime import datetime
from http.server import HTTPServer, BaseHTTPRequestHandler

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
    channel.queue_declare(
        queue=queue,                    # название
        durable=True,                   # объявить устойчивой
    )

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


# Ниже пример на приложении:
# Сервер принимает запросы и отправляет адрес клиента в очередь RabbitMQ
class RequestHandler(BaseHTTPRequestHandler):
    """Обработчик HTTP-запроса"""

    def do_GET(self):
        """На GET-запрос возвращаем текущее время и отправляем в очередь адрес клиента"""
        self.send_response(200)
        self.end_headers()
        now = datetime.now().isoformat(timespec='seconds').encode()
        self.wfile.write(b'<h1>This request time: %s</h1>' % now)
        self.send_to_rabbit()

    def send_to_rabbit(self):
        """Отправляем в очередь сообщение о совершенном на сервер запросе"""
        self.server.rabbit.basic_publish(
            exchange='',
            routing_key=self.server.queue,
            body=self.address_string(),
            properties=BasicProperties(delivery_mode=2)
        )


class AppServer(HTTPServer):
    """Сервер, имитирующий реальное приложение"""

    queue = 'app_log'

    def __init__(self, host, port):
        self.rabbit = self.connect_to_rabbit()
        super().__init__((host, port), RequestHandler)

    def connect_to_rabbit(self):
        """Подключаемся в RabbitMQ и создаём очередь"""
        connection = BlockingConnection(ConnectionParameters('localhost'))
        channel = connection.channel()
        channel.queue_declare(queue=self.queue, durable=True)
        return channel

    def run(self):
        host, port = self.server_address[:2]
        print(f"Serving HTTP on {host} port {port}\n")
        try:
            self.serve_forever()
        except KeyboardInterrupt:
            print('Shutting down the server...')
            self.rabbit.close()
            self.server_close()


if __name__ == '__main__':
    AppServer('localhost', 8000).run()
