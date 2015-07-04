package gorandom

import (
	"bytes"
	"io"
	"math/big"
)

// ReadSeed consumes a reader to generate a single seed value
func ReadSeed(r io.Reader) (n int64, err error) {
	var seeds []int64
	if seeds, err = ReadSeeds(r); err != nil {
		return
	}

	// Loop through and XOR seeds until all processed
	for _, x := range seeds {
		n ^= x
	}

	return
}

// ReadSeeds consumes a reader, generating seeds as it goes
func ReadSeeds(r io.Reader) (seeds []int64, err error) {
	var b bytes.Buffer
	for {
		if _, err = io.CopyN(&b, r, 8); err != nil {
			if err != io.EOF {
				return
			}
		}

		// Convert our bytes to an int64
		var num big.Int
		num.SetBytes(b.Next(8))
		seeds = append(seeds, num.Int64())

		if err == io.EOF {
			err = nil
			break
		}
	}

	return seeds, nil
}
