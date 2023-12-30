# Microservice-Playground
Microservice Playground

## Explanation

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
│   ├── node_modules
│   ├── package-lock.json
│   ├── package.json
│   ├── server.js
│   └── test
├── kubeconfig
└── kubernetes
    ├── apps
    └── infra
```


## Tech Stacks
- Kubernetes using `k3s`: v1.28.4
- Golang: v1.21.1
- Node js: v16.20
- DockerHub Registry
- CertManager: 
- NGINX Ingress Controller: