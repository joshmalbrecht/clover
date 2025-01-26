package stack

type Stack[T comparable] interface {
	Push(T)
	Pop() (T, bool)
	Peek() (T, bool)
	Size() int
}

type node[T comparable] struct {
	value    T
	previous *node[T]
}

type stack[T comparable] struct {
	tail *node[T]
	size int
}

func New[T comparable]() Stack[T] {
	return &stack[T]{nil, 0}
}

func (s *stack[T]) Push(value T) {
	newNode := node[T]{
		value: value,
	}

	s.size = s.size + 1

	if s.tail == nil {
		s.tail = &newNode
		return
	}

	newNode.previous = s.tail

	s.tail = &newNode
}

func (s *stack[T]) Pop() (T, bool) {
	if s.tail == nil {
		var t T
		return t, false
	}

	val := s.tail.value
	s.tail = s.tail.previous
	s.size = s.size - 1

	return val, true
}

func (s *stack[T]) Peek() (T, bool) {
	if s.tail == nil {
		var t T
		return t, false
	}

	return s.tail.value, true
}

func (s *stack[T]) Size() int {
	return s.size
}
