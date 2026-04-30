# 📂 DIRECTORY STRUCTURE (FILE TREE)

```
go-mini-runner/
├── cmd/
│   └── runner/
│       └── main.go              # Entry point: Initialize config, dependencies, wire up and run server.
├── internal/                    # Contains business logic code (external projects cannot import)
│   ├── models/
│   │   └── job.go               # Define Job struct, Task, JobStatus enum (Pending, Running...).
│   ├── queue/
│   │   ├── memory_queue.go      # Basic queue using Go Channel.
│   │   └── priority_queue.go    # Priority queue (Reuse container/heap from your CP).
│   ├── worker/
│   │   ├── pool.go              # Manage number of workers (Goroutines).
│   │   └── executor.go          # Job execution logic (run shell script, calculations...).
│   ├── state/
│   │   └── memory_store.go      # Store Job state (use Map + sync.RWMutex to prevent Data Race).
│   └── api/
│       ├── handlers.go          # Contains HTTP functions (Receive submit job request, check status).
│       └── router.go            # Initialize REST API (use Go 1.22+ default net/http).
├── pkg/                         # Self-written libraries, reusable for other projects
│   └── logger/
│       └── logger.go            # Custom Logger (print to console with colors, write to file).
├── scripts/
│   └── dummy_task.sh            # Sample bash script for testing (E.g.: sleep 5, echo "Hello").
├── Dockerfile                   # Package app into extremely lightweight binary file.
├── Makefile                     # Shortcut commands (make run, make build, make test).
├── go.mod                       # Module management
└── README.md
```

## 🚀 MODULE 1: SETUP & CORE MODELS (Foundation & Data Structures)

- [ ] TSK-101 [Setup] Initialize Project & Logger. (Estimate: 2h)
    - Description:
        - Run `go mod init go-mini-runner`.
        - Create `pkg/logger` package. Write a function to log to console with levels
          [INFO], [WARN], [ERROR], [DEBUG]. (Use Go's default log package
          or slog from Go 1.21+).

- [ ] TSK-102 [Core] Define Job Struct and State Management. (Estimate: 3h)
    - Description:
        - In `internal/models/job.go`: Create Job struct containing: ID (uuid), Name,
          Command (string), Status (enum: Pending, Running, Success, Failed),
          RetryCount, MaxRetries, Timeout.
        - In `internal/state/memory_store.go`: Write a Thread-safe Map using
          sync.RWMutex. Include methods: SaveJob(), UpdateJobStatus(),
          GetJobByID(). Note: Don't use regular map since multiple Goroutines accessing
          it will cause Data Race errors.

## ⚙️ MODULE 2: CONCURRENCY ENGINE (System's Heart - Queue & Worker)

- [ ] TSK-201 [Concurrency] Implement In-memory Queue using Channels. (Estimate: 3h)
    - Description:
        - In `internal/queue/memory_queue.go`: Initialize a Buffered Channel
          `chan *models.Job`.
        - Write `Enqueue(job *Job)` method to push job into channel.
        - Write `Dequeue() <-chan *Job` method for workers to pull jobs out.

- [ ] TSK-202 [Concurrency] Build Worker Pool. (Estimate: 4h)
    - Description:
        - In `internal/worker/pool.go`: Initialize WorkerPool that accepts
          workerCount.
        - Use for loop and go keyword to spawn N Goroutines (Workers).
        - Each worker runs an infinite loop (for-select) listening to channel from
          Dequeue(). Upon receiving a Job, call `executor.Execute(job)`.

- [ ] TSK-203 [DevOps] Write Executor to run actual Shell Commands. (Estimate: 3h)
    - Description:
        - In `internal/worker/executor.go`: Use Go's `os/exec` package.
        - Function receives Job's Command (E.g.: `bash scripts/dummy_task.sh`), executes it,
          captures Output (stdout/stderr), saves to Job struct, and updates
          Status to Success or Failed in Memory Store.

## ⏱ MODULE 3: RESILIENCY & CONTROL (Context, Timeout, Retry)

- [ ] TSK-301 [Core] Apply context.Context to handle Job Timeout. (Estimate: 4h)
    - Description:
        - Upgrade Execute() function. Based on Job's Timeout field (e.g., 10 seconds).
        - Use `ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)`.
        - Pass ctx to `exec.CommandContext()`. If shell script runs over 10s,
          Go will automatically send kill signal to that process (Mimic CI/CD
          Timeout exactly).

- [ ] TSK-302 [Core] Exponential Backoff Retry Mechanism. (Estimate: 3h)
    - Description:
        - If Job fails, check if `RetryCount < MaxRetries`.
        - If satisfied, increment RetryCount, sleep for a duration (E.g.: 2s) and
          push Job back to queue tail (Enqueue) for another worker to retry.
          Update Status to Retrying.

- [ ] TSK-303 [API/Control] Cancel Job Feature (Abort running task). (Estimate: 4h)
    - Description:
        - Store running Jobs' `context.CancelFunc` in a Map
          (`map[string]context.CancelFunc`).
        - Write CancelJob(jobID) function. When called, retrieve CancelFunc and execute
          `cancel()`. Worker will immediately abort the running Job.

## 🌐 MODULE 4: HTTP SERVER & GRACEFUL SHUTDOWN (External Communication)

- [ ] TSK-401 [API] Build REST API for Submit & Check Job. (Estimate: 4h)
    - Description:
        - In `internal/api`: Use net/http (use new `http.NewServeMux()` from
          Go 1.22 that easily supports GET/POST method separation).
        - `POST /api/v1/jobs`: Receive JSON payload (Task name, command, retry count).
          Parse JSON, create Job → Save to Store → Push to Queue. Return JobID.
        - `GET /api/v1/jobs/{id}`: Query Store, return current status
          (Running, Completed, Job output log).

- [ ] TSK-402 [API] Wire API Handler to Main. (Estimate: 2h)
    - Description:
        - In `cmd/runner/main.go`: Initialize Queue, Store, WorkerPool (Start
          pool).
        - Initialize HTTP Server on port 8080.

- [ ] TSK-403 [DevOps] Implement Graceful Shutdown (Prevent data loss when stopping server). (Estimate: 4h)
    - Description:
        - Use `os/signal` to listen for SIGINT, SIGTERM (When you press Ctrl+C or
          Docker stop).
        - When shutdown signal received:
          1. Stop accepting new API requests (Shutdown HTTP Server).
          2. Close Queue Channel.
          3. Use `sync.WaitGroup` to wait for Workers to finish processing running Jobs
              before shutting down the program completely.

## 🏆 MODULE 5: ADVANCED FEATURES (Bonus Level - After completing Module 4)

- [ ] TSK-501 [Adv-Data-Structure] Convert Channel Queue to Priority Queue. (Estimate: 4h)
    - Description:
        - Reuse Priority-Queue.go code you wrote before.
        - Jobs with Priority = HIGH get placed at the front of the queue. (Note: Now you
          must combine Heap struct with `sync.Cond` or loop pattern
          for workers to listen instead of regular Channel).

- [ ] TSK-502 [DevOps] Package Docker & Setup Makefile. (Estimate: 2h)
    - Description:
        - Write Dockerfile Multi-stage build (Use golang alpine builder, then
          copy binary to lightweight ubuntu image with bash shell to execute jobs).
        - Write Makefile containing: `make build`, `make run`, `make test`.