package core

import "github.com/golang-jwt/jwt"

type JWTClaim struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
