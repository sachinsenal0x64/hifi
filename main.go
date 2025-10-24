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

	handler := middleware.RecoveryMiddleware(mux)

	fmt.Printf("Hifi API server running at http://%s:%s\n", config.Host, config.Port)
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", config.Host, config.Port), handler); err != nil {
		log.Fatal(err)
	}

}
