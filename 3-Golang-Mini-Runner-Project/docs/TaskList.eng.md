## 📂 DIRECTORY STRUCTURE (FILE TREE)

```
go-mini-runner/
├── cmd/
│   └── runner/
│       └── main.go              # Entry point: Parse Flags, Load Env, Initialize Server & Worker.
├── internal/
│   ├── config/
│   │   └── config.go            # Load .env, configure DB, Port, Auth Key.
│   ├── models/
│   │   └── job.go               # Define Job Struct (with GORM tags & JSON tags).
│   ├── repository/              
│   │   ├── store.go             # Interface defining storage functions (Save, Get, Update).
│   │   ├── memory_store.go      # Implementation using Map + Mutex (for quick testing).
│   │   └── postgres_store.go    # Implementation using GORM to connect PostgreSQL.
│   ├── queue/
│   │   └── memory_queue.go      # Channel-based queue (Enqueue/Dequeue).
│   ├── worker/
│   │   ├── pool.go              # Manage Goroutines (Worker Pool).
│   │   └── executor.go          # Use os/exec to run bash scripts with context timeout.
│   ├── controllers/             
│   │   └── job_controller.go    # (NEW) Handle HTTP Requests, Encode/Decode JSON.
│   ├── middlewares/             
│   │   └── middlewares.go       # (NEW) Pipeline: Logger, Auth, Panic Recover.
│   └── routes/                  
│       └── routes.go            # (NEW) Register APIs with http.NewServeMux (Go 1.22+).
├── pkg/
│   └── logger/
│       └── logger.go            # Custom Logger.
├── artifacts/                   # Directory containing Job output log files after execution.
└── .env                         # Environment variables file.
```

## 🚀 MODULE 1: SETUP, CLI & CORE MODELS

- [x] TSK-101 [Setup] Initialize Project, Flags & Environment. (Estimate: 2h)
  - Initialize `go mod init`.
  - Write `internal/config/config.go` using `godotenv` to load the `.env` file (containing `DB_DSN`, `PORT`, `AUTH_KEY`).
  - In `cmd/runner/main.go`, use the `flag` package to accept startup parameters: `--workers=5` (number of workers) and `--port=8080`.

- [x] TSK-102 [Models] Define Job Struct with GORM & JSON Tags. (Estimate: 2h)
  - Create a `Job` struct with: `ID (uuid/uint)`, `Name`, `Command`, `Status`, `LogsPath`, `CreatedAt`, `FinishedAt`. Set up complete `json:"name"` and `gorm:"primaryKey"` tags.

- [x] TSK-103 [Interface & Mutex] Design Store Interface & Memory Store. (Estimate: 3h)
  - Create a `JobStore` interface (with functions `Save()`, `GetByID()`, `Update()`).
  - Write `memory_store.go` implementing the interface using `map[string]*models.Job`. Must use `sync.RWMutex` to prevent Data Races.

## ⚙️ MODULE 2: CONCURRENCY ENGINE & DEVOPS

- [x] TSK-201 [Queue] Implement In-memory Queue using Channels. (Estimate: 2h)
  - Create a buffered channel `chan *models.Job`. Write `Enqueue` and `Dequeue` functions.

- [x] TSK-202 [Worker] Build Worker Pool. (Estimate: 3h)
  - In `pool.go`, use a `for` loop to spawn N Goroutines (count based on the `--workers` flag).
  - Each worker listens on the channel. When a Job is received, change its status to `RUNNING`, persist it to the Store, then hand it off to the `executor` for processing.
****
- [x] TSK-203 [DevOps] Execute Shell Commands & Write Logs to File. (Estimate: 4h)
  - In `executor.go`, use `os/exec` to run the Job's command.
  - Apply **Read-Write-Files** knowledge: Capture all output (Stdout/Stderr) and use `os.WriteFile` to save it as a `.txt` file in the `artifacts/` directory (e.g., `artifacts/job_123_log.txt`). Update this file path in the Job's `LogsPath` field.

## ⏱ MODULE 3: RESILIENCY & CONTROL (Context, Timeout, Panic)

- [ ] TSK-301 [Context] Enforce Timeout for Long-running Jobs. (Estimate: 3h)
  - Use `context.WithTimeout(context.Background(), 10*time.Minute)` attached to `exec.CommandContext`. If a script (e.g., `sleep 9999`) exceeds the time limit, automatically kill the process and mark the status as `FAILED_TIMEOUT`.

- [ ] TSK-302 [Recover] Protect Workers from Panic-induced Crashes. (Estimate: 2h)
  - Wrap the worker function with `defer func() { recover() }`. If a shell file causes a critical panic in Go, the Worker must recover itself, mark the Job as failed, and continue accepting new Jobs.

- [ ] TSK-303 [Retry] Exponential Backoff Retry Logic. (Estimate: 2h)
  - If a Job fails and `RetryCount < MaxRetries`, use `time.Sleep()` to delay, then re-enqueue the Job back into the Queue.

## 🌐 MODULE 4: HTTP SERVER, MVC & MIDDLEWARES

- [ ] TSK-401 [Controllers] Write Submit & Check Job APIs. (Estimate: 4h)
  - Write `job_controller.go`. Use `json.NewDecoder(r.Body)` to receive `{name, command}` from the user.
  - Initialize Job -> Save to Store -> Push to Queue. Return `201 Created` with the `job_id`.
  - Write a Get Job by ID API. Return the current information and status.

- [ ] TSK-402 [Middleware] Server Protection Pipeline. (Estimate: 3h)
  - In `middlewares.go`, write `CreatePipeline`.
  - Create `LoggerMiddleware` (measure API response time) and `AuthMiddleware` (check if the `X-API-Key` header matches the `AUTH_KEY` in `.env`).

- [ ] TSK-403 [Routes] Assemble Router & Graceful Shutdown. (Estimate: 3h)
  - Register APIs in `http.NewServeMux()` (syntax: `POST /api/jobs`).
  - In `main.go`, implement Graceful Shutdown (`os/signal`, `srv.Shutdown`). Stop accepting new requests, but wait for the WorkerPool's WaitGroup to finish any in-progress Jobs before shutting down.

## 🗄 MODULE 5: DATABASE PERSISTENCE

- [ ] TSK-501 [Database] Implement Postgres Store with GORM. (Estimate: 4h)
  - Create `postgres_store.go`. Set up `ConnectDatabase()` as previously done.
  - Implement the `Save()`, `GetByID()`, `Update()` functions of the `JobStore` interface, calling `config.DB.Create()` and `config.DB.Save()`.
  - Feature: In `main.go`, simply change one initialization line from `NewMemoryStore()` to `NewPostgresStore()`, and the entire system immediately switches to using a real Database without modifying any Worker or Controller code!

## 🏆 MODULE 6: DOCKER

- [ ] TSK-601 [Docker] Package Application with Multi-stage Build. (Estimate: 2h)
  - Write a Dockerfile using `golang:alpine` to build, then copy the binary into an `ubuntu` image (so the server has bash/shell commands available to execute Jobs).

- [ ] TSK-602 [Generics] Upgrade to Generic Queue. (Estimate: 2h)
  - Use `Generics` knowledge to refactor the Channel Queue into `chan T any`, allowing the Queue to later accept `Job`, `Notification`, or `EmailTask` types.