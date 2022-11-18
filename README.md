# Задача

Необходимо реализовать микросервис для работы с балансом пользователей (зачисление средств, списание средств, перевод средств от пользователя к пользователю, а также метод получения баланса пользователя). Сервис должен предоставлять HTTP API и принимать/отдавать запросы/ответы в формате JSON.

# Реализация

- "Чистая архитектура"
- REST API
- Использование gin-gonic
- Использование PostgreSQL

# Endpoints

- POST /accrual - Метод начисления средств на баланс. Принимает id пользователя (user_id) и сколько средств зачислить (amount)

- POST /reserveTransaction - Метод резервирования средств с основного баланса на отдельном счете. Принимает id пользователя (user_id, ИД услуги (service_id), ИД заказа (order_id), стоимость (amount)

- GET /getBalance - Метод получения баланса пользователя. Принимает id пользователя (user_id)

- POST /revenueRecognition - Метод признания выручки – списывает из резерва деньги, добавляет данные в отчет для бухгалтерии. Принимает id пользователя (user_id), ИД услуги (service_id), ИД заказа (order_id), сумму (amount)

# Запуск

Для того чтобы запустить приложение необходимо сперва выполнить скрипт, который создает таблицы в БД. После этого воспользоваться командой:
```
docker compose up
```
Это необходимо для развертывания приложения и БД в контейнерах

# Результаты работы

Для проверки работоспособности был использован Postman

### 1. POST /accrual

**Запрос:**
```
{
    "user_id": 1,
    "amount": 1000
}
```

**Ответ:**
```
{
    "id": 60
}
```

### 2. POST /reserveTransaction

**Запрос:**
```
{
    "user_id": 1,
    "service_id": 4,
    "order_id": 4,
    "amount": 88
}
```

**Ответ:**
```
{
    "id": 61
}
```

### 3. GET /getBalance

**Запрос:**
```
{
    "user_id": 1
}
```

**Ответ:**
```
{
    "balance": 2720
}
```

### 4. POST /revenueRecognition

**Запрос:**
```
{
    "user_id": 1,
    "service_id": 1,
    "order_id": 1,
    "amount": 999
}
```

**Ответ:**
```
{
    "id": 66
}
```
