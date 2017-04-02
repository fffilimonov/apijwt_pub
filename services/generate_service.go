package services

import (
	"../core/authentication"
	"../services/models"
	"encoding/json"
	"net/http"
)

func Generate(requestUser *models.User) (int, []byte) {
	authBackend := authentication.InitJWTAuthenticationBackend()

	token, err := authBackend.GenerateToken(requestUser)
	if err != nil {
		if token != "" {
			return http.StatusUnauthorized, []byte(token)
		} else {
			return http.StatusInternalServerError, []byte("")
		}
	} else {
		response, _ := json.Marshal(models.TokenAuthentication{token})
		return http.StatusOK, response
	}

	return http.StatusUnauthorized, []byte("")
}
