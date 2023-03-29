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
	kubermaticappsv1 "k8c.io/api/v3/pkg/apis/apps.kubermatic/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=appinstall

// ApplicationInstallation describes a single installation of an Application.
type ApplicationInstallation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   kubermaticappsv1.ApplicationInstallationSpec   `json:"spec,omitempty"`
	Status kubermaticappsv1.ApplicationInstallationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationInstallationList is a list of ApplicationInstallations.
type ApplicationInstallationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ApplicationInstallation `json:"items"`
}
