package jwtx

import (
	"github.com/golang-jwt/jwt/v4"
)

type (
	Claims struct {
		UserID int64
		jwt.StandardClaims
	}
)

func GenToken(secretKey string, iat, seconds, uid int64) (string, error) {
	claims := &Claims{
		UserID: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: iat + seconds,
		},
	}
	token := jwt.New(jwt.SigningMethodHS384)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func ParseToken(token, secretKey string) (*Claims, error) {
	c, err := jwt.ParseWithClaims(token, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if c != nil {
		if mapClaims, ok := c.Claims.(*Claims); ok && c.Valid {
			return mapClaims, err
		}
	}
	return nil, err
}

func GetUserId(token, secretKey string) (userID int64, err error) {
	claims, err := ParseToken(token, secretKey)
	if err != nil {
		return 0, err
	}
	userID = claims.UserID
	return
}
