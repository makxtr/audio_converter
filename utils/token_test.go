package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenToken(t *testing.T) {
	assert.Len(t, GenToken(), 26)
}
