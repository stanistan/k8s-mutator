# k8s mutators

Experimental library for composing mutators of kubernetes objects (containers & pods).

---

This doesn't have all of the things built in but you should be able
to mutate a pod by composing low-level to high-level behaviors as a
sort of "plan" for mutation.

There are a couple of things at play here:

1. `container.Mutator` and `pod.Mutator` -- these are an interface that can mutate a kubernets
  container and pod (respectively).

2. `container.Container` and `pod.Pod` are wrapping structs for `*corev1.Container` and `*corev1.Pod`.

The original concept was you would mutate a container `Mutate(*corev1.Container) error` or fail,
but then the question arises: what about the difference between init containers and non-init
containers in the pod spec? Instead of changing it to `Mutate(*corev1.Container, bool) error`
I wanted to try an options struct `Mutate(*corev1.Container, Options)` but then if you had nothing
to pass in what was the point?

I opted for `Container{}` which has a field of `Init` along with embedding `*corev1.Container` and to keep
the design consistent, did the same thing for `Pod`. There are no other options for `Pod` at this time.

---

### Fake Example

Given a corev1.Pod I want to add a security context & env var to all non-init container.

```go
err := pod.New(&pod).ApplyContainerMutator(container.Filtered(
    func(c container.Container) bool { return !c.Init },
    container.Mutators{
        container.SetSecurityContext(&corev1.SecurityContext{
           // ...
        }),
        container.AppendEnvVar(corev1.EnvVar{
            Name: "MUTATION_STATUS",
            Value: "security_context",
        }),
    },
))
```
