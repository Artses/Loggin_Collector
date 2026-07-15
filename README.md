# Log Colector (Go)

> 🔗 **Main API Repository (C#):** [Artses/Loggin](https://github.com/Artses/Loggin)

This repository contains the **Log Collector** developed in Go. It is part of a complete observability solution for auditing, contextualizing, and centrally managing logs in distributed applications.

---

## 📋 About the General Project

Development of an observability solution for auditing, contextualizing, and centrally managing logs in distributed applications. The project consists of a Go library integrated into applications via the **Gin** framework, responsible for collecting and contextualizing logs at runtime, and a main API developed in **C# using ASP.NET Core**, responsible for orchestrating the collection process, defining the rest of the lifecycle of the data present in the logs (processing, interpretation, storage/display, etc.).

---

## 🛠️ Technologies Used in the Collector

- **Language:** Go (v1.25)
- **Web Framework:** Gin-gonic (for providing log collection APIs)
- **Dynamic Reading:** `github.com/nxadm/tail` (for high-performance reading and monitoring of log files)
- **Database (Auxiliary):** PostgreSQL (provided via Docker Compose in the `docker-compose.yml` file)

---

## ⚙️ Directory Structure

The project follows the recommended structure for Go applications:

- `cmd/main.go`: Service entry point, configuring the Gin HTTP server and endpoints.
- `internal/dto/`: Data Transfer Objects (DTOs) for request validation (such as log path and filtering).
- `internal/handlers/`: Controllers responsible for handling HTTP requests and returning JSON responses.
- `internal/model/`: Definitions of log domain entities (`Log` and `LogItem`).
- `internal/repository/`: Log file access layer (using tailing to scan the requested file).
- `internal/service/`: Business logic to filter log lines starting from a specific identifier/order (`Order`).

---

## 🚀 How to Run

### Prerequisites

- **Go** (version 1.25 or higher) installed.

### Run the Go Log Collector

In the root directory of the project, run:

```bash
go run cmd/main.go
```

The application will start on port `8000`:
`Servidor rodando em http://localhost:8000`

> Yes it's simple like that :D
---

## 📡 API Endpoints

### 1. Log Ingestion/Retrieval
* **URL:** `/api/v1/logs`
* **Method:** `POST`
* **Description:** Reads the log file specified in `path` and returns lines whose order number is greater than the provided `order` parameter.
* **Request Body (JSON):**
  ```json
  {
    "path": "path/to/your/file.log",
    "order": 0
  }
  ```
* **Example Response (JSON):**
  ```json
  {
    "content": [
      {
        "order": 1,
        "timestamp": "2026-07-06T05:30:00Z",
        "line": "[INFO] Starting payment service"
      },
      {
        "order": 2,
        "timestamp": "2026-07-06T05:31:12Z",
        "line": "[SUCCESS] Payment processed successfully"
      }
    ]
  }
  ```

### 2. Health Check
* **URL:** `/api/v1/healthstatus`
* **Method:** `GET`
* **Description:** Checks if the log collector service is online.
* **Example Response (JSON):**
  ```json
  {
    "message": "i'm alive ;D"
  }
  ```

---

## 🤝 Contributing

If you wish to contribute or report issues regarding log orchestration and management, please visit the main C# repository: [Artses/Loggin](https://github.com/Artses/Loggin).
