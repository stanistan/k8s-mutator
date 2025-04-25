package lens

type ListLens[C, V any, F ~func(C) error] struct {
	Lens    Lens[C, []V, F]
	Prepend bool
}

type ListUpdate[V any] struct {
	Apply   UpdateFunc[V]
	Matches func(V) bool
}

func (l ListLens[C, V, F]) Updator(f ListUpdate[V]) F {
	return l.Lens.Updator(func(items []V) ([]V, error) {
		for idx, value := range items {
			if f.Matches(value) {
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
