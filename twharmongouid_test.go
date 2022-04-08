package uuid_shootout_test

import (
	"testing"

	"github.com/fgm/uuid_shootout"
)

func twharmonGouidBenchStrategy(b *testing.B, strategy func() string) {
	count := 0
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		count += len(strategy())
	}
	b.StopTimer()
	// Ensure no unused data elimination.
	uuid_shootout.UseInt(count)
}

func BenchmarkTwharmonGouidLCA(b *testing.B) {
	twharmonGouidBenchStrategy(b, uuid_shootout.TwharmonGouidGen1LCA)
}

func BenchmarkTwharmonGouidLCAN(b *testing.B) {
	twharmonGouidBenchStrategy(b, uuid_shootout.TwharmonGouidGen1LCAN)
}

func BenchmarkTwharmonGouidMCA(b *testing.B) {
	twharmonGouidBenchStrategy(b, uuid_shootout.TwharmonGouidGen1MCA)
}

func BenchmarkTwharmonGouidMCAN(b *testing.B) {
	twharmonGouidBenchStrategy(b, uuid_shootout.TwharmonGouidGen1MCAN)
}

func BenchmarkTwharmonGouidUCA(b *testing.B) {
	twharmonGouidBenchStrategy(b, uuid_shootout.TwharmonGouidGen1UCA)
}

func BenchmarkTwharmonGouidUCAN(b *testing.B) {
	twharmonGouidBenchStrategy(b, uuid_shootout.TwharmonGouidGen1UCAN)
}
