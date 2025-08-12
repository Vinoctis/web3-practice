package utils 

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
	"fmt"
)

var secretKey    = []bytes("web3-practice")

const signMethod = jwt.SigningMethodHS256

func GenerateToken(userID int, userName string) (string, error) {
	claims := jwt.MapClaims{
		"user_id"   : userID,
		"exp"       : time.Now().Add(time.Hour * 24).Unix(),
		"user_name" : userName,
		"iss"       : "web3-blog",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token)(any, error) {
		if token.Method != signMethod {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	}) 
}