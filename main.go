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

	mux := http.NewServeMux()

	// Define user credentials and excluded paths from config

	person := config.Person{
		UserName: config.UserAdmin,
		PassWord: config.UserPassword,
	}

	excluded := config.ExcludedPaths

	targetHost := config.TargetHost
	sessionWrapped := middleware.Session(person.UserName, person.PassWord, targetHost, excluded)(mux)

	handler := middleware.Recovery(sessionWrapped)

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
