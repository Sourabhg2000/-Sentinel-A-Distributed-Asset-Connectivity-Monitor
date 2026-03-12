# Golang Test API

A production-ready REST API built with Go, Gin framework, and MongoDB.

## 📁 Project Structure

```
Golang_Test/
├── main.go              # Application entry point
├── routes/
│   └── routes.go        # API route definitions
├── handlers/
│   ├── user_handler.go  # User HTTP handlers
│   └── health_handler.go # Health check handler
├── service/
│   └── user_service.go  # Business logic layer
├── dao/
│   └── user_dao.go      # Database access layer
├── model/
│   └── user.go          # Data models and DTOs
├── go.mod               # Go module dependencies
└── .env.example         # Environment variables template
```

## 🏗️ Architecture

This project follows a **layered architecture** pattern:

1. **Handlers Layer** (`handlers/`) - HTTP request/response handling
2. **Service Layer** (`service/`) - Business logic and validation
3. **DAO Layer** (`dao/`) - Database operations
4. **Model Layer** (`model/`) - Data structures and DTOs

## ✨ Features

- ✅ Clean layered architecture
- ✅ RESTful API endpoints
- ✅ MongoDB integration
- ✅ Input validation
- ✅ Error handling
- ✅ Health check endpoint
- ✅ Environment-based configuration
- ✅ Context timeout management

## 🚀 Getting Started

### Prerequisites

- Go 1.21 or higher
- MongoDB running locally or remote instance

### Installation

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod download
   ```

3. Create `.env` file from example:
   ```bash
   copy .env.example .env
   ```

4. Update `.env` with your MongoDB connection string:
   ```
   MONGODB_URI=mongodb://localhost:27017
   PORT=8000
   ```

5. Run the application:
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8000`

## 📡 API Endpoints

### Health Check
```
GET /api/v1/health
```

### User Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST   | `/api/v1/users` | Create a new user |
| GET    | `/api/v1/users` | Get all users |
| GET    | `/api/v1/users/:id` | Get user by ID |
| PUT    | `/api/v1/users/:id` | Update user |
| DELETE | `/api/v1/users/:id` | Delete user |

### Example Requests

**Create User:**
```bash
curl -X POST http://localhost:8000/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "phone": "1234567890",
    "age": 30
  }'
```

**Get All Users:**
```bash
curl http://localhost:8000/api/v1/users
```

**Get User by ID:**
```bash
curl http://localhost:8000/api/v1/users/{user_id}
```

**Update User:**
```bash
curl -X PUT http://localhost:8000/api/v1/users/{user_id} \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Jane Doe",
    "age": 31
  }'
```

**Delete User:**
```bash
curl -X DELETE http://localhost:8000/api/v1/users/{user_id}
```

## 📦 Dependencies

- **[Gin](https://github.com/gin-gonic/gin)** - HTTP web framework
- **[MongoDB Go Driver](https://go.mongodb.org/mongo-driver)** - Official MongoDB driver

## 🔧 Configuration

Environment variables (in `.env` file):

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | Server port | `8000` |
| `MONGODB_URI` | MongoDB connection string | `mongodb://localhost:27017` |
| `DB_NAME` | Database name | `golang_test_db` |

## 📝 Code Structure

### Handlers
Handle HTTP requests and responses, call service layer methods.

### Services
Contain business logic, validation, and orchestrate DAO operations.

### DAO (Data Access Object)
Direct database operations using MongoDB driver.

### Models
Define data structures, request/response DTOs, and validation rules.

## 🛡️ Error Handling

All API responses follow a consistent format:

```json
{
  "success": true,
  "message": "Operation successful",
  "data": { ... }
}
```

Error response:
```json
{
  "success": false,
  "error": "Error message here"
}
```

## 🧪 Testing

Run tests:
```bash
go test ./...
```

## 📄 License

This project is licensed under the MIT License.
