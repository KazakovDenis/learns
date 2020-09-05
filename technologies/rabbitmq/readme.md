# RabbitMQ
Hello-world пример для работы с RabbitMQ.

## Подготовка
Запускаем сервер в docker-контейнере (обязательно указать имя хоста):
```
docker run -d --hostname my-rabbit --name rabbit-server --network host rabbitmq
```
По умолчанию сервер работает на порту 5672.  
Устанавливаем библиотеку для работы с AMQP: `pip install pika`  

## Работа
Отправка сообщений описана в *producer.py*, подписка - *consumer.py*.

## Прочее
Проверить, существует ли очередь:
```
docker exec -it rabbit-server rabbitmqctl list_queues
```
