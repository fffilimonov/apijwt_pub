package authentication

import (
	"../../services/models"
	"../redis"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	request "github.com/dgrijalva/jwt-go/request"
	"github.com/gorilla/context"
	"net/http"
)

func RequireTokenAuthentication(rw http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	authBackend := InitJWTAuthenticationBackend()

	token, err := request.ParseFromRequestWithClaims(req, request.OAuth2Extractor, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		} else {
			return authBackend.PublicKey, nil
		}
	})

	if err == nil && token.Valid {
		claims := token.Claims.(*models.CustomClaims)
		redisConn := redis.Pconnect(2)
		defer redisConn.Close()
		hashMap, reer := redisConn.HgetAll(claims.Username)
		fmt.Printf("Middle Error: %v %v\n", reer, hashMap)
		if password, ok := hashMap["Password"]; !ok || password != claims.Password {
			rw.WriteHeader(http.StatusUnauthorized)
		} else {
			context.Set(req, "Username", claims.Username)
			next(rw, req)
		}
	} else {
		rw.WriteHeader(http.StatusUnauthorized)
	}
}
