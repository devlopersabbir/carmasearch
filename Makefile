.PHONY: web server scraper infrastructure

# ---------- Infra ----------
docker-build:
	docker compose -f infrastructure/docker/docker-compose.yaml --profile prod build --no-cache

docker-up:
	docker compose -f infrastructure/docker/docker-compose.yaml --profile prod up -d

docker-down:
	docker compose -f infrastructure/docker/docker-compose.yaml --profile prod down

docker-dev:
	docker compose -f infrastructure/docker/docker-compose.dev.yaml --profile dev up -d

# ---------- Apps ----------
web-%:
	$(MAKE) -C apps/web $*

server-%:
	$(MAKE) -C apps/server $*

scraper-%:
	$(MAKE) -C apps/scraper $*
