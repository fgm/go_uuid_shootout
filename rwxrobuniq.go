package uuid_shootout

import (
	"gitlab.com/rwxrob/uniq"
)

func RwxrobUniqGen1() string {
	return uniq.UUID()
}
