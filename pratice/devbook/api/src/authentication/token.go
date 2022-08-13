package authentication

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateToken(userID uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["user_id"] = userID
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte("Secret"))
}
