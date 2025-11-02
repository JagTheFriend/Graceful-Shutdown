# Graceful Shutdown

## Description

This is a simple web server written in Go that demonstrates graceful shutdown.
It listens for OS signals (like SIGINT and SIGTERM) and performs cleanup operations before exiting.

## Usage

To run the application, execute the following command:

```bash
go run ./cmd/main.go
```

The application will start and listen on port 8000.\
You can access the application by opening a web browser and navigating to `http://localhost:8000`.
