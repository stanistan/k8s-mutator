package pod

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/stanistan/mutator/container"
)

type Mutator interface {
	MutatePod(pod Pod) error
}

type MutatorFunc func(Pod) error

func (f MutatorFunc) MutatePod(p Pod) error {
	return f(p)
}

func WithInitContainer(c corev1.Container, cs container.Mutator) Mutator {
	return MutatorFunc(initContainersListLens.Mutator(
		func(in corev1.Container) bool { return in.Name == c.Name },
		func(in corev1.Container) (corev1.Container, error) {
			if err := container.NewInit(&c).Apply(cs); err != nil {
				return in, err
			}
			return c, nil
		},
	))
}
