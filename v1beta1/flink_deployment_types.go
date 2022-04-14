package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//+kubebuilder:object:root=true
//TODO +kubebuilder:subresource:status

type FlinkDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FlinkDeploymentSpec   `json:"spec"`
	Status FlinkDeploymentStatus `json:"status"`
}

type FlinkDeploymentSpec struct {
	Image string `json:"image"`

	FlinkVersion string `json:"flinkVersion"`

	FlinkConfiguration map[string]string `json:"flinkConfiguration"`

	ServiceAccount string `json:"serviceAccount"`

	JobManager JobManagerSpec `json:"jobManager"`

	TaskManager TaskManagerSpec `json:"taskManager"`
}

const (
	JobManagerStatusReady = "READY"

	JobManagerStatusDeployedNotReady = "DEPLOYED_NOT_READY"

	JobManagerStatusDeploying = "DEPLOYING"

	// TODO: currently a mix of SUSPENDED and ERROR, needs cleanup
	JobManagerStatusMissing = "MISSING"

	JobManagerStatusError = "ERROR"
)

type FlinkDeploymentStatus struct {
	JobManagerStatus string `json:"jobManagerDeploymentStatus"`
}

type JobManagerSpec struct {
	Replicas int32        `json:"replicas"`
	Resource ResourceSpec `json:"resource"`
}

type TaskManagerSpec struct {
	Resource ResourceSpec `json:"resource"`
}

type ResourceSpec struct {
	Memory string  `json:"memory"`
	CPU    float32 `json:"cpu"`
}

//+kubebuilder:object:root=true

type FlinkDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []FlinkDeployment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FlinkDeployment{}, &FlinkDeploymentList{})
}
