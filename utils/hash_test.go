package utils

import (
	"bytes"
	"crypto/rand"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashPass(t *testing.T) {
	salt := make([]byte, 8)
	if _, err := rand.Read(salt); err != nil {
		t.Fatalf("Ошибка генерации соли: %v", err)
	}

	password := "password123"
	hashedPass := hashPass(salt, password)

	t.Logf("Salt: %x", salt)
	t.Logf("HashedPass: %x", hashedPass)
	t.Logf("HashedPass: %x", hashPass(salt, "asdsad"))

	assert.Equal(t, len(hashedPass), 40)
	assert.True(t, bytes.Equal(hashedPass[:8], salt))

	assert.Equal(t, hashPass(salt, "one")[:8], hashPass(salt, "two")[:8])
	assert.NotEqual(t, hashPass(salt, "one"), hashPass(salt, "two"))

	t.Logf("checkPass: %t", checkPass(hashedPass, "111"))

}

func TestCheckPass(t *testing.T) {
	salt := make([]byte, 8)
	if _, err := rand.Read(salt); err != nil {
		t.Fatalf("Ошибка генерации соли: %v", err)
	}

	password := "password123"
	hashedPass := hashPass(salt, password)

	// Отладочная информация
	t.Logf("Salt: %x", salt)
	t.Logf("HashedPass: %x", hashedPass)

	assert.True(t, checkPass(hashedPass, password))
	assert.False(t, checkPass(hashedPass, "wrongpassword"))
}
