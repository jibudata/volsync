package mover

import (
	volsyncv1alpha1 "github.com/backube/volsync/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func AddNodeConfig(podSpec *corev1.PodSpec, c *volsyncv1alpha1.GatewayNodeConfig) {
	if c == nil {
		return
	}

	if c.NodeSelector != nil {
		podSpec.NodeSelector = c.NodeSelector
	}

	if c.Tolerations != nil {
		podSpec.Tolerations = append(podSpec.Tolerations, c.Tolerations...)
	}
}

func GetGatewayNodeConfig(owner client.Object, isSource bool) *volsyncv1alpha1.GatewayNodeConfig {
	if isSource {
		source, ok := owner.(*volsyncv1alpha1.ReplicationSource)
		if ok {
			return source.Spec.GatewayNodeConfig
		}
	} else {
		dst, ok := owner.(*volsyncv1alpha1.ReplicationDestination)
		if ok {
			return dst.Spec.GatewayNodeConfig
		}
	}
	return nil
}
