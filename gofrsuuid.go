package uuid_shootout

import (
	"github.com/gofrs/uuid"
)

func GofrsUuidGen1() string {
	id, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return id.String()
}
