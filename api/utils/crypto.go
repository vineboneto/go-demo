package utils

import "golang.org/x/crypto/bcrypt"

func GenerateHash(payload string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(payload), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	return string(hash)
}
