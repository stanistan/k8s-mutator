package pod

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/stanistan/mutator/container"
	"github.com/stanistan/mutator/internal/lens"
)

type Mutator interface {
	MutatePod(pod Pod) error
}

type MutatorFunc func(Pod) error

func (f MutatorFunc) MutatePod(p Pod) error {
	return f(p)
}

func WithInitContainer(c corev1.Container, cs container.Mutator) Mutator {
	return initContainersListLens.Updator(lens.ListUpdate[corev1.Container]{
		Matches: func(in corev1.Container) bool { return in.Name == c.Name },
		Apply: func(_ corev1.Container) (corev1.Container, error) {
			if err := container.NewInit(&c).Apply(cs); err != nil {
				return c, err
			}
			return c, nil
		},
	})
}
