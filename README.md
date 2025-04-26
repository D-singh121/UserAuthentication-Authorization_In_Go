# User Authentication and Authorization API (Golang)

A production-grade, clean architecture based authentication and authorization API built using **Golang**, **Gin**, **PostgreSQL**, and **JWT**.

GitHub Repo: [UserAuthentication-Authorization_In_Go](https://github.com/D-singh121/UserAuthentication-Authorization_In_Go.git)

---

## Project Structure

```
AUTH_BACKEND/
├── cmd/
│   └── main.go             # Application entry point
├── internals/
│   ├── controllers/        # User Controllers
│   │   └── user_controller.go
│   ├── dto/                # Data Transfer Objects (DTOs)
│   │   └── user_dto.go
│   ├── middlewares/        # Authentication Middleware
│   │   └── auth_middleware.go
│   ├── models/             # Database Models
│   │   └── user.go
│   ├── repositories/       # Database Queries Layer
│   │   ├── user_repository.go
│   │   └── user_repository_test.go
│   ├── routes/             # API Route Definitions
│   │   └── user_routes.go
│   ├── services/           # Business Logic Layer
│   │   ├── user_service.go
│   │   └── user_service_test.go
│   └── utils/              # Utility functions
│       └── jwt.go          # JWT generation and validation
├── pkg/
│   └── config/             # Database and environment config
│       ├── config.go
│       └── db.go
├── .env                    # Environment variables
├── .gitignore              # Git ignore file
├── apiDoc.md               # API documentation (separate)
├── envSample.txt           # Environment sample
├── go.mod                  # Go module file
├── go.sum                  # Go checksum file
```

---

## Installation and Setup

```bash
# Clone the repository
git clone https://github.com/D-singh121/UserAuthentication-Authorization_In_Go.git

# Navigate into the project directory
cd UserAuthentication-Authorization_In_Go

# Install dependencies
go mod tidy

# Setup PostgreSQL Database

# Run the application
go run cmd/main.go
```

Make sure to configure your **.env** file based on **envSample.txt**.

---

## API Endpoints

### Public Routes

| Method | Endpoint                  | Description             |
|:------:|:---------------------------|:-------------------------|
| POST   | `/api/v1/users/register`    | Register a new user      |
| POST   | `/api/v1/users/login`       | Login and receive a token |
| POST   | `/api/v1/users/logout`      | Logout user (handled client-side) |

### Protected Routes (Require JWT Token)

| Method | Endpoint                  | Description             |
|:------:|:---------------------------|:-------------------------|
| GET    | `/api/v1/users/`            | Get all users            |
| GET    | `/api/v1/users/:id`         | Get user by ID           |
| POST   | `/api/v1/users/email`       | Get user by email        |
| PUT    | `/api/v1/users/:id`         | Update user by ID        |
| DELETE | `/api/v1/users/:id`         | Delete user by ID        |


---

## Key Features

- JWT Based Authentication (Token Generation and Validation)
- Middleware for Protected Routes
- Clean Architecture (Controller, Service, Repository)
- PostgreSQL Database
- Gin Framework for routing
- Environment based Configurations


---

## Tech Stack

- Go (Golang)
- Gin Web Framework
- PostgreSQL
- JWT (JSON Web Tokens)
- Clean Architecture

---

## How JWT Works

- User logs in -> Server generates a JWT token -> Client stores token -> Client sends token in Authorization Header for protected routes.

JWT functions are managed inside:
```
internals/utils/jwt.go
```

---

## Author

- [Devesh Choudhary](https://github.com/D-singh121)

---

## License

This project is licensed under the MIT License - see the LICENSE file for details.

---

# 🚀 Happy Coding!

