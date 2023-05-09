# 1.Build a small image with latex
FROM ubuntu:lunar AS runner

    # Install latex packages
    RUN apt update 
    RUN apt install -y latexmk texlive-fonts-extra

# 2. Build Golang Application
FROM golang:1.19.1-alpine3.16 AS builder

    # Set environmet variables needed for our image
    ENV GO111MODULE=on \
        CGO_ENABLED=0 \
        GOOS=linux \
        GOARCH=amd64

    # Move to working directory /build
    WORKDIR /build

    # Copy and download dependency using go mod
    COPY go.mod .
    COPY go.sum .
    RUN go mod download

    # Copy the code into the container
    COPY . .

    # Build the application
    RUN go build -o main ./cmd/main.go

# 3.Copy binary into runner
FROM runner

    COPY --from=builder /build/main /
    COPY --from=builder /build/template /template

    # Command to run
    # ENTRYPOINT ["/main", "-email=","dockeremail","-phone=","dockerphone"]