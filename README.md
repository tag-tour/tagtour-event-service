# как поднять

1. ```git clone git@github.com:gleblagov/tagtour-events.git && cd tagtour-events```
2. поменять данные в:
 - ```.env-pgadmin```
 - ```.env-postgres```
 - ```config.yaml```
3. ```docker compose up --build``` или ```docker compose up --build -d``` -- запустить в фоне
4. ```docker compose down``` -- потушить

## параметры в конфиге
- ```username``` -- имя юзера от которого подключаться постгрес
- ```password``` -- пароль юзера
- ```db_name``` -- название бд, к которой подключаться

# контейнеры
1. ```localhost:2000``` -- api
2. ```localhost:2001``` -- pg-admin (логин ```admin@admin.admin``` пароль ```admin``` хост ```postgres```)
3. ```localhost:5432``` -- торчит постгрес (вольюм примонтирован в ./postgres)

# ручки 

1. ```GET /``` -- все события
2. ```GET /:id``` -- событие по id
3. ```POST /``` -- новое событие
4. ```PATCH /:id``` -- обновление события
5. ```DELETE /:id``` -- удаление события
6. ```GET /version``` -- проверить подключение к бд, должно вернуть версию


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
- логи