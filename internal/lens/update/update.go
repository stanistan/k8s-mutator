package update

type Apply[T any] func(T) (T, error)

type Maybe[T any] struct {
	Match func(T) bool
	Apply Apply[T]
}

func Infallible[T any](val T) Apply[T] {
	return func(_ T) (T, error) {
		return val, nil
	}
}
