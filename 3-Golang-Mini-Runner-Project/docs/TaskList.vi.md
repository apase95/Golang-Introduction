## 📂 CẤU TRÚC THƯ MỤC (FILE TREE)

```
go-mini-runner/
├── cmd/
│   └── runner/
│       └── main.go              # Entry point: Parse Flags, Load Env, Khởi tạo Server & Worker.
├── internal/
│   ├── config/
│   │   └── config.go            # Load .env, cấu hình DB, Port, Auth Key.
│   ├── models/
│   │   └── job.go               # Định nghĩa Struct Job (kèm GORM tags & JSON tags).
│   ├── repository/              
│   │   ├── store.go             # Interface định nghĩa các hàm lưu trữ (Save, Get, Update).
│   │   ├── memory_store.go      # Implement dùng Map + Mutex (Để test nhanh).
│   │   └── postgres_store.go    # Implement dùng GORM kết nối PostgreSQL.
│   ├── queue/
│   │   └── memory_queue.go      # Hàng đợi Channel (Enqueue/Dequeue).
│   ├── worker/
│   │   ├── pool.go              # Quản lý Goroutines (Worker Pool).
│   │   └── executor.go          # Dùng os/exec chạy bash script, context timeout.
│   ├── controllers/             
│   │   └── job_controller.go    # (MỚI) Nhận HTTP Request, Encode/Decode JSON.
│   ├── middlewares/             
│   │   └── middlewares.go       # (MỚI) Pipeline: Logger, Auth, Panic Recover.
│   └── routes/                  
│       └── routes.go            # (MỚI) Đăng ký API với http.NewServeMux (Go 1.22+).
├── pkg/
│   └── logger/
│       └── logger.go            # Custom Logger.
├── artifacts/                   # Thư mục chứa các file log output của Job sau khi chạy.
└── .env                         # File biến môi trường.
```
****
## 🚀 MODULE 1: SETUP, CLI & CORE MODELS

- [x] TSK-101 [Setup] Khởi tạo Project, Flags & Environment. (Estimate: 2h)
  - Khởi tạo `go mod init`.
  - Viết `internal/config/config.go` dùng `godotenv` để load file `.env` (chứa `DB_DSN`, `PORT`, `AUTH_KEY`).
  - Tại `cmd/runner/main.go`, dùng package `flag` để nhận tham số khởi động: `--workers=5` (Số lượng worker) và `--port=8080`.

- [x] TSK-102 [Models] Định nghĩa Job Struct với GORM & JSON Tags. (Estimate: 2h)
  - Tạo struct `Job` gồm: `ID (uuid/uint)`, `Name`, `Command`, `Status`, `LogsPath`, `CreatedAt`, `FinishedAt`. Cài đặt đầy đủ `json:"name"` và `gorm:"primaryKey"`.

- [x] TSK-103 [Interface & Mutex] Thiết kế Store Interface & Memory Store. (Estimate: 3h)
  - Tạo interface `JobStore` (gồm hàm `Save()`, `GetByID()`, `Update()`).
  - Viết `memory_store.go` implement interface trên bằng `map[string]*models.Job`. Bắt buộc dùng `sync.RWMutex` để chống Data Race.

## ⚙️ MODULE 2: CONCURRENCY ENGINE & DEVOPS
- [ ] TSK-201 [Queue] Triển khai In-memory Queue bằng Channels. (Estimate: 2h)
  - Tạo buffered channel `chan *models.Job`. Viết hàm `Enqueue` và `Dequeue`.

- [ ] TSK-202 [Worker] Xây dựng Worker Pool. (Estimate: 3h)
  - Tại `pool.go`, dùng vòng lặp `for` spawn ra N Goroutines (số lượng dựa vào `--workers` cờ flag).
  - Mỗi worker lắng nghe channel. Bắt được Job thì đổi status thành `RUNNING`, lưu xuống Store, rồi chuyển cho `executor` xử lý.

