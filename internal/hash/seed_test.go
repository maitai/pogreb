package hash

import (
	"testing"

	"github.com/maitai/pogreb/internal/assert"
)

func TestRandSeed(t *testing.T) {
	_, err := RandSeed()
	assert.Nil(t, err)
}
