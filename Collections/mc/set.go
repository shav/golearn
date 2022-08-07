package mc

import (
	"fmt"
	"strings"
)

type Set[T comparable] struct {
	items map[T]bool
}

func NewSet[T comparable](values ...T) *Set[T] {
	set := &Set[T]{items: make(map[T]bool, len(values))}
	set.Add(values...)
	return set
}

func (set *Set[T]) Contains(value T) bool {
	exists, ok := set.items[value]
	return ok && exists
}

func (set *Set[T]) Add(values ...T) {
	for _, value := range values {
		set.items[value] = true
	}
}

func (set *Set[T]) Remove(values ...T) {
	for _, value := range values {
		delete(set.items, value)
	}
}

func (set *Set[T]) String() string {
	// TOOD: Optimize string concatenation
	numStr := make([]string, len(set.items))
	i := 0
	for item, _ := range set.items {
		numStr[i] = fmt.Sprintf("%d", item)
		i++
	}
	return fmt.Sprintf("{%s}", strings.Join(numStr, ", "))
}
