# RabbitMQ
"Hello world" пример для работы с Redis.  
  
## Источники
- [Get started](https://redislabs.com/get-started-with-redis/)
- [Документация](https://redis.io/documentation)
- [Best practices](https://habr.com/ru/post/485672/)
- [Шпаргалка](https://habr.com/ru/post/204354/)

## Подготовка
Запускаем сервер в docker-контейнере:
```
docker run -d --name redis -p 6379:6379 redis
```
По умолчанию сервер работает на порту 6379.  
Устанавливаем библиотеку: `pip install redis`  

## Работа

## CLI
Войти в CLI
```
docker exec -it redis redis-cli
```
