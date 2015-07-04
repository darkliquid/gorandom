package gorandom

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadSeeds(t *testing.T) {
	buffer := bytes.NewBuffer([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	seeds, err := ReadSeeds(buffer)
	assert.NoError(t, err)
	assert.Equal(t, []int64{283686952306183, 2057}, seeds)
}

func TestReadSeed(t *testing.T) {
	buffer := bytes.NewBuffer([]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	seed, err := ReadSeed(buffer)
	assert.NoError(t, err)
	assert.Equal(t, 283686952308238, seed)
}
