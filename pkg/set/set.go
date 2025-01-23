package set

type Set[T comparable] interface {
	Add(T)
	Remove(T)
	Contains(T) bool
	Size() int
}
