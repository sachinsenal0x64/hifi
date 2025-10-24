package middleware

import (
	"encoding/json"
	"hifi/config"
	"log"
	"net/http"
)

// -------------------- RECOVERY --------------------

type ErrorResponse struct {
	Message string `json:"message"`
}

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				log.Printf("Recovered from panic: %v\n", rec)
				if w.Header().Get(config.HeaderContentType) == "" {
					w.Header().Set(config.HeaderContentType, config.ContentTypeJSON)
				}
				w.WriteHeader(config.StatusInternalServerError)
				_ = json.NewEncoder(w).Encode(ErrorResponse{
					Message: "internal server error",
				})
			}
		}()
		next.ServeHTTP(w, r)
	})
}
