.PHONY: test-backend-go test-backend-node 

docker-compose-test-file = "docker-compose-test.yml"

start-backend-node-local:
	cd backend-node && docker compose up -d

stop-backend-node-local:
	cd backend-node && docker compose down

start-backend-go-local:
	cd backend-go && docker compose up -d

stop-backend-go-local:
	cd backend-go && docker compose down

test-backend-node:
	cd backend-node && docker compose --file=$(docker-compose-test-file) run --rm backend-node-test

test-backend-go:
	cd backend-go && docker compose --file=$(docker-compose-test-file) run --rm backend-go-test