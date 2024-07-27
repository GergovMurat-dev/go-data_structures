package structures

type Comparable interface {
	comparable
}

type Set[T Comparable] struct {
	collection []T
}

func (s *Set[T]) has(item T) bool {
	for _, value := range s.collection {
		if item == value {
			return true
		}
	}
	return false
}

func (s *Set[T]) Add(item T) bool {
	if !s.has(item) {
		s.collection = append(s.collection, item)
		return true
	}

	return false
}

func (s *Set[T]) Remove(item T) bool {
	if s.has(item) {
		slice := make([]T, 0, len(s.collection))

		for _, value := range s.collection {
			if value != item {
				slice = append(slice, value)
			}
		}

		s.collection = slice

		return true
	}

	return false
}

func (s *Set[T]) Values() []T {
	return s.collection
}

func (s *Set[T]) Size() int {
	return len(s.collection)
}

func (s *Set[T]) Union(anotherSet Set[T]) *Set[T] {
	unionSet := Set[T]{}

	firstSetValues := s.Values()
	secondSetValues := anotherSet.Values()

	for _, item := range firstSetValues {
		unionSet.Add(item)
	}
	for _, item := range secondSetValues {
		unionSet.Add(item)
	}

	return &unionSet
}
