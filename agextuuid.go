package uuid_shootout

import (
	"github.com/agext/uuid"
)

func AgextUuidGen1Crypto() string {
	return uuid.NewCrypto().String()
}

func AgextUuidGen1Math() string {
	return uuid.New().String()
}
