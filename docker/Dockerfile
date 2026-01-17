# Use the official Go image with the latest version
FROM golang:latest

# Set the working directory
WORKDIR /go/src/app

# Install useful Go tools for learning
RUN go install golang.org/x/tools/cmd/goimports@latest && \
    go install golang.org/x/vuln/cmd/govulncheck@latest && \
    go install honnef.co/go/tools/cmd/staticcheck@latest

# Set environment variables for Go
ENV GOPATH=/go
ENV GOBIN=/go/bin

# Create a sample directory for your Go learning projects
RUN mkdir -p /go-learning

# Set the default command to bash
CMD ["/bin/bash"]

# Instructions:
# 1. Build the Docker image: docker build -t go-learning .
# 2. Run the container: docker run -it -v $(pwd):/go-learning go-learning
# 3. Inside the container, you can:
#    - Create new Go files: touch main.go
#    - Run Go programs: go run main.go
#    - Format code: gofmt -w main.go or goimports -w main.go
#    - Check for issues: staticcheck ./...
#    - Follow the Effective Go guide at https://go.dev/doc/effective_go