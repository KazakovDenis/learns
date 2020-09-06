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

По умолчанию, RabbitMQ хранит сообщения в оперативной памяти. Для сохранения сообщений на диске, необходимо объявить 
очередь устойчивой при создании: `durable=True`, а также указать `delivery_mode=2` для сообщений. 
Для 100% устойчивости необходимо оборачивать операции в транзакции.

## Прочее
Проверить, существует ли очередь:
```
docker exec -it rabbit-server rabbitmqctl list_queues
```
Вывести неподтвержденные сообщения:
```
docker exec -it rabbit-server rabbitmqctl list_queues name messages_ready messages_unacknowledged
```
