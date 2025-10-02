package network

import (
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func CheckGatewayMiddleware(next http.Handler) http.Handler {
	addr := os.Getenv("GATEWAY_ADDR")
	if addr == "" {
		panic("Env var GATEWAY_ADDR not founded")
	}
	splittedAddr := strings.Split(addr, ":")
	if len(splittedAddr) != 2 {
		panic("Failed to parse GATEWAY_ADDR")
	}
	var isLocalHost bool
	gatewayPort, err := strconv.Atoi(splittedAddr[1])
	if err != nil {
		panic("Failed to parse GATEWAY_ADDR port")
	}
	gatewayHost := splittedAddr[0]
	isLocalHost = checkLocalHost(gatewayHost)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		remoteAddr := r.RemoteAddr
		remoteHost, portStr, err := net.SplitHostPort(remoteAddr)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		remotePort, err := strconv.Atoi(portStr)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if isLocalHost && checkLocalHost(remoteHost) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if !isLocalHost && remoteHost != gatewayHost {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		if remotePort != gatewayPort {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func checkLocalHost(host string) bool {
	isLocalHost := false
	for _, v := range []string{"", "::1", "127.0.0.1", "localhost"} {
		isLocalHost = v == host
		if isLocalHost {
			break
		}
	}
	return isLocalHost
}
