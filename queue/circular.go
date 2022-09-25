package queue

type CircularQueue[T comparable] struct {
	contents []T
	head     int
	tail     int
	count    int
}

type Queuer[T comparable] interface {
	EnQueue(value T) bool
	DeQueue() (T, bool)
	Front() (T, bool)
	Rear() (T, bool)
	IsEmpty() bool
	IsFull() bool
}

func NewQueue[T comparable](k int) Queuer[T] {
	return NewCircularQueue[T](k)
}

func NewCircularQueue[T comparable](k int) Queuer[T] {
	return &CircularQueue[T]{
		contents: make([]T, k),
		head:     0,
		tail:     0,
		count:    0,
	}
}

func (q *CircularQueue[T]) EnQueue(v T) bool {
	if q.IsFull() {
		return false
	}
	q.contents[q.tail] = v
	if q.tail+1 == len(q.contents) {
		q.tail = 0
	} else {
		q.tail++
	}
	q.count++
	return true
}

func (q *CircularQueue[T]) DeQueue() (T, bool) {
	if q.IsEmpty() {
		return zero[T](), false
	}
	v := q.contents[q.head]
	if q.head+1 == len(q.contents) {
		q.head = 0
	} else {
		q.head++
	}

	q.count--
	return v, true
}

func (q *CircularQueue[T]) Front() (T, bool) {
	if q.count == 0 {
		return zero[T](), false
	}
	return q.contents[q.head], true
}

func (q *CircularQueue[T]) Rear() (T, bool) {
	if q.count == 0 {
		return zero[T](), false
	}
	if q.tail == 0 && !q.IsEmpty() {
		return q.contents[len(q.contents)-1], true
	}
	return q.contents[q.tail-1], true
}

func (q *CircularQueue[T]) IsEmpty() bool {
	return q.count == 0
}

func (q *CircularQueue[T]) IsFull() bool {
	return q.count == len(q.contents)
}

func zero[T comparable]() T { var r T; return r }
