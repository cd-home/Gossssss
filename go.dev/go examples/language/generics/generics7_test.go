package generics_test

type Stack[T any] []T

func (s *Stack[T]) push(elem T) {
	*s = append(*s, elem)
}

func (s *Stack[T]) pop() {
	if len(*s) > 0 {
		*s = (*s)[:len(*s)-1]
	}
}

func (s *Stack[T]) top() *T {
	if len(*s) > 0 {
		return &(*s)[len(*s)-1]
	}
	return nil
}

func (s *Stack[T]) Len() int {
	return len(*s)
}
