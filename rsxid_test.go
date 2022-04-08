package uuid_shootout_test

import (
	"testing"

	"github.com/fgm/uuid_shootout"
)

func BenchmarkRxXidGen1(b *testing.B) {
	count := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		count += len(uuid_shootout.RsXidGen1())
	}
	b.StopTimer()
	// Ensure no unused data elimination.
	uuid_shootout.UseInt(count)
}
