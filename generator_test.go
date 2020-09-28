package generator_test

import (
	"testing"

	"github.com/hadv/generator"
	"github.com/stretchr/testify/assert"
)

func TestGenerator(t *testing.T) {
	gen := generator.New(20, "tnx", "_")
	id, _ := gen.Get()
	assert.Equal(t, 24, len(id))
	assert.Equal(t, "tnx_", id[0:4])
}

func benchmarkGenerator(n int, b *testing.B) {
	gen := generator.New(20, "tnx", "_")
	for i := 0; i < n; i++ {
		_, _ = gen.Get()
	}
}

func BenchmarkGenerator10(b *testing.B) {
	benchmarkGenerator(10, b)
}

func BenchmarkGenerator100(b *testing.B) {
	benchmarkGenerator(100, b)
}

func BenchmarkGenerator1000(b *testing.B) {
	benchmarkGenerator(1000, b)
}
