package middleware

import (
	"api/config"
	"encoding/json"
	"log"
	"net/http"
)

// -------------------- RECOVERY --------------------

type ErrorResponse struct {
	Message string `json:"message"`
}

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("Recovered from panic: %v\n", rec)
				w.Header().Set(config.HeaderContentType, config.ContentTypeJSON)
				w.WriteHeader(config.StatusInternalServerError)
				_ = json.NewEncoder(w).Encode(ErrorResponse{
					Message: "Internal server error",
				})
			}
		}()
		next.ServeHTTP(w, r)
	})
}
