package lens

// UpdateFunc is a function that takes in a current
// value T and produces a new one, or an error. Returning
// the inital value and nil is effectively a noop.
type UpdateFunc[T any] func(T) (T, error)

// InfallibleUpdate returns an UpdateFunc that always
// returns the given value regardless of the current
// value provided.
func InfallibleUpdate[T any](val T) UpdateFunc[T] {
	return func(_ T) (T, error) {
		return val, nil
	}
}

// Lens is a datatype that abstracts over getting a value V
// and setting a value V in a container C.
type Lens[C, V any, F ~func(C) error] struct {
	Set func(C, V)
	Get func(C) V
}

func (l Lens[C, V, F]) Updator(fn UpdateFunc[V]) F {
	return func(c C) error {
		newVal, err := fn(l.Get(c))
		if err != nil {
			return err
		}
		l.Set(c, newVal)
		return nil
	}
}

func (l Lens[C, V, F]) InfallibleUpdator(val V) F {
	return l.Updator(InfallibleUpdate(val))
}
