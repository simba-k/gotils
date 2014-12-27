package containers

import (
	"testing"
)

func TestIntStack(t *testing.T) {
	stack := NewIntStack(5)
	err := stack.Push(10)
	if err!=nil {
		t.Error(err)
	}
	if stack.Size() != 1 {
		t.Error("Incorrect size")
	}
	val, err := stack.Pop()
	if err!=nil {
		t.Error(err)
	}
	if val != 10 {
		t.Error("wrong pop value expected 10 got ", val)
	}
	for i := 0; i < 5; i++ {
		err := stack.Push(i)
		if err!=nil {
			t.Error(err)
		}
	}
	val, err = stack.Pop()
	if err!=nil {
		t.Error(err)
	}
	if val != 4 {
		t.Error("Wrong Value")
	}
	t.Log(stack)
}
