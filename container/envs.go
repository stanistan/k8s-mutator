package container

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/stanistan/mutator/internal/lens"
)

var (
	envVarsLens = containerLens(
		func(c *corev1.Container) []corev1.EnvVar { return c.Env },
		func(c *corev1.Container, envs []corev1.EnvVar) { c.Env = envs },
	)
	envVarsListLens = lens.ListLens[corev1.EnvVar, Container, *corev1.Container, MutatorFunc]{
		Lens:    envVarsLens,
		Prepend: true,
	}
)

func WithEnvVar(name string, updateFn lens.UpdateFunc[corev1.EnvVar]) Mutator {
	return envVarsListLens.Mutator(
		func(in corev1.EnvVar) bool { return in.Name == name },
		updateFn,
	)
}
