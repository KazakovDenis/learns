# RabbitMQ
"Hello world" пример для работы с RabbitMQ.  
  
![](https://avatars.mds.yandex.net/get-zen_doc/1591747/pub_5e00562d1e8e3f00b0e56b24_5e0059d6b477bf00af3fd8d9/scale_1200)

## Источники
- [Официальный сайт](https://www.rabbitmq.com/getstarted.html)
- [Хабр](https://habr.com/ru/post/149694/)
- [Дзен](https://zen.yandex.ru/media/id/5de8a5395d636200b075e410/amqp-na-primere-rabbitmq-kak-je-gotovit-krolika-5e00562d1e8e3f00b0e56b24)

## Подготовка
Запускаем сервер в docker-контейнере (обязательно указать имя хоста):
```
docker run -d --hostname my-rabbit --name rabbit-server --network host rabbitmq
```
По умолчанию сервер работает на порту 5672.  
Устанавливаем библиотеку для работы с AMQP: `pip install pika`  

## Работа
Отправка сообщений описана в *producer.py*, подписка - *consumer.py*.

### По умолчанию
- RabbitMQ хранит сообщения в оперативной памяти. Для сохранения сообщений на диске, необходимо объявить 
очередь устойчивой при создании: `durable=True`, а также указать `delivery_mode=2` для сообщений. 
Для 100% устойчивости необходимо оборачивать операции в транзакции.
- Сообщения приходят в безымянный обменник, имеющий тип direct, т.е. сообщение направляется только в очередь
 с точным соответствием routing_key

## Прочее
- binding_key у exchange представляет собой список слов, разделённых точкой, может быть не длинее 255 байт и
иметь подстановки для типа topic: * - одно слово, \# - 0 или более слов. 

## CLI
Включить веб-интерфейс (доступен на порту 15672):
```
docker exec -it rabbit-server rabbitmq-plugins enable rabbitmq_management
```
Проверить, существует ли очередь:
```
docker exec -it rabbit-server rabbitmqctl list_queues
```
Вывести неподтвержденные сообщения:
```
docker exec -it rabbit-server rabbitmqctl list_queues name messages_ready messages_unacknowledged
```
Просмотреть обменники:
```
docker exec -it rabbit-server rabbitmqctl list_exchanges
```