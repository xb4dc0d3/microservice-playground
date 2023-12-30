# Microservice-Playground
Microservice Playground

## Explanation

### Infrastructure
- 


### Application

#### Backend Node js
1. Endpoint `GET`
   * Version deployed: https://node-js.bonestealer.xyz/version
   * Healthcheck: https://node-js.bonestealer.xyz/healthz

2. Endpoint `POST`
   * Add operation
    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"num1": 3, "num2": 5}' https://node-js.bonestealer.xyz/api/add
    ```
   * Subtract operation
    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"num1": 5, "num2": 10}' https://node-js.bonestealer.xyz/api/subtract
    ```
   * Multipy operation
    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"num1": 5, "num2": 10}' https://node-js.bonestealer.xyz/api/multiply
    ```
   * Divide operation 
    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"num1": 10, "num2": 5}' https://node-js.bonestealer.xyz/api/divide
    ```

#### Backend Golang
1. Endpoint `GET`
   * Version deployed: https://go.bonestealer.xyz/version
   * Healthcheck: https://go.bonestealer.xyz/healthz

2. Endpoint `POST`
   * Add operation
    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"num1": 3, "num2": 5}' https://go.bonestealer.xyz/call-node-backend/add
    ```
   * Subtract operation
    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"num1": 5, "num2": 10}' https://go.bonestealer.xyz/call-node-backend/subtract
    ```
   * Multipy operation
    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"num1": 5, "num2": 10}' https://go.bonestealer.xyz/call-node-backend/multiply
    ```
   * Divide operation 
    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"num1": 10, "num2": 5}' https://go.bonestealer.xyz/call-node-backend/divide
    ```

## Repo Directory Structure
```bash
.
├── Makefile
├── README.md
├── backend-go
│   ├── Dockerfile
│   ├── docker-compose-test.yml
│   ├── docker-compose.yml
│   ├── go.mod
│   ├── go.sum
│   ├── main.go
│   └── main_test.go
├── backend-node-js
│   ├── Dockerfile
│   ├── docker-compose-test.yml
│   ├── docker-compose.yml
│   ├── package-lock.json
│   ├── package.json
│   ├── server.js
│   └── test
│       └── tests.js
├── kubeconfig
└── kubernetes
    ├── apps
    │   ├── backend-go
    │   └── backend-node-js
    └── infra
        ├── README.md
        ├── certs
        └── ingress-nginx
```

## Tech Stacks
- Kubernetes using `k3s`: v1.28.4
- Golang: v1.21.1
- Node js: v16.20
- DockerHub Registry
- NGINX Ingress Controller: v3.4.0