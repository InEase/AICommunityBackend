package Authorization

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("a_secret_key_yes")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user Users) (string, error) {
	exprireTime := time.Now().Add(3 * time.Hour)
	claim := Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exprireTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "DoneServer",
			Subject:   "DoneToken",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	return token, claims, err
}
