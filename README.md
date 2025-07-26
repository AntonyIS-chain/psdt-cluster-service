# 🏢 PSDT Cluster Service

The **PSDT Cluster Service** is a foundational backend service for managing the **organizational structure** of the [PSDT Connect Platform](https://github.com/your-org). It handles data and operations related to:

- **Clusters** – Major transformation programs
- **Tribes** – Functional or thematic sub-divisions
- **Squads** – Agile delivery teams
- **Developers** – Team members within squads

Built with **Go**, **gRPC**, and **PostgreSQL**, it follows **Hexagonal Architecture** principles and integrates with other services using gRPC.

---

## 📐 Architecture

This service follows **Hexagonal (Ports and Adapters)** architecture, ensuring clean separation of core domain logic and external integrations (DB, gRPC transport, HTTP).

![Architecture Diagram](https://link-to-diagram.com) <!-- Optional -->

---

## 🚀 Features

- CRUD operations for Clusters, Tribes, Squads
- Developer-to-squad assignment APIs
- gRPC APIs for internal services
- PostgreSQL storage with clean domain boundaries
- Optional HTTP/REST or GraphQL gateway adapter

---

## 📦 Tech Stack

| Layer         | Technology        |
|---------------|-------------------|
| Language       | Golang             |
| Database       | PostgreSQL         |
| Internal API   | gRPC               |
| External API   | REST / GraphQL (optional) |
| CI/CD          | GitHub Actions / GitLab CI |
| Auth           | JWT / OAuth2 (via Gateway) |

---

## 🧰 Installation & Local Development
cp .env.example .env
go mod tidy
go run ./cmd/servercp .env.example .env
go mod tidy
go run ./cmd/server
### 🖥️ Requirements

- Go 1.21+
- Docker + Docker Compose
- `protoc` (Protocol Buffers compiler)
- `buf` (for proto linting and generation)

### 🚀 Clone and Run

```bash
git clone https://github.com/your-org/psdt-cluster-service.git
cd psdt-cluster-service

# Run with Docker Compose
docker-compose up --build

###🛠️ Run Locally (without Docker)
```cp .env.example .env
go mod tidy
go run ./cmd/server
```

### 🔌 API Documentation
####📡 gRPC
Proto files: psdt-shared-protos

Service: ClusterService

Endpoints:

CreateCluster, ListClusters, GetTribe, AssignDeveloperToSquad, etc.

```🌐 REST / GraphQL (optional)
REST: GET /clusters, POST /tribes, etc.

GraphQL: query { clusters { id name } }

```

### 🧪 Testing

```go test ./...```
To run tests in a container:
```
docker-compose -f docker-compose.test.yml up --build
```

### 🧬 Code Structure
```
.
├── cmd/
│   └── server/           # App entry point
├── internal/
│   ├── domain/           # Core business logic and models
│   ├── service/          # Use cases
│   ├── grpc/             # gRPC handlers
│   ├── rest/             # (optional) HTTP handlers
│   └── db/               # PostgreSQL adapters
├── proto/                # Local .proto definitions
├── scripts/              # Dev & deployment scripts
└── Dockerfile

```

#### 🤝 Contributing
We welcome contributions!

Fork the repo

Create a new branch

Open a pull request with context

Follow code style & lint rules (golangci-lint run)

🧠 Related Repositories
psdt-api-management

psdt-auth-service

psdt-shared-protos

psdt-infra

### 📬 Contact
Questions or suggestions? Open an issue or reach us at dev@psdtconnect.org



---

Would you like me to:
- Generate `docker-compose.yml`, `.env.example`, or CI config templates?
- Help bootstrap the actual service code using Go (e.g., initial `main.go`, gRPC server)?
- Generate a matching Docusaurus docs site for PSDT services?

Let me know how deep you'd like to go next.
