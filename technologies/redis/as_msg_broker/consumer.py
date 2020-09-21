#!/usr/bin/env python3
"""Модуль с примерами подписки на сообщения

Для подписки необходимо:
- создать подключение к Redis
- подписаться на канал с помощью PubSub.subscribe / PubSub.psubscribe
- извлекать сообщения через методы PubSub.get_message или PubSub.listen

get_message - возвращает 1 сообщение
listen      - возвращает генератор

PubSub.subscribe/psubscribe позволяют передавать в аргументах собственный обработчик сообщения.
В таком случае методы get_message и listen будут возвращать None, всю работу с сообщением
необходимо реализовать в обработчике.
"""
import os
import sys
from typing import Callable

import redis


file = os.path.basename(__file__)
if len(sys.argv) < 2:
    raise Exception(f'Usage: python {file} [channel]')

CHANNEL = sys.argv[1]


def on_message(msg: dict) -> None:
    """Функция для обработки сообщений

    :param msg: словарь, передаваемый PubSub.handle_message
    """
    print('[HANDLER]', msg['data'])


def using_get_message(consumer):
    """Получает сообщения поштучно с помощью PubSub.get_message()

    Инорирует служебные сообщения. Если используется обработчик, msg = None
    """
    while True:
        msg = consumer.get_message(
            ignore_subscribe_messages=True,
            timeout=1
        )
        if msg:
            print('[NO HANDLER]', msg['data'])


def using_listen(consumer):
    """Получает сообщения с помощью генератора PubSub.listen()

    Берёт только сообщения с типом message, пропуская служебные.
    """
    for msg in consumer.listen():
        if msg['type'] == 'message':
            print('[NO HANDLER]', msg['data'])


def main(listen: Callable, channel: str, handler: Callable = None):
    """Запускает прослушку указанного канала

    :param listen: функция для прослушки с помощью get_message / listen
    :param channel: канал для прослушки
    :param handler: обработчик сообщения
    """
    conn = redis.Redis()
    sub = conn.pubsub()

    if handler:
        sub.subscribe(**{channel: handler})
    else:
        sub.subscribe(channel)

    print('Listen for messages on channel:', channel)
    try:
        listen(sub)
    except KeyboardInterrupt:
        print('Shutting down...')
        sub.close()


if __name__ == '__main__':
    # без обработчика
    # main(using_get_message, CHANNEL)
    # или
    # main(using_listen, CHANNEL)

    # с обработчиком
    # main(using_get_message, CHANNEL, on_message)
    # или
    main(using_listen, CHANNEL, on_message)
