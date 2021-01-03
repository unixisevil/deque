package deque

var _ Deque = (*deque)(nil)

type deque struct {
	links [2]*deque
	value interface{}
}

func NewLinkListDeque() *deque {
	d := &deque{}
	d.links[Front] = d
	d.links[Back] = d

	return d
}

func (d *deque) push(dir int, val interface{}) {
	e := &deque{}
	e.links[dir] = d.links[dir]
	e.links[dir^1] = d
	e.value = val

	d.links[dir] = e
	e.links[dir].links[dir^1] = e
}

func (d *deque) pop(dir int) (interface{}, bool) {
	e := d.links[dir]
	if e == d {
		return nil, false
	}
	d.links[dir] = e.links[dir]
	e.links[dir].links[dir^1] = d

	return e.value, true
}

func (d *deque) PushFront(val interface{}) {
	d.push(Front, val)
}

func (d *deque) PushBack(val interface{}) {
	d.push(Back, val)
}

func (d *deque) PopFront() (interface{}, bool) {
	return d.pop(Front)
}

func (d *deque) PopBack() (interface{}, bool) {
	return d.pop(Back)
}

func (d *deque) IsEmpty() bool {
	return d.links[Front] == d
}

func (d *deque) Destroy() {
	for !d.IsEmpty() {
		d.pop(Front)
	}
}
