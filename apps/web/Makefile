APP_NAME=nextjs-app
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

build:
	$(PACKAGE_MANAGER) build

start:
	$(PACKAGE_MANAGER) start --port $(PORT)

lint:
	$(PACKAGE_MANAGER) lint

format:
	$(PACKAGE_MANAGER) prettier --write .

clean:
	rm -rf .next out node_modules

docker-build:
	docker build -t $(APP_NAME):latest .

docker-run:
	docker run -p $(PORT):3000 $(APP_NAME):latest
