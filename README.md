# GoAuth

A secure authentication service built with Go, featuring JWT tokens, refresh tokens, and user management.

## Features

- Basic CLI
- User registration and authentication
- JWT token-based authentication
- Refresh token mechanism (WIP)
- Secure password hashing
- HTTPS support
- SQLite database with GORM
- RESTful API endpoints

## Tech Stack

- **Go** - Backend language
- **Cobra** - CLI framework
- **Gin** - HTTP web framework
- **GORM** - ORM for database operations
- **SQLite** - Database
- **JWT** - JSON Web Tokens for authentication
- **bcrypt** - Password hashing

## Project Structure

```
.
├── bin
├── cert.pem
├── cmd
│   ├── root.go
│   ├── serve.go
│   └── version.go
├── goauth.db
├── go.mod
├── go.sum
├── http-client.env.json
├── internal
│   ├── api
│   │   ├── http.go
│   │   ├── router.go
│   │   └── server.go
│   ├── auth
│   │   ├── middleware.go
│   │   ├── service.go
│   │   └── utils.go
│   ├── models
│   │   ├── refresh_token.go
│   │   └── user.go
│   ├── storage
│   │   └── db.go
│   └── version
│       └── version.go
├── key.pem
├── LICENSE.md
├── main.go
├── Makefile
└── README.md
```
## Getting Started

### Prerequisites

- Go 1.21 or higher
- Make (optional)

### Installation

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod tidy
3. Create the SSL certificates:
```make
make generate-cert
```
4. Generate app secret:
```make
make generate-secret
```
5. Build the application:
```make 
make build-[linux|macos]
```
6. Run the application:
```bash 
./bin/goauth-[linux|macos]
```

## Usage

### API Endpoints

The application provides the following API endpoints:

- `GET /hc`: Health check endpoint
```curl
curl -k -X GET https://localhost:8080/hc
```

- `POST /register`: Register a new user
```curl
curl -k -X POST https://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"username": "john", "password": "securepass123"}'
  -c yourcookiefile.txt
```

- `POST /login`: Login a user
```curl
curl -k -X POST https://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username": "john", "password": "securepass123"}'
  -c yourcookiefile.txt
```

- `GET /api/test`: Test endpoint
```curl
curl -k -X GET https://localhost:8080/api/test \
  -H "Content-Type: application/json" \
  -b yourcookiefile.txt
```

- `POST /api/logout`: Logout endpoint
```curl
curl -k -X POST https://localhost:8080/api/logout \
  -H "Content-Type: application/json" \
  -b yourcookiefile.txt
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE.md) file for details.
