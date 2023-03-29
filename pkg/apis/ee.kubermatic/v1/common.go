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
	"fmt"
	"net"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	netutils "k8s.io/utils/net"
)

const (
	// AuthZRoleLabel is the label used by rbac-controller and group-rbac-controller to identify the KKP role a ClusterRole or Role were created for.
	AuthZRoleLabel = "authz.k8c.io/role"

	// AuthZGroupProjectBindingLabel references the GroupProjectBinding resource that a ClusterRole/Role was created for.
	AuthZGroupProjectBindingLabel = "authz.k8c.io/group-project-binding"
)

// +kubebuilder:validation:Pattern:=`^((\d{1,3}\.){3}\d{1,3}\/([0-9]|[1-2][0-9]|3[0-2]))$`
type CIDR string

// +kubebuilder:validation:Pattern="((^((([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]))/([0-9]|[1-2][0-9]|3[0-2])$)|(^(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))/([0-9]|[0-9][0-9]|1[0-1][0-9]|12[0-8])$))"

// SubnetCIDR is used to store IPv4/IPv6 CIDR.
type SubnetCIDR string

const (
	InitialMachineDeploymentRequestAnnotation        = "kubermatic.io/initial-machinedeployment-request"
	InitialApplicationInstallationsRequestAnnotation = "kubermatic.io/initial-application-installations-request"
	InitialCNIValuesRequestAnnotation                = "kubermatic.io/initial-cni-values-request"
)

type MachineFlavorFilter struct {
	// +kubebuilder:default=0
	// +kubebuilder:validation:Minimum:=0

	// Minimum number of vCPU
	MinCPU int `json:"minCPU"`

	// +kubebuilder:default=0
	// +kubebuilder:validation:Minimum:=0

	// Maximum number of vCPU
	MaxCPU int `json:"maxCPU"`

	// +kubebuilder:default=0
	// +kubebuilder:validation:Minimum:=0

	// Minimum RAM size in GB
	MinRAM int `json:"minRAM"`

	// +kubebuilder:default=0
	// +kubebuilder:validation:Minimum:=0

	// Maximum RAM size in GB
	MaxRAM int `json:"maxRAM"`

	// Include VMs with GPU
	EnableGPU bool `json:"enableGPU"` //nolint:tagliatelle
}

// NetworkRanges represents ranges of network addresses.
type NetworkRanges struct {
	CIDRBlocks []string `json:"cidrBlocks,omitempty"`
}

// Validate validates the network ranges. Returns nil if valid, error otherwise.
func (r *NetworkRanges) Validate() error {
	if r == nil {
		return nil
	}

	for _, cidr := range r.CIDRBlocks {
		if _, _, err := net.ParseCIDR(cidr); err != nil {
			return fmt.Errorf("unable to parse CIDR %q: %w", cidr, err)
		}
	}

	return nil
}

// GetIPv4CIDR returns the first found IPv4 CIDR in the network ranges, or an empty string if no IPv4 CIDR is found.
func (r *NetworkRanges) GetIPv4CIDR() string {
	for _, cidr := range r.CIDRBlocks {
		if netutils.IsIPv4CIDRString(cidr) {
			return cidr
		}
	}

	return ""
}

// GetIPv4CIDRs returns all IPv4 CIDRs in the network ranges, or an empty string if no IPv4 CIDR is found.
func (r *NetworkRanges) GetIPv4CIDRs() (res []string) {
	for _, cidr := range r.CIDRBlocks {
		if netutils.IsIPv4CIDRString(cidr) {
			res = append(res, cidr)
		}
	}

	return
}

// HasIPv4CIDR returns true if the network ranges contain any IPv4 CIDR, false otherwise.
func (r *NetworkRanges) HasIPv4CIDR() bool {
	return r.GetIPv4CIDR() != ""
}

// GetIPv6CIDR returns the first found IPv6 CIDR in the network ranges, or an empty string if no IPv6 CIDR is found.
func (r *NetworkRanges) GetIPv6CIDR() string {
	for _, cidr := range r.CIDRBlocks {
		if netutils.IsIPv6CIDRString(cidr) {
			return cidr
		}
	}

	return ""
}

// GetIPv6CIDRs returns all IPv6 CIDRs in the network ranges, or an empty string if no IPv6 CIDR is found.
func (r *NetworkRanges) GetIPv6CIDRs() (res []string) {
	for _, cidr := range r.CIDRBlocks {
		if netutils.IsIPv6CIDRString(cidr) {
			res = append(res, cidr)
		}
	}

	return
}

// HasIPv6CIDR returns true if the network ranges contain any IPv6 CIDR, false otherwise.
func (r *NetworkRanges) HasIPv6CIDR() bool {
	return r.GetIPv6CIDR() != ""
}

// ResourceDetails holds the CPU, Memory and Storage quantities.
type ResourceDetails struct {
	// CPU holds the quantity of CPU. For the format, please check k8s.io/apimachinery/pkg/api/resource.Quantity.
	CPU *resource.Quantity `json:"cpu,omitempty"`
	// Memory represents the quantity of RAM size. For the format, please check k8s.io/apimachinery/pkg/api/resource.Quantity.
	Memory *resource.Quantity `json:"memory,omitempty"`
	// Storage represents the disk size. For the format, please check k8s.io/apimachinery/pkg/api/resource.Quantity.
	Storage *resource.Quantity `json:"storage,omitempty"`
}

func emptyQuantity(q *resource.Quantity) bool {
	return q == nil || q.IsZero()
}

func (r *ResourceDetails) IsEmpty() bool {
	return r == nil || (emptyQuantity(r.CPU) && emptyQuantity(r.Memory) && emptyQuantity(r.Storage))
}

// GlobalObjectKeySelector is needed as we can not use v1.SecretKeySelector
// because it is not cross namespace.
type GlobalObjectKeySelector struct {
	corev1.ObjectReference `json:",inline"`
	Key                    string `json:"key,omitempty"`
}

type GlobalSecretKeySelector GlobalObjectKeySelector
type GlobalConfigMapKeySelector GlobalObjectKeySelector