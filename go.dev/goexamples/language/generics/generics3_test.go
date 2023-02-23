package generics_test

type Element interface {
	~int64 | ~float64 | ~string
}

type Container[T Element] interface {
	Add(T)
	Len() int
}

type MyArray[T Element] struct {
	data []T
}

func (m *MyArray[T]) Add(e T) {
	m.data = append(m.data, e)
}

func (m *MyArray[T]) Len() int {
	return len(m.data)
}
