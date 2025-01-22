package queue

type Queue[T any] interface {
	Enqueue(T)
	Dequeue() (T, bool)
	Peek() (T, bool)
	Size() int
}

type node[T any] struct {
	value T
	next  *node[T]
}

type queue[T any] struct {
	head *node[T]
	tail *node[T]
	size int
}

func New[T any]() Queue[T] {
	return &queue[T]{nil, nil, 0}
}

func (q *queue[T]) Enqueue(value T) {
	var newNode node[T]
	newNode.value = value

	q.size = q.size + 1

	if q.head == nil {
		q.head = &newNode
		q.tail = &newNode
		return
	}

	q.tail.next = &newNode
	q.tail = &newNode
}

func (q *queue[T]) Dequeue() (T, bool) {
	if q.head == nil {
		var t T
		return t, false
	}

	var head = q.head
	q.head = head.next
	q.size = q.size - 1

	return head.value, true
}

func (q *queue[T]) Peek() (T, bool) {
	if q.head == nil {
		var t T
		return t, false
	}

	return q.head.value, true
}

func (q *queue[T]) Size() int {
	return q.size
}
