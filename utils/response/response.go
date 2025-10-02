package response_utils

import (
	"encoding/json"
	"net/http"

	response_models "github.com/mbilarusdev/fool_base/v2/utils/response/models"
)

func WriteResponse(w http.ResponseWriter, code int, data any) {
	response := response_models.Response{
		Code: code,
		Data: data,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}

func WriteError(w http.ResponseWriter, code int, errMsg string) {
	WriteResponse(w, code, response_models.Response{
		Error: errMsg,
	})
}
