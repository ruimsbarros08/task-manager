package services

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

func hashAndSalt(pwd string) string {
	saltedBytes := []byte(pwd)
	hash, err := bcrypt.GenerateFromPassword(saltedBytes, bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

func compare(hash string, pwd string) error {
	incoming := []byte(pwd)
	existing := []byte(hash)
	return bcrypt.CompareHashAndPassword(existing, incoming)
}