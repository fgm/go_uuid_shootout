package uuid_shootout

import (
	"github.com/rs/xid"
)

func RsXidGen1() string {
	return xid.New().String()
}
