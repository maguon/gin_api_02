package request

import (
	jwt "github.com/golang-jwt/jwt/v4"
)

// Custom claims structure
type CustomClaims struct {
	BaseClaims
	BufferTime int64
	jwt.StandardClaims
}

type AdminCustomClaims struct {
	AdminBaseClaims
	BufferTime int64
	jwt.StandardClaims
}
type BaseClaims struct {
	ID    int64
	Phone string
}

type AdminBaseClaims struct {
	ID       int64
	Username string
	Type     int16
}
