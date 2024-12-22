# calc_server

calc_server представлеяет HTTP API, на который можно отправлять простые арифметические выражения в формате JSON и получать в ответ результат

## EndPoints
#### 1. Проверка сервера

**Endpoint**: `/health`
**Метод**: `GET`
**Описание**: `Позволяет проверить состояние сервера`

**Успешный ответ**:
"It's alive!"

**Возможные статусы**:
- 200 - Успешно

#### 2. Вычисление

**Endpoint**: `/calculate`
**Метод**: `POST`
**Описание**: `Вычисляет передданое арифметическое выражение`

**Тело запроса**:
```json
{
    "expression": "string",
}
```

**Успешный ответ**:
```json
{
    "result": "number",
}
```

**Неуспешный ответ**:
```json
{
    "error": "string",
}
```

**Возможные статусы**:
- 200 - Успешно
- 400 - Неправильное тело запроса
- 422 - Неправильное арифметическое выражение

## Шаги для запуска проекта
#### 1. Клонирование репозитория с GitHub

```bash
git clone https://github.com/m3owmurrr/calc_server
```

#### 2. Запуск сервера

```bash
go run .\cmd 
```

## Переменные окружения
В проекте используются следующие перемеенные окружения для конфигурации системы:
- **RUN_TYPE** - позволяет запустить серверверный или консольный вариант приложения ("Server"\\"Local")
- **HOST** - Хост для запуска сервера
- **PORT** - Порт, который прослушивает сервер

По умолчанию используется режим сервера на адрессе "localhost:8080"

## Примеры запросов

1. Корректная работа:
```bash
curl --location -X POST 'localhost:8080/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2"
}'
```

Результат:
{
    "result":4
}

Статус: 200


2. Ошибочное выражение:
```bash
curl --location -X POST 'localhost:8080/calculate' \
--header 'Content-Type: application/json' \
--data '{
  "expression": "2+2+"
}'
```

Результат:
{
    "error":"expression is not valid"
}

Статус: 422

3. Ошибочный запрос:
```bash
curl --location -X POST 'localhost:8080/calculate' \
--header 'Content-Type: application/json' \
--data '
  "expression": "2+2"
'
```

Результат:
{
    "error":"sended data is invalid"
}

Статус: 400

4. Ошибочный метод:
```bash
curl --location -X GET 'localhost:8080/calculate' \
--header 'Content-Type: application/json' \
--data '
  "expression": "2+2"
'
```

Результат:
"Method Not Allowed"

Статус: 405

5. Проверка сервера:
```bash
curl --location -X GET 'localhost:8080/health'
```

Результат:
"It's alive!"

Статус: 200
