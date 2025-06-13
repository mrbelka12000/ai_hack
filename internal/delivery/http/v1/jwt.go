package v1

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"

	aihack "github.com/mrbelka12000/ai_hack"
	"github.com/mrbelka12000/ai_hack/internal"
)

type Token struct {
	JWT string `json:"jwt"`
}

func (h *Handler) buildJWT(user internal.User) (string, error) {
	payload := jwt.MapClaims{
		"email": user.Email,
		"id":    user.ID,
		"role":  user.Role,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	t, err := token.SignedString(jwtSecretKey)
	if err != nil {
		h.log.With("error", err).Error("error signing token")
		return "", err
	}

	return t, nil
}

func (h *Handler) jwtPayloadFromRequest(tokenString string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})

	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	payload, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, err
	}

	return payload, nil
}

func castClaims(claims map[string]interface{}) (internal.User, error) {
	id, ok := claims["id"].(float64)
	if !ok {
		return internal.User{}, errors.New("invalid token")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return internal.User{}, errors.New("invalid token")
	}

	role, ok := claims["role"].(aihack.Role)
	if !ok {
		return internal.User{}, errors.New("invalid role")
	}

	return internal.User{
		ID:    int64(id),
		Email: email,
		Role:  role,
	}, nil
}
