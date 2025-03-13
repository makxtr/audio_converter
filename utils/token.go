package utils

import (
	"github.com/oklog/ulid"
	"math/rand"
	"time"
)

func GenToken() string {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))

	return ulid.MustNew(ulid.Timestamp(time.Now()), entropy).String()
}
