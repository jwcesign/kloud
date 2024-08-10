package analyzer

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1beta1 "sigs.k8s.io/karpenter/pkg/apis/v1beta1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories={cloudpilot-ai}
// +kubebuilder:metadata:labels=cloudpilot.ai/crd-install=true

type ClusterMigration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// +required
	Spec ClusterMigrationSpec `json:"spec,omitempty"`

	// +optional
	Status ClusterMigrationStatus `json:"status,omitempty"`
}

type ClusterMigrationSpec struct {
	// +required
	CloudProvider string `json:"clusterName"`
	// +required
	Region string `json:"region"`
	// +required
	CloudPilotAIEnabled bool `json:"cloudPilotAIEnabled"`
	// +required
	NodePoolConfig corev1beta1.NodePoolSpec `json:"nodePoolConfig"`
}

type ClusterMigrationStatus struct {
	// +optional
	PodErrors []string `json:"podErrors,omitempty"`
}

type PodError struct {
	// +required
	Namespace string
	// +required
	Name string
	// +required
	Reason string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterMigrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterMigration `json:"items"`
}
