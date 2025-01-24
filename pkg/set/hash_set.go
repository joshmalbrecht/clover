package set

type hashSet[T comparable] struct {
	hashMap map[T]struct{}
}

func NewHashSet[T comparable]() Set[T] {
	return &hashSet[T]{
		hashMap: make(map[T]struct{}),
	}
}

func (hs hashSet[T]) Add(value T) {
	hs.hashMap[value] = struct{}{}
}

func (hs hashSet[T]) Remove(value T) {
	delete(hs.hashMap, value)
}

func (hs hashSet[T]) Contains(value T) bool {
	_, existing := hs.hashMap[value]

	return existing
}

func (hs hashSet[T]) Size() int {
	return len(hs.hashMap)
}
