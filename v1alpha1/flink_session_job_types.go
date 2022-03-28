package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

func init() {
	SchemeBuilder.Register(&FlinkSessionJob{}, &FlinkSessionJobList{})
}

//+kubebuilder:object:root=true
//TODO +kubebuilder:subresource:status

type FlinkSessionJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec FlinkSessionJobSpec `json:"spec"`
}

//+kubebuilder:object:root=true

type FlinkSessionJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []FlinkSessionJob `json:"items"`
}

type FlinkSessionJobSpec struct {
	ClusterId string  `json:"clusterId"`
	Job       JobSpec `json:"job"`
}

const (
	UpgradeModeSavepoint = "savepoint"
	UpgradeModeLastState = "last-state"
	UpgradeModeStateless = "stateless"
)

type JobSpec struct {
	JarURI      string `json:"jarURI"`
	Parallelism int32  `json:"parallelism"`
	EntryClass  string `json:"entryClass"`
	UpgradeMode string `json:"upgradeMode"`
}
