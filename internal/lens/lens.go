package lens

import (
	"github.com/stanistan/mutator/internal/lens/update"
)

// Lens is a datatype that abstracts over getting a value V
// and setting a value V in a container C.
type Lens[C, V any, F ~func(C) error] struct {
	Set func(C, V)
	Get func(C) V
}

func (l Lens[C, V, F]) Do(fn update.Apply[V]) F {
	return func(c C) error {
		newVal, err := fn(l.Get(c))
		if err != nil {
			return err
		}
		l.Set(c, newVal)
		return nil
	}
}

func Infallible[C, V any, F ~func(C) error](l Lens[C, V, F]) func(V) F {
	return func(v V) F {
		return l.Do(update.Infallible(v))
	}
}
