package container

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/stanistan/mutator/internal/lens"
	"github.com/stanistan/mutator/internal/lens/update"
)

var (
	UpdateSecurityContext = lens.Updater[MutatorFunc](securityContextLens)
	SetSecurityContext    = lens.InfallibleUpdater[MutatorFunc](securityContextLens)
	UpdateResources       = lens.Updater[MutatorFunc](resourceLens)
	SetResources          = lens.InfallibleUpdater[MutatorFunc](resourceLens)
	WithEnvVar            = lens.ListUpdater[MutatorFunc](envVarsLens, false)
)

func AppendEnvVar(v corev1.EnvVar) Mutator {
	return WithEnvVar(envVarUpdate(v))
}

func containerLens[T any](
	get func(*corev1.Container) T,
	set func(*corev1.Container, T),
) lens.Lens[Container, T] {
	return lens.Lens[Container, T]{
		Set: func(c Container, v T) { set(c.Container, v) },
		Get: func(c Container) T { return get(c.Container) },
	}
}

var (
	securityContextLens = containerLens(
		func(c *corev1.Container) *corev1.SecurityContext { return c.SecurityContext },
		func(c *corev1.Container, val *corev1.SecurityContext) { c.SecurityContext = val },
	)

	resourceLens = containerLens(
		func(c *corev1.Container) corev1.ResourceRequirements { return c.Resources },
		func(c *corev1.Container, val corev1.ResourceRequirements) { c.Resources = val },
	)

	envVarsLens = containerLens(
		func(c *corev1.Container) []corev1.EnvVar { return c.Env },
		func(c *corev1.Container, envs []corev1.EnvVar) { c.Env = envs },
	)
)

func envVarUpdate(v corev1.EnvVar) update.Maybe[corev1.EnvVar] {
	return update.Maybe[corev1.EnvVar]{
		Match: func(in corev1.EnvVar) bool { return in.Name == v.Name },
		Apply: func(_ corev1.EnvVar) (corev1.EnvVar, error) { return v, nil },
	}
}
