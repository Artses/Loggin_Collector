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

### Option 1: Docker (Recommended for production/cross-platform)

The Docker installation mounts a folder containing host log files into the container so the API can read them.

#### 1. Setup Environment
Define where your host log files are stored by creating a `.env` file in the root of the project, or setting the environment variable `HOST_LOGS_DIR`.
Example `.env` file:
```env
# On Windows
HOST_LOGS_DIR=C:\path\to\your\logs

# On Linux
# HOST_LOGS_DIR=/var/log/your_app
```
*Note: If no path is specified, it will default to a `./logs` directory in the project root.*

#### 2. Run with Docker Compose
In the root directory, run:
```bash
docker-compose up -d --build
```
This builds the multi-stage image and mounts your logs to `/var/log/collector` inside the container as **read-only**.

#### 3. Querying logs in Docker
When querying the API via Docker, you must refer to the files using their container-internal path (`/var/log/collector/...`).
For example, to read `C:\path\to\your\logs\app.log` (on Windows host) or `/var/log/your_app/app.log` (on Linux host), send:
```json
{
  "path": "/var/log/collector/app.log",
  "order": 0
}
```

---

### Option 2: Native Compilation

#### Windows
You can use the PowerShell helper script:
```powershell
# Build the binary (log-collector.exe)
.\build.ps1 build

# Build and run the service
.\build.ps1 run

# Clean up build artifacts
.\build.ps1 clean
```

#### Linux / macOS
You can use the `Makefile`:
```bash
# Build the binary (log-collector)
make build

# Build and run the service
make run

# Clean up build artifacts
make clean
```

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
