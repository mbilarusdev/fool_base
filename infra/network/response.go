package network

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code  int    `json:"code"`
	Data  any    `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

func WriteResponse(w http.ResponseWriter, code int, data any) {
	response := Response{
		Code: code,
		Data: data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}

func WriteError(w http.ResponseWriter, code int, errMsg string) {
	WriteResponse(w, code, Response{
		Error: errMsg,
	})
}
