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
// +kubebuilder:printcolumn:JSONPath=".spec.email",name="Email",type="string"
// +kubebuilder:printcolumn:JSONPath=".spec.name",name="HumanReadableName",type="string"
// +kubebuilder:printcolumn:JSONPath=".spec.admin",name="Admin",type="boolean"
// +kubebuilder:printcolumn:JSONPath=".metadata.creationTimestamp",name="Age",type="date"

// User specifies a KKP user.
type User struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   UserSpec   `json:"spec,omitempty"`
	Status UserStatus `json:"status,omitempty"`
}

// UserStatus stores status information about a user.
type UserStatus struct {
	// LastSeen is the date and time this user last used the KKP dashboard.
	LastSeen metav1.Time `json:"lastSeen,omitempty"`
}

// UserSpec specifies a user.
type UserSpec struct {
	// Name is the full name of this user.
	Name string `json:"name"`
	// Email is the email address of this user. Emails must be globally unique across all KKP users.
	Email string `json:"email"`
	// IsAdmin defines whether this user is an administrator with additional permissions.
	// +kubebuilder:default=false
	IsAdmin bool `json:"admin"`
	// Groups holds the information to which groups the user belongs to. Set automatically when logging in to the
	// KKP API, and used by the KKP API.
	Groups []string `json:"groups,omitempty"`

	// Settings contains both user-configurable and system-owned configuration for the KKP dashboard.
	DashboardSettings *UserDashboardSettings `json:"dashboardSettings,omitempty"`

	// InvalidTokensReference is a reference to a Secret that contains invalidated
	// login tokens. The tokens are used to provide a safe logout mechanism.
	InvalidTokensReference *GlobalSecretKeySelector `json:"invalidTokensReference,omitempty"`
}

// UserDashboardSettings represent the settings for a user within the KKP dashboard.
type UserDashboardSettings struct {
	SelectedTheme            string `json:"selectedTheme,omitempty"`
	ItemsPerPage             int8   `json:"itemsPerPage,omitempty"`
	SelectProjectTableView   bool   `json:"selectProjectTableView,omitempty"`
	CollapseSidenav          bool   `json:"collapseSidenav,omitempty"`
	LastSeenChangelogVersion string `json:"lastSeenChangelogVersion,omitempty"`
	UseClustersView          bool   `json:"useClustersView,omitempty"`
}

// +kubebuilder:object:generate=true
// +kubebuilder:object:root=true

// UserList is a list of users.
type UserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []User `json:"items"`
}
