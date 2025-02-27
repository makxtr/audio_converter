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
	hashedPass := HashPass(salt, password)

	assert.Equal(t, len(hashedPass), 40)
	assert.True(t, bytes.Equal(hashedPass[:8], salt))

	assert.Equal(t, HashPass(salt, "one")[:8], HashPass(salt, "two")[:8])
	assert.NotEqual(t, HashPass(salt, "one"), HashPass(salt, "two"))

	t.Logf("checkPass: %t", CheckPass(hashedPass, "111"))
}

func TestCheckPass(t *testing.T) {
	salt := make([]byte, 8)
	if _, err := rand.Read(salt); err != nil {
		t.Fatalf("Ошибка генерации соли: %v", err)
	}

	password := "password123"
	hashedPass := HashPass(salt, password)

	assert.True(t, CheckPass(hashedPass, password))
	assert.False(t, CheckPass(hashedPass, "wrongpassword"))
}
