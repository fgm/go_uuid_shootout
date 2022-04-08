package uuid_shootout_test

import (
	"testing"

	"github.com/google/uuid"

	"github.com/fgm/uuid_shootout"
)

func BenchmarkGoogleUuidGen1Raw(b *testing.B) {
	count := 0
	uuid.DisableRandPool()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		count += len(uuid_shootout.GoogleUuidGen1())
	}
	b.StopTimer()
	// Ensure no unused data elimination.
	uuid_shootout.UseInt(count)
}

func BenchmarkGoogleUuidGen1Pool(b *testing.B) {
	count := 0
	uuid.EnableRandPool()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		count += len(uuid_shootout.GoogleUuidGen1())
	}
	b.StopTimer()
	// Ensure no unused data elimination.
	uuid_shootout.UseInt(count)
}
