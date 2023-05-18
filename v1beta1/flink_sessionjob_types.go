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
	DeploymentName     string            `json:"deploymentName"`
	Job                JobSpec           `json:"job"`
	RestartNonce       int64             `json:"restartNonce,omitempty"`
	FlinkConfiguration map[string]string `json:"flinkConfiguration,omitempty"`
}

func (sessionJob FlinkSessionJobSpec) Equals(other FlinkSessionJobSpec) (bool, error) {
	if sessionJob.DeploymentName != other.DeploymentName {
		return false, nil
	}

	return sessionJob.Job.Equals(other.Job)
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

func (job JobSpec) Equals(other JobSpec) (bool, error) {
	if job.JarURI != other.JarURI {
		return false, nil
	}
	if job.EntryClass != other.EntryClass {
		return false, nil
	}
	if job.Parallelism != other.Parallelism {
		return false, nil
	}
	if job.State != other.State {
		return false, nil
	}
	if job.UpgradeMode != other.UpgradeMode {
		return false, nil
	}
	if job.SavepointTriggerNonce != other.SavepointTriggerNonce {
		return false, nil
	}
	if job.InitialSavepointPath != other.InitialSavepointPath {
		return false, nil
	}
	if job.AllowNonRestoredState != other.AllowNonRestoredState {
		return false, nil
	}

	return IsArgsEqual(job.Args, other.Args)
}

type FlinkSessionJobStatus struct {
	// reconciliationStatus string `json:"reconciliationStatus"`
	LifecycleState string             `json:"lifecycleState"`
	Error          string             `json:"error"`
	FlinkJobStatus FlinkJobStatusSpec `json:"jobStatus"`
}

const (
	FlinkJobStateRunning = "RUNNING"
)

type FlinkJobStatusSpec struct {
	JobId      string `json:"jobId"`
	JobName    string `json:"jobName"`
	State      string `json:"state"`
	StartTime  string `json:"startTime"`
	UpdateTime string `json:"updateTime"`
	// SavepointInfo SavepointInfo `json:"savepointInfo"`
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
