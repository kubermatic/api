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

// +kubebuilder:validation:Pattern:=`^((\d{1,3}\.){3}\d{1,3}\/([0-9]|[1-2][0-9]|3[0-2]))$`
type CIDR string

// +kubebuilder:validation:Pattern="((^((([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5]))/([0-9]|[1-2][0-9]|3[0-2])$)|(^(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:))/([0-9]|[0-9][0-9]|1[0-1][0-9]|12[0-8])$))"

// SubnetCIDR is used to store IPv4/IPv6 CIDR.
type SubnetCIDR string

// Finalizers should be kept to their controllers. Only if a finalizer is
// used by multiple controllers should it be placed here.

const (
	// NodeDeletionFinalizer indicates that the nodes still need cleanup.
	NodeDeletionFinalizer = "kubermatic.k8c.io/delete-nodes"
	// NamespaceCleanupFinalizer indicates that the cluster namespace still exists and the owning Cluster object
	// must not yet be deleted.
	NamespaceCleanupFinalizer = "kubermatic.k8c.io/cleanup-namespace"
	// InClusterPVCleanupFinalizer indicates that the PVs still need cleanup.
	InClusterPVCleanupFinalizer = "kubermatic.k8c.io/cleanup-in-cluster-pv"
	// InClusterLBCleanupFinalizer indicates that the LBs still need cleanup.
	InClusterLBCleanupFinalizer = "kubermatic.k8c.io/cleanup-in-cluster-lb"
	// CredentialsSecretsCleanupFinalizer indicates that secrets for credentials still need cleanup.
	CredentialsSecretsCleanupFinalizer = "kubermatic.k8c.io/cleanup-credentials-secrets"
	// ExternalClusterKubeOneNamespaceCleanupFinalizer indicates that kubeone cluster namespace still need cleanup.
	ExternalClusterKubeOneNamespaceCleanupFinalizer = "kubermatic.k8c.io/cleanup-kubeone-namespace"
	// ExternalClusterKubeconfigCleanupFinalizer indicates that secrets for kubeconfig still need cleanup.
	ExternalClusterKubeconfigCleanupFinalizer = "kubermatic.k8c.io/cleanup-kubeconfig-secret"
	// ExternalClusterKubeOneCleanupFinalizer indicates that secrets for kubeone cluster still need cleanup.
	ExternalClusterKubeOneSecretsCleanupFinalizer = "kubermatic.k8c.io/cleanup-kubeone-secret"
	// EtcdBackConfigCleanupFinalizer indicates that EtcdBackupConfigs for the cluster still need cleanup.
	EtcdBackupConfigCleanupFinalizer = "kubermatic.k8c.io/cleanup-etcdbackupconfigs"
	// GatekeeperConstraintCleanupFinalizer indicates that gatkeeper constraints on the user cluster need cleanup.
	GatekeeperConstraintCleanupFinalizer = "kubermatic.k8c.io/cleanup-gatekeeper-constraints"
	// KubermaticConstraintCleanupFinalizer indicates that Kubermatic constraints for the cluster need cleanup.
	KubermaticConstraintCleanupFinalizer = "kubermatic.k8c.io/cleanup-kubermatic-constraints"
)

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

// ProxySettings allow configuring a HTTP proxy for the control planes and nodes.
type ProxySettings struct {
	// Optional: If set, this proxy will be configured for both HTTP and HTTPS.
	HTTPProxy *string `json:"httpProxy,omitempty"`
	// Optional: If set this will be set as NO_PROXY environment variable on the node;
	// The value must be a comma-separated list of domains for which no proxy
	// should be used, e.g. "*.example.com,internal.dev".
	// Note that the in-cluster apiserver URL will be automatically prepended
	// to this value.
	NoProxy *string `json:"noProxy,omitempty"`
}

func emptyStrPtr(s *string) bool {
	return s == nil || *s == ""
}

// Empty returns true if p or all of its children are nil or empty strings.
func (p *ProxySettings) Empty() bool {
	return p == nil || (emptyStrPtr(p.HTTPProxy) && emptyStrPtr(p.NoProxy))
}

// Merge applies the settings from p into dst if the corresponding setting
// in dst is nil or an empty string.
func (p *ProxySettings) Merge(dst *ProxySettings) {
	if emptyStrPtr(dst.HTTPProxy) {
		dst.HTTPProxy = p.HTTPProxy
	}
	if emptyStrPtr(dst.NoProxy) {
		dst.NoProxy = p.NoProxy
	}
}

// ClusterReference is a struct that allows referencing a single Cluster object.
type ClusterReference struct {
	// Name of the Cluster object.
	Name string `json:"name"`
}
