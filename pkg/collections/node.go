package collections

type Node[T any] struct {
	Current  string
	Next     string
	Previous string
	Action   func(params *T) string
}

const (
	NEXT     = "next"
	PREVIOUS = "previous"
	CURRENT  = "current"
	SKIP     = "skip"
)
