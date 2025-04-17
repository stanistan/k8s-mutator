package pod

import (
	corev1 "k8s.io/api/core/v1"

	c "github.com/stanistan/mutator/container"
)

func New(pod *corev1.Pod) Pod {
	return Pod{Pod: pod}
}

type Pod struct {
	*corev1.Pod
}

func (p Pod) Apply(m Mutator) error {
	return m.MutatePod(p)
}

func (p Pod) AsInner() *corev1.Pod {
	return p.Pod
}

func (p Pod) ApplyContainerMutator(m c.Mutator) error {
	if err := c.NewInitContainers(&p.Spec.InitContainers).Apply(m); err != nil {
		return err
	}

	if err := c.NewContainers(&p.Spec.Containers).Apply(m); err != nil {
		return err
	}

	return nil
}
