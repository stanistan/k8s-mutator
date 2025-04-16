package lens

type Lens[T, Inner, Outer any, Mutator ~func(Outer) error] struct {
	Set     func(Inner, T)
	Get     func(Inner) T
	ToInner func(Outer) Inner
}

type UpdateFunc[T any] func(T) (T, error)

func InfallibleUpdate[T any](val T) UpdateFunc[T] {
	return func(_ T) (T, error) {
		return val, nil
	}
}

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
	return l.Mutator(InfallibleUpdate(val))
}
