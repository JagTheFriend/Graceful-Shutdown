package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func createServer() *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/example", func(w http.ResponseWriter, r *http.Request) {
		log.Default().Println("Slow request started")
		time.Sleep(8 * time.Second)
		log.Default().Println("Slow request finished")
		_, err := w.Write([]byte("Hello, World!"))
		if err != nil {
			log.Fatal(err.Error())
		}
	})

	return &http.Server{
		Addr:    ":8000",
		Handler: mux,
	}
}

func runServer(
	ctx context.Context,
	server *http.Server,
	shutdownDuration time.Duration,
) error {
	// Use buffered channel to handle server shutdown errors and not block main thread
	serverError := make(chan error, 1)

	// Start server in go route to not block
	go func() {
		log.Default().Println("Starting server on", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			// Send error to channel if server fails
			serverError <- err
		}
		close(serverError)
	}()

	// Listen for termination signal to gracefully shutdown the server
	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(stopSignal)
	defer close(stopSignal)

	select {
	case err := <-serverError:
		return err

	case <-stopSignal:
		log.Default().Println("Shutting down server -- Received termination signal")
		return shutDownServer(server, shutdownDuration)

	case <-ctx.Done():
		log.Default().Println("Shutting down server -- Received context cancellation")
		return shutDownServer(server, shutdownDuration)
	}
}

func shutDownServer(server *http.Server, shutdownDuration time.Duration) error {
	// Create context with timeout for server shutdown
	ctx, cancel := context.WithTimeout(context.Background(), shutdownDuration)
	defer cancel()

	// Attempt graceful server shutdown and fallback to hard close if it fails
	if shutdownErr := server.Shutdown(ctx); shutdownErr != nil {
		if closeErr := server.Close(); closeErr != nil {
			return errors.Join(shutdownErr, closeErr)
		}
	}
	log.Default().Println("Server gracefully stopped")
	return nil
}

func main() {
	server := createServer()
	shutdownDuration := 5 * time.Second

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := runServer(ctx, server, shutdownDuration); err != nil {
		log.Fatal(err.Error())
	}
}
