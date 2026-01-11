# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

This is a Go-based web service called `mh-ui-service` built with the Gin web framework. The service is designed to be containerized and deployed via Docker.

## Build and Run Commands

### Local Development
```bash
# Build the application
go build -o mh-ui-service ./src/main.go

# Run the application
go run ./src/main.go

# Run with production environment
ENV=production go run ./src/main.go
```

### Docker
```bash
# Build Docker image
docker build -t tushark1704/mh-ui-service:latest .

# Run Docker container
docker run -d -p 8080:8080 tushark1704/mh-ui-service:latest

# View running containers
docker ps
docker ps -a

# Stop container
docker stop <CONTAINER_ID>
```

## Architecture

### Module Structure
- **Module name**: `morhaat.com/mh-ui-service`
- **Go version**: 1.24.4
- **Main dependencies**: Gin (web framework), Logrus (logging)

### Directory Layout
```
src/
├── main.go          # Application entry point, server initialization, route definitions
├── models/          # Data models and business entities
└── utils/           # Shared utilities (logging, etc.)
```

### Key Components

**Application Entry Point** ([src/main.go](src/main.go))
- Initializes logger via `utils.InitLogger()`
- Creates Gin server with default middleware
- Defines HTTP routes
- Server runs on port 8080

**Logging** ([src/utils/logger.go](src/utils/logger.go))
- Uses Logrus for structured logging
- Environment-aware configuration:
  - **Development** (default): Debug level, colored text formatter
  - **Production** (ENV=production): Info level, JSON formatter
- Singleton pattern: `utils.Log` is a global logger instance
- Must call `utils.InitLogger()` before using `utils.Log`

**Models** ([src/models/](src/models/))
- Contains domain models like `Product`
- Models include JSON tags for API serialization
- Implement String() methods for debugging

### Docker Build Strategy
- Multi-stage build using Alpine Linux
- Builder stage: golang:1.25.0-alpine3.22
- Runtime stage: alpine:latest (minimal footprint)
- Builds from `src/main.go` with all dependencies downloaded via go.mod