package middleware

import (
	"api/config"
	"fmt"
	"log/slog"
	"net"
	"os"
)

func PortRotate() string {
	for i, p := range config.Port {
		if ln, err := net.Listen("tcp", fmt.Sprintf("%s:%s", config.Host, p)); err == nil {
			_ = ln.Close()
			return p
		}

		if i+1 < len(config.Port) {
			slog.Warn(fmt.Sprintf("%s port in use, trying next... port=%s", p, config.Port[i+1]))
		} else {
			slog.Error(fmt.Sprintf("%s port in use, no more ports available", p))
			os.Exit(1)
		}
	}
	return ""
}
