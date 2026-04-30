📂 CẤU TRÚC THƯ MỤC (FILE TREE)

```
go-mini-runner/
├── cmd/
│   └── runner/
│       └── main.go              # Entry point: Khởi tạo config, dependencies, gài đạn và chạy server.
├── internal/                    # Chứa code logic nghiệp vụ (không cho project ngoài import)
│   ├── models/
│   │   └── job.go               # Định nghĩa struct Job, Task, enum JobStatus (Pending, Running...).
│   ├── queue/
│   │   ├── memory_queue.go      # Hàng đợi cơ bản dùng Go Channel.
│   │   └── priority_queue.go    # Hàng đợi ưu tiên (Dùng lại container/heap từ CP của bạn).
│   ├── worker/
│   │   ├── pool.go              # Quản lý số lượng worker (Goroutines).
│   │   └── executor.go          # Logic thực thi Job (chạy shell script, tính toán...).
│   ├── state/
│   │   └── memory_store.go      # Nơi lưu trạng thái Job (dùng Map + sync.RWMutex để tránh Data Race).
│   └── api/
│       ├── handlers.go          # Chứa các HTTP function (Nhận request submit job, check status).
│       └── router.go            # Khởi tạo REST API (sử dụng net/http mặc định của Go 1.22+).
├── pkg/                         # Thư viện tự viết, có thể tái sử dụng cho project khác
│   └── logger/
│       └── logger.go            # Custom Logger (in ra console có màu sắc, ghi file).
├── scripts/
│   └── dummy_task.sh            # Script bash mẫu để test (VD: sleep 5, echo "Hello").
├── Dockerfile                   # Đóng gói app thành file binary cực nhẹ.
├── Makefile                     # Các lệnh gõ tắt (make run, make build, make test).
├── go.mod                       # Quản lý module
└── README.md
```

## 🚀 MODULE 1: SETUP & CORE MODELS (Nền tảng & Cấu trúc dữ liệu)

- [ ] TSK-101 [Setup] Khởi tạo Project & Logger. (Estimate: 2h)
    - Description:
        - Chạy go mod init go-mini-runner.
        - Tạo package pkg/logger. Viết hàm log in ra console với các mức độ
          [INFO], [WARN], [ERROR], [DEBUG]. (Sử dụng package log mặc định của Go
          hoặc slog của Go 1.21+).
- [ ] TSK-102 [Core] Định nghĩa Job Struct và State Management. (Estimate: 3h)
    - Description:
        - Trong internal/models/job.go: Tạo struct Job gồm: ID (uuid), Name,
          Command (string), Status (enum: Pending, Running, Success, Failed),
          RetryCount, MaxRetries, Timeout.
        - Trong internal/state/memory_store.go: Viết một Thread-safe Map sử dụng
          sync.RWMutex. Bao gồm các method: SaveJob(), UpdateJobStatus(),
          GetJobByID(). Lưu ý: Không dùng map thường vì nhiều Goroutine truy cập
          sẽ gây lỗi Data Race.

## ⚙️ MODULE 2: CONCURRENCY ENGINE (Trái tim hệ thống - Queue & Worker)

- [ ] TSK-201 [Concurrency] Triển khai In-memory Queue bằng Channels.
  (Estimate: 3h)
    - Description:
        - Trong internal/queue/memory_queue.go: Khởi tạo một Buffered Channel
          chan *models.Job.
        - Viết method Enqueue(job *Job) để đẩy job vào channel.
        - Viết method Dequeue() <-chan *Job để worker kéo job ra. -[ ] TSK-202
          [Concurrency] Xây dựng Worker Pool. (Estimate: 4h)
    - Description:
        - Trong internal/worker/pool.go: Khởi tạo WorkerPool nhận vào số lượng
          workerCount.
        - Sử dụng vòng lặp for và từ khóa go để spawn ra N Goroutines (Workers).
        - Mỗi worker sẽ chạy 1 vòng lặp vô tận (for-select) lắng nghe channel từ
          Dequeue(). Nhận được Job thì gọi hàm executor.Execute(job).
- [ ] TSK-203 [DevOps] Viết Executor chạy Shell Command thực tế. (Estimate: 3h)
    - Description:
        - Trong internal/worker/executor.go: Dùng package os/exec của Go.
        - Hàm nhận vào Command của Job (VD: bash scripts/dummy_task.sh), thực
          thi nó, hứng lấy Output (stdout/stderr) lưu vào struct Job và cập nhật
          Status thành Success hoặc Failed xuống Memory Store.

