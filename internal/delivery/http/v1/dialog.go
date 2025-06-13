package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	aihack "github.com/mrbelka12000/ai_hack"
	"github.com/mrbelka12000/ai_hack/internal"
	"github.com/mrbelka12000/ai_hack/pkg/validator"
)

// DialogCreate godoc
// @Summary      Create dialog
// @Tags         dialog
// @Accept       json
// @Produce      json
// @Param        data body internal.DialogCU    true "Dialog object"
// @Success      201
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /dialog [post]
// @Security Bearer
func (h *Handler) DialogCreate(w http.ResponseWriter, r *http.Request) {
	var obj internal.DialogCU
	if err := json.NewDecoder(r.Body).Decode(&obj); err != nil {
		h.errorResponse(w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := validator.ValidateStruct(obj); err != nil {
		h.errorResponse(w, err, http.StatusBadRequest)
		return
	}

	fmt.Println(r.Context().Value(userKey))
	user, ok := r.Context().Value(userKey).(internal.User)
	if !ok {
		h.errorResponse(w, errors.New(http.StatusText(http.StatusUnauthorized)), http.StatusUnauthorized)
		return
	}

	obj.ClientID = user.ID

	id, err := h.uc.DialogCreate(r.Context(), obj)
	if err != nil {
		h.errorResponse(w, err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(internal.DialogCreateResp{
		ID: id,
	}); err != nil {
		h.errorResponse(w, err, http.StatusInternalServerError)
		return
	}
}

// DialogGet godoc
// @Summary      Get dialog
// @Tags         dialog
// @Accept       json
// @Produce      json
// @Success      200  {object}	internal.Dialog
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /dialog/{id} [get]
func (h *Handler) DialogGet(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]

	id, err := uuid.Parse(idStr)
	if err != nil {
		h.errorResponse(w, err, http.StatusBadRequest)
		return
	}

	response, err := h.uc.DialogGet(r.Context(), id)
	if err != nil {
		h.errorResponse(w, err, http.StatusInternalServerError)
		return
	}

	body, err := json.Marshal(response)
	if err != nil {
		h.errorResponse(w, err, http.StatusInternalServerError)
		return
	}
	w.Write(body)
}

// DialogAddMessage godoc
// @Summary      Continue dialog
// @Tags         dialog
// @Accept       json
// @Produce      json
// @Param        data body internal.DialogMessage    true "DialogMessage object"
// @Success      200  {object}  internal.DialogMessageResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /dialog_message [post]
// @Security Bearer
func (h *Handler) DialogAddMessage(w http.ResponseWriter, r *http.Request) {
	var obj internal.DialogMessage
	if err := json.NewDecoder(r.Body).Decode(&obj); err != nil {
		h.errorResponse(w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	user, ok := r.Context().Value(userKey).(internal.User)
	role := user.Role
	if !ok {
		role = aihack.RoleClient
	}

	obj.Role = role

	if err := validator.ValidateStruct(obj); err != nil {
		h.errorResponse(w, err, http.StatusBadRequest)
		return
	}

	resp, err := h.uc.DialogAddMessage(r.Context(), obj)
	if err != nil {
		h.errorResponse(w, err, http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(resp); err != nil {
		h.errorResponse(w, err, http.StatusInternalServerError)
		return
	}
}

// DialogUpdate godoc
// @Summary      Update dialog
// @Tags         dialog
// @Accept       json
// @Produce      json
// @Param        data body internal.DialogCU    true "Dialog object"
// @Success      200
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /dialog/{id} [put]
func (h *Handler) DialogUpdate(w http.ResponseWriter, r *http.Request) {
	var obj internal.Dialog
	if err := json.NewDecoder(r.Body).Decode(&obj); err != nil {
		h.errorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	defer r.Body.Close()

	if err := validator.ValidateStruct(obj); err != nil {
		h.errorResponse(w, err, http.StatusBadRequest)
		return
	}

	if err := h.uc.DialogUpdate(r.Context(), obj); err != nil {
		h.errorResponse(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DialogDelete godoc
// @Summary      Delete dialog
// @Tags         dialog
// @Accept       json
// @Produce      json
// @Success      204
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /dialog/{id} [delete]
func (h *Handler) DialogDelete(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := uuid.Parse(idStr)
	if err != nil {
		h.errorResponse(w, err, http.StatusBadRequest)
		return
	}

	if err := h.uc.DialogDelete(r.Context(), id); err != nil {
		h.errorResponse(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DialogList godoc
// @Summary      List dialogs
// @Tags         dialog
// @Accept       json
// @Produce      json
// @Param        client_id    query     int  false  "search by client_id"
// @Param        operator_id    query     int  false  "search by operator_id"
// @Param        status    query     string  false  "search by status"
// @Param        limit    query     int  false  "search by limit"
// @Param        offset    query     string  false  "search by offset"
// @Success      200  {object}	internal.DialogListResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /dialog [get]
func (h *Handler) DialogList(w http.ResponseWriter, r *http.Request) {
	var pars internal.DialogPars
	if err := h.decoder.Decode(&pars, r.URL.Query()); err != nil {
		h.errorResponse(w, err, http.StatusBadRequest)
		return
	}

	response, err := h.uc.DialogList(r.Context(), pars)
	if err != nil {
		h.errorResponse(w, err, http.StatusInternalServerError)
		return
	}

	result := internal.Response{
		Result: response,
		PaginationParams: internal.PaginationParams{
			Limit:  pars.Limit,
			Offset: pars.Offset,
			Page:   pars.PaginationParams.CalculatePage(),
		},
	}

	body, err := json.Marshal(result)
	if err != nil {
		h.errorResponse(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
