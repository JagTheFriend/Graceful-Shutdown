# Graceful Shutdown

## Description

This is a simple web server written in Go that demonstrates graceful shutdown.
It listens for OS signals (like SIGINT and SIGTERM) and performs cleanup operations before exiting.


## Signal Handling

The application listens for OS signals (SIGINT and SIGTERM) to perform cleanup operations before exiting.

### Handling Signals

When the application receives SIGINT/SIGTERM, it will perform the following cleanup operations:

1. Log a message indicating that the server is shutting down.
2. Send a SIGINT signal to the server's goroutine to gracefully shutdown the server.
3. Wait for the server's goroutine to finish shutting down.
4. Log a message indicating that the server has shut down.
5. Exit the application.

## Usage

To run the application, execute the following command:

```bash
go run ./cmd/main.go
```

The application will start and listen on port 8000.\
You can access the application by opening a web browser and navigating to `http://localhost:8000/example`.

This route will simulate a slow request that takes 8 seconds to complete.

### Testing Graceful Shutdown

To test graceful shutdown, you can send a SIGINT signal to the application.\
This will cause the application to perform the cleanup operations and exit.\
You can send a SIGINT signal by pressing `Ctrl+C` in the terminal where the application is running.\