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
)

// +kubebuilder:validation:Enum=Healthy;Unhealthy;Invalid;Terminating;Paused

type SeedPhase string

// These are the valid phases of a seed.
const (
	// SeedPhaseHealthy means the seed is reachable and was successfully reconciled.
	SeedPhaseHealthy SeedPhase = "Healthy"

	// SeedPhaseUnhealthy means the KKP resources on the seed cluster could not be
	// successfully reconciled.
	SeedPhaseUnhealthy SeedPhase = "Unhealthy"

	// SeedPhaseInvalid means the seed kubeconfig is defunct.
	SeedPhaseInvalid SeedPhase = "Invalid"

	// SeedPhaseTerminating means the seed is currently being deleted.
	SeedPhaseTerminating SeedPhase = "Terminating"

	// SeedPhasePaused means the seed is not being reconciled because the SkipReconciling
	// annotation is set.
	SeedPhasePaused SeedPhase = "Paused"
)

// +genclient
// +kubebuilder:object:generate=true
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:JSONPath=".status.clusters",name="Clusters",type="integer"
// +kubebuilder:printcolumn:JSONPath=".spec.location",name="Location",type="string"
// +kubebuilder:printcolumn:JSONPath=".status.versions.kubermatic",name="KKP Version",type="string"
// +kubebuilder:printcolumn:JSONPath=".status.versions.cluster",name="Cluster Version",type="string"
// +kubebuilder:printcolumn:JSONPath=".status.phase",name="Phase",type="string"
// +kubebuilder:printcolumn:JSONPath=".metadata.creationTimestamp",name="Age",type="date"

// Seed is the type representing a Seed cluster. Seed clusters host the the control planes
// for KKP user clusters. Seedlets are responsible for registering a seed cluster in the
// KKP management system, similar to how a kubelet registers a node.
type Seed struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec SeedSpec `json:"spec"`

	//nolint:staticcheck
	//lint:ignore SA5008 omitgenyaml is used by the example-yaml-generator
	Status SeedStatus `json:"status,omitempty,omitgenyaml"`
}

// The spec for a seed cluster.
type SeedSpec struct {
	// Optional: Country of the seed as ISO-3166 two-letter code, e.g. DE or UK.
	// For informational purposes in the Kubermatic dashboard only.
	Country string `json:"country,omitempty"`
	// Optional: Detailed location of the cluster, like "Hamburg" or "Datacenter 7".
	// For informational purposes in the Kubermatic dashboard only.
	Location string `json:"location,omitempty"`
	// Datacenters contains a map of all datacenters (DCs) on this seed. The datacenter
	// names are not globally unique, i.e. two seeds can both have a "test" datacenter.
	Datacenters map[string]kubermaticv1.DatacenterSpec `json:"datacenters,omitempty"`
}

// SeedStatus contains runtime information regarding the seed.
type SeedStatus struct {
	// Phase contains a human readable text to indicate the seed cluster status. No logic should be tied
	// to this field, as its content can change in between KKP releases.
	Phase SeedPhase `json:"phase,omitempty"`

	// +kubebuilder:default=0
	// +kubebuilder:validation:Minimum:=0

	// Clusters is the total number of user clusters that exist on this seed, the sum across all its
	// datacenters.
	Clusters int `json:"clusters"`

	// Versions contains information regarding versions of components in the cluster and the cluster
	// itself.
	// +optional
	Versions SeedVersionsStatus `json:"versions,omitempty"`

	// Datacenters contains a map of all datacenter statuses on this seed.
	Datacenters map[string]kubermaticv1.DatacenterStatus `json:"datacenters,omitempty"`

	// Conditions contains conditions the seed is in, its primary use case is status signaling
	// between controllers or between controllers and the API.
	// +optional
	Conditions map[SeedConditionType]SeedCondition `json:"conditions,omitempty"`
}

// +kubebuilder:validation:Enum=KubeconfigValid;ResourcesReconciled;ClusterInitialized

// SeedConditionType is used to indicate the type of a seed condition. For all condition
// types, the `true` value must indicate success.
type SeedConditionType string

const (
	// SeedConditionKubeconfigValid indicates that the configured kubeconfig for the seed is valid.
	// The seed-sync controller manages this condition.
	SeedConditionKubeconfigValid SeedConditionType = "KubeconfigValid"
	// SeedConditionResourcesReconciled indicates that the KKP operator has finished setting up the
	// resources inside the seed cluster.
	SeedConditionResourcesReconciled SeedConditionType = "ResourcesReconciled"
	// SeedConditionClusterInitialized indicates that the KKP operator has finished setting up the
	// CRDs and other prerequisites on the Seed cluster. After this condition is true, other
	// controllers can begin to create watches and reconcile resources (i.e. this condition is
	// a precondition to ResourcesReconciled). Once this condition is true, it is never set to false
	// again.
	SeedConditionClusterInitialized SeedConditionType = "ClusterInitialized"
)

type SeedCondition struct {
	// Status of the condition, one of True, False, Unknown.
	Status corev1.ConditionStatus `json:"status"`
	// Last time we got an update on a given condition.
	LastHeartbeatTime metav1.Time `json:"lastHeartbeatTime"`
	// Last time the condition transit from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
	// (brief) reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty"`
	// Human readable message indicating details about last transition.
	// +optional
	Message string `json:"message,omitempty"`
}

// SeedVersionsStatus contains information regarding versions of components in the cluster
// and the cluster itself.
type SeedVersionsStatus struct {
	// Kubermatic is the version of the currently deployed KKP components. Note that a permanent
	// version skew between master and seed is not supported and KKP setups should never run for
	// longer times with a skew between the clusters.
	Kubermatic string `json:"kubermatic,omitempty"`
	// Cluster is the Kubernetes version of the cluster's control plane.
	Cluster string `json:"cluster,omitempty"`
}

// +kubebuilder:object:generate=true
// +kubebuilder:object:root=true

// SeedList is the type representing a SeedList.
type SeedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	// List of seeds
	Items []Seed `json:"items"`
}
