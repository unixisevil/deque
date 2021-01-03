package deque

const (
	Front = iota
	Back
)

type Deque interface {
	PushFront(interface{})
	PushBack(interface{})
	PopFront() (interface{}, bool)
	PopBack() (interface{}, bool)
	IsEmpty() bool
}
