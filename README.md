# AdminIn API

AdminIn API is a RESTful backend service built with Go. The project is designed with a strong emphasis on Clean Architecture, scalability, and security to manage system administration data.

## Key Features

* **Layered Architecture**: Clear separation of concerns utilizing Handler, Service, Repository, and Model layers, inspired by Domain-Driven Design (DDD).
* **Admin Registration**: Secure onboarding for administrators with built-in email uniqueness validation.
* **Security First**: Implements bcrypt for cryptographic password hashing prior to database persistence.
* **GORM ORM**: Robust and structured database operations utilizing GORM with an SQLite dialect for seamless local development.
* **Go 1.22 Standard Router**: Leverages the enhanced HTTP method routing capabilities introduced in the Go 1.22 `net/http` standard library.

## Tech Stack

* **Language**: [Go](https://go.dev/) (v1.22+)
* **Database**: [SQLite](https://sqlite.org/) (Local file-based)
* **ORM**: [GORM](https://gorm.io/)
* **Security**: [Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)
* **Architecture**: Layered Architecture

## Roadmap

    [x] Implement JWT (JSON Web Token) authentication mechanism.

    [x] Add /login endpoint for session generation.

    [ ] Integrate structured logging (e.g., Zap or Logrus).

    [ ] Write comprehensive Unit Tests for the Service Layer.

    [ ] Create Dockerfile and docker-compose for containerized deployment.

## Project Structure

```text
adminin-api/
├── cmd/
│   └── api/
│       └── main.go         # Application entry point & dependency injection
├── internal/
│   ├── auth/ 
│   ├── handler/            # HTTP transport layer (Request/Response)
│   ├── service/            # Business logic layer (Validation, hashing, etc.)
│   ├── repository/         # Data access layer (Database operations)
│   └── model/              # Domain layer (Structs & entities)
├── adminin.db              # SQLite database (Auto-generated)
├── go.mod                  # Go module dependencies
└── README.md
