package v3

import (
	"github.com/rancher/norman/condition"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type MachineTemplate struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the desired behavior of the the cluster. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Spec MachineTemplateSpec `json:"spec"`
	// Most recent observed status of the cluster. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Status MachineTemplateStatus `json:"status"`
}

type MachineTemplateStatus struct {
	Conditions []MachineTemplateCondition `json:"conditions"`
}

type MachineTemplateCondition struct {
	// Type of cluster condition.
	Type string `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status v1.ConditionStatus `json:"status"`
	// The last time this condition was updated.
	LastUpdateTime string `json:"lastUpdateTime,omitempty"`
	// Last time the condition transitioned from one status to another.
	LastTransitionTime string `json:"lastTransitionTime,omitempty"`
	// The reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`
}

type MachineTemplateSpec struct {
	DisplayName         string `json:"displayName"`
	Description         string `json:"description"`
	Driver              string `json:"driver"`
	MachineCommonParams `json:",inline"`
}

type Machine struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the desired behavior of the the cluster. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Spec MachineSpec `json:"spec"`
	// Most recent observed status of the cluster. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Status MachineStatus `json:"status"`
}

type MachineStatus struct {
	Conditions          []MachineCondition   `json:"conditions,omitempty"`
	NodeStatus          v1.NodeStatus        `json:"nodeStatus,omitempty"`
	NodeName            string               `json:"nodeName,omitempty"`
	ClusterName         string               `json:"clusterName,omitempty" norman:"type=reference[cluster]"`
	Requested           v1.ResourceList      `json:"requested,omitempty"`
	Limits              v1.ResourceList      `json:"limits,omitempty"`
	MachineTemplateSpec *MachineTemplateSpec `json:"machineTemplateSpec,omitempty"`
	NodeConfig          *RKEConfigNode       `json:"rkeNode,omitempty"`
	SSHUser             string               `json:"sshUser,omitempty"`
	MachineDriverConfig string               `json:"machineDriverConfig,omitempty"`
}

var (
	MachineConditionInitialized condition.Cond = "Initialized"
	MachineConditionProvisioned condition.Cond = "Provisioned"
	MachineConditionConfigSaved condition.Cond = "Saved"
	MachineConditionConfigReady condition.Cond = "Ready"
)

type MachineCondition struct {
	// Type of cluster condition.
	Type condition.Cond `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status v1.ConditionStatus `json:"status"`
	// The last time this condition was updated.
	LastUpdateTime string `json:"lastUpdateTime,omitempty"`
	// Last time the condition transitioned from one status to another.
	LastTransitionTime string `json:"lastTransitionTime,omitempty"`
	// The reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`
	// Human-readable message indicating details about last transition
	Message string `json:"message,omitempty"`
}

type MachineSpec struct {
	NodeSpec             v1.NodeSpec `json:"nodeSpec"`
	DisplayName          string      `json:"displayName,omitempty"`
	RequestedHostname    string      `json:"requestedHostname,omitempty" norman:"noupdate"`
	RequestedClusterName string      `json:"requestedClusterName,omitempty" norman:"type=reference[cluster],noupdate"`
	RequestedRoles       []string    `json:"requestedRoles,omitempty" norman:"noupdate"`
	MachineTemplateName  string      `json:"machineTemplateName,omitempty" norman:"type=reference[machineTemplate],noupdate"`
	Description          string      `json:"description,omitempty"`
}

type MachineCommonParams struct {
	AuthCertificateAuthority string            `json:"authCertificateAuthority,omitempty"`
	AuthKey                  string            `json:"authKey,omitempty"`
	EngineInstallURL         string            `json:"engineInstallURL,omitempty"`
	DockerVersion            string            `json:"dockerVersion,omitempty"`
	EngineOpt                map[string]string `json:"engineOpt,omitempty"`
	EngineInsecureRegistry   []string          `json:"engineInsecureRegistry,omitempty"`
	EngineRegistryMirror     []string          `json:"engineRegistryMirror,omitempty"`
	EngineLabel              map[string]string `json:"engineLabel,omitempty"`
	EngineStorageDriver      string            `json:"engineStorageDriver,omitempty"`
	EngineEnv                map[string]string `json:"engineEnv,omitempty"`
}

type MachineDriver struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the desired behavior of the the cluster. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Spec MachineDriverSpec `json:"spec"`
	// Most recent observed status of the cluster. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Status MachineDriverStatus `json:"status"`
}

type MachineDriverStatus struct {
	Conditions []MachineDriverCondition `json:"conditions"`
}

type MachineDriverCondition struct {
	// Type of cluster condition.
	Type string `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status v1.ConditionStatus `json:"status"`
	// The last time this condition was updated.
	LastUpdateTime string `json:"lastUpdateTime,omitempty"`
	// Last time the condition transitioned from one status to another.
	LastTransitionTime string `json:"lastTransitionTime,omitempty"`
	// The reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`
}

type MachineDriverSpec struct {
	Description string `json:"description"`
	URL         string `json:"url"`
	ExternalID  string `json:"externalId"`
	Builtin     bool   `json:"builtin"`
	Active      bool   `json:"active"`
	Checksum    string `json:"checksum"`
	UIURL       string `json:"uiUrl"`
}
