package deque

const bootStrapSize = 8

var _ Deque = (*ring)(nil)

type ring struct {
	buf  []interface{}
	base int
	len  int
	size int
}

func NewRingBufferDeque() *ring {
	return newRingBufferDeque(bootStrapSize)
}

func newRingBufferDeque(size int) *ring {
	return &ring{
		size: size,
		buf:  make([]interface{}, size),
	}
}

func (r *ring) IsEmpty() bool {
	return r.len == 0
}

func (r *ring) pop(dir int) (interface{}, bool) {
	if r.IsEmpty() {
		return nil, false
	}
	if dir == Front {
		ret := r.buf[r.base]
		r.base = (r.base + 1) % r.size
		r.len--
		return ret, true
	}
	r.len--
	return r.buf[(r.base+r.len)%r.size], true
}

func (r *ring) push(dir int, val interface{}) {
	if r.len == r.size {
		nr := newRingBufferDeque(2 * r.size)
		for !r.IsEmpty() {
			val, _ := r.pop(Front)
			nr.push(Back, val)
		}
		*r = *nr
	}
	if dir == Front {
		if r.base == 0 {
			r.base = r.size - 1
		} else {
			r.base--
		}
		r.len++
		r.buf[r.base] = val
	} else {
		r.buf[(r.base+r.len)%r.size] = val
		r.len++
	}
}

func (r *ring) PushFront(val interface{}) {
	r.push(Front, val)
}

func (r *ring) PushBack(val interface{}) {
	r.push(Back, val)
}

func (r *ring) PopFront() (interface{}, bool) {
	return r.pop(Front)
}

func (r *ring) PopBack() (interface{}, bool) {
	return r.pop(Back)
}
