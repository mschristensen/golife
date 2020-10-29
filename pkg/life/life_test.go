package life_test

import (
	"fmt"
	"testing"

	"github.com/mschristensen/golife/pkg/life"
)

type benchmark struct {
	numGoroutines int
	height        int
}

var benchmarks = []benchmark{
	{1, 128}, {1, 1024}, {1, 8192},
	{2, 128}, {2, 1024}, {2, 8192},
	{4, 128}, {4, 1024}, {4, 8192},
	{8, 128}, {8, 1024}, {8, 8192},
	{16, 128}, {16, 1024}, {16, 8192},
}

func benchmarkUpdate(partitions, height int, b *testing.B) {
	world := life.NewWorld(height, height)
	for n := 0; n < b.N; n++ {
		world.Update(partitions)
	}
}

func BenchmarkUpdate(b *testing.B) {
	for _, bb := range benchmarks {
		b.Run(fmt.Sprintf("benchmark_%d_%d", bb.numGoroutines, bb.height), func(b *testing.B) {
			benchmarkUpdate(bb.numGoroutines, bb.height, b)
		})
	}
}
