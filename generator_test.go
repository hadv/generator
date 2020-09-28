package generator_test

import (
	"testing"

	"github.com/hadv/generator"
	"github.com/stretchr/testify/assert"
)

func TestGenerator(t *testing.T) {
	gen, _ := generator.New(20, "tnx", "_")
	assert.Equal(t, 24, len(gen.Get()))
	assert.Equal(t, "tnx_", gen.Get()[0:4])
}
