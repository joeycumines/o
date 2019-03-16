package o

type maskRing struct {
	cap, read, write uint
}

// NewPowerOfTwo returns a ring accounting structure that is optimized
// for rings with capacities of 2^n.
func NewPowerOfTwo(n uint) Ring {
	return &maskRing{cap: 1 << n}
}

func (r *maskRing) mask(val uint) uint {
	return val & (r.cap - 1)
}

func (r *maskRing) start() uint {
	return r.mask(r.read)
}

func (r *maskRing) capacity() uint {
	return r.cap
}

func (r *maskRing) Full() bool {
	return r.Size() == r.cap
}

func (r *maskRing) Empty() bool {
	return r.read == r.write
}

func (r *maskRing) Push() (uint, error) {
	if r.Full() {
		return 0, ErrFull
	}
	i := r.write
	r.write++

	return r.mask(i), nil
}

func (r *maskRing) Shift() (uint, error) {
	if r.Empty() {
		return 0, ErrEmpty
	}
	i := r.read
	r.read++
	return i, nil
}

func (r *maskRing) Size() uint {
	return r.write - r.read
}