# GO Realworld

This project implements JSON API as specified in RealWorld spec.

```
/cmd/
  └── main.go                  # точка входа
/internal/
  /article/
    ├── handler.go             # HTTP (gin) обработчики
    ├── service.go             # бизнес-логика
    ├── repository.go          # доступ к данным
    ├── model.go               # DTO / бизнес-сущности
    ├── validator.go           # валидация
    └── serializer.go          # маппинг DTO → JSON
  /user/
    ├── handler.go
    ├── service.go
    ├── repository.go
    ├── model.go
    ├── validator.go
    ├── serializer.go
    └── middleware.go
/common/
    ├── database.go
    └── logger.go
/pkg/
    └── utils.go               # общие утилиты
/config/
    └── config.go              # env / yaml загрузка
/migrations/

```