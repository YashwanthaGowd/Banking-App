.PHONY: build-mongo

run-mongo:
	@if docker ps -a | grep -q mongodb; then \
		echo "MongoDB container already exists"; \
	else \
		docker run -d --name banking-mongodb -p 27017:27017 mongo; \
	fi

test:
	go test -v ./...