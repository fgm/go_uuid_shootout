package uuid_shootout

import (
	crand "crypto/rand"
	mrand "math/rand"
	"time"

	"github.com/oklog/ulid"
)

var (
	OklogUlidCryptoEntropy = crand.Reader
	OklogUlidMathEntropy   = ulid.Monotonic(mrand.New(mrand.NewSource(time.Now().Unix())), 0)
)

func OklogUlidGen1Math() string {
	id, err := ulid.New(ulid.Timestamp(time.Now()), OklogUlidMathEntropy)
	if err != nil {
		panic(err)
	}
	return id.String()
}

func OklogUlidGen1Crypto() string {
	id, err := ulid.New(ulid.Timestamp(time.Now()), OklogUlidCryptoEntropy)
	if err != nil {
		panic(err)
	}
	return id.String()
}
