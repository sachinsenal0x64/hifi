package routes

import (
	"api/middleware"
	"net/http"
)

func Handle() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/signin", middleware.SigninUser)
	mux.HandleFunc("/signup", middleware.SignupUser)
	mux.HandleFunc("/validate", middleware.ValidateHandler)
	mux.HandleFunc("/delete", middleware.DeleteUser)
	mux.HandleFunc("/update", middleware.UpdateUser)

	return mux
}
