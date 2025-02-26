package utils

import (
	"bytes"
	"golang.org/x/crypto/argon2"
)

func hashPass(salt []byte, plainPassword string) []byte {
	// Делаем копию salt, чтобы избежать изменения исходного слайса
	saltCopy := make([]byte, len(salt))
	copy(saltCopy, salt)

	hashedPass := argon2.IDKey([]byte(plainPassword), saltCopy, 1, 64*1024, 4, 32)
	return append(saltCopy, hashedPass...) // Теперь append не изменит оригинальный salt
}

func checkPass(passHash []byte, plainPassword string) bool {
	salt := passHash[0:8]
	userPassHash := hashPass(salt, plainPassword)

	return bytes.Equal(userPassHash, passHash)
}
