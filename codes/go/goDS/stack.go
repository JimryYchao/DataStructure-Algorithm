package goDS

type Stack[T any] struct {
	len      int
	capacity int
	data     []T
}

func NewStack[T any](capacity int) Stack[T] {
	if capacity < 4 {
		capacity = 4
	}
	return Stack[T]{0, capacity, make([]T, capacity)}
}

func (s *Stack[T]) Push(elem T) {
	if s.len >= s.capacity {
		s.capacity *= 2
		data := make([]T, s.capacity)
		copy(data, s.data)
		s.data = data
	}
	s.data[s.len] = elem
	s.len++
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.len == 0 {
		return (T)(*new(T)), false
	}
	s.len--
	return s.data[s.len], true
}
