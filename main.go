package main

import (
	"context"
	"fmt"
	"hifi/config"
	"hifi/middleware"
	"hifi/routes/rest"
	"hifi/types"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// Define subsonic user credentials
	person := types.Person{
		UserName: "",
		PassWord: "",
	}

	// Hifi proxy
	validPaths := config.ValidPaths

	if config.MODE == "managed" {
		slog.Info("Hifi Running in MANAGED mode")
	} else {
		go middleware.StartTidalRefresher()
	}
	go middleware.RecentAlbum()
	go middleware.StartFreshRefresher()

	// HTTP server setup
	mux := http.NewServeMux()

	mux.HandleFunc(rest.Fresh(), middleware.FreshHandler)

	// Middleware setup
	session := middleware.Session(person.UserName, person.PassWord, validPaths)(mux)

	cors := middleware.CORS(session)

	port := middleware.PortRotate()

	handler := middleware.Recovery(cors)

	// Server setup

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", config.Host, port),
		Handler: handler,
	}

	slog.Info("Hifi API server running",
		"host", config.Host,
		"port", port,
		"url", fmt.Sprintf("%s://%s:%s", config.HifiScheme, config.Host, port),
	)

	// Run server in background
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Failed to start server", "error", err)
			os.Exit(1)
		}
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	slog.Info("Shutting down gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Server forced to shutdown", "error", err)
	}

	slog.Info("Shutdown complete")

}
