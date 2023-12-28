package helper

import (
	"net/http"
	"os"
	"peken-be/models/domain"
	"peken-be/models/errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user *domain.User) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS512,
		jwt.MapClaims{
			"userId": user.Id,
			"exp":    time.Now().Add(time.Hour * 24 * 30).Unix(),
		},
	)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	return tokenString, err
}

func DecodeToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.NewLudesError(http.StatusUnauthorized, "Invalid token")
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.NewLudesError(http.StatusUnauthorized, "Invalid token")
	}
	return claims, nil
}
