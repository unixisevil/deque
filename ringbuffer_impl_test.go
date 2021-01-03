package deque

import (
	"testing"
)

func TestRingBufferEmpty(t *testing.T) {
	rbd := NewRingBufferDeque()

	if expect, actual := rbd.IsEmpty(), true; expect != actual {
		t.Errorf("expect empty deque, got nonempty")
	}
	_, ok1 := rbd.PopBack()
	_, ok2 := rbd.PopFront()
	if ok1 == true || ok2 == true {
		t.Errorf("pop something from empty deque")
	}
}

func TestRingBufferPushPopBack(t *testing.T) {
	rbd := NewRingBufferDeque()

	for i := 10; i >= 1; i-- {
		rbd.PushBack(i)
	}
	for i := 1; i <= 10; i++ {
		val, ok := rbd.PopBack()
		if !ok {
			t.Errorf("return nothing")
		}
		t.Logf("got val %v", val)
		if expect, actual := val.(int), i; expect != actual {
			t.Errorf("expect %d, got %d", expect, actual)
		}
	}
}

func TestRingBufferPushPopFront(t *testing.T) {
	rbd := NewRingBufferDeque()

	for i := 10; i >= 1; i-- {
		rbd.PushFront(i)
	}
	for i := 1; i <= 10; i++ {
		val, ok := rbd.PopFront()
		if !ok {
			t.Errorf("return nothing")
		}
		t.Logf("got val %v", val)
		if expect, actual := val.(int), i; expect != actual {
			t.Errorf("expect %d, got %d", expect, actual)
		}
	}
}

func TestRingBufferPushBackPopFront(t *testing.T) {
	rbd := NewRingBufferDeque()

	for i := 1; i <= 10; i++ {
		rbd.PushBack(i)
	}

	for i := 1; i <= 10; i++ {
		val, ok := rbd.PopFront()
		if !ok {
			t.Errorf("return nothing")
		}
		t.Logf("got val = %v", val)
		if expect, actual := val.(int), i; expect != actual {
			t.Errorf("expect %d, got %d", expect, actual)
		}
	}
}
