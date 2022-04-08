package uuid_shootout

import (
	"github.com/google/uuid"
)

func GoogleUuidGen1() string {
	id, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}
	return id.String()
}
