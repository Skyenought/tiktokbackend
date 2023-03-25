package logic

import (
	"github.com/Skyenought/tiktokbackend/pkg/jwtx"
)

func GenToken(secret string, now int64, expire int64, userID int64) (string, error) {
	token, err := jwtx.GenToken(secret, now, expire, userID)
	if err != nil {
		return "", err
	}
	return token, nil
}
