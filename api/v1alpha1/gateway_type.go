package v1alpha1

import corev1 "k8s.io/api/core/v1"

type GatewayNodeConfig struct {
	NodeSelector map[string]string   `json:"nodeSelector,omitempty"`
	Tolerations  []corev1.Toleration `json:"tolerations,omitempty"`
}
