DOCKER_COMPOSE_TEST_FILE = "docker-compose-test.yml"
GIT_SHA="$(shell git rev-parse HEAD)"

start-backend-node-js-local:
	cd backend-node-js && docker compose build --build-arg GIT_SHA=$(GIT_SHA) && docker compose up -d

stop-backend-node-js-local:
	cd backend-node-js && docker compose down

start-backend-go-local:
	cd backend-go && docker compose build --build-arg GIT_SHA=$(GIT_SHA) && docker compose up -d

stop-backend-go-local:
	cd backend-go && docker compose down

test-backend-node-js:
	cd backend-node-js && docker compose --file=$(DOCKER_COMPOSE_TEST_FILE) run --rm backend-node-js-test

test-backend-go:
	cd backend-go && docker compose --file=$(DOCKER_COMPOSE_TEST_FILE) run --rm backend-go-test