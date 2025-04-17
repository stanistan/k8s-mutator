package lens

type As[T any] interface {
	AsInner() T
}

type Lens[T any, Outer As[Inner], Inner any, Mutator ~func(Outer) error] struct {
	Set func(Inner, T)
	Get func(Inner) T
}

type UpdateFunc[T any] func(T) (T, error)

func InfallibleUpdate[T any](val T) UpdateFunc[T] {
	return func(_ T) (T, error) {
		return val, nil
	}
}

func (l Lens[T, Outer, Inner, Mutator]) Mutator(with UpdateFunc[T]) Mutator {
	return Mutator(func(o Outer) error {
		inner := o.AsInner()
		newVal, err := with(l.Get(inner))
		if err != nil {
			return err
		}
		l.Set(inner, newVal)
		return nil
	})
}

func (l Lens[T, Outer, Inner, Mutator]) InfallibleMutator(val T) Mutator {
	return l.Mutator(InfallibleUpdate(val))
}
