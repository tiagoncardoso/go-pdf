# PDF Maker (Wip)

    .
    ├── cmd/
    │   └── pdf-maker/
    │          └── main.go
    │   └── webserver/
    │          └── main.go
    │   └── config/
    │          └── env-config.go
    ├── internal/
    │   ├── domain/
    │   │   ├── entity/
    │   │   ├── repository/
    │   │   └── service/
    │   ├── usecase/
    │   ├── adapter/
    │   │   ├── http/
    │   │   ├── storage/
    │   │   └── pdf/
    │   └── bootstrap/
    ├── pkg/
    │   ├── http/
    │   ├── logger/
    │   └── pdf-generator/
    ├── test/
    ├── go.mod
    ├── go.sum
    └── README.md

Keep all business rules inside `internal/domain` + `internal/usecase`. Adapters depend inward; never the opposite.

## 🧱 Architecture (Clean Architecture Layers)

1. Domain
    - Pure entities & invariants.
    - No external dependencies.
2. Use Cases
    - Application-specific orchestration.
    - Input/Output request/response models (internal).
3. Interface Adapters
    - HTTP handlers, storage, PDF generators.
    - Maps external formats to internal models.
4. Frameworks & Drivers (edge)
    - `cmd/server`, DB client, PDF libraries, cloud SDKs.

Dependency Rule: `cmd` → adapter → usecase → domain (one direction inward).

## 🧩 Options Pattern (Example Concept)

Encapsulates PDF generation parameters (page size, orientation, margins, storage target) so the API surface remains clean and extensible.

## ⚙️ Configuration

Environment variables (suggested):

| Variable | Description | Default |
|----------|-------------|---------|
| `APP_PORT` | HTTP server port | `8080` |
| `APP_ENV` | Environment (`dev`, `prod`) | `dev` |
| `STORAGE_DRIVER` | `local`, `s3`, `azure` | `local` |
| `S3_BUCKET` | S3 bucket name | - |
| `AWS_REGION` | AWS region | - |
| `AZURE_CONTAINER` | Azure Blob container | - |
| `PDF_ENGINE` | Backend engine (`wkhtml`, `gopdf`, etc.) | `gopdf` |
| `LOG_LEVEL` | `debug`, `info`, etc. | `info` |

Example `.env` (not committed):

## 🚀 Getting Started

Clone and install dependencies:

```sh
git clone git@github.com:tiagoncardoso/go-pdf.git
cd go-pdf
go mod tidy
```

Run the server:
```sh
go run ./cmd/pdf-maker/main.go
```

Build binary:
```sh
go build -o bin/pdf-maker cmd/server/main.go
```

Run all tests:
```sh
go test ./...
```

## 🧪 Testing Strategy (Planned)
- Unit tests: entities, use cases (no external I/O).
- Adapter tests: mock repositories + HTTP handlers.
- Integration tests: storage + PDF generation (tagged).
- Add CI workflow using actions/setup-go.

Example (future) test split:
```sh
go test ./internal/domain/...
go test -tags=integration ./test/...
```

## 🔐 Security / Validation (Planned)
- HTML sanitization for unsafe input.
- Restrict external resource loading in PDF engine.
- Size limits for input payloads.
- Authentication middleware (Basic).


## 🌐 HTTP API (Draft)

Generate PDF:
```
POST /v1/pdf
Content-Type: application/json
{
  "content": "<h1>Hello</h1>",
  "type": "html",
  "options": {
    "filename": "example.pdf",
    "storage": "s3"
  }
}
```

Response (example):
```
{
  "id": "f4c2a1",
  "filename": "example.pdf",
  "status": "stored",
  "location": "s3://bucket/example.pdf"
}
```

- Funcionalidades para ser implementadas:
  - [x] Criar PDF a partir de HTML string
  - [x] Criar PDF a partir de texto simples
  - [ ] Reader e Footer customizáveis
  - [ ] Criação de template para PDF
  - [x] Implementar Options Pattern para facilitar chamadas
  - [ ] Implementar testes unitários
  - [x] Implementar output para storage (S3, Azure Blob, etc)
  - [ ] Implementar output para email
  - [ ] Implementar output para impressão direta
  - [ ] Implementar output para visualização no navegador
  - [ ] Implementar output para download
  - [ ] Implementar EDA para html em string
  - [ ] Implementar log de erros e eventos
  - [ ] Implementar header com logo e título
  - [ ] Adicionar testes unitário a CI/CD
  - [ ] Adicionar autenticação na API