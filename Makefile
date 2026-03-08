.PHONY: Carmasearch infrastructure

# ---------- Infra ----------
docker-build:
	docker compose -f infrastructure/docker/docker-compose.yaml --profile prod build --no-cache

docker-up:
	docker compose -f infrastructure/docker/docker-compose.yaml --profile prod up -d

docker-down:
	docker compose -f infrastructure/docker/docker-compose.yaml --profile prod down

docker-dev:
	docker compose -f infrastructure/docker/docker-compose.dev.yaml --profile dev up -d

k8s-es-apply:
	kubectl apply -f infrastructure/k8s/elasticsearch-statefulset.yaml

k8s-es-delete:
	kubectl delete -f infrastructure/k8s/elasticsearch-statefulset.yaml

# ---------- Apps ----------
web-%:
	$(MAKE) -C apps/web $*

server-%:
	$(MAKE) -C apps/server $*

in-%:
	$(MAKE) -C apps/intelligence $*

scraper-%:
	$(MAKE) -C apps/scraper $*
