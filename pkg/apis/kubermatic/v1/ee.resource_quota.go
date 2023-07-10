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

const (
	ResourceQuotaSubjectNameLabelKey = "subject-name"
	ResourceQuotaSubjectKindLabelKey = "subject-kind"
)

// +kubebuilder:validation:Enum=project

type ResourceQuotaSubjectName string

const (
	ResourceQuotaSubjectProject ResourceQuotaSubjectName = "project"
)

// +genclient
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:resource:categories=kkpee
// +kubebuilder:object:generate=true
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:JSONPath=".spec.cluster.name",name="Cluster",type="string"
// +kubebuilder:printcolumn:JSONPath=".spec.subject.name",name="Subject Name",type="string"
// +kubebuilder:printcolumn:JSONPath=".spec.subject.kind",name="Subject Kind",type="string"
// +kubebuilder:printcolumn:JSONPath=".metadata.creationTimestamp",name="Age",type="date"

// ResourceQuota specifies the amount of cluster resources a project can use.
//
// Note that this resource is part of a KKP Enterprise feature and is not used in the Community Edition.
type ResourceQuota struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ResourceQuotaSpec   `json:"spec,omitempty"`
	Status ResourceQuotaStatus `json:"status,omitempty"`
}

// ResourceQuotaSpec describes the desired state of a resource quota.
type ResourceQuotaSpec struct {
	// Subject specifies to which entity the quota applies to.
	Subject ResourceQuotaSubject `json:"subject"`
	// Quota specifies the current maximum allowed usage of resources.
	Quota ResourceDetails `json:"quota"`
}

// ResourceQuotaStatus describes the current state of a resource quota.
type ResourceQuotaStatus struct {
	// GlobalUsage is holds the current usage of resources for all seeds.
	GlobalUsage *ResourceDetails `json:"globalUsage,omitempty"`
	// LocalUsage is holds the current usage of resources for the local seed.
	LocalUsage *ResourceDetails `json:"localUsage,omitempty"`
}

// ResourceQuotaSubject describes the entity to which the quota applies to.
type ResourceQuotaSubject struct {
	// Name of the quota subject.
	Name string `json:"name"`

	// +kubebuilder:default=project

	// Kind of the quota subject.
	Kind ResourceQuotaSubjectName `json:"kind"`
}

// +kubebuilder:object:generate=true
// +kubebuilder:object:root=true

// ResourceQuotaList is a collection of resource quotas.
type ResourceQuotaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []ResourceQuota `json:"items"`
}
