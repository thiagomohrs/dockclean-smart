# Guidelines for dockclean-smart

## Architecture & Code Standards
- **Language:** Go 1.22+
- **Architecture:** Clean package structure under `internal/` (`internal/docker`, `internal/i18n`).
- **Testing:** Mandatory unit tests for any new feature in `internal/`. Tests must execute without a live Docker daemon.
- **Docker SDK:** Use `github.com/docker/docker/client` through the `DockerAPIClient` interface.
- **Environment:** Respect `DOCKER_HOST` environment variable for macOS Docker Desktop compatibility.
