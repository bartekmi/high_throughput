.DEFAULT_GOAL := help
.PHONY: help dev


up: ## Run all the services
	docker compose up

down: ## Shut down the services
	docker compose down

# sh-client: ## Open a shell with all dependencies
# 	docker compose run client sh

sh-write: ## Open a shell in write container
	docker compose run write sh

sh-read: ## Open a shell in read container
	docker compose run read sh

build: ## Build the docker images
	docker compose build

# open-app: ## Open the React app
# 	open http://localhost:3000

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
