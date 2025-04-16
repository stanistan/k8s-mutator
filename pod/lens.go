package pod

import (
	corev1 "k8s.io/api/core/v1"

	"github.com/stanistan/mutator/internal/lens"
)

var (
	initContainersLens = lens.Lens[[]corev1.Container, *corev1.Pod, Pod, MutatorFunc]{
		Get:     func(pod *corev1.Pod) []corev1.Container { return pod.Spec.InitContainers },
		Set:     func(pod *corev1.Pod, cs []corev1.Container) { pod.Spec.InitContainers = cs },
		ToInner: func(p Pod) *corev1.Pod { return p.Pod },
	}
	initContainersListLens = lens.ListLens[corev1.Container, *corev1.Pod, Pod, MutatorFunc]{
		Lens:    initContainersLens,
		Prepend: true,
	}
)
