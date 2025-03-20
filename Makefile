DEV_COMPOSE_ARGS=--env-file .env.dev -f docker-compose.dev.yaml
DEV_COMPOSE=docker compose $(DEV_COMPOSE_ARGS)

dev-build:
	$(DEV_COMPOSE) build --no-cache

dev-up:
	$(DEV_COMPOSE) up -d --no-deps --build

dev-migrate-up:
	$(DEV_COMPOSE) up -d migrations-up

dev-migrate-down:
	$(DEV_COMPOSE) --profile migrations-down up -d migrations-down