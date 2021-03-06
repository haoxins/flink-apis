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

func (self FlinkSessionJobSpec) Equals(other FlinkSessionJobSpec) (bool, error) {
	if self.DeploymentName != other.DeploymentName {
		return false, nil
	}

	return self.Job.Equals(other.Job)
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
	JarURI                string   `json:"jarURI"`
	EntryClass            string   `json:"entryClass"`
	Args                  []string `json:"args,omitempty"`
	Parallelism           int32    `json:"parallelism"`
	State                 string   `json:"state"`
	UpgradeMode           string   `json:"upgradeMode"`
	SavepointTriggerNonce int64    `json:"savepointTriggerNonce,omitempty"`
	InitialSavepointPath  string   `json:"initialSavepointPath,omitempty"`
	AllowNonRestoredState bool     `json:"allowNonRestoredState,omitempty"`
}

func (self JobSpec) Equals(other JobSpec) (bool, error) {
	if self.JarURI != other.JarURI {
		return false, nil
	}
	if self.EntryClass != other.EntryClass {
		return false, nil
	}
	if self.Parallelism != other.Parallelism {
		return false, nil
	}
	if self.State != other.State {
		return false, nil
	}
	if self.UpgradeMode != other.UpgradeMode {
		return false, nil
	}
	if self.SavepointTriggerNonce != other.SavepointTriggerNonce {
		return false, nil
	}
	if self.InitialSavepointPath != other.InitialSavepointPath {
		return false, nil
	}
	if self.AllowNonRestoredState != other.AllowNonRestoredState {
		return false, nil
	}

	return IsArgsEqual(self.Args, other.Args)
}

type FlinkSessionJobStatus struct {
	Error          string             `json:"error"`
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
