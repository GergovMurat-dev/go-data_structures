package structures

type Stack[T any] struct {
	size int
	data []T
}

func (s *Stack[T]) Push(item T) {
	s.data = append(s.data, item)
	s.size++
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.size == 0 {
		var zero T
		return zero, false
	}

	item := s.data[s.size-1]
	s.data = s.data[:s.size-1]

	return item, true
}
