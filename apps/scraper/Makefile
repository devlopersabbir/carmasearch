APP_MODULE=app.main:app
PORT=8000
PYTHON=python3
PIP=pip3
IMAGE=devlopersabbir/carmasearch-scraper
TAG=latest

.PHONY: help install dev run run-prod build clean docker-build docker-run
help:
	@echo "Available commands:"
	@echo "  make install        Install python dependencies"
	@echo "  make dev            Run FastAPI in dev mode (auto-reload)"
	@echo "  make run            Run with uvicorn (dev-style)"
	@echo "  make run-prod       Run uvicorn in production mode"
	@echo "  make build          Build FastAPI app"
	@echo "  make docker-build   Build docker image"
	@echo "  make docker-run     Run docker container"
	@echo "  make clean          Remove __pycache__ files"


install:
	$(PIP) install --upgrade pip
	$(PIP) install -r requirements.txt

dev: install
	fastapi dev $(APP_MODULE) --reload --port $(PORT)

run:
	uvicorn $(APP_MODULE) --reload --port $(PORT)

run-prod:
	uvicorn $(APP_MODULE) --host 0.0.0.0 --port $(PORT) --workers 2

build:
	fastapi build $(APP_MODULE)

clean:
	find . -type d -name "__pycache__" -exec rm -rf {} +


docker-build:
	docker build -f Dockerfile.scraper -t $(IMAGE):$(TAG) .

docker-run:
	docker run --env-file .env -p $(PORT):8000 $(IMAGE):$(TAG)