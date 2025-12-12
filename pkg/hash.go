package pkg

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/argon2"
)

type Argon2Params struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

var secretKey = []byte("my_secret_key")

func HashPassword(password string, p *Argon2Params) (string, error) {
	salt := make([]byte, p.SaltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, p.Iterations, p.Memory, p.Parallelism, p.KeyLength)

	// Кодируем всё в base64 для хранения в SQL
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// Стандартный формат хранения
	encoded := fmt.Sprintf("$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s",
		p.Memory, p.Iterations, p.Parallelism,
		b64Salt, b64Hash,
	)

	return encoded, nil
}

func Verify(password, encodedHash string) (bool, error) {
	var p Argon2Params
	var saltBase64, hashBase64 string

	_, err := fmt.Sscanf(encodedHash, "$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s",
		&p.Memory, &p.Iterations, &p.Parallelism,
		&saltBase64, &hashBase64,
	)
	if err != nil {
		return false, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(saltBase64)
	if err != nil {
		return false, err
	}

	hash, err := base64.RawStdEncoding.DecodeString(hashBase64)
	if err != nil {
		return false, err
	}

	newHash := argon2.IDKey([]byte(password), salt, p.Iterations, p.Memory, p.Parallelism, uint32(len(hash)))

	if subtle.ConstantTimeCompare(newHash, hash) == 1 {
		return true, nil
	}
	return false, nil
}

func GenerateToken(email string) (map[string]string, error) {
	tokens := make(map[string]string, 2)

	acess_claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Minute * 1).Unix(),
		"role":  "acess_token",
	}

	refresh_claims := jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Minute * 1).Unix(),
		"role":  "refresh_token",
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
