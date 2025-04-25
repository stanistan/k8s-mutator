package lens

import (
	"github.com/stanistan/mutator/internal/lens/update"
)

type ListLens[C, V any, F ~func(C) error] struct {
	Lens    Lens[C, []V, F]
	Prepend bool
}

func (l ListLens[C, V, F]) Do(f update.Maybe[V]) F {
	return l.Lens.Do(func(items []V) ([]V, error) {
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

		if l.Prepend {
			return append([]V{out}, items...), nil
		}

		return append(items, out), nil
	})
}
