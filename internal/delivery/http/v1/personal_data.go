package v1

import (
	"encoding/json"
	"net/http"

	"github.com/mrbelka12000/ai_hack/internal"
)

// GetPersonalData godoc
// @Summary      Get personal data call only from ai sufler
// @Tags         personal_data
// @Accept       json
// @Produce      json
// @Param        data body internal.PersonalDataRequest    true "Personal Data object"
// @Success      201
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /personal_data [post]
func (h *Handler) GetPersonalData(w http.ResponseWriter, r *http.Request) {
	var obj internal.PersonalDataRequest
	if err := json.NewDecoder(r.Body).Decode(&obj); err != nil {
		h.errorResponse(w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	resp, err := h.uc.GetPersonalData(r.Context(), obj)
	if err != nil {
		h.errorResponse(w, err, http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		h.errorResponse(w, err, http.StatusInternalServerError)
		return
	}
}
