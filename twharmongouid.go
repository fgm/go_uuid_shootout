package uuid_shootout

import (
	"github.com/twharmon/gouid"
)

// UUIDs are 128 bits, this generator returns strings of ASCII chars in bytes.
const size = 128 / 8

func TwharmonGouidGen1LCA() string {
	return gouid.String(size, gouid.LowerCaseAlpha)
}

func TwharmonGouidGen1LCAN() string {
	return gouid.String(size, gouid.LowerCaseAlphaNum)
}

func TwharmonGouidGen1MCA() string {
	return gouid.String(size, gouid.MixedCaseAlpha)
}

func TwharmonGouidGen1MCAN() string {
	return gouid.String(size, gouid.MixedCaseAlphaNum)
}

func TwharmonGouidGen1UCA() string {
	return gouid.String(size, gouid.UpperCaseAlpha)
}

func TwharmonGouidGen1UCAN() string {
	return gouid.String(size, gouid.UpperCaseAlphaNum)
}
