---
name: go-best-practices
description: Standards, testing patterns, and error handling conventions for Go CLI development.
---

# Go Best Practices & CLI Development

## 1. Project Commands
- Run tests: `go test -v ./...`
- Check test coverage: `go test -cover ./...`
- Build binary: `go build -o dockclean-smart main.go`
- Run binary with flags: `./dockclean-smart -lang=pt` or `./dockclean-smart -lang=en`

## 2. Docker Socket on macOS
When running locally on macOS with Docker Desktop:
- The default socket path is `unix://$HOME/.docker/run/docker.sock`.
- Ensure `DOCKER_HOST="unix://$HOME/.docker/run/docker.sock"` is set when executing the CLI directly.

## 3. Testing with Mocks
- Never hit the real Docker socket in unit tests.
- Always implement test cases against the mock implementation of `DockerAPIClient`.
- Test both success and error handling scenarios.
