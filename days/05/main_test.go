package main

import (
	"reflect"
	"testing"
)

func TestStack_PushAndPop(t *testing.T) {
	stack := NewStack()

	items := []string{"a", "b", "c", "d"}

	for _, item := range items {
		stack.Push(item)
	}

	for i := len(items) - 1; i > 0; i-- {
		item := items[i]
		if got := stack.Pop(); !reflect.DeepEqual(got, item) {
			t.Errorf("stack.Pop() = %v, want %v", got, item)
		}

	}
}
