package container

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/stanistan/mutator/internal/lens"
)

var (
	securityContextLens = containerLens(
		func(c *corev1.Container) *corev1.SecurityContext { return c.SecurityContext },
		func(c *corev1.Container, val *corev1.SecurityContext) { c.SecurityContext = val },
	)

	UpdateSecurityContext = securityContextLens.Updator
	SetSecurityContext    = securityContextLens.InfallibleUpdator
)

var (
	resourceLens = containerLens(
		func(c *corev1.Container) corev1.ResourceRequirements { return c.Resources },
		func(c *corev1.Container, val corev1.ResourceRequirements) { c.Resources = val },
	)

	UpdateResources = resourceLens.Updator
	SetResources    = resourceLens.InfallibleUpdator
)

var (
	envVarsLens = containerLens(
		func(c *corev1.Container) []corev1.EnvVar { return c.Env },
		func(c *corev1.Container, envs []corev1.EnvVar) { c.Env = envs },
	)
	envVarsListLens = lens.ListLens[Container, corev1.EnvVar, MutatorFunc]{
		Lens:    envVarsLens,
		Prepend: true,
	}
)

var (
	WithEnvVar = envVarsListLens.Updator
)

func InjectEnvVar(v corev1.EnvVar) Mutator {
	return WithEnvVar(lens.ListUpdate[corev1.EnvVar]{
		Matches: func(in corev1.EnvVar) bool { return in.Name == v.Name },
		Apply:   func(_ corev1.EnvVar) (corev1.EnvVar, error) { return v, nil },
	})
}

func containerLens[T any](
	get func(*corev1.Container) T,
	set func(*corev1.Container, T),
) lens.Lens[Container, T, MutatorFunc] {
	return lens.Lens[Container, T, MutatorFunc]{
		Set: func(c Container, v T) { set(c.Container, v) },
		Get: func(c Container) T { return get(c.Container) },
	}
}
