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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:object:generate=true
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:JSONPath=".spec.cluster.name",name="Cluster",type="string"
// +kubebuilder:printcolumn:JSONPath=".metadata.creationTimestamp",name="Age",type="date"

type Alertmanager struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AlertmanagerSpec   `json:"spec,omitempty"`
	Status AlertmanagerStatus `json:"status,omitempty"`
}

type AlertmanagerSpec struct {
	// Cluster is the reference to the cluster that this Alertmanager belongs to.
	Cluster ClusterReference `json:"cluster"`
	// ConfigSecret refers to the Secret in the same namespace as the Alertmanager object,
	// which contains configuration for this Alertmanager.
	ConfigSecret corev1.LocalObjectReference `json:"configSecret"`
}

// AlertmanagerStatus stores status information about the AlertManager.
type AlertmanagerStatus struct {
	ConfigStatus AlertmanagerConfigurationStatus `json:"configStatus,omitempty"`
}

// AlertmanagerConfigurationStatus stores status information about the AlertManager configuration.
type AlertmanagerConfigurationStatus struct {
	// LastUpdated stores the last successful time when the configuration was successfully applied
	LastUpdated metav1.Time `json:"lastUpdated,omitempty"`
	// Status of whether the configuration was applied, one of True, False
	Status corev1.ConditionStatus `json:"status"`
	// ErrorMessage contains a default error message in case the configuration could not be applied.
	// Will be reset if the error was resolved and condition becomes True
	ErrorMessage string `json:"errorMessage,omitempty"`
}

// +kubebuilder:object:generate=true
// +kubebuilder:object:root=true

type AlertmanagerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Alertmanager `json:"items"`
}
