docker-build:
	docker compose -f infrastructure/docker/docker-compose.yaml --profile prod up --build

docker-up:
	docker compose -f infrastructure/docker/docker-compose.yaml --profile prod up -d

docker-push:
	docker compose -f infrastructure/docker/docker-compose.yaml --profile prod push

docker-down:
	docker compose -f infrastructure/docker/docker-compose.yaml --profile prod down

docker-up_dev:
	docker compose -f infrastructure/docker/docker-compose.yaml --profile dev up --build

docker-down_dev:
	docker compose -f infrastructure/docker/docker-compose.yaml --profile dev down


docker-logs:
	docker compose logs -f

