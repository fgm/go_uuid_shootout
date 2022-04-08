package uuid_shootout_test

import (
	"testing"

	"github.com/fgm/uuid_shootout"
)

func BenchmarkAgextUuidGen1Crypto(b *testing.B) {
	count := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		count += len(uuid_shootout.AgextUuidGen1Crypto())
	}
	b.StopTimer()
	// Ensure no unused data elimination.
	uuid_shootout.UseInt(count)
}

func BenchmarkAgextUuidGen1Math(b *testing.B) {
	count := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		count += len(uuid_shootout.AgextUuidGen1Math())
	}
	b.StopTimer()
	// Ensure no unused data elimination.
	uuid_shootout.UseInt(count)
}
