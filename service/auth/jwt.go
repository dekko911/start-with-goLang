package auth

import (
	"strconv"
	"time"

	"github.com/dekko911/start-with-goLang/config"
	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(secret []byte, userID int) (string, error) {
	loc, err := time.LoadLocation("Asia/Kuala_Lumpur")
	if err != nil {
		return "", err
	}

	expiration := time.Second * time.Duration(config.Env.JWTExpirationInSeconds)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":     strconv.Itoa(userID),
		"expired_at": time.Now().In(loc).Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
