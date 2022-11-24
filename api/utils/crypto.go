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

func GenerateJWT(payload, secretKey string, expires time.Duration) (string, error) {
	var mySigningKey = []byte(secretKey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["sub"] = payload
	claims["exp"] = time.Now().Add(expires).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		log.Fatalf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func SignJWT(digest string, secretKey string) (jwt.MapClaims, error) {
	var mySigningKey = []byte(secretKey)

	token, err := jwt.Parse(digest, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("there was an error in parsing")
		}
		return mySigningKey, nil
	})

	if err != nil {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid claims")

}
