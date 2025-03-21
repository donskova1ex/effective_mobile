## Songs API Service for effective_mobile

REST API сервис для управления музыкальной библиотекой. Позволяет хранить и управлять информацией о песнях, включая название группы, название песни, дату релиза, текст песни и ссылку.

## Технологии

- Go 1.23
- PostgreSQL 15
- Docker & Docker Compose
- Goose (миграции)
- Gorilla Mux (роутинг)
- OpenAPI/Swagger (документация API)

## Требования

- Docker
- Docker Compose
- Make

## Быстрый старт

1. Клонируйте репозиторий:
```bash
git clone https://github.com/donskova1ex/effective_mobile
```

2. Создайте файл конфигурации:
```bash
cp config/.env.example config/.env.dev
```

3. Запустите сервис:
```bash
make dev-up
```

Сервис будет доступен по адресу: http://localhost:8080

## Структура проекта

- `api/` - HTTP хендлеры и роутинг
- `cmd/` - точки входа приложения
- `config/` - конфигурационные файлы
- `internal/` - внутренняя логика приложения
- `migrations/` - SQL миграции
- `openapi/` - OpenAPI спецификации
- `scripts/` - вспомогательные скрипты

## Доступные команды

- `make dev-up` - запуск всех сервисов
- `make dev-migrate-up` - применить миграции
- `make dev-migrate-down` - откатить миграции
- `make dev-build` - сборка сервисов
- `make dev-api-up` - запуск только API сервиса

## База данных

Сервис использует PostgreSQL для хранения данных. Схема базы данных:

### Таблица songs
- `id` - уникальный идентификатор
- `group_name` - название группы
- `song_name` - название песни
- `release_date` - дата релиза
- `text` - текст песни
- `link` - ссылка на песню

### Тестовые данные
- The Beatles "Hey Jude" 1968-08-26
- Queen "Bohemian Rhapsody" 1975-10-31