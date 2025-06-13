package v1

import (
	"context"
	"net/http"
	"strings"
)

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Expose-Headers", "Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (h *Handler) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwtRaw, ok := r.Header["Authorization"]
		if !ok {
			h.log.Info("no Authorization header")
			next.ServeHTTP(w, r)
			return
		}

		jwtToken := strings.TrimPrefix(jwtRaw[0], "Bearer ")

		claims, err := h.jwtPayloadFromRequest(jwtToken)
		if err != nil {
			h.log.With("error", err).Info("jwt payload invalid")
			next.ServeHTTP(w, r)
			return
		}

		user, err := castClaims(claims)
		if err != nil {
			h.log.With("error", err).Info("cast claims")
			next.ServeHTTP(w, r)
			return
		}

		ctx := context.WithValue(r.Context(), userKey, user)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
