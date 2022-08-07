package mc

import (
	"errors"
	"fmt"
)

type Queue[T comparable] struct {
	items []T
}

func NewQueue[T comparable](capacity uint) *Queue[T] {
	queue := &Queue[T]{items: make([]T, 0, capacity)}
	return queue
}

func NewEmptyQueue[T comparable]() *Queue[T] {
	queue := &Queue[T]{items: make([]T, 0, defaultCapacity)}
	return queue
}

func (queue *Queue[T]) Enqueue(value T) {
	queue.items = append(queue.items, value)
}

func (queue *Queue[T]) Dequeue() (T, error) {
	var defaultT T
	var length = len(queue.items)
	if length == 0 {
		return defaultT, errors.New("Queue is empty")
	}

	var first = queue.items[0]
	queue.items[0] = defaultT
	queue.items = queue.items[1:]
	return first, nil
}

func (queue *Queue[T]) Peek() (T, error) {
	var length = len(queue.items)
	if length == 0 {
		return defaultOf[T](), errors.New("Queue is empty")
	}

	return queue.items[0], nil
}

func (queue *Queue[T]) Length() int {
	return len(queue.items)
}

func (queue *Queue[T]) IsEmpty() bool {
	return queue.Length() == 0
}

func (queue *Queue[T]) Contains(value T) bool {
	for _, item := range queue.items {
		if item == value {
			return true
		}
	}
	return false
}

func (queue *Queue[T]) Clear() {
	queue.items = nil
	queue.items = make([]T, 0, defaultCapacity)
}

func (queue *Queue[T]) String() string {
	return fmt.Sprintf("%v", queue.items)
}
