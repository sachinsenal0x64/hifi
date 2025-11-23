package middleware

import (
	"api/config"
	"api/types"
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	users = make(map[string]*types.User)
	mu    sync.RWMutex

	registerJobs = make(chan types.RegisterJob, 10)

	tokenHashes = make(map[string][32]byte)
)

func RegistrationWorker() {
	for job := range registerJobs {
		creds := job.Creds

		mu.Lock()
		// if _, exists := users[creds.Username]; exists {
		// 	mu.Unlock()
		// 	job.Reply <- types.RegisterResult{
		// 		Success: false,
		// 		Message: "⚠️ User already exists",
		// 		User:    nil,
		// 	}
		// 	continue
		// }

		user := &types.User{
			ID:       uuid.NewString(),
			Username: creds.Username,
			Password: creds.Password,
		}

		users[creds.Username] = user
		mu.Unlock()

		job.Reply <- types.RegisterResult{
			Success: true,
			Message: "✅ User registered successfully",
			User:    user,
		}
	}
}

func registerUser(username, password string) (types.RegisterResult, error) {
	replyCh := make(chan types.RegisterResult)

	job := types.RegisterJob{
		Creds: types.SignupRequest{
			Username: username,
			Password: password,
		},
		Reply: replyCh,
	}

	select {
	case registerJobs <- job:
		res := <-replyCh
		if !res.Success {
			return res, errors.New(res.Message)
		}
		return res, nil

	default:
		return types.RegisterResult{}, fmt.Errorf("server busy, try again later")
	}
}

func ValidateHandler(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
		return
	}

	var tokenString string
	fmt.Sscanf(authHeader, "Bearer %s", &tokenString)

	token, err := jwt.ParseWithClaims(tokenString, &types.Claims{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return config.JwtSecret, nil
	})

	if err != nil || !token.Valid {
		http.Error(w, "Invalid or expired token", http.StatusBadRequest)
		return
	}

	claims, ok := token.Claims.(*types.Claims)
	if !ok {
		http.Error(w, "Invalid claims", http.StatusBadRequest)
		return
	}

	computed := sha256.Sum256([]byte(claims.RegisteredClaims.ID))

	stored, ok := tokenHashes[claims.RegisteredClaims.ID]
	if !ok || !bytes.Equal(stored[:], computed[:]) {
		http.Error(w, "Invalid token", http.StatusBadRequest)
		return
	}

	mu.RLock()
	user, exists := users[claims.Username]
	mu.RUnlock()

	if !exists || user.ID != claims.ID {
		http.Error(w, "User does not exist", http.StatusBadRequest)
		return
	}

	resp := map[string]string{
		"host":     config.Scheme + "://" + config.ProxyHost,
		"username": user.Username,
		"password": user.Password,
	}

	w.Header().Set(config.HeaderContentType, config.ContentTypeJSON)
	json.NewEncoder(w).Encode(resp)
}
