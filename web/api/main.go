package main

import (
	"api/config"
	"api/middleware"
	"fmt"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	handler := middleware.Recovery(mux)

	fmt.Printf("Hifi web running at http://%s:%s\n", config.Host, config.Port)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", config.Host, config.Port), handler); err != nil {
		log.Fatal(err)
	}
}
