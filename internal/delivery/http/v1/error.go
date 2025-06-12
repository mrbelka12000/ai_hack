package v1

import (
	"encoding/json"
	"net/http"
)

type (
	ErrorResponse struct {
		Message string `json:"message"`
	}
)

func (h *Handler) errorResponse(w http.ResponseWriter, err error, code int) {
	h.log.With("error", err).Error("error response")
	body, err := json.Marshal(ErrorResponse{Message: err.Error()})
	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
