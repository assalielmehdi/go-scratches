package data_structures

type Stack[T any] interface {
	Push(T) Stack[T]
	Pop() (*T, error)
	Size() int
}
