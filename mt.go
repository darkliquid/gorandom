package gorandom

import "sync"

// MersenneTwister generates pseudo-random ints (MT19937)
type MersenneTwister struct {
	mt    [624]int
	index int
	lock  sync.Mutex
}

const mersenne1 int = 1812433253
const mersenne2 int = 2636928640
const mersenne3 int = 4022730752
const mersenne4 int = 0x80000000
const mersenne5 int = 0x7fffffff
const mersenne6 int = 0x9908b0df

// NewMersenneTwister returns a new MersenneTwister
func NewMersenneTwister(seed int64) (twister *MersenneTwister) {
	twister = &MersenneTwister{}
	twister.Seed(seed)

	return
}

// Seed seeds the generator with the given value
func (mt *MersenneTwister) Seed(seed int64) {
	mt.lock.Lock()
	defer mt.lock.Unlock()

	mt.mt[0] = int(seed)

	for i := 1; i < 624; i++ {
		mt.mt[i] = int(mersenne1*(mt.mt[i-1]^mt.mt[i-1]>>30) + i)
	}
}

// Int63 returns a pseudo-random int64
func (mt *MersenneTwister) Int63() int64 {
	// We need to lock to make this concurrency-safe
	mt.lock.Lock()
	defer mt.lock.Unlock()

	if mt.index == 0 {
		mt.generate()
	}

	num := mt.mt[mt.index]
	num = num ^ num>>11
	num = num ^ num<<7&mersenne2
	num = num ^ num<<15&mersenne3
	num = num ^ num>>18

	mt.index = (mt.index + 1) % 624

	return int64(num)
}

func (mt *MersenneTwister) generate() {
	for i := range mt.mt {
		num := int((mt.mt[i] & mersenne4) + mt.mt[(i+1)%624]&mersenne5)
		mt.mt[i] = mt.mt[(i+397)%624] ^ num>>1
		if num%2 != 0 {
			mt.mt[i] = mt.mt[i] ^ mersenne6
		}
	}
}
