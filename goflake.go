package uuid_shootout

import (
	"github.com/hart87/GoFlake/generator"
)

func GoflakeGen1() string {
	return generator.GenerateIdentifier()
}
