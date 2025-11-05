# Go `net/http` Authentication & OAuth2 API

![Go Version](https://img.shields.io/badge/Go-1.21%2B-blue.svg)
![Database](https://img.shields.io/badge/Database-PostgreSQL-blue.svg)
![License](https://img.shields.io/badge/License-MIT-green.svg)

A complete user authentication API (Signup, Login, Protected Routes) built from scratch using **only** Go's standard `net/http` library.

This project was built to demonstrate core Go concepts (middleware, context, handlers, project structure) and authentication principles (JWT, bcrypt, OAuth2) *without* the magic of a framework like Gin or Echo.

## ✨ Features

* **Framework-Free:** Built entirely with the `net/http` standard library.
* **Modular Structure:** Clean code separation into `handlers`, `models`, `database`, `auth`, and `middleware` packages.
* **Standard Signup:** User registration with input validation (`go-playground/validator`) and password hashing (`bcrypt`).
* **Standard Login:** Email & password login that returns a JSON Web Token (JWT).
* **Social Login:** "Login with Google" using the `golang.org/x/oauth2` package.
* **JWT Authentication:** Secure token generation (`golang-jwt/jwt/v5`) and validation.
* **Protected Routes:** Custom `AuthMiddleware` to protect endpoints and check for credentials.
* **Role-Based Access Control (RBAC):** Example of an admin-only route (`/users`).
* **Modern Database:** Uses `sqlx` for clean database queries to PostgreSQL.
* **Configuration:** Securely configured using `.env` files (`joho/godotenv`).

## 🛠 Tech Stack

* **Core:** `net/http` (Server & Routing)
* **Database:** `github.com/jmoiron/sqlx`
* **Postgres Driver:** `github.com/lib/pq`
* **JWT:** `github.com/golang-jwt/jwt/v5`
* **OAuth2:** `golang.org/x/oauth2`
* **Hashing:** `golang.org/x/crypto/bcrypt`
* **Validation:** `github.com/go-playground/validator/v10`
* **Config:** `github.com/joho/godotenv`
* **UUIDs:** `github.com/google/uuid`

## 📂 Project Structure

```bash
go-auth-manual/
├── main.go               # Entry point, server setup, and routing
├── go.mod
├── .env                  # (Must be created by you)
├── .env.example          # Example config file
├── auth/
│   ├── generate_jwt.go            # JWT generation & validation
│   └── oauth.go          # OAuth2 configuration
│   └── init.go
│   └── validate_jwt.go
├── database/
│   └── database.go       # PostgreSQL (sqlx) connection
├── handlers/
│   ├── auth_handler.go   # Signup, Login, GetUsers handlers
│   └── oauth_handlers.go  # GoogleLogin, GoogleCallback handlers
│   └── handler.go
│   └── login handler.go
│   └── oauth_callback.go
│   └── sign_up.go
├── middleware/
│   └── auth_middleware.go # JWT validation middleware
├── models/
│   └── user.go           # User, LoginRequest structs
└── validator/
    └── validator.go      # Global validator instance
```

## 🚀 Getting Started

### 1. Prerequisites

* Go 1.21 or newer
* PostgreSQL
* Git

### 2. Clone the Project

```bash
git clone [https://github.com/rzhbadhon/go-nethttp-oauth2-jwt-api.git](https://github.com/rzhbadhon/go-nethttp-oauth2-jwt-api.git)
cd user-create-login-logout-authentication-go
```

### 3. Install Dependencies

```bash
go mod tidy
```

### 4. Database Setup

Log in to PostgreSQL and create a database (e.g., `auth_db`). Then, run the following SQL to create the `users` table:

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

### 5. Configuration (Most Important Step)

You must set up your environment variables.

#### a. Google OAuth2 Credentials
1.  Go to the [Google Cloud Console](https://console.cloud.google.com/).
2.  Create a new project.
3.  Go to "APIs & Services" > "Credentials".
4.  Click "Create Credentials" > "OAuth client ID".
5.  Select "Web application".
6.  Under **"Authorized redirect URIs"**, add: `http://localhost:9000/auth/google/callback`
7.  Click "Create". Copy your **Client ID** and **Client Secret**.

#### b. Create `.env` file
Copy the example file:
```bash
cp .env.example .env
```
Now, edit your new `.env` file with your credentials:

```.env
# Your PostgreSQL connection string
DB_URL="user=postgres password=1212 dbname=auth_db sslmode=disable"

# A strong, random secret key for signing JWTs
JWT_SECRET="your_very_strong_random_secret_key_here"

# Credentials from Google Cloud Console (Step 5a)
GOOGLE_CLIENT_ID="your_google_client_id.apps.googleusercontent.com"
GOOGLE_CLIENT_SECRET="your_google_client_secret"
```

### 6. Run the Server

```bash
go run main.go
```
The server will start on `http://localhost:9000`.

## 🔑 API Endpoints

| Method | Endpoint | Description | Auth Required |
| :--- | :--- | :--- | :--- |
| `POST` | `/signup` | Registers a new user. | No |
| `POST` | `/login` | Logs in with email/password, returns a JWT. | No |
| `GET` | `/auth/google/login` | Redirects you to the Google login page. | No |
| `GET` | `/auth/google/callback`| Google redirects here after login. Returns a JWT. | No |
| `GET` | `/users` | Gets a list of all users. | **Yes** (Admin Role) |

## 🧪 How to Test

### Test 1: Standard Signup & Login (Postman)

**1. Signup (as Admin)**
```bash
# Request (POST http://localhost:9000/signup)
curl -X POST http://localhost:9000/signup \
-H "Content-Type: application/json" \
-d '{
    "first_name": "Admin",
    "last_name": "User",
    "email": "admin@example.com",
    "password": "adminpassword123",
    "role": "admin"
}'
```

**2. Login**
```bash
# Request (POST http://localhost:9000/login)
curl -X POST http://localhost:9000/login \
-H "Content-Type: application/json" \
-d '{
    "email": "admin@example.com",
    "password": "adminpassword123"
}'

# Response (Save this token)
# { "token": "eyJhbGciOiJIUzI1Ni..." }
```

**3. Access Protected Route**
(Replace `<TOKEN>` with the token from Step 2)
```bash
# Request (GET http://localhost:9000/users)
curl -X GET http://localhost:9000/users \
-H "Authorization: Bearer <TOKEN>"

# Response (Success! You will see the list of users)
```

### Test 2: Google OAuth2 Login (Browser)

1.  Start your server: `go run main.go`.
2.  In your **browser** (not Postman), go to:
    `http://localhost:9000/auth/google/login`
3.  You will be redirected to Google. Log in with your Google account.
4.  Google will redirect you back to the callback URL.
5.  Your browser will display the final JSON response, including your new JWT:
    ```json
    {
        "message": "Login successful via Google",
        "token": "eyJhbGciOiJIUzI1Ni..."
    }
    ```

## 📜 License

This project is licensed under the MIT License.