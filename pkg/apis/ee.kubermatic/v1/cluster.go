/*
Copyright 2023 The Kubermatic Kubernetes Platform contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	kubermaticv1 "k8c.io/api/v3/pkg/apis/kubermatic/v1"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
)

const (
	// PresetNameAnnotation is key of the annotation used to hold preset name if was used for the cluster creation.
	PresetNameAnnotation = "presetName"

	// PresetInvalidatedAnnotation is key of the annotation used to indicate why the preset was invalidated.
	PresetInvalidatedAnnotation = "presetInvalidated"

	// ProjectIDLabelKey is the label on a Cluster resource that points to the project it belongs to.
	ProjectIDLabelKey = "project-id"

	IsCredentialPresetLabelKey = "is-credential-preset"

	UpdatedByVPALabelKey = "updated-by-vpa"

	ExternalClusterIDLabelKey = "external-cluster-id"
)

const (
	// ClusterFeatureEncryptionAtRest enables the experimental "encryption-at-rest" feature, which allows encrypting
	// Kubernetes data in etcd with a user-provided encryption key or KMS service.
	ClusterFeatureEncryptionAtRest = "encryptionAtRest"
)

// ProtectedClusterLabels is a set of labels that must not be set on clusters manually by users,
// as they are relevant for the correct functioning of and security in KKP.
var ProtectedClusterLabels = sets.New(ProjectIDLabelKey, IsCredentialPresetLabelKey)

// +kubebuilder:validation:Enum=deleted;changed
type PresetInvalidationReason string

const (
	PresetInvalidationReasonDeleted PresetInvalidationReason = "deleted"
	PresetInvalidationReasonChanged PresetInvalidationReason = "changed"
)

// +genclient
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:object:generate=true
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:JSONPath=".spec.humanReadableName",name="HumanReadableName",type="string"
// +kubebuilder:printcolumn:JSONPath=".status.userEmail",name="Owner",type="string"
// +kubebuilder:printcolumn:JSONPath=".spec.version",name="Version",type="string"
// +kubebuilder:printcolumn:JSONPath=".spec.cloud.providerName",name="Provider",type="string"
// +kubebuilder:printcolumn:JSONPath=".spec.cloud.datacenter",name="Datacenter",type="string"
// +kubebuilder:printcolumn:JSONPath=".status.phase",name="Phase",type="string"
// +kubebuilder:printcolumn:JSONPath=".spec.pause",name="Paused",type="boolean"
// +kubebuilder:printcolumn:JSONPath=".metadata.creationTimestamp",name="Age",type="date"

// Cluster represents a Kubermatic Kubernetes Platform user cluster.
// Cluster objects exist on Seed clusters and each user cluster consists
// of a namespace containing the Kubernetes control plane and additional
// pods (like Prometheus or the machine-controller).
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec describes the desired cluster state.
	Spec ClusterSpec `json:"spec,omitempty"`

	// Status contains reconciliation information for the cluster.
	Status ClusterStatus `json:"status,omitempty"`
}

// IsEncryptionConfigurationEnabled returns whether encryption-at-rest is configured on this cluster.
func (cluster *Cluster) IsEncryptionEnabled() bool {
	return cluster.Spec.Features[ClusterFeatureEncryptionAtRest] && cluster.Spec.EncryptionConfiguration != nil && cluster.Spec.EncryptionConfiguration.Enabled
}

// ClusterSpec describes the desired state of a user cluster.
type ClusterSpec struct {
	kubermaticv1.ClusterSpec `json:",inline"`

	// Optional: Deploys the UserSSHKeyAgent to the user cluster. This field is immutable.
	// If enabled, the agent will be deployed and used to sync user ssh keys attached by users to the cluster.
	// No SSH keys will be synced after node creation if this is disabled.
	EnableUserSSHKeyAgent *bool `json:"enableUserSSHKeyAgent,omitempty"`

	// Optional: AuditLogging configures Kubernetes API audit logging (https://kubernetes.io/docs/tasks/debug-application-cluster/audit/)
	// for the user cluster.
	AuditLogging *AuditLoggingSettings `json:"auditLogging,omitempty"`

	// Optional: Configures encryption-at-rest for Kubernetes API data. This needs the `encryptionAtRest` feature gate.
	EncryptionConfiguration *EncryptionConfiguration `json:"encryptionConfiguration,omitempty"`
}

// EncryptionConfiguration configures encryption-at-rest for Kubernetes API data.
type EncryptionConfiguration struct {
	// Enables encryption-at-rest on this cluster.
	Enabled bool `json:"enabled"`

	// +kubebuilder:validation:MinItems=1

	// List of resources that will be stored encrypted in etcd.
	Resources []string `json:"resources"`
	// Configuration for the `secretbox` static key encryption scheme as supported by Kubernetes.
	// More info: https://kubernetes.io/docs/tasks/administer-cluster/encrypt-data/#providers
	Secretbox *SecretboxEncryptionConfiguration `json:"secretbox,omitempty"`
}

// SecretboxEncryptionConfiguration defines static key encryption based on the 'secretbox' solution for Kubernetes.
type SecretboxEncryptionConfiguration struct {
	// +kubebuilder:validation:MinItems=1

	// List of 'secretbox' encryption keys. The first element of this list is considered
	// the "primary" key which will be used for encrypting data while writing it. Additional
	// keys will be used for decrypting data while reading it, if keys higher in the list
	// did not succeed in decrypting it.
	Keys []SecretboxKey `json:"keys"`
}

// SecretboxKey stores a key or key reference for encrypting Kubernetes API data at rest with a static key.
type SecretboxKey struct {
	// Identifier of a key, used in various places to refer to the key.
	Name string `json:"name"`
	// Value contains a 32-byte random key that is base64 encoded. This is the key used
	// for encryption. Can be generated via `head -c 32 /dev/urandom | base64`, for example.
	Value string `json:"value,omitempty"`
	// Instead of passing the sensitive encryption key via the `value` field, a secret can be
	// referenced. The key of the secret referenced here needs to hold a key equivalent to the `value` field.
	SecretRef *corev1.SecretKeySelector `json:"secretRef,omitempty"`
}

// ClusterStatus stores status information about a cluster.
type ClusterStatus struct {
	kubermaticv1.ClusterStatus `json:",inline"`

	// Encryption describes the status of the encryption-at-rest feature for encrypted data in etcd.
	// +optional
	Encryption *ClusterEncryptionStatus `json:"encryption,omitempty"`

	// ResourceUsage shows the current usage of resources for the cluster.
	ResourceUsage *ResourceDetails `json:"resourceUsage,omitempty"`
}

// ClusterEncryptionStatus holds status information about the encryption-at-rest feature on the user cluster.
type ClusterEncryptionStatus struct {
	// The current "primary" key used to encrypt data written to etcd. Secondary keys that can be used for decryption
	// (but not encryption) might be configured in the ClusterSpec.
	ActiveKey string `json:"activeKey"`

	// List of resources currently encrypted.
	EncryptedResources []string `json:"encryptedResources"`

	// The current phase of the encryption process. Can be one of `Pending`, `Failed`, `Active` or `EncryptionNeeded`.
	// The `encryption_controller` logic will process the cluster based on the current phase and issue necessary changes
	// to make sure encryption on the cluster is active and updated with what the ClusterSpec defines.
	Phase ClusterEncryptionPhase `json:"phase,omitempty"`
}

// +kubebuilder:validation:Enum=Pending;Failed;Active;EncryptionNeeded
type ClusterEncryptionPhase string

const (
	ClusterEncryptionPhasePending          ClusterEncryptionPhase = "Pending"
	ClusterEncryptionPhaseFailed           ClusterEncryptionPhase = "Failed"
	ClusterEncryptionPhaseActive           ClusterEncryptionPhase = "Active"
	ClusterEncryptionPhaseEncryptionNeeded ClusterEncryptionPhase = "EncryptionNeeded"
)

// +kubebuilder:validation:Enum=HealthStatusDown;HealthStatusUp;HealthStatusProvisioning

type HealthStatus string

const (
	HealthStatusDown         HealthStatus = "HealthStatusDown"
	HealthStatusUp           HealthStatus = "HealthStatusUp"
	HealthStatusProvisioning HealthStatus = "HealthStatusProvisioning"
)

// ExtendedClusterHealth stores health information of a cluster.
type ExtendedClusterHealth struct {
	Apiserver                    HealthStatus  `json:"apiserver,omitempty"`
	Scheduler                    HealthStatus  `json:"scheduler,omitempty"`
	Controller                   HealthStatus  `json:"controller,omitempty"`
	MachineController            HealthStatus  `json:"machineController,omitempty"`
	Etcd                         HealthStatus  `json:"etcd,omitempty"`
	CloudProviderInfrastructure  HealthStatus  `json:"cloudProviderInfrastructure,omitempty"`
	UserClusterControllerManager HealthStatus  `json:"userClusterControllerManager,omitempty"`
	ApplicationController        HealthStatus  `json:"applicationController,omitempty"`
	OpenVPN                      *HealthStatus `json:"openvpn,omitempty"`
	Konnectivity                 *HealthStatus `json:"konnectivity,omitempty"`
	GatekeeperController         *HealthStatus `json:"gatekeeperController,omitempty"`
	GatekeeperAudit              *HealthStatus `json:"gatekeeperAudit,omitempty"`
	Monitoring                   *HealthStatus `json:"monitoring,omitempty"`
	Logging                      *HealthStatus `json:"logging,omitempty"`
	AlertmanagerConfig           *HealthStatus `json:"alertmanagerConfig,omitempty"`
	MLAGateway                   *HealthStatus `json:"mlaGateway,omitempty"`
	OperatingSystemManager       *HealthStatus `json:"operatingSystemManager,omitempty"`
	KubernetesDashboard          *HealthStatus `json:"kubernetesDashboard,omitempty"`
}

// ControlPlaneHealthy returns if all Kubernetes control plane components are healthy.
func (h *ExtendedClusterHealth) ControlPlaneHealthy() bool {
	return h.Etcd == HealthStatusUp &&
		h.Controller == HealthStatusUp &&
		h.Apiserver == HealthStatusUp &&
		h.Scheduler == HealthStatusUp
}

// AllHealthy returns true if all components are healthy. Gatekeeper components not included as they are optional and not
// crucial for cluster functioning.
func (h *ExtendedClusterHealth) AllHealthy() bool {
	return h.ControlPlaneHealthy() &&
		h.MachineController == HealthStatusUp &&
		h.CloudProviderInfrastructure == HealthStatusUp &&
		h.UserClusterControllerManager == HealthStatusUp
}

// ApplicationControllerHealthy checks for health of all essential components and the ApplicationController.
func (h *ExtendedClusterHealth) ApplicationControllerHealthy() bool {
	return h.AllHealthy() &&
		h.ApplicationController == HealthStatusUp
}

// +kubebuilder:object:generate=true
// +kubebuilder:object:root=true

// ClusterList specifies a list of user clusters.
type ClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Cluster `json:"items"`
}
