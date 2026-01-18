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

docker exec -it go-learning bash
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

## Go Modules and Packages

### Local Packages vs Separate Modules

#### Local Packages
- Same module (single `go.mod` file)
- Directory structure: `module/package/`
- Import: `"module-name/package-name"`
- No additional setup needed
- Example: `math/` and `hello/` packages

#### Separate Modules  
- Each has its own `go.mod` file
- Independent versioning and dependencies
- Must add as dependency first
- Example: `goodbye/` module

### Adding Dependencies

#### For Separate Modules
1. **Add the module to go.mod:**
   ```bash
   # Inside container
   go mod edit -replace module-name=./local-path
   ```

2. **Download and add the dependency:**
   ```bash
   go get module-name
   ```

3. **Then import in your code:**
   ```go
   import "module-name"
   ```

#### For External Dependencies
```bash
# GitHub packages
go get github.com/user/package

# Specific versions
go get github.com/user/package@v1.2.3
```

### Project Structure Example
```
project/
├── go.mod              # Main module: test-package
├── main.go
├── math/               # Local package
│   └── math.go
├── hello/              # Local package  
│   └── hello.go
└── goodbye/            # Separate module
    ├── go.mod          # Module: goodbye-module
    └── goodbye.go
```

### Additional Resources
Follow the Effective Go guide at https://go.dev/doc/effective_go