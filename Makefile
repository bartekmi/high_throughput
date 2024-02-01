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

deploy: ## Build and deploy application to EKS
	# Authenticate
	aws ecr get-login-password --region us-east-2 | docker login --username AWS --password-stdin 975049909939.dkr.ecr.us-east-2.amazonaws.com
	# Build
	docker build -t ht-write -f Dockerfile.write .
	docker build -t ht-read -f Dockerfile.read .
	# Tag Images
	docker tag ht-write:latest 975049909939.dkr.ecr.us-east-2.amazonaws.com/ht:write-latest
	docker tag ht-read:latest 975049909939.dkr.ecr.us-east-2.amazonaws.com/ht:read-latest
	# Push Images
	docker push 975049909939.dkr.ecr.us-east-2.amazonaws.com/ht:write-latest
	docker push 975049909939.dkr.ecr.us-east-2.amazonaws.com/ht:read-latest
	# Restart the Pods
	kubectl -n ht rollout restart deployment/ht-deployment-write
	kubectl -n ht rollout restart deployment/ht-deployment-read

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
