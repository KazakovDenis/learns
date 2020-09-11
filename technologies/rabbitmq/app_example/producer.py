"""
Модуль с приложением-источником сообщений для RabbitMQ.

Представляет собой http-сервер, который при GET-запросе отправляет сообщение в очередь.
"""
from datetime import datetime
from http.server import HTTPServer, BaseHTTPRequestHandler

from .common import RabbitMQMixin


class HTTPRequestHandler(BaseHTTPRequestHandler):
    """Обработчик HTTP-запроса.

    Использует метод send_to_rabbit миксина сервера для отправки сообщения в очередь.
    """

    def do_GET(self):
        """На GET-запрос возвращаем текущее время и отправляем в очередь URL и адрес клиента"""
        self.send_response(200)
        self.end_headers()
        now = datetime.now().isoformat(timespec='seconds').encode()
        self.wfile.write(b'<h1>This request\'s time: %s</h1>' % now)

        msg = '%s %s' % (self.address_string(), self.path)
        queue = 'app.index' if self.path == '/' else 'app.other'
        self.server.send_to_rabbit(msg, self.server.exchange, queue)


class AppServer(HTTPServer, RabbitMQMixin):
    """Сервер, имитирующий реальное приложение"""

    exchange = 'app'

    def __init__(self, host, port):
        super().__init__((host, port), HTTPRequestHandler)
        self.rabbit = self.connect_to_rabbit()
        self.rabbit.exchange_declare(self.exchange, exchange_type='topic')

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
