package container

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/stanistan/mutator/internal/lens"
	"github.com/stanistan/mutator/internal/lens/update"
)

var (
	securityContextLens = containerLens(
		func(c *corev1.Container) *corev1.SecurityContext { return c.SecurityContext },
		func(c *corev1.Container, val *corev1.SecurityContext) { c.SecurityContext = val },
	)

	UpdateSecurityContext = securityContextLens.Do
	SetSecurityContext    = lens.Infallible(securityContextLens)
)

var (
	resourceLens = containerLens(
		func(c *corev1.Container) corev1.ResourceRequirements { return c.Resources },
		func(c *corev1.Container, val corev1.ResourceRequirements) { c.Resources = val },
	)

	UpdateResources = resourceLens.Do
	SetResources    = lens.Infallible(resourceLens)
)

var (
	envVarsLens = containerLens(
		func(c *corev1.Container) []corev1.EnvVar { return c.Env },
		func(c *corev1.Container, envs []corev1.EnvVar) { c.Env = envs },
	)
	envVarsListLens = lens.ListLens[Container, corev1.EnvVar, MutatorFunc]{
		Lens: envVarsLens,
	}
)

var (
	WithEnvVar = envVarsListLens.Do
)

func envVarUpdate(v corev1.EnvVar) update.Maybe[corev1.EnvVar] {
	return update.Maybe[corev1.EnvVar]{
		Match: func(in corev1.EnvVar) bool { return in.Name == v.Name },
		Apply: func(_ corev1.EnvVar) (corev1.EnvVar, error) { return v, nil },
	}
}

func AppendEnvVar(v corev1.EnvVar) Mutator {
	return WithEnvVar(envVarUpdate(v))
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
