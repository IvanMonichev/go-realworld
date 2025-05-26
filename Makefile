# Имя бинарного файла
BINARY_NAME=real_world

# Путь к SQLite БД
DB_PATH=real_world.db

# Запуск приложения
run:
	go run ./cmd/main.go

# Сборка бинарника
build:
	go build -o $(BINARY_NAME) ./cmd/main.go

# Форматирование кода
fmt:
	go fmt ./...

# Проверка линтером (нужен golangci-lint)
lint:
	golangci-lint run

# Тесты с покрытием
test:
	go test ./... -v -coverprofile=coverage.out

# Очистка бинарников и кэша
clean:
	rm -f $(BINARY_NAME) coverage.out

# Запуск миграции, если ты подключишь migrate tool
migrate:
	go run ./internal/infrastructure/database/migrate.go --db $(DB_PATH)
