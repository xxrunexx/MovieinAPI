package middleware

import (
	"movie-api/config"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateToken(userId uint, name string) (string, error) {
	claims := jwt.MapClaims{}

	claims["userId"] = userId
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 6).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.JWT_KEY))
}

// func ExtractClaim(e echo.Context) (claims map[string]interface{}) {
// 	user := e.Get("user").(*jwt.Token)
// 	if user.Valid {
// 		claims = user.Claims.(jwt.MapClaims)
// 	}
// 	return
// }
