package utils

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func GenerateHash(payload string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(payload), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	return string(hash)
}

func CompareHash(hasher string, payload string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hasher), []byte(payload))

	return err == nil
}

func GenerateJWT(sub, secretKey string, expires time.Duration) (string, error) {
	claims := jwt.MapClaims{}
	claims["sub"] = sub
	claims["exp"] = time.Now().Add(expires).Unix()

	to := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := to.SignedString([]byte(secretKey))

	if err != nil {
		log.Fatalf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return token, nil
}

func SignJWT(digest, secretKey string) (jwt.RegisteredClaims, error) {

	claims := jwt.RegisteredClaims{}

	parser, _ := jwt.ParseWithClaims(digest, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if parser.Valid {
		return claims, nil
	}

	return claims, errors.New("token expired")

}
