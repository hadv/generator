package generator_test

import (
	"fmt"
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

func TestGeneratorPrintOut(t *testing.T) {
	gen := generator.New(20, "tnx", "_")
	for i := 0; i < 1000; i++ {
		id, _ := gen.Get()
		fmt.Println(id)
	}
}
