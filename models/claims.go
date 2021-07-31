package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	UserID   int
	UserName string
	jwt.StandardClaims
}
