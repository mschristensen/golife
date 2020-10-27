package life_test

import (
	"testing"

	"github.com/mschristensen/golife/pkg/life"
)

func benchmarkUpdate(sideLength int, b *testing.B) {
	world := life.NewWorld(sideLength, sideLength)
	for n := 0; n < b.N; n++ {
		world.Update()
	}
}

func BenchmarkUpdate10(b *testing.B)    { benchmarkUpdate(10, b) }
func BenchmarkUpdate100(b *testing.B)   { benchmarkUpdate(100, b) }
func BenchmarkUpdate1000(b *testing.B)  { benchmarkUpdate(1000, b) }
func BenchmarkUpdate10000(b *testing.B) { benchmarkUpdate(10000, b) }
