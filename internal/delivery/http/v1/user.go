package v1

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/mrbelka12000/ai_hack/internal"
	"github.com/mrbelka12000/ai_hack/pkg/validator"
)

// UserCreate godoc
// @Summary      Create user
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        data body internal.UserCU    true "User object"
// @Success      201
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /user [post]
func (h *Handler) UserCreate(w http.ResponseWriter, r *http.Request) {
	var obj internal.UserCU

	if err := json.NewDecoder(r.Body).Decode(&obj); err != nil {
		h.errorResponse(w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := validator.ValidateStruct(obj); err != nil {
		h.errorResponse(w, err, http.StatusBadRequest)
		return
	}

	if err := h.uc.UserCreate(r.Context(), obj); err != nil {
		h.errorResponse(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Login godoc
// @Summary      Login
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        data body internal.UserLogin    true "User login object"
// @Success      201  {object}  Token
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /login [post]
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var obj internal.UserLogin
	if err := json.NewDecoder(r.Body).Decode(&obj); err != nil {
		h.errorResponse(w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := validator.ValidateStruct(obj); err != nil {
		h.errorResponse(w, err, http.StatusBadRequest)
		return
	}

	user, err := h.uc.UserLogin(r.Context(), obj)
	if err != nil {
		h.errorResponse(w, err, http.StatusBadRequest)
		return
	}

	jwt, err := h.buildJWT(user)
	if err != nil {
		h.errorResponse(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Token{
		JWT: jwt,
	})
}

// UsersList godoc
// @Summary      List users
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        email    query     string  false  "search by client_id"
// @Param        role    query     string  false  "search by operator_id"
// @Param        limit    query     int  false  "search by limit"
// @Param        offset    query     string  false  "search by offset"
// @Success      200  {object}	internal.UserListResponse
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /users [get]
func (h *Handler) UsersList(w http.ResponseWriter, r *http.Request) {
	var pars internal.UserPars
	if err := h.decoder.Decode(&pars, r.URL.Query()); err != nil {
		h.errorResponse(w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := validator.ValidateStruct(pars); err != nil {
		h.errorResponse(w, err, http.StatusBadRequest)
		return
	}

	result, err := h.uc.UserList(r.Context(), pars)
	if err != nil {
		h.errorResponse(w, err, http.StatusInternalServerError)
		return
	}

	response := internal.Response{
		Result: result,
		PaginationParams: internal.PaginationParams{
			Limit:  pars.Limit,
			Offset: pars.Offset,
			Page:   pars.CalculatePage(),
		},
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.errorResponse(w, err, http.StatusInternalServerError)
		return
	}
}

// Profile godoc
// @Summary      GetProfile
// @Tags         user
// @Accept       json
// @Produce      json
// @Success      200  {object}  internal.User
// @Failure      400  {object}  ErrorResponse
// @Failure      404  {object}  ErrorResponse
// @Failure      500  {object}  ErrorResponse
// @Router       /profile [post]
// @Security Bearer
func (h *Handler) Profile(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(userKey).(internal.User)
	if !ok {
		h.errorResponse(w, errors.New(http.StatusText(http.StatusUnauthorized)), http.StatusUnauthorized)
		return
	}

	response, err := h.uc.UserGet(r.Context(), internal.UserGetPars{
		ID: user.ID,
	})
	if err != nil {
		h.errorResponse(w, err, http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		h.errorResponse(w, err, http.StatusInternalServerError)
		return
	}
}
