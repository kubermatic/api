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
// +kubebuilder:printcolumn:JSONPath=".keyRef.name",name="Key",type="string"
// +kubebuilder:printcolumn:JSONPath=".clusterRef.name",name="Cluster",type="string"
// +kubebuilder:printcolumn:JSONPath=".metadata.creationTimestamp",name="Age",type="date"

// SSHKeyBinding sets the relationship between a UserSSHKey and a Cluster.
type SSHKeyBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	SSHKeyRef  corev1.LocalObjectReference `json:"sshKeyRef"`
	ClusterRef corev1.LocalObjectReference `json:"clusterRef"`
}

// +kubebuilder:object:generate=true
// +kubebuilder:object:root=true

// SSHKeyBindingList specifies a users SSHKeyBinding.
type SSHKeyBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []SSHKeyBinding `json:"items"`
}
