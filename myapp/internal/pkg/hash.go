package pkg

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("my_secret_key")

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ComparePassword(hash string, password string) error {
	bhash := []byte(hash)
	bpassword := []byte(password)
	return bcrypt.CompareHashAndPassword(bhash, bpassword)
}

func GenerateToken(user_login string) (map[string]string, error) {
	tokens := make(map[string]string, 2)

	acess_claims := jwt.MapClaims{
		"user_login": user_login,
		"exp":        time.Now().Add(time.Hour * 1).Unix(),
		"role":       "acess_token",
	}

	refresh_claims := jwt.MapClaims{
		"user_login": user_login,
		"exp":        time.Now().Add(time.Hour * 1).Unix(),
		"role":       "refresh_token",
	}

	acess_token := jwt.NewWithClaims(jwt.SigningMethodHS256, acess_claims)

	refresh_token := jwt.NewWithClaims(jwt.SigningMethodHS256, refresh_claims)

	acess_token_str, err := acess_token.SignedString(secretKey)
	if err != nil {
		return nil, nil
	}

	refresh_token_str, err := refresh_token.SignedString(secretKey)
	if err != nil {
		return nil, nil
	}

	tokens["acess_token"] = acess_token_str
	tokens["refresh_token"] = refresh_token_str

	return tokens, nil
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
}
