package goDS

type Queue[T any] struct {
	len      int
	head     int
	index    int
	capacity int
	data     []T
}

func NewQueue[T any](capacity int) Queue[T] {
	if capacity < 4 {
		capacity = 4
	}
	return Queue[T]{0, 0, 0, capacity, make([]T, capacity)}
}

func (q *Queue[T]) Enqueue(elem T) {
	if q.len >= q.capacity {
		data := make([]T, q.capacity*2)
		copy(data, q.data[q.head:q.capacity])
		copy(data[q.capacity-q.head:], q.data[:q.head])

		// for i, j := 0, q.head; i < q.len; i++ {
		// 	data[i] = q.data[j]
		// 	j = (j + 1) % q.len
		// }

		q.head = 0
		q.index = q.len
		q.capacity *= 2
		q.data = data
	}
	q.data[q.index] = elem
	q.index = (q.index + 1) % q.capacity
	q.len++
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if q.len < 1 {
		return (T)(*new(T)), false
	}
	rt := q.data[q.head]
	q.head = (q.head + 1) % q.capacity
	q.len--
	return rt, true
}

func (q *Queue[T]) Len() int {
	return q.len
}
