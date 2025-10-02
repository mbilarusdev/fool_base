package network

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/mbilarusdev/fool_base/v2/log"
	"go.uber.org/zap"
)

func Run(r *mux.Router, serviceName string) {
	addr := os.Getenv("ADDR")

	log.Info("Server starting on", zap.String("ADDR", addr))

	if err := http.ListenAndServe(addr, r); err != nil {
		log.Err(err, fmt.Sprintf("Listen and serve %v on %v failed!", serviceName, addr))
	}
}
