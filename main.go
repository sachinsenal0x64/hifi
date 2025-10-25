package main

import (
	"fmt"
	"hifi/config"
	"hifi/middleware"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	person := config.Person{
		UserName: "admin",
		PassWord: "admin",
	}

	excluded := []string{"/admin/home"}

	targetHost := "http://localhost:4747"
	sessionWrapped := middleware.Session(person.UserName, person.PassWord, targetHost, excluded)(mux)

	handler := middleware.Recovery(sessionWrapped)

	fmt.Printf("Hifi API server running at http://%s:%s\n", config.Host, config.Port)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", config.Host, config.Port), handler); err != nil {
		log.Fatal(err)
	}

}