## ⏱ MODULE 3: RESILIENCY & CONTROL (Context, Timeout, Retry)

- [ ] TSK-301 [Core] Áp dụng context.Context để xử lý Timeout Job.
  (Estimate: 4h)
    - Description:
        - Nâng cấp hàm Execute(). Dựa vào trường Timeout của Job (ví dụ: 10
          giây).
        - Sử dụng ctx, cancel :=
          context.WithTimeout(context.Background(), 10*time.Second).
        - Truyền ctx vào exec.CommandContext(). Nếu shell script chạy quá 10s,
          Go sẽ tự động gửi signal kill process đó (Mô phỏng y hệt CI/CD
          Timeout).
- [ ] TSK-302 [Core] Cơ chế Retry Exponential Backoff. (Estimate: 3h)
    - Description:
        - Nếu Job failed, kiểm tra RetryCount < MaxRetries.
        - Nếu thoả mãn, tăng RetryCount, sleep một khoảng thời gian (VD: 2s) và
          đẩy Job ngược lại vào đuôi Queue (Enqueue) để worker khác lấy xử lý
          lại. Cập nhật Status là Retrying.
- [ ] TSK-303 [API/Control] Chức năng Cancel Job (Hủy tác vụ đang chạy).
  (Estimate: 4h)
    - Description:
        - Lưu trữ các context.CancelFunc của từng Job đang chạy vào một Map
          (map[string]context.CancelFunc).
        - Viết một hàm CancelJob(jobID). Khi gọi hàm này, lấy CancelFunc ra và
          thực thi cancel(). Worker sẽ lập tức ngắt Job đang chạy giữa chừng.

## 🌐 MODULE 4: HTTP SERVER & GRACEFUL SHUTDOWN (Giao tiếp bên ngoài)

- [ ] TSK-401 [API] Xây dựng REST API Submit & Kiểm tra Job. (Estimate: 4h)
    - Description:
        - Trong internal/api: Dùng net/http (sử dụng http.NewServeMux() mới của
          Go 1.22 hỗ trợ chia method GET/POST dễ dàng).
        - API POST /api/v1/jobs: Nhận payload JSON (Tên tác vụ, câu lệnh, số lần
          retry). Parse JSON tạo Job -> Lưu vào Store -> Push vào Queue. Trả về
          JobID.
        - API GET /api/v1/jobs/{id}: Truy vấn Store, trả về trạng thái hiện tại
          (Đang chạy, Đã xong, Output log của tác vụ).
- [ ] TSK-402 [API] Gắn API Handler vào Main. (Estimate: 2h)
    - Description:
        - Tại cmd/runner/main.go: Khởi tạo Queue, Store, WorkerPool (Start
          pool).
        - Khởi tạo HTTP Server mở port 8080.
- [ ] TSK-403 [DevOps] Cài đặt Graceful Shutdown (Tránh tắt nghẽn dữ liệu khi
  Stop Server). (Estimate: 4h)
    - Description:
        - Sử dụng os/signal lắng nghe SIGINT, SIGTERM (Khi bạn bấm Ctrl+C hoặc
          Docker stop).
        - Khi nhận tín hiệu tắt:
          1.  Không nhận API mới (Shutdown HTTP Server).
          2.  Đóng Queue Channel.
          3.  Dùng sync.WaitGroup đợi các Worker xử lý nốt các Job đang chạy dở
              rồi mới tắt hẳn chương trình.

## 🏆 MODULE 5: ADVANCED FEATURES (Bonus Level - Khi bạn đã làm xong Mod 4)

- [ ] TSK-501 [Adv-Data-Structure] Chuyển đổi Channel Queue sang Priority Queue.
  (Estimate: 4h)
    - Description:
        - Tái sử dụng code Priority-Queue.go bạn từng viết.
        - Job nào gán Priority = HIGH sẽ được nhét vào đầu hàng đợi. (Lưu ý: Lúc
          này bạn phải kết hợp Heap struct với sync.Cond hoặc pattern vòng lặp
          để worker lắng nghe thay vì Channel thông thường).
- [ ] TSK-502 [DevOps] Đóng gói Docker & Setup Makefile. (Estimate: 2h)
    - Description:
        - Viết Dockerfile Multi-stage build (Dùng golang alpine builder, sau đó
          copy binary sang image ubuntu nhỏ gọn rỗng để có bash shell thực thi
          job).
        - Viết Makefile chứa: make build, make run, make test.
