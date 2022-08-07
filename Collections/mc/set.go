package mc

import (
	"fmt"
	"strings"
)

type Set struct {
	items map[int]bool
}

func NewSet(values ...int) *Set {
	set := &Set{items: make(map[int]bool)}
	set.Add(values...)
	return set
}

func (set *Set) Contains(value int) bool {
	exists, ok := set.items[value]
	return ok && exists
}

func (set *Set) Add(values ...int) {
	for _, value := range values {
		set.items[value] = true
	}
}

func (set *Set) Remove(values ...int) {
	for _, value := range values {
		delete(set.items, value)
	}
}

func (set *Set) String() string {
	numStr := make([]string, len(set.items))
	i := 0
	for item, _ := range set.items {
		numStr[i] = fmt.Sprintf("%d", item)
		i++
	}
	return fmt.Sprintf("{%s}", strings.Join(numStr, ", "))
}
