package gorandom

import "sync"

// XORShift128Plus is a xorshift+ generator
type XORShift128Plus struct {
	state [2]uint64
	lock  sync.Mutex
}

// NewXORShift128Plus creates a XORShift128Plus generator
func NewXORShift128Plus(seed int64) *XORShift128Plus {
	xor := &XORShift128Plus{}
	xor.Seed(seed)
	return xor
}

// Seed seeds the generator
func (x *XORShift128Plus) Seed(seed int64) {
	x.lock.Lock()
	defer x.lock.Unlock()

	x.state = [2]uint64{uint64(seed), uint64(seed)}
}

// Int63 returns a pseudo-random int64
func (x *XORShift128Plus) Int63() int64 {
	x.lock.Lock()
	defer x.lock.Unlock()

	tmp := x.state[0]
	tmp2 := x.state[1]
	x.state[0] = tmp2
	tmp = tmp ^ tmp<<23
	tmp = tmp ^ tmp>>17
	tmp = tmp ^ tmp2 ^ (tmp2 >> 26)
	x.state[1] = tmp
	return int64(tmp + tmp2)
}
