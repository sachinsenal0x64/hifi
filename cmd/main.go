package main

import (
	"fmt"
	"hifi/config"
	"hifi/middleware"
	"log/slog"
	"net/http"
	"os"
)

func main() {

	// Define subsonic user credentials
	person := config.Person{
		UserName: "",
		PassWord: "",
	}

	// Hifi proxy
	excluded := config.ExcludedPaths
	targetHost := config.TargetHost

	// HTTP server setup
	mux := http.NewServeMux()

	// Middleware setup
	session := middleware.Session(person.UserName, person.PassWord, targetHost, excluded)(mux)

	cors := middleware.CORS(session)

	handler := middleware.Recovery(cors)

	slog.Info("Hifi API server running",
		"host", config.Host,
		"port", config.Port,
		"url", "http://"+config.Host+":"+config.Port,
	)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", config.Host, config.Port), handler); err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}

}
