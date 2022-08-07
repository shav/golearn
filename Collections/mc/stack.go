package mc

import (
	"errors"
	"fmt"
	"strings"
)

type Stack[T comparable] struct {
	items []T
}

func NewStack[T comparable](capacity uint) *Stack[T] {
	stack := &Stack[T]{items: make([]T, 0, capacity)}
	return stack
}

func NewEmptyStack[T comparable]() *Stack[T] {
	stack := &Stack[T]{items: make([]T, 0, defaultCapacity)}
	return stack
}

func (stack *Stack[T]) Push(value T) {
	stack.items = append(stack.items, value)
}

func (stack *Stack[T]) Pop() (T, error) {
	var length = len(stack.items)
	if length == 0 {
		return defaultOf[T](), errors.New("Stack is empty")
	}

	var topIndex = length - 1
	var top = stack.items[topIndex]
	stack.items[topIndex] = defaultOf[T]()
	stack.items = stack.items[:topIndex]
	return top, nil
}

func (stack *Stack[T]) Peek() (T, error) {
	var length = len(stack.items)
	if length == 0 {
		return defaultOf[T](), errors.New("Stack is empty")
	}

	return stack.items[len(stack.items)-1], nil
}

func (stack *Stack[T]) Length() int {
	return len(stack.items)
}

func (stack *Stack[T]) IsEmpty() bool {
	return stack.Length() == 0
}

func (stack *Stack[T]) Contains(value T) bool {
	for _, item := range stack.items {
		if item == value {
			return true
		}
	}
	return false
}

func (stack *Stack[T]) Clear() {
	stack.items = nil
	stack.items = make([]T, 0, defaultCapacity)
}

func (stack *Stack[T]) String() string {
	// TOOD: Optimize string concatenation
	numStr := make([]string, len(stack.items))
	var length = len(stack.items)
	for i := length - 1; i >= 0; i-- {
		numStr[length-i-1] = fmt.Sprintf("%d", stack.items[i])
	}
	return fmt.Sprintf("[%s]", strings.Join(numStr, ", "))
}
