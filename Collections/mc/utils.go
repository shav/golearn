package mc

func defaultOf[T any]() T {
	var defaultT T
	return defaultT
}

func deleteInPlace[T any](slice *[]T, index uint) {
	copy((*slice)[index:], (*slice)[index+1:]) // Shift slice[i+1:] left one index.
	(*slice)[len(*slice)-1] = defaultOf[T]()   // Erase last element (write zero value).
	*slice = (*slice)[:len(*slice)-1]          // Truncate slice.
}
