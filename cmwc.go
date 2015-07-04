package gorandom

import "sync"

// CMWC is a xorshift+ generator
type CMWC struct {
	state [4096]uint32
	c     uint32
	lock  sync.Mutex
}

const cmwcA uint32 = 18705
const cmwcC uint32 = 362
const cmwcF uint64 = 0xffffffff

// NewCMWC creates a CMWC generator
func NewCMWC(seed int64) *CMWC {
	gen := &CMWC{}
	gen.Seed(seed)
	return gen
}

// Seed seeds the generator
func (gen *CMWC) Seed(seed int64) {
	gen.lock.Lock()
	defer gen.lock.Unlock()

	gen.c = cmwcC
	s := uint32(seed)
	for i := 0; i < 4096; i++ {
		s = s ^ s<<uint(i%63)
		gen.state[i] = s ^ uint32(seed)
	}
}

// Int63 returns a pseudo-random int64
func (gen *CMWC) Int63() int64 {
	gen.lock.Lock()
	defer gen.lock.Unlock()

	tmp := uint32(4095)
	tmp = (tmp + 1) & 4095
	tmp2 := uint64(cmwcA * gen.state[tmp])
	gen.c = uint32((tmp2 + uint64(gen.c)) >> 32)
	gen.state[tmp] = uint32(cmwcF - tmp2)

	return int64(gen.state[tmp])
}
