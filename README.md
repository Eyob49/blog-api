# Blog API (Go + PostgreSQL)

A RESTful blog API built with Go and PostgreSQL, supporting full CRUD operations with clean layered architecture and validation.

---

## 🛠 Tech Stack
- Go (net/http)
- PostgreSQL
- pgx
- godotenv

---

## ⚙️ Environment Setup

Create a `.env` file in the root directory:

```env
DATABASE_URL=postgres://username:password@localhost:5432/blogdb


▶️ Run the Project
go mod tidy
go run .


📁 Project Structure
handlers/   → HTTP request handling
services/   → business logic
store/      → database layer
models/     → data structures


📡 API Endpoints

Get all posts
GET /posts
Create post
POST /posts
Get post by ID
GET /posts/{id}
Update post
PUT /posts/{id}
Delete post
DELETE /posts/{id}

📥 Example Request
POST /posts
{
  "title": "My Post",
  "content": "Hello world"
}

📤 Example Response
{
  "id": 1,
  "title": "My Post",
  "content": "Hello world",
  "created_at": "2026-04-23T10:00:00Z"
}


✅ Features
Full CRUD operations
PostgreSQL persistence
Clean architecture (handlers/services/store)
Input validation
Environment-based configuration


📌 Notes
Ensure PostgreSQL is running before starting the app
.env file is required for database connection
