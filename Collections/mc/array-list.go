package mc

import (
	"errors"
	"fmt"
)

type ArrayList[T comparable] struct {
	items []T
}

func NewArrayList[T comparable](capacity uint) *ArrayList[T] {
	list := &ArrayList[T]{items: make([]T, 0, capacity)}
	return list
}

func NewEmptyArrayList[T comparable]() *ArrayList[T] {
	list := &ArrayList[T]{items: make([]T, 0, defaultCapacity)}
	return list
}

func (list *ArrayList[T]) Add(values ...T) {
	list.items = append(list.items, values...)
}

func (list *ArrayList[T]) Remove(value T) {
	for index, item := range list.items {
		if item == value {
			deleteInPlace(&list.items, uint(index))
		}
	}
}

func (list *ArrayList[T]) GetByIndex(index int) (T, error) {
	if index < 0 || index >= len(list.items) {
		return defaultOf[T](), errors.New("Index out of range")
	}
	return list.items[index], nil
}

func (list *ArrayList[T]) Length() int {
	return len(list.items)
}

func (list *ArrayList[T]) IsEmpty() bool {
	return list.Length() == 0
}

func (list *ArrayList[T]) Contains(value T) bool {
	for _, item := range list.items {
		if item == value {
			return true
		}
	}
	return false
}

func (list *ArrayList[T]) Clear() {
	list.items = nil
	list.items = make([]T, 0, defaultCapacity)
}

func (list *ArrayList[T]) String() string {
	return fmt.Sprintf("%v", list.items)
}
