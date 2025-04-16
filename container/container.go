package container

import (
	corev1 "k8s.io/api/core/v1"
)

func New(c *corev1.Container) Container {
	return Container{
		Container: c,
	}
}

func NewInit(c *corev1.Container) Container {
	return Container{
		Container: c,
		Init:      true,
	}
}

type Container struct {
	*corev1.Container

	Init bool
}

func (c Container) Apply(m Mutator) error {
	return m.MutateContainer(c)
}

func NewContainers(ls *[]corev1.Container) Containers {
	return Containers{
		List: ls,
	}
}

func NewInitContainers(ls *[]corev1.Container) Containers {
	return Containers{
		List: ls,
		Init: true,
	}
}

type Containers struct {
	List *[]corev1.Container

	Init bool
}

func (l Containers) Apply(m Mutator) error {
	if l.List == nil {
		return nil
	}

	var (
		list      = *l.List
		container = Container{Init: l.Init}
	)

	for idx, c := range list {
		container.Container = &c
		if err := m.MutateContainer(container); err != nil {
			return err
		}
		list[idx] = c
	}

	l.List = &list
	return nil
}
