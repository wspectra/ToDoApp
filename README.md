# :heavy_check_mark: ToDoApp
**ToDoApp** - Приложение по созданию списков дел 

# API documentation
https://app.swaggerhub.com/apis-docs/DMGRISHINA/todo_app/1.0

# В проекте используются следующие концепции:
- REST API.
- Работа с фреймворком <a href="https://github.com/gin-gonic/gin">gin-gonic/gin</a>.
- Подход Чистой Архитектуры в построении структуры приложения. Техника внедрения зависимости.
- Работа с БД Postgres. Запуск из Docker.
- Работа с БД, используя библиотеку <a href="https://github.com/jmoiron/sqlx">sqlx</a>.
- Регистрация и аутентификация. Работа с JWT. Middleware.
- Написание SQL запросов.
- Graceful Shutdown

# Usage
По умолчанию поднимается контейнер в котором работает сервис и база данных

    make

Выполняются тесты

    make test

Очистить все 

    make fclean
  
Посмотреть логи 

    make logs
    
Посмотреть статус 

    make status
    
## Other
**Author:**
:pig:**[wspectra](https://github.com/wspectra)**
