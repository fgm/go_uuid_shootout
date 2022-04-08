package uuid_shootout

import (
	"github.com/jakehl/goid"
)

func JakehlGoidGen1() string {
	return goid.NewV4UUID().String()
}
