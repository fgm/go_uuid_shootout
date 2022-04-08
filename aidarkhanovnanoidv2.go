package uuid_shootout

import (
	"github.com/aidarkhanov/nanoid/v2"
)

func AidarkhanovNanoidv2Gen1() string {
	s, err := nanoid.New()
	if err != nil {
		panic(err)
	}
	return s
}
