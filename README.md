# Hexagonal Architecture in Go

A product management application built with Go to demonstrate **Hexagonal Architecture (Ports & Adapters)**. The same business logic is exposed through both a CLI and an HTTP/REST API, with SQLite persistence — all with clean layer separation.

## Architecture

```
                    ┌──────────────┐     ┌──────────────┐
                    │   CLI (Cobra)│     │  HTTP (Mux)  │
                    │   Adapter    │     │   Adapter    │
                    └──────┬───────┘     └──────┬───────┘
                           │  Primary Adapters  │
                    ═══════╪════════════════════╪════════
                           │                    │
                    ┌──────▼────────────────────▼──────┐
                    │        Application Core          │
                    │   Product Entity + Service       │
                    │   Ports (Interfaces)             │
                    └──────────────┬───────────────────┘
                                  │
                    ══════════════╪══════════════════════
                                  │  Secondary Adapter
                    ┌─────────────▼───────────────────┐
                    │     SQLite Adapter (go-sqlite3)  │
                    └─────────────────────────────────┘
```

### Layer Breakdown

| Layer | Location | Responsibility |
|-------|----------|----------------|
| **Domain** | `application/product.go` | Product entity, validation, business rules |
| **Ports** | `application/` (interfaces) | `ProductServiceInterface`, `ProductPersistenceInterface` |
| **Application Service** | `application/product_service.go` | Orchestrates create, get, enable, disable |
| **CLI Adapter** | `adapters/cli/` | Cobra-based command-line interface |
| **HTTP Adapter** | `adapters/web/` | REST API with Gorilla Mux + Negroni middleware |
| **DB Adapter** | `adapters/db/` | SQLite persistence via go-sqlite3 |
| **DTOs** | `adapters/dto/` | JSON serialization for the HTTP layer |

### Project Structure

```
├── application/              # Core: entity, service, ports, mocks
│   ├── product.go           # Domain entity with validation
│   ├── product_service.go   # Business operations
│   └── mocks/               # Generated mocks (golang/mock)
├── adapters/
│   ├── cli/                 # Primary: CLI adapter
│   ├── web/                 # Primary: HTTP handlers + server
│   ├── db/                  # Secondary: SQLite persistence
│   └── dto/                 # Data transfer objects
├── cmd/                     # Entry points (Cobra commands)
│   ├── root.go             # DI setup
│   ├── cli.go              # CLI subcommand
│   └── http.go             # HTTP server subcommand
├── main.go
├── Dockerfile
└── docker-compose.yaml
```

## Business Rules

- Products are created in **disabled** status
- **Enable** requires `price > 0`
- **Disable** requires `price == 0`
- IDs are UUID v4, validated with `govalidator`

## Tech Stack

| Technology | Purpose |
|------------|---------|
| Go 1.19 | Language |
| [Cobra](https://github.com/spf13/cobra) | CLI framework |
| [Gorilla Mux](https://github.com/gorilla/mux) | HTTP router |
| [Negroni](https://github.com/urfave/negroni) | HTTP middleware |
| [go-sqlite3](https://github.com/mattn/go-sqlite3) | Database driver |
| [golang/mock](https://github.com/golang/mock) | Mock generation for testing |
| [testify](https://github.com/stretchr/testify) | Test assertions |

## Getting Started

### Docker (recommended)

```bash
docker-compose up
# HTTP server available at http://localhost:9000
```

### CLI

```bash
go build -o app

# Create a product
./app cli --action create --product "My Product" --price 49.99

# Get a product
./app cli --action get --id <product-id>

# Enable / Disable
./app cli --action enable --id <product-id>
./app cli --action disable --id <product-id>
```

### HTTP API

```bash
go run main.go http
# Server starts on :9000
```

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/products` | Create product (`{"name": "...", "price": 50.0}`) |
| `GET` | `/products/{id}` | Get product by ID |
| `PATCH` | `/products/{id}/enable` | Enable product |
| `PATCH` | `/products/{id}/disable` | Disable product |

## Testing

Tests at every layer with mocked dependencies:

```bash
go test ./...
```

- **Domain tests** — entity validation and business rules
- **Service tests** — mocked persistence via `golang/mock`
- **CLI tests** — mocked service, output verification
- **DB tests** — in-memory SQLite (`:memory:`)
- **Handler tests** — HTTP response formatting

## Key Takeaways

This project demonstrates how hexagonal architecture enables:

1. **Swappable adapters** — replace SQLite with PostgreSQL, or HTTP with gRPC, without changing domain logic
2. **Testability** — every layer tested independently with interface mocks
3. **Multiple entry points** — CLI and HTTP share the exact same business logic
4. **Dependency inversion** — the domain defines ports (interfaces); adapters implement them
