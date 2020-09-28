package generator

import (
	"crypto/rand"
	"math/big"
)

const alphaNum = "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ"

// Generator random id generator
type Generator struct {
	len       int
	prefix    string
	separator string
}

// New create new random id generator
func New(len int, prefix string, separator string) *Generator {
	return &Generator{
		len:       len,
		prefix:    prefix,
		separator: separator,
	}
}

// Get generate a new random id
func (id *Generator) Get() (string, error) {
	ret := make([]byte, 0)

	for i := 0; i < id.len; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(alphaNum))))
		if err != nil {
			return "", err
		}
		ret = append(ret, alphaNum[num.Int64()])
	}

	return id.prefix + id.separator + string(ret), nil
}
