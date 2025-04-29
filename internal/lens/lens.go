package lens

import (
	"github.com/stanistan/k8s-mutator/internal/lens/update"
)

// Lens is a datatype that abstracts over getting a value V
// and setting a value V in a container C.
type Lens[C, V any] struct {
	Set func(C, V)
	Get func(C) V
}

func Updater[F ~func(C) error, C, V any](l Lens[C, V]) func(update.Apply[V]) F {
	return func(fn update.Apply[V]) F {
		return func(c C) error {
			newVal, err := fn(l.Get(c))
			if err != nil {
				return err
			}
			l.Set(c, newVal)
			return nil
		}
	}
}

func InfallibleUpdater[F ~func(C) error, C, V any](l Lens[C, V]) func(V) F {
	return func(v V) F {
		return Updater[F, C, V](l)(update.Infallible(v))
	}
}

func ListUpdater[F ~func(C) error, C, V any](l Lens[C, []V], prepend bool) func(update.Maybe[V]) F {
	listLens := Updater[F, C, []V](l)
	return func(f update.Maybe[V]) F {
		return listLens(func(items []V) ([]V, error) {
			for idx, value := range items {
				if f.Match(value) {
					out, err := f.Apply(value)
					if err != nil {
						return nil, err
					}

					items[idx] = out
					return items, nil
				}
			}

			var empty V
			out, err := f.Apply(empty)
			if err != nil {
				return nil, err
			}

			if prepend {
				return append([]V{out}, items...), nil
			}

			return append(items, out), nil
		})
	}
}
