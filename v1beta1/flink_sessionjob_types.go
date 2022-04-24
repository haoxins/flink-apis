package v1beta1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

//+kubebuilder:object:root=true
//TODO +kubebuilder:subresource:status

type FlinkSessionJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FlinkSessionJobSpec   `json:"spec"`
	Status FlinkSessionJobStatus `json:"status"`
}

type FlinkSessionJobSpec struct {
	DeploymentName string  `json:"deploymentName"`
	Job            JobSpec `json:"job"`
}

const (
	UpgradeModeSavepoint = "savepoint"
	UpgradeModeLastState = "last-state"
	UpgradeModeStateless = "stateless"
)

const (
	SessionJobStateRunning   = "running"
	SessionJobStateSuspended = "suspended"
)

type JobSpec struct {
	JarURI      string `json:"jarURI"`
	Parallelism int32  `json:"parallelism"`
	EntryClass  string `json:"entryClass"`
	State       string `json:"state"`
	UpgradeMode string `json:"upgradeMode"`
}

type FlinkSessionJobStatus struct {
	FlinkJobStatus FlinkJobStatusSpec `json:"jobStatus"`
}

const (
	FlinkJobStateRunning = "RUNNING"
)

type FlinkJobStatusSpec struct {
	JobId   string `json:"jobId"`
	JobName string `json:"jobName"`
	State   string `json:"state"`
}

//+kubebuilder:object:root=true

type FlinkSessionJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []FlinkSessionJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FlinkSessionJob{}, &FlinkSessionJobList{})
}
