APP_NAME=carmasearch-web
PORT=3000
PACKAGE_MANAGER=pnpm

.PHONY: help install dev build start lint format clean docker-build docker-run

help:
	@echo "Available commands:"
	@echo "  make install        Install dependencies"
	@echo "  make dev            Run Next.js dev server"
	@echo "  make build          Build Next.js app"
	@echo "  make start          Start production server"
	@echo "  make lint           Run eslint"
	@echo "  make format         Run prettier"
	@echo "  make clean          Remove build artifacts"
	@echo "  make docker-build   Build docker image"
	@echo "  make docker-run     Run docker container"

install:
	$(PACKAGE_MANAGER) install

dev: install
	$(PACKAGE_MANAGER) dev --port $(PORT)

build: install
	$(PACKAGE_MANAGER) build

start: install
	$(PACKAGE_MANAGER) start --port $(PORT)

lint: install
	$(PACKAGE_MANAGER) lint

format: install
	$(PACKAGE_MANAGER) prettier --write .

clean: install
	rm -rf .next out node_modules

docker-build:
	docker build -t $(APP_NAME):latest .

docker-run:
	docker run -p $(PORT):3000 $(APP_NAME):latest
