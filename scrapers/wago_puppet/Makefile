start: ## Start service in background
	docker compose up -d

stop: ## Stop and remove containers (does NOT remove volumes)
	docker compose down

logs: ## Follow logs
	docker compose logs -f

rebuild: ## Rebuild image and restart
	docker compose build --no-cache
	docker compose up -d

.PHONY: start stop logs rebuild