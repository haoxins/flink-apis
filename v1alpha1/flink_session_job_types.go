package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type FlinkSessionJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec FlinkSessionJobSpec `json:"spec"`
}

type FlinkSessionJobSpec struct {
	ClusterId string  `json:"clusterId"`
	Job       JobSpec `json:"job"`
}

type JobSpec struct {
	JarURI      string `json:"jarURI"`
	Parallelism int32  `json:"parallelism"`
	EntryClass  string `json:"entryClass"`
	UpgradeMode string `json:"upgradeMode"`
}
