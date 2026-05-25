# ⚙️ Golang Backend & DevOps Core

Welcome to the **Golang Backend & DevOps Core** repository! This module bridges the gap between basic Go syntax and real-world Software Engineering. It provides an extensive deep dive into building scalable backend services, mastering advanced concurrency patterns, adopting layered architectures (MVC), and developing CLI tools for DevOps.

This workspace represents the ultimate transition from writing "scripts that work" to building "systems that scale."

---

## ✨ Core Topics Covered

This module is structurally divided into 5 progressive phases:

### 🧩 Go Core Mechanics:
   - Mastering Go's unique features: Implicit Interfaces, Generics (Type Parameters), Struct Embedding (Composition over Inheritance), and Error Handling.
   
### 🚦 Advanced Concurrency & Patterns:
   - **Goroutines & Channels:** Deep understanding of Goroutine scheduling (`runtime.Gosched`), anonymous functions, closures, channel deadlocks, and buffered channels.
   - **Industry-Standard Patterns:** Implementation of high-level concurrency patterns used in production:
     - Confinement (Adhoc & Lexical)
     - For-Select & WaitGroup Synchronization
     - Goroutine Leak Prevention
     - The `Or-Channel` Pattern
     - Error Handling in Concurrent Streams
     - Data Pipelines
     - High-throughput Worker Queues (e.g., Crawling 10K URLs).

### 🌐 Backend Fundamentals:
   - **Routing:** Building HTTP servers using the standard `net/http` library (utilizing Go 1.22+ routing).
   - **Data Handling:** JSON Encoding/Decoding, File Uploading, and Serialization.
   - **Middleware:** Creating scalable Middleware Pipelines (Logging, Auth, Panic Recovery).
   - **Database:** Raw `database/sql` queries and Connection Pooling.
   - **Production-Ready:** Configuration loaders (`.env`) and Graceful Shutdown mechanisms.

### 🏗️ Basic MVC Architecture & ORM:
   - Structuring a complete RESTful API using a Layered Architecture (Router ➔ Controller ➔ Service ➔ Repository ➔ DB).
   - Connecting to PostgreSQL and automating schema migrations using **GORM**.
   - Implementing Dependency Injection and separating concerns for high maintainability.

### 🛠️ DevOps CLI Tools:
   - Interacting directly with the Operating System.
   - Executing Bash shell commands directly from Go (`os/exec`).
   - Parsing CLI flags (`flag` package) and reading/writing files efficiently.

---

## ⚙️ Prerequisites

- **Language:** [Go](https://golang.org/dl/) (version 1.22 or higher is recommended for the new `ServeMux` routing syntax).
- **Database:** PostgreSQL.
- **Tools:** [Docker](https://www.docker.com/) (Used to easily spin up local PostgreSQL instances) and `curl` or Postman for API testing.

---

## 🚀 How to Run & Test

### 1. Exploring Concurrency Patterns (Module 2)
The concurrency scripts are designed to be run directly. Many of them accept CLI arguments to test different behaviors (e.g., deadlocks vs. safe channels).
```bash
go run 2-Concurrency/1-Goroutines/4-Goroutine-Gosched.go schedule
```

### 2. Running the MVC Application
**Step 2.1: Start PostgreSQL via Docker**
```bash
docker run --name pg-mvc -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=123456 -e POSTGRES_DB=user_management -p 5432:5432 -d postgres
```
**Step 2.2: Setup Environment Variables**
```bash
cd 4-Basic-MVC-Architecture-ORM
cp .env.example .env
go mod tidy
go run cmd/main.go
```

**Step 2.3: Test the API**
```bash
curl -X POST http://localhost:8081/api/v1/users/register \
-H "Content-Type: application/json" \
-d '{"name": "Developer", "email": "dev@golang.org", "password": "supersecretpassword"}'
```

---
## Happy Coding!