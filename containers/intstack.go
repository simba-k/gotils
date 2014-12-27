package containers

import (
	"errors"
)

//todo add a resizable stack

type IntStack interface{
	Push(val int) error
	Pop() (int, error)
	Size() int
}

type intstack []int

func NewIntStack(capacity int) (val IntStack){
	stk := new(intstack)
	*stk = intstack(make([]int, 0, capacity))
	val = stk
	return
}

func (stack *intstack) Push(val int) error {
	if stack.Size() < cap(*stack) {
		*stack = append(*stack, val)
		return nil
	} else {
		return errors.New("Stack Overflow")
	}
}

func (stack *intstack) Pop() (val int, err error) {
	if(stack.Size() > 0) {
		err = nil
		lent := len(*stack) - 1
		val = (*stack)[lent]
		*stack = (*stack)[:lent]
		return
	} else {
		err = errors.New("Stack Underflow")
		val = -1
		return
	}
}

func (stack *intstack) Size() int {
	return len(*stack)
}
