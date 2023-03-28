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

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +kubebuilder:object:generate=true
// +kubebuilder:object:root=true
// +kubebuilder:printcolumn:JSONPath=".metadata.creationTimestamp",name="Age",type="date"

// MLAAdminSetting is the object representing cluster-specific administrator settings
// for KKP user cluster MLA (monitoring, logging & alerting) stack.
type MLAAdminSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec kubermaticv1.MLAAdminSettingSpec `json:"spec,omitempty"`
}

// +kubebuilder:object:generate=true
// +kubebuilder:object:root=true

// MLAAdminSettingList specifies a list of administrtor settings for KKP
// user cluster MLA (monitoring, logging & alerting) stack.
type MLAAdminSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []MLAAdminSetting `json:"items"`
}
