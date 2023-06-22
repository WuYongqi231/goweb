package model

import "github.com/dgrijalva/jwt-go"

type User struct {
	Id       int
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
type Question struct {
}
type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
