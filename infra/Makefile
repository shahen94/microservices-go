BROKER_BINARY=broker-service
AUTH_BINARY=auth-service

up: build_broker build_auth
	@echo "Starting Docker images..."
	docker compose up -d
	@echo "Docker images started."
up_build: build_broker build_auth
	@echo "Stopping Docker images (if running)..."
	docker compose down
	@echo "Building (when required) and starting Docker images..."
	docker compose up --build -d
	@echo "Docker images built started."
down:
	@echo "Stopping Docker images (if running)..."
	docker compose down
	@echo "Docker images stopped."
build_broker:
	@echo "Building Docker image for broker..."
	cd ../broker-service  && env GOOS=linux CGO_ENABLED=0 go build -o ${BROKER_BINARY} ./cmd/api
	@echo "Docker image for broker built."

build_auth:
	@echo "Building Docker image for broker..."
	cd ../auth-service  && env GOOS=linux CGO_ENABLED=0 go build -o ${AUTH_BINARY} ./cmd/api
	@echo "Docker image for broker built."