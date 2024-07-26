# Messagio project
Проект для Messagio на позицию Golang разработчика


## Project Stack
- Golang
- Golang/Echo
- Golang/Gorm
- Postgres
- Docker
- Docker Compose
- Kafka

## Dependencies to run
- Docker
- Docker compose


## Запуск
### linux

`cp .env.example .env`

Чтобы запустить проект для разработки `start.dev.sh`

Чтобы запустить проект для продакшена `start.prod.sh`

## Endpoints
Доступ к серверу получается по ссылке `http://localhost:1234`

Префикс апи: `/api/v1`

```ts 
1. POST /message // Создает новый message 
Body: {
    content: string
}

Response: NoContent

2. GET /message // Запрос всех сообщений, которые есть в БД
Query: {
    limit: number; // max = 200, required
    page: number; // default = 0
}

Response: {
    values: []{
        id: number;
        content: string;
        createdAt: string;
    }
    total: number;
}

3. GET /message/:id  // Запрос одного сообщения по его айдишнику
Params: {
    id: number; // required
}

Response: {
    id: number;
    content: string;
    createdAt: string;
}

4. GET /message/process // Запустить обработку сообщений, которые еще не были обрабатаны из кафки
Response: []{
    id: number;
    content: string;
    createdAt: string;
}

5. GET /statistic/message/:id // 
Params: {
    id: number; // required
}

Response: {
    id: number;
    readTimes: []{ // Массив каждого раза, когда сообщение было отправлено клиенту
        readAt: Date;
    } ,
    kafkaProcessed: boolean; // Была ли запущена проверка и обработано ли это сообщения
    createdAt: Date;
}   
```

## Структура ДБ
https://dbdiagram.io/d/66a3a45f8b4bb5230e72b4fd