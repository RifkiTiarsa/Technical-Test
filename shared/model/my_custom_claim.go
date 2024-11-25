package model

import (
	"github.com/golang-jwt/jwt/v5"
)

type MyCustomClaims struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}
