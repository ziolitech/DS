package stack

import (
	"errors"
	"sync"
)

type Stack[T any] struct {
	elements []T
}

type ConcurrentStack[T any] struct {
	mu      sync.Mutex
	element []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		elements: make([]T, 0),
	}
}

func (st *Stack[T]) IsEmpty() bool {
	if len(st.elements) == 0 {
		return true
	}

	return false
}

// Push - method on stack
func (st *Stack[T]) Push(item T) {
	st.elements = append(st.elements, item)
}

func (st *Stack[T]) Pop() (T, error) {
	if st.IsEmpty() {
		var zero T
		return zero, errors.New("can not pop, stack is empty")
	}

	topIndex := len(st.elements) - 1
	top := st.elements[topIndex]
	//reslice
	st.elements = st.elements[:topIndex]

	return top, nil
}

// TODO :
// Other methods to implement:
// Peek, Size, etc.
