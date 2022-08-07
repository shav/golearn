package mc

type IStack[T comparable] interface {
	Push(value T)
	Pop() (T, error)
	Peek() (T, error)
	Length() int
	IsEmpty() bool
	Contains(value T) bool
	Clear()
}

type IList[T comparable] interface {
	Add(values ...T)
	Remove(value T)
	GetByIndex(index int) (T, error)
	Length() int
	IsEmpty() bool
	Contains(value T) bool
	Clear()
}
