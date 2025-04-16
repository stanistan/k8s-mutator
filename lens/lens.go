package lens

type Lens[T, Inner, Outer any, Mutator ~func(Outer) error] struct {
	Set     func(Inner, T)
	Get     func(Inner) T
	ToInner func(Outer) Inner
}

type UpdateFunc[T any] func(T) (T, error)

func (l Lens[T, Inner, Outer, Mutator]) Mutator(with UpdateFunc[T]) Mutator {
	return Mutator(func(o Outer) error {
		inner := l.ToInner(o)
		newVal, err := with(l.Get(inner))
		if err != nil {
			return err
		}
		l.Set(inner, newVal)
		return nil
	})
}

func (l Lens[T, Inner, Outer, Mutator]) InfallibleMutator(val T) Mutator {
	return func(o Outer) error {
		l.Set(l.ToInner(o), val)
		return nil
	}
}

type ListLens[T, Inner, Outer any, Mutator ~func(Outer) error] struct {
	Lens    Lens[[]T, Inner, Outer, Mutator]
	Prepend bool
}

func (l ListLens[T, Inner, Outer, Mutator]) Mutator(matches func(T) bool, f UpdateFunc[T]) Mutator {
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
