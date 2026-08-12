# dockclean-smart (Smart Docker Cleanup)

[Português](README.md) | [English](README_EN.md)

CLI tool written in Go for smart identification of unused Docker images and disk space recovery estimation.

## 🚀 Features & Architecture

- **Unused Image Detection:** Identifies images not associated with running or stopped containers (`Containers == 0`).
- **Reclaimable Space Calculation:** Displays individual image sizes in MB and total cumulative space in MB and GB.
- **Graceful Context Cancellation:** OS signal handling (`SIGINT`, `SIGTERM` / `Ctrl+C`).
- **Testable Architecture (Mocks):** Docker SDK API abstraction via `DockerAPIClient` interface for isolated unit tests without requiring a running Docker daemon.
- **Built-in i18n:** Portuguese and English output via `-lang=en|pt` CLI flag or OS environment auto-detection (`LANG`/`LC_ALL`).

## 📁 Project Structure

```plaintext
dockclean-smart/
├── go.mod
├── main.go
├── README.md
├── README_EN.md
└── internal/
    ├── docker/
    │   ├── cliente.go
    │   └── cliente_test.go
    └── i18n/
        ├── i18n.go
        └── i18n_test.go
```

## 🛠️ Usage Guide

### Prerequisites
- [Go 1.22+](https://go.dev/dl/) installed.
- [Docker Daemon](https://www.docker.com/) running locally (for CLI execution).

### Direct Execution

```bash
# English output
go run main.go -lang=en

# Portuguese output
go run main.go -lang=pt
```

### Running Unit Tests

Run unit tests using isolated mocks (no Docker daemon required):

```bash
go test -v ./...
```

### Build Executable

```bash
go build -o dockclean-smart main.go
./dockclean-smart -lang=en
```

### CLI Flags

- `-lang`: Output language (`en` or `pt`). Auto-detected if omitted.
