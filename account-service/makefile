.PHONY: run deps build-mongo build-postgres clean run-docker

# Default target
run:
	@echo "Starting account-service..."
	go run main.go

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go get -u github.com/joho/godotenv
	go get -u github.com/gin-gonic/gin
	go get -u github.com/lib/pq
	go get -u github.com/google/uuid
	go mod tidy
	go mod download


# Build PostgreSQL container
build-postgres:
	@if docker ps -a | grep -q postgres-banking; then \
		echo "PostgreSQL container already exists"; \
	else \
		docker run -d --name postgres-banking \
			-p 5432:5432 \
			-e POSTGRES_PASSWORD=postgres \
			-e POSTGRES_USER=postgres \
			-e POSTGRES_DB=banking \
			postgres && \
		echo "Waiting for PostgreSQL to start..." && \
		sleep 5 && \
		docker exec -it postgres-banking psql -U postgres -d banking -c "CREATE USER accounts WITH PASSWORD 'accounts';" && \
		docker exec -it postgres-banking psql -U postgres -d banking -c "GRANT ALL PRIVILEGES ON DATABASE banking TO accounts;"; \
	fi



run-postgres:
	@if docker ps -a | grep -q postgres-banking; then \
		echo "PostgreSQL container already exists"; \
	else \
		docker run -d --name banking-postgres \
			-p 5432:5432 \
			-e POSTGRES_PASSWORD=postgres \
			-e POSTGRES_USER=postgres \
			-e POSTGRES_DB=banking \
			postgres; \
	fi

run-kafka:
	docker compose down zookeeper kafka
	docker compose up -d --build zookeeper kafka
	docker compose up zookeeper kafka
	sleep 5  # Give Kafka time to start
	docker exec banking-app-kafka-1 kafka-topics \
		--bootstrap-server banking-kafka-1:9092 \
		--create \
		--if-not-exists \
		--topic banking.transactions \
		--partitions 3 \
		--replication-factor 1

test:
	go test -v ./...
