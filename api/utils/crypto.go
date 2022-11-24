package utils

import "golang.org/x/crypto/bcrypt"

func GenerateHash(payload string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(payload), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	return string(hash)
}

func CompareHash(hasher string, payload string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hasher), []byte(payload))

	if err == nil {
		return true
	}

	return false
}
