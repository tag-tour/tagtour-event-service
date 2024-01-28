# как собрать

1. ```git clone git@github.com:gleblagov/tagtour-events.git && cd tagtour-events```
2. ```docker compose up --build``` или ```docker compose up --build -d``` -- запустить в фоне

# контейнеры
1. ```localhost:2000``` -- api
2. ```localhost:2001``` -- pg-admin (логин ```admin@admin.admin``` пароль ```admin``` хост ```postgres```)
3. ```localhost:5432``` -- торчит постгрес

# ручки 

1. ```GET /``` -- все события
2. ```GET /:id``` -- событие по id
3. ```POST /``` -- новое событие
4. ```PATCH /:id``` -- обновление события
5. ```DELETE /:id``` -- удаление события


# примеры запроса/ответа

## тело запроса POST /
```
{
  "title": "title",
  "media": [
    "media1",
    "media2"
  ],
  "author": 280205,
  "date": "2024-01-28T17:49:38.774008Z",
  "description": "description",
  "members": [
    1,
    2,
    3
  ]
}
```

## тело ответа
```
[
  {
    "id": 1,
    "title": "title",
    "likes": 0,
    "media": [
      "media1",
      "media2"
    ],
    "author": 280205,
    "createdAt": "2024-01-28T18:00:00.774008Z",
    "date": "2024-01-28T17:49:38.774008Z",
    "description": "description",
    "members": [
      1,
      2,
      3
    ]
  }
```

# доделать
- нормальные хттп коды и тела в ответах
- вынести всё чувствительное в конфиг файл и .env