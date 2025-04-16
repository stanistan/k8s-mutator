package container

import (
	"testing"

	corev1 "k8s.io/api/core/v1"
)

func ptr[T any](in T) *T {
	return &in
}

func TestMutateContainer(t *testing.T) {
	container := corev1.Container{Name: "foo"}

	c := New(&container)
	securityContext := &corev1.SecurityContext{
		AllowPrivilegeEscalation: ptr(false),
	}
	err := c.Apply(SetSecurityContext(securityContext))
	if err != nil {
		t.Fatal("expected no error")
	}

	if container.SecurityContext != securityContext {
		t.Fatal("expected security context set")
	}

	containers := []corev1.Container{
		{Name: "foo"},
		{Name: "bar"},
	}

	cs := NewInitContainers(&containers)
	err = cs.Apply(Filtered(
		func(c Container) bool { return !c.Init },
		SetSecurityContext(securityContext),
	))
	if err != nil {
		t.Fatal("expected no error")
	}

	for _, c := range containers {
		if c.SecurityContext != nil {
			t.Fatal("securityContext must be nil")
		}
	}

}
