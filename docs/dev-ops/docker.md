## List of docker service

```yaml
services:
  web:
    build: apps/web
    ports: ["3000:3000"]

  api:
    build: apps/api
    ports: ["8080:8080"]

  scraper:
    build: apps/scraper
    ports: ["8000:8000"]

  postgres:
    image: postgres:16

  redis:
    image: redis:7
```
