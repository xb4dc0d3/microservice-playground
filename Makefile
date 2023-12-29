docker-compose-test-file = "docker-compose-test.yml"
kubernetes-app-path = "kubernetes/apps"

start-backend-node-js-local:
	cd backend-node-js && docker compose up -d

stop-backend-node-js-local:
	cd backend-node-js && docker compose down

start-backend-go-local:
	cd backend-go && docker compose up -d

stop-backend-go-local:
	cd backend-go && docker compose down

test-backend-node-js:
	cd backend-node-js && docker compose --file=$(docker-compose-test-file) run --rm backend-node-js-test

test-backend-go:
	cd backend-go && docker compose --file=$(docker-compose-test-file) run --rm backend-go-test