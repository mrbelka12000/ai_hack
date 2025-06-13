package v1

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	httpSwagger "github.com/swaggo/http-swagger"

	_ "github.com/mrbelka12000/ai_hack/docs"
	"github.com/mrbelka12000/ai_hack/internal/usecase"
)

var jwtSecretKey = []byte("very-secret-key")

const (
	apiPath = "/api/v1"
	userKey = "user_key"
)

type Handler struct {
	uc      *usecase.UseCase
	log     *slog.Logger
	decoder *schema.Decoder
}

func Init(uc *usecase.UseCase, mx *mux.Router, log *slog.Logger) {
	h := &Handler{
		uc:      uc,
		log:     log,
		decoder: schema.NewDecoder(),
	}
	mx.Use(CORSMiddleware)
	mx.Use(h.AuthMiddleware)

	mx.PathPrefix("/documentation/").Handler(CORSMiddleware(httpSwagger.WrapHandler))

	mx.HandleFunc(apiPath+"/user", h.UserCreate).Methods(http.MethodPost, http.MethodOptions)
	mx.HandleFunc(apiPath+"/users", h.UsersList).Methods(http.MethodGet, http.MethodOptions)
	mx.HandleFunc(apiPath+"/login", h.Login).Methods(http.MethodPost, http.MethodOptions)
	mx.HandleFunc(apiPath+"/profile", h.Profile).Methods(http.MethodGet, http.MethodOptions)

	mx.HandleFunc(apiPath+"/dialog", h.DialogCreate).Methods(http.MethodPost, http.MethodOptions)
	mx.HandleFunc(apiPath+"/dialog", h.DialogList).Methods(http.MethodGet, http.MethodOptions)
	mx.HandleFunc(apiPath+"/dialog_message", h.DialogAddMessage).Methods(http.MethodPost, http.MethodOptions)
	mx.HandleFunc(apiPath+"/dialog/{id}", h.DialogDelete).Methods(http.MethodDelete, http.MethodOptions)
	mx.HandleFunc(apiPath+"/dialog/{id}", h.DialogUpdate).Methods(http.MethodPut, http.MethodOptions)
	mx.HandleFunc(apiPath+"/dialog/{id}", h.DialogGet).Methods(http.MethodGet, http.MethodOptions)

	mx.HandleFunc(apiPath+"/personal_data", h.GetPersonalData).Methods(http.MethodPost, http.MethodOptions)
}
