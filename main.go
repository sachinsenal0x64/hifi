package main

import (
	"fmt"
	"hifi/config"
	"hifi/middleware"
	"log"
	"net/http"
)

func main() {

	fmt.Println("Access Token:", config.AccessToken)

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

	fmt.Printf("Hifi API server running at http://%s:%s\n", config.Host, config.Port)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", config.Host, config.Port), handler); err != nil {
		log.Fatal(err)
	}

}
