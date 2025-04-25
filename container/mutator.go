package container

// Mutator mutates a *corev1.Container or error.
//
// There is guarantee that a container _was not_ mutated
// if there was a failure here.
type Mutator interface {
	MutateContainer(Container) error
}

type Mutators []Mutator

func (ms Mutators) Mutate(c Container) error {
	for _, m := range ms {
		if err := m.MutateContainer(c); err != nil {
			return err
		}
	}

	return nil
}

type MutatorFunc func(Container) error

func (f MutatorFunc) MutateContainer(c Container) error {
	return f(c)
}

type Filter func(Container) bool

func Filtered(f Filter, m Mutator) Mutator {
	return MutatorFunc(func(c Container) error {
		if f != nil && !f(c) {
			return nil
		}

		return m.MutateContainer(c)
	})
}
