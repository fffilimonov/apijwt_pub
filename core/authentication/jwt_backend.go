package authentication

import (
	"../../services/models"
	"../../settings"
	"../redis"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

type JWTAuthenticationBackend struct {
	privateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

var authBackendInstance *JWTAuthenticationBackend = nil

func InitJWTAuthenticationBackend() *JWTAuthenticationBackend {
	if authBackendInstance == nil {
		authBackendInstance = &JWTAuthenticationBackend{
			privateKey: getPrivateKey(),
			PublicKey:  getPublicKey(),
		}
	}

	return authBackendInstance
}

func (backend *JWTAuthenticationBackend) GenerateToken(user *models.User) (string, error) {
	redisConn := redis.Pconnect(2)
	defer redisConn.Close()
	hashMap, reer := redisConn.HgetAll(user.Username)
	fmt.Printf("Error: %v %v\n", reer, hashMap)

	if password, ok := hashMap["Password"]; !ok || password != user.GetPassword() {
		return "nu", errors.New("nu")
	}

	if active, ok := hashMap["Active"]; !ok || active != "1" {
		return "na", errors.New("na")
	}

	fmt.Printf("Username JWT: %v\n", user.Username)
	fmt.Printf("Password JWT: %v\n", user.Password)

	token := jwt.New(jwt.SigningMethodRS512)

	claimUser := user
	claimUser.Password = user.GetPassword()

	token.Claims = &models.CustomClaims{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(settings.Get().JWTExpirationDelta)).Unix(),
		},
		claimUser,
	}

	tokenString, err := token.SignedString(backend.privateKey)
	if err != nil {
		panic(err)
		return "", err
	}
	return tokenString, nil
}

func getPrivateKey() *rsa.PrivateKey {
	data, _ := pem.Decode([]byte(private_key))

	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)

	if err != nil {
		panic(err)
	}

	return privateKeyImported
}

func getPublicKey() *rsa.PublicKey {
	data, _ := pem.Decode([]byte(public_key))

	publicKeyImported, err := x509.ParsePKIXPublicKey(data.Bytes)

	if err != nil {
		panic(err)
	}

	rsaPub, ok := publicKeyImported.(*rsa.PublicKey)

	if !ok {
		panic(err)
	}

	return rsaPub
}