- [ ] TSK-203 [DevOps] Thực thi lệnh Shell & Ghi Log ra file. (Estimate: 4h)
  - Tại `executor.go`, dùng `os/exec` chạy lệnh của Job.
  - Dùng kiến thức **Read-Write-Files**: Chụp toàn bộ Output (Stdout/Stderr) và dùng `os.WriteFile` lưu thành file `.txt` vào thư mục `artifacts/` (VD: `artifacts/job_123_log.txt`). Cập nhật đường dẫn file này vào trường `LogsPath` của Job.

## ⏱ MODULE 3: RESILIENCY & CONTROL (Context, Timeout, Panic)
- [ ] TSK-301 [Context] Ràng buộc Timeout cho các Job chạy quá lâu. (Estimate: 3h)
  - Dùng `context.WithTimeout(context.Background(), 10*time.Minute)` gắn vào `exec.CommandContext`. Nếu script (VD: `sleep 9999`) chạy lố giờ, tự động kill tiến trình và đánh dấu status `FAILED_TIMEOUT`.

- [ ] TSK-302 [Recover] Chống sập Worker do Panic. (Estimate: 2h)
  - Bọc hàm worker bằng `defer func() { recover() }`. Nếu file shell lỗi nghiêm trọng gây panic Go, Worker phải tự cứu sống chính mình, đánh dấu Job lỗi và tiếp tục nhận Job mới.

- [ ] TSK-303 [Retry] Logic Retry Exponential Backoff. (Estimate: 2h)
  - Nếu Job failed và `RetryCount < MaxRetries`, dùng `time.Sleep()` để delay, sau đó đẩy ngược Job vào Queue.

## 🌐 MODULE 4: HTTP SERVER, MVC & MIDDLEWARES

- [ ] TSK-401 [Controllers] Viết API Submit & Check Job. (Estimate: 4h)
  - Viết `job_controller.go`. Dùng `json.NewDecoder(r.Body)` nhận `{name, command}` từ user.
  - Khởi tạo Job -> Lưu vào Store -> Nhét vào Queue. Trả về `201 Created` kèm `job_id`.
  - Viết API Get Job theo ID. Trả về thông tin và trạng thái hiện tại.

- [ ] TSK-402 [Middleware] Pipeline bảo vệ Server. (Estimate: 3h)
  - Tại `middlewares.go`, viết `CreatePipeline`.
  - Tạo `LoggerMiddleware` (đo thời gian API) và `AuthMiddleware` (check Header `X-API-Key` có khớp với `AUTH_KEY` trong `.env` không).

- [ ] TSK-403 [Routes] Lắp ráp Router & Graceful Shutdown. (Estimate: 3h)
  - Đăng ký API vào `http.NewServeMux()` (Cú pháp `POST /api/jobs`).
  - Tại `main.go`, cài đặt Graceful Shutdown (`os/signal`, `srv.Shutdown`). Chặn nhận request mới, nhưng đợi WaitGroup của WorkerPool chạy xong nốt Job đang dở rồi mới tắt máy.

## 🗄 MODULE 5: DATABASE PERSISTENCE

- [ ] TSK-501 [Database] Implement Postgres Store bằng GORM. (Estimate: 4h)
  - Tạo file `postgres_store.go`. Setup `ConnectDatabase()` y hệt bạn đã làm.
  - Implement các hàm `Save()`, `GetByID()`, `Update()` của interface `JobStore` nhưng gọi qua `config.DB.Create()` và `config.DB.Save()`.
  - Tính năng: Tại `main.go`, chỉ cần đổi 1 dòng khởi tạo từ `NewMemoryStore()` sang `NewPostgresStore()`, toàn bộ hệ thống ngay lập tức chuyển sang dùng Database thật mà không cần sửa dòng code nào của Worker hay Controller!

## 🏆 MODULE 6: DOCKER
- [ ] TSK-601 [Docker] Đóng gói ứng dụng Multi-stage. (Estimate: 2h)
  - Viết Dockerfile dùng `golang:alpine` để build, sau đó copy file nhị phân (binary) sang image `ubuntu` (để server có sẵn các lệnh bash/shell thực thi Job).

- [ ] TSK-602 [Generics] Nâng cấp Generic Queue. (Estimate: 2h)
  - Dùng kiến thức `Generics` để sửa Queue Channel thành `chan T any`, giúp Queue này sau này có thể nhận cả `Job`, `Notification`, hoặc `EmailTask`.