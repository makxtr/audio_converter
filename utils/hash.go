package utils

import (
	"bytes"
	"crypto/rand"

	"golang.org/x/crypto/argon2"
)

const saltLen = 8

func GetSalt() []byte {
	salt := make([]byte, saltLen)
	if _, err := rand.Read(salt); err != nil {
		return nil
	}
	return salt
}

func HashPass(salt []byte, plainPassword string) []byte {
	saltCopy := make([]byte, len(salt))
	copy(saltCopy, salt)

	hashedPass := argon2.IDKey([]byte(plainPassword), saltCopy, 1, 64*1024, 4, 32)
	return append(saltCopy, hashedPass...)
}

func CheckPass(passHash []byte, plainPassword string) bool {
	salt := passHash[0:saltLen]
	userPassHash := HashPass(salt, plainPassword)

	return bytes.Equal(userPassHash, passHash)
}
