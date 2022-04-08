package uuid_shootout

import (
	"github.com/muyo/sno"
)

func MuyoSnoGen1() string {
	return sno.New(0).String()
}
