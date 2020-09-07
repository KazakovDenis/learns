"""
Hello-world пример для работы с RabbitMQ.
Producer - отправитель сообщения.
"""
from datetime import datetime
from http.server import HTTPServer, BaseHTTPRequestHandler

from pika import BlockingConnection, BasicProperties, ConnectionParameters

from .common import RabbitMQMixin


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
class HTTPRequestHandler(BaseHTTPRequestHandler):
    """Обработчик HTTP-запроса.
    Использует метод send_to_rabbit миксина сервера для отправки сообщения в очередь.
    """

    def do_GET(self):
        """На GET-запрос возвращаем текущее время и отправляем в очередь адрес клиента"""
        self.send_response(200)
        self.end_headers()
        now = datetime.now().isoformat(timespec='seconds').encode()
        self.wfile.write(b'<h1>This request\'s time: %s</h1>' % now)
        self.server.send_to_rabbit(self.address_string())


class AppServer(HTTPServer, RabbitMQMixin):
    """Сервер, имитирующий реальное приложение"""

    def __init__(self, host, port):
        super().__init__((host, port), HTTPRequestHandler)
        RabbitMQMixin.__init__(self)

    def run(self):
        host, port = self.server_address[:2]
        try:
            print(f'[AppServer] Serving HTTP on {host} port {port}\n')
            self.serve_forever()
        except KeyboardInterrupt:
            print('[AppServer] Shutting down...')
        finally:
            self.rabbit.close()
            self.server_close()


if __name__ == '__main__':
    AppServer('localhost', 8000).run()
