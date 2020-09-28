package generator

import (
	"math/rand"
	"strings"
	"time"
)

const alphaNum = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Generator random id generator
type Generator struct {
	len       int
	prefix    string
	separator string
}

// New create new random id generator
func New(len int, prefix string, separator string) (*Generator, error) {
	return &Generator{
		len:       len,
		prefix:    prefix,
		separator: separator,
	}, nil
}

// Get generate a new random id
func (id *Generator) Get() string {
	rand.Seed(time.Now().UnixNano())
	rnd := make([]byte, id.len)
	rand.Read(rnd)
	value := make([]string, id.len)
	charsLength := byte(len(alphaNum))

	for i := 0; i < id.len; i++ {
		value[i] = alphaNum[rnd[i]%charsLength : rnd[i]%charsLength+1]
	}

	return id.prefix + id.separator + strings.Join(value, "")
}
