.PHONY: run test deps

# build and run
build:
	docker compose up --build

# run banking-app
run:
	docker compose up