package middleware

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/valkey-io/valkey-go"
)

type Router struct {
	Valkey valkey.Client
	Mem    map[string]string
	MemMu  sync.Mutex
}

func NewRouter(valkeyAddr string) (*Router, error) {
	client, err := valkey.NewClient(valkey.ClientOption{
		InitAddress: []string{valkeyAddr},
	})
	if err != nil {
		return nil, err
	}

	return &Router{
		Valkey: client,
		Mem:    make(map[string]string),
	}, nil
}

// ------------------------- CLOUD HANDLER -------------------------

func SendToCloud(action, key, value string) {
	fmt.Printf("[CLOUD] action=%s key=%s value=%s\n", action, key, value)
}

// ------------------------- GET -------------------------

func (r *Router) Get(ctx context.Context, key string) (string, error) {
	if os.Getenv("CLOUD_HOST") != "" {
		r.MemMu.Lock()
		v, ok := r.Mem[key]
		r.MemMu.Unlock()

		if !ok {
			return "", fmt.Errorf("cloud key missing")
		}

		go SendToCloud("get", key, v)
		return v, nil
	}

	// cmd := r.Valkey.B().Get().Key(key).Build()
	// r.Valkey.Do(ctx, cmd).ToString()
	return "", nil
}

// ----------------------- SET -----------------------

func (r *Router) Set(ctx context.Context, key, val string) (bool, error) {
	if os.Getenv("CLOUD_HOST") != "" {

		r.MemMu.Lock()
		r.Mem[key] = val
		r.MemMu.Unlock()

		go SendToCloud("set", key, val)
		return true, nil
	}

	// cmd := r.Valkey.B().Set().Key(key).Value(val).Build()
	// err := r.Valkey.Do(ctx, cmd).Error()
	// if err != nil {
	// 	return false, err
	// }

	return false, nil
}

// ----------------------- DEL -----------------------

func (r *Router) Del(ctx context.Context, key string) (bool, error) {
	if os.Getenv("CLOUD_HOST") != "" {
		r.MemMu.Lock()
		_, existed := r.Mem[key]
		if existed {
			delete(r.Mem, key)
		}
		r.MemMu.Unlock()

		go SendToCloud("del", key, "")
		return !existed, nil

	}

	// cmd := r.Valkey.B().Del().Key(key).Build()
	// n, err := r.Valkey.Do(ctx, cmd).AsInt64()
	// if err != nil {
	// 	return false, err
	// }

	// n == 0

	return false, nil
}
