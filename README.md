# Microservice-Playground
Microservice Playground

### CI/CD
Our CI/CD pipeline is strategically divided into two distinct workflows:

1. Continuous Integration (CI):
    - Triggered when there are `pull requests` and when changes occur in the `backend-go/**` and `backend-node-js/**` application code.
    - Flows:
      ```mermaid
      graph LR;
      PullRequests --> Trigger-Github-Workflow
      Trigger-Github-Workflow --> Build
      Build --> Test
      Test --> Build-Images
      Build-Images --> Image-Scan-Trivy
      ```
  
2. Continuous Deployment (CD):
    - Triggered by `push only` and when changes occur in the `backend-go/**` and `backend-node-js/**` application code.
    - Employs `workflow_dispatch:` for manual run workflow via the GitHub UI
    - Flows:
      ```mermaid
      graph LR;
      Push --> Trigger-Github-Workflow
      Trigger-Github-Workflow --> Build
      Build --> Test
      Test --> Build-Images
      Build-Images --> Push-Images
      Push-Images --> Deploy-Apps
      Deploy-Apps --> Verify-Deployment
      ```

### Infrastructure
1. Infra Manifests
   - Setup `NGINX Ingress Controller` to expose the service to the public, refer to the provided manifests [here](kubernetes/infra/ingress-nginx)
   - For the TLS currently using Cerbot Let's Encrypt instead of CertManager, because there's no DNS01 challange provider for namecheap domain. List of supported DNS01 [provider](https://cert-manager.io/docs/configuration/acme/dns01/#supported-dns01-providers)
2. Apps Manifests
   - Setup Ingress and Deployment for `backend-go` and `backend-node-js`, refer to the provided manifests [here](kubernetes/apps)

### Application

#### Backend Node js
##### Endpoints
1. Endpoint `GET`
   * Version deployed: https://node-js.bonestealer.xyz/version
   * Healthcheck: https://node-js.bonestealer.xyz/healthz

2. Endpoint `POST`
   * Add operation
    ```bash
    $ curl -X POST -H "Content-Type: application/json" -d '{"num1": 3, "num2": 5}' https://node-js.bonestealer.xyz/api/add
    {"result":8}
    ```
   * Subtract operation
    ```bash
    $ curl -X POST -H "Content-Type: application/json" -d '{"num1": 5, "num2": 10}' https://node-js.bonestealer.xyz/api/subtract
    {"result":-5}     
    ```
   * Multipy operation
    ```bash
    $ curl -X POST -H "Content-Type: application/json" -d '{"num1": 5, "num2": 10}' https://node-js.bonestealer.xyz/api/multiply
    {"result":50}
    ```
   * Divide operation 
    ```bash
    $ curl -X POST -H "Content-Type: application/json" -d '{"num1": 10, "num2": 5}' https://node-js.bonestealer.xyz/api/divide
    {"result":2}
    ```

#### Backend Golang
##### Endpoints
1. Endpoint `GET`
   * Version deployed: https://go.bonestealer.xyz/version
   * Healthcheck: https://go.bonestealer.xyz/healthz

2. Endpoint `POST`
   * Add operation
    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"num1": 3, "num2": 5}' https://go.bonestealer.xyz/call-node-backend/add
    {"result":8}
    ```
   * Subtract operation
    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"num1": 5, "num2": 10}' https://go.bonestealer.xyz/call-node-backend/subtract
    {"result":-5}
    ```
   * Multipy operation
    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"num1": 5, "num2": 10}' https://go.bonestealer.xyz/call-node-backend/multiply
    {"result":50}
    ```
   * Divide operation 
    ```bash
    curl -X POST -H "Content-Type: application/json" -d '{"num1": 10, "num2": 5}' https://go.bonestealer.xyz/call-node-backend/divide
    {"result":2}
    ```
##### Environment Variables
   | General Variables                | Required? | Descriptions                                 | Default Value              | Type       |
   |----------------------------------|-----------|----------------------------------------------|----------------------------|-------------
   | BACKEND_NODE_URL                 | ✔         | Backend Node URL                             | https://go.bonestealer.xyz | Config     |

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