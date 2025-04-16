package pod

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/stanistan/mutator/internal/lens"
)

func podLens[T any](
	get func(*corev1.Pod) T,
	set func(*corev1.Pod, T),
) lens.Lens[T, *corev1.Pod, Pod, MutatorFunc] {
	return lens.Lens[T, *corev1.Pod, Pod, MutatorFunc]{
		Get:     get,
		Set:     set,
		ToInner: func(p Pod) *corev1.Pod { return p.Pod },
	}
}

var (
	initContainersLens = podLens(
		func(pod *corev1.Pod) []corev1.Container { return pod.Spec.InitContainers },
		func(pod *corev1.Pod, cs []corev1.Container) { pod.Spec.InitContainers = cs },
	)

	initContainersListLens = lens.ListLens[corev1.Container, *corev1.Pod, Pod, MutatorFunc]{
		Lens:    initContainersLens,
		Prepend: true,
	}
)
