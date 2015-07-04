package gorandom

import "sync"

// MersenneTwister64 generates pseudo-random uint64s (MT19937_64)
type MersenneTwister64 struct {
	mt    [312]uint64
	index uint64
	lock  sync.Mutex
}

const nn uint64 = 312
const mm uint64 = 156
const matrixA uint64 = 0xB5026F5AA96619E9
const um uint64 = 0xFFFFFFFF80000000
const lm uint64 = 0x7FFFFFFF

const m64a uint64 = 6364136223846793005
const m64b uint64 = 0x5555555555555555
const m64c uint64 = 0x71D67FFFEDA60000
const m64d uint64 = 0xFFF7EEE000000000
const m64seed int64 = 5489

// NewMersenneTwister64 returns a new MersenneTwister64
func NewMersenneTwister64(seed int64) (twister *MersenneTwister64) {
	twister = &MersenneTwister64{}
	twister.Seed(seed)

	return
}

// Seed seeds the generator
func (mt *MersenneTwister64) Seed(seed int64) {
	mt.lock.Lock()
	defer mt.lock.Unlock()

	mt.mt[0] = uint64(seed)

	for i := uint64(1); i < nn; i++ {
		mt.mt[i] = m64a*(mt.mt[i-1]^mt.mt[i-1]>>62) + i
	}
}

// Int63 returns a pseudo-random int64
func (mt *MersenneTwister64) Int63() int64 {
	// We need to lock to make this concurrency-safe
	mt.lock.Lock()
	defer mt.lock.Unlock()

	if mt.index >= nn {
		mt.generate()
	}

	num := mt.mt[mt.index]
	mt.index++
	num = num ^ (num>>29)&m64b
	num = num ^ (num>>17)&m64c
	num = num ^ (num>>37)&m64d
	num = num ^ (num >> 43)

	return int64(num)
}

func (mt *MersenneTwister64) generate() {
	var i uint64
	var x uint64
	ma := []uint64{0, matrixA}

	if mt.index == nn+1 {
		mt.Seed(m64seed)
	}

	for i = 0; i < nn-mm; i++ {
		x = (mt.mt[i] & um) | (mt.mt[i+1] & lm)
		mt.mt[i] = mt.mt[i+mm] ^ (x >> 1) ^ ma[int(x&uint64(1))]
	}
	for ; i < nn-1; i++ {
		x = (mt.mt[i] & um) | (mt.mt[i+1] & lm)
		mt.mt[i] = mt.mt[i-(nn-mm)] ^ (x >> 1) ^ ma[int(x&uint64(1))]
	}

	x = (mt.mt[nn-1] & um) | (mt.mt[0] & lm)
	mt.mt[nn-1] = mt.mt[mm-1] ^ (x >> 1) ^ ma[int(x&uint64(1))]
	mt.index = 0
}
