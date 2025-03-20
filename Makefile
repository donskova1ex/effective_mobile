include scripts/*.mk

DEV_COMPOSE_ARGS=--env-file config/.env.dev -f docker-compose.dev.yaml
DEV_COMPOSE_ENV=docker compose $(DEV_COMPOSE_ARGS)
DEV_COMPOSE=docker compose $(DEV_COMPOSE_ARGS)

dev-build: dev-api-build
	$(DEV_COMPOSE) build

dev-up: dev-build dev-api-up
	$(DEV_COMPOSE) --env-file config/.env.dev up -d

dev-migrate-up:
	docker-compose -f docker-compose.dev.yaml up -d migrations-up

dev-migrate-down:
	docker compose --profile migrations-down -f docker-compose.dev.yaml up -d migrations-down

dev-api-build: api_docker_build

dev-api-up:
	$(DEV_COMPOSE) --profile api-up up -d api-up