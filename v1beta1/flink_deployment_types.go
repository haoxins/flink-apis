package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
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

	ImagePullPolicy string `json:"imagePullPolicy,omitempty"`

	FlinkVersion string `json:"flinkVersion"`

	FlinkConfiguration map[string]string `json:"flinkConfiguration,omitempty"`

	LogConfiguration map[string]string `json:"logConfiguration,omitempty"`

	ServiceAccount string `json:"serviceAccount"`

	JobManager JobManagerSpec `json:"jobManager"`

	TaskManager TaskManagerSpec `json:"taskManager"`

	JobSpec JobSpec `json:"job,omitempty"`

	PodTemplate corev1.PodTemplate `json:"podTemplate,omitempty"`

	RestartNonce int64 `json:"restartNonce,omitempty"`
}

type FlinkDeploymentStatus struct {
	// ReconciliationStatus string            `json:"reconciliationStatus,omitempty"`
	JobManagerStatus string            `json:"jobManagerDeploymentStatus"`
	ClusterInfo      map[string]string `json:"clusterInfo"`
	TaskManager      TaskManagerInfo   `json:"taskManager"`
}

type JobManagerSpec struct {
	Replicas    int32              `json:"replicas"`
	Resource    ResourceSpec       `json:"resource"`
	PodTemplate corev1.PodTemplate `json:"podTemplate,omitempty"`
}

type TaskManagerSpec struct {
	Resource    ResourceSpec       `json:"resource"`
	Replicas    int32              `json:"replicas,omitempty"`
	PodTemplate corev1.PodTemplate `json:"podTemplate,omitempty"`
}

type ResourceSpec struct {
	Memory           string  `json:"memory"`
	CPU              float32 `json:"cpu"`
	EphemeralStorage string  `json:"ephemeralStorage,omitempty"`
}

type TaskManagerInfo struct {
	Replicas      int32  `json:"replicas"`
	LabelSelector string `json:"labelSelector"`
}

const (
	JobManagerStatusReady = "READY"

	JobManagerStatusDeployedNotReady = "DEPLOYED_NOT_READY"

	JobManagerStatusDeploying = "DEPLOYING"

	// TODO: currently a mix of SUSPENDED and ERROR, needs cleanup
	JobManagerStatusMissing = "MISSING"

	JobManagerStatusError = "ERROR"
)

//+kubebuilder:object:root=true

type FlinkDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []FlinkDeployment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FlinkDeployment{}, &FlinkDeploymentList{})
}
