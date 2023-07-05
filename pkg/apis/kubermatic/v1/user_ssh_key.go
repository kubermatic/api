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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:object:generate=true
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:JSONPath=".spec.name",name="HumanReadableName",type="string"
// +kubebuilder:printcolumn:JSONPath=".status.fingerprint",name="Fingerprint",type="string"
// +kubebuilder:printcolumn:JSONPath=".metadata.creationTimestamp",name="Age",type="date"

// UserSSHKey specifies a users UserSSHKey.
type UserSSHKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UserSSHKeySpec `json:"spec,omitempty"`
	Status SSHKeyStatus   `json:"status,omitempty"`
}

type UserSSHKeySpec struct {
	// Name is the human readable name for this SSH key.
	Name string `json:"name"`
	// PublicKey is the SSH public key.
	PublicKey string `json:"publicKey"`
}

type SSHKeyStatus struct {
	// Fingerprint is calculated server-side based on the supplied public key.
	Fingerprint string `json:"fingerprint"`
}

// +kubebuilder:object:generate=true
// +kubebuilder:object:root=true

// UserSSHKeyList specifies a users UserSSHKey.
type UserSSHKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []UserSSHKey `json:"items"`
}
