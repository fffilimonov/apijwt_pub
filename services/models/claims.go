package models

import (
	jwt "github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	*jwt.StandardClaims
	*User
}
