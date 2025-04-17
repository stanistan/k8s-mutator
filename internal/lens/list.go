package lens

type ListLens[T any, Outer As[Inner], Inner any, Mutator ~func(Outer) error] struct {
	Lens    Lens[[]T, Outer, Inner, Mutator]
	Prepend bool
}

func (l ListLens[T, Outer, Inner, Mutator]) Mutator(matches func(T) bool, f UpdateFunc[T]) Mutator {
	return l.Lens.Mutator(func(items []T) ([]T, error) {
		for idx, value := range items {
			if matches(value) {
				out, err := f(value)
				if err != nil {
					return nil, err
				}

				items[idx] = out
				return items, nil
			}
		}

		var empty T
		out, err := f(empty)
		if err != nil {
			return nil, err
		}

		if l.Prepend {
			return append([]T{out}, items...), nil
		}

		return append(items, out), nil
	})
}

func (l ListLens[T, Outer, Inner, Mutator]) InfallibleMutator(matches func(T) bool, val T) Mutator {
	return l.Mutator(matches, InfallibleUpdate(val))
}
