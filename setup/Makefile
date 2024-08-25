include .env

WORKERZ_BINARY=workerz

frontend:
	@echo "starting frontend"
	cd ../frontend && npm run dev

## up: starts all containers in the background without forcing build
start:
	@echo "Starting Docker images..."
	docker compose up -d
	@echo "Docker images started!"

down:
	@echo "killing docker containers"
	docker compose down

## up_build: stops docker-compose (if running), builds all projects and starts docker compose
build: build_workerz 
	@echo "Stopping docker images (if running...)"
	docker compose down
	@echo "Building (when required) and starting docker images..."
	docker compose up --build -d
	@echo "Docker images built and started!"

## workerz: builds the workerz binary as a linux executable
build_workerz:
	@echo "Building workerz binary..."
	cd ../server && env GOOS=linux CGO_ENABLED=0 go build -o ${WORKERZ_BINARY} .
	@echo "Done!"

m-up-d:
	#migrate db up
	migrate -path ../server/database/migrations -database "postgresql://postgres:password@localhost:5432/workerz?sslmode=disable" up

m-down-d:
	#migrate db down
	migrate -path ../server/database/migrations -database "postgresql://postgres:password@localhost:5432/workerz?sslmode=disable" down

#create migrations folder
sqlxm:
	cd ../server && sqlx migrate add -r init

sqlc:
	#generate sql commands
	cd ../server && sqlc generate

test:
	#test go
	go test -v -cover ./...