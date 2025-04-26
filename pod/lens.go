package pod

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/stanistan/mutator/internal/lens"
)

func podLens[T any](
	get func(*corev1.Pod) T,
	set func(*corev1.Pod, T),
) lens.Lens[Pod, T] {
	return lens.Lens[Pod, T]{
		Get: func(p Pod) T { return get(p.Pod) },
		Set: func(p Pod, v T) { set(p.Pod, v) },
	}
}

var (
	initContainersLens = podLens(
		func(pod *corev1.Pod) []corev1.Container { return pod.Spec.InitContainers },
		func(pod *corev1.Pod, cs []corev1.Container) { pod.Spec.InitContainers = cs },
	)
)
