package uuid_shootout

import (
	"github.com/satori/go.uuid"
)

func SatoriGouuidGen1() string {
	return uuid.NewV4().String()
}
