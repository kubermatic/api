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

// +kubebuilder:validation:Enum=prefix;range

// IPAMPoolAllocationType defines the type of allocation to be used.
type IPAMPoolAllocationType string

const (
	// IPAMPoolAllocationTypePrefix corresponds to prefix allocation type.
	IPAMPoolAllocationTypePrefix IPAMPoolAllocationType = "prefix"
	// IPAMPoolAllocationTypeRange corresponds to range allocation type.
	IPAMPoolAllocationTypeRange IPAMPoolAllocationType = "range"
)

// +genclient
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:object:generate=true
// +kubebuilder:object:root=true
// +kubebuilder:printcolumn:JSONPath=".metadata.creationTimestamp",name="Age",type="date"

// IPAMPool is the object representing Multi-Cluster IP Address Management (IPAM)
// configuration for KKP user clusters.
type IPAMPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec IPAMPoolSpec `json:"spec,omitempty"`
}

// IPAMPoolSpec specifies the  Multi-Cluster IP Address Management (IPAM)
// configuration for KKP user clusters.
type IPAMPoolSpec struct {
	// Datacenters contains a map of datacenters (DCs) for the allocation.
	Datacenters map[string]IPAMPoolDatacenterSettings `json:"datacenters"`
}

// IPAMPoolDatacenterSettings contains IPAM Pool configuration for a datacenter.
type IPAMPoolDatacenterSettings struct {
	// Type is the allocation type to be used.
	Type IPAMPoolAllocationType `json:"type"`

	// PoolCIDR is the pool CIDR to be used for the allocation.
	PoolCIDR SubnetCIDR `json:"poolCidr"`

	// +kubebuilder:validation:Minimum:=1
	// +kubebuilder:validation:Maximum:=128
	// AllocationPrefix is the prefix for the allocation.
	// Used when "type=prefix".
	AllocationPrefix int `json:"allocationPrefix,omitempty"`

	// Optional: ExcludePrefixes is used to exclude particular subnets for the allocation.
	// NOTE: must be the same length as allocationPrefix.
	// Can be used when "type=prefix".
	ExcludePrefixes []SubnetCIDR `json:"excludePrefixes,omitempty"`

	// +kubebuilder:validation:Minimum:=1
	// AllocationRange is the range for the allocation.
	// Used when "type=range".
	AllocationRange int `json:"allocationRange,omitempty"`

	// Optional: ExcludeRanges is used to exclude particular IPs or IP ranges for the allocation.
	// Examples: "192.168.1.100-192.168.1.110", "192.168.1.255".
	// Can be used when "type=range".
	ExcludeRanges []string `json:"excludeRanges,omitempty"`
}

// +kubebuilder:object:generate=true
// +kubebuilder:object:root=true

type IPAMPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []IPAMPool `json:"items"`
}
