package deque

import (
	"testing"
)

func TestLinkListEmpty(t *testing.T) {
	lld := NewLinkListDeque()

	if expect, actual := lld.IsEmpty(), true; expect != actual {
		t.Errorf("expect empty deque, got nonempty")
	}
	_, ok1 := lld.PopBack()
	_, ok2 := lld.PopFront()
	if ok1 == true || ok2 == true {
		t.Errorf("pop something from empty deque")
	}
}

func TestLinkListPushPopBack(t *testing.T) {
	lld := NewLinkListDeque()

	for i := 10; i >= 1; i-- {
		lld.PushBack(i)
	}
	for i := 1; i <= 10; i++ {
		val, ok := lld.PopBack()
		if !ok {
			t.Errorf("return nothing")
		}
		t.Logf("got val %v", val)
		if expect, actual := val.(int), i; expect != actual {
			t.Errorf("expect %d, got %d", expect, actual)
		}
	}
}

func TestLinkListPushPopFront(t *testing.T) {
	lld := NewLinkListDeque()

	for i := 10; i >= 1; i-- {
		lld.PushFront(i)
	}
	for i := 1; i <= 10; i++ {
		val, ok := lld.PopFront()
		if !ok {
			t.Errorf("return nothing")
		}
		t.Logf("got val %v", val)
		if expect, actual := val.(int), i; expect != actual {
			t.Errorf("expect %d, got %d", expect, actual)
		}
	}
}

func TestLinkListPushBackPopFront(t *testing.T) {
	lld := NewLinkListDeque()

	for i := 1; i <= 10; i++ {
		lld.PushBack(i)
	}

	for i := 1; i <= 10; i++ {
		val, ok := lld.PopFront()
		if !ok {
			t.Errorf("return nothing")
		}
		t.Logf("got val = %v", val)
		if expect, actual := val.(int), i; expect != actual {
			t.Errorf("expect %d, got %d", expect, actual)
		}
	}
}
