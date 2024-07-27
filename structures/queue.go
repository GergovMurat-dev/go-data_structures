package structures

type Queue[T any] struct {
	size int
	data []T
}

func (q *Queue[T]) Push(item T) {
	q.data = append(q.data, item)
	q.size++
}

func (q *Queue[T]) Pop() (T, bool) {
	if len(q.data) == 0 {
		var zero T
		return zero, false
	}

	item := q.data[0]
	q.data = q.data[1:]

	return item, true
}
