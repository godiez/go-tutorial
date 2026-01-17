# Go Learning Docker Project

## Requirements

- Docker
- Docker Compose

## Getting Started

### Option 1: Using Docker Compose (Recommended)

1. Make sure you have Docker and Docker Compose installed on your system
2. Run the container with:
   ```bash
   docker-compose -f docker/docker-compose.yml up
   ```

### Option 2: Using Docker directly

1. Build the Docker image:
   ```bash
   docker build -t go-learning docker/
   ```

2. Run the container:
   ```bash
   docker run -it -v $(pwd):/go-learning go-learning
   ```

## Using the Container

### Entering the container
```bash
docker-compose -f docker/docker-compose.yml exec go-learning bash
```

### Running Go programs
Inside the container, you can:
- Create new Go files: `touch main.go`
- Run Go programs: `go run main.go`
- Format code: `gofmt -w main.go` or `goimports -w main.go`
- Check for issues: `staticcheck ./...`

### Available Tools
The container includes:
- goimports - for managing imports
- govulncheck - for security vulnerability checking
- staticcheck - for static analysis

### Additional Resources
Follow the Effective Go guide at https://go.dev/doc/effective_go