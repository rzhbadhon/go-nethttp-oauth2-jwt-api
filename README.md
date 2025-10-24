# Go `net/http` JWT Authentication API

![Go Version](https://img.shields.io/badge/Go-1.21%2B-blue.svg)
![Database](https://img.shields.io/badge/Database-PostgreSQL-blue.svg)
![License](https://img.shields.io/badge/License-MIT-green.svg)

A user authentication (Signup/Login) API built from scratch. This project intentionally avoids frameworks like Gin or Echo, using only Go's standard `net/http` library.

The primary goal is to deeply understand Go's core packages, modular project structure, dependency injection, and the fundamentals of a secure API (like JWT and bcrypt).

## ✨ Features

* **Standard Library:** Routing and server handling using only `net/http`.
* **User Signup:** Input validation using `go-playground/validator`.
* **User Login:** Email and password verification.
* **Secure Passwords:** Password hashing and comparison using `golang.org/x/crypto/bcrypt`.
* **JWT Authentication:** Token generation and validation using `golang-jwt/jwt/v5`.
* **Modular Structure:** Clean separation of concerns into `handlers`, `models`, `database`, and `auth` packages.
* **Database:** PostgreSQL integration using `sqlx` for cleaner database interactions.

## 🛠 Tech Stack

* **Core:** `net/http` (Server & Routing)
* **Database:** `github.com/jmoiron/sqlx` (PostgreSQL)
* **Driver:** `github.com/lib/pq` (Postgres Driver)
* **Authentication:** `github.com/golang-jwt/jwt/v5`
* **Security:** `golang.org/x/crypto/bcrypt`
* **Validation:** `github.com/go-playground/validator/v10`
* **Configuration:** `github.com/joho/godotenv`
* **Utilities:** `github.com/google/uuid`

## 📂 Project Structure

```bash
go-auth-manual/
├── main.go               # Entry point, server setup, and routing
├── go.mod
├── go.sum
├── .env.example          # Example environment variables
├── auth/
│   └── jwt.go            # JWT generation and validation logic
├── database/
│   └── database.go       # PostgreSQL connection setup (sqlx)
├── handlers/
│   └── auth_handler.go   # HTTP handlers (Signup, Login)
├── models/
│   └── user.go           # Data structures (User, LoginRequest)
└── validator/
    └── validator.go      # Global validator instance
```

## 🚀 Getting Started

### Prerequisites

* [Go (1.21 or newer)](https://go.dev/dl/)
* [PostgreSQL](https://www.postgresql.org/download/)
* [Git](https://git-scm.com/downloads/)

### 1. Clone the Project

```bash
git clone [https://github.com/rzhbadhon/user-create-login-logout-authentication-go.git](https://github.com/rzhbadhon/user-create-login-logout-authentication-go.git)
cd user-create-login-logout-authentication-go
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Database Setup

Log in to your PostgreSQL instance and create a new database (e.g., `auth_db`). Then, run the following SQL query to create the `users` table:

```sql
CREATE TABLE users (
    id UUID PRIMARY KEY,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    role TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);
```

### 4. Configuration

Create a `.env` file by copying the example file. This file will hold your database credentials and secret keys.

```bash
cp .env.example .env
```

Now, edit the `.env` file with your information:

```.env
# Your PostgreSQL connection string
DB_URL="user=postgres password=yourpassword dbname=auth_db sslmode=disable"

# A strong, random secret key for signing JWTs
JWT_SECRET="your_very_strong_and_random_secret_key"
```

### 5. Run the Server

```bash
go run main.go
```
The server will start on `http://localhost:9000`.

## 🔑 API Endpoints

| Method | Endpoint | Description | Body (Request) | Response (Success) |
| :--- | :--- | :--- | :--- | :--- |
| `POST` | `/signup` | Registers a new user. | `models.User` | `models.User` (password omitted) |
| `POST` | `/login` | Logs in a user and returns a JWT. | `models.LoginRequest` | `{ "token": "..." }` |
| `GET` | `/` | A simple welcome message. | N/A | `string` |

## 🧪 Testing with cURL

### 1. Signup

```bash
curl -X POST http://localhost:9000/signup \
-H "Content-Type: application/json" \
-d '{
    "first_name": "Test",
    "last_name": "User",
    "email": "test@example.com",
    "password": "password123"
}'
```

### 2. Login

```bash
curl -X POST http://localhost:9000/login \
-H "Content-Type: application/json" \
-d '{
    "email": "test@example.com",
    "password": "password123"
}'
```

**Success Response:**
```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoi..."
}
```

## 📜 License

This project is licensed under the [MIT License](LICENSE).
