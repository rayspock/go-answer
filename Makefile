.PHONY: dev test docs run

dev:
	go run main.go

test:
	go test -v ./...

docs: ## Generate swagger api documentation
	swag init -g routes/routes.go

# Docker Tasks
run: ## Spin up the app and database
	docker-compose -f docker-compose.yml -p go-answer up -d

start-dev-db: ## Spin up database for development
	# run your container
	docker-compose -f docker-compose-dev.yml -p dev up -d

stop-dev-db: ## Stop database
	docker-compose -p dev stop
