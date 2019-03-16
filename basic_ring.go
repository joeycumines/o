package o

// BasicRing contains the accounting data for a ring buffer or other
// data structure of arbitrary length. It uses three variables (insert
// index, length of buffer, ring capacity) to keep track of the
// state.
//
// The index wrap-around operation is implemented with modulo division.
type basicRing struct {
	cap, read, length uint
}

// NewBasic returns a BasicRing with the capacity given as an absolute
// number of elements.
func NewBasic(cap uint) RingAccountant {
	return &basicRing{cap: cap}
}

func (r *basicRing) Mask(val uint) uint {
	return val % r.cap
}

func (r *basicRing) Full() bool {
	return r.cap == r.length
}

func (r *basicRing) Empty() bool {
	return r.length == 0
}

func (r *basicRing) Push() (uint, error) {
	if r.Full() {
		return 0, ErrFull
	}
	l := r.length
	r.length++

	return r.Mask(r.read + l), nil
}

func (r *basicRing) Shift() (uint, error) {
	if r.Empty() {
		return 0, ErrEmpty
	}
	r.length--
	i := r.read
	r.read = r.Mask(r.read + 1)
	return i, nil
}

func (r *basicRing) Size() uint {
	return r.length
}