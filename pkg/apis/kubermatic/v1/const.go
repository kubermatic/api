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
	"k8s.io/apimachinery/pkg/util/sets"
)

// +kubebuilder:validation:Enum=alibaba;anexia;aws;azure;bringyourown;digitalocean;gcp;hetzner;kubevirt;nutanix;openstack;packet;vmwareclouddirector;vsphere

// CloudProvider defines the cloud provider where the a cluster's nodes are running.
// Note that these constants may match the machine-controller's constant, but don't
// have to. Use the functions in the helper package to translate between the two.
type CloudProvider string

func (p CloudProvider) String() string {
	return string(p)
}

const (
	CloudProviderFake                CloudProvider = "fake"
	CloudProviderAlibaba             CloudProvider = "alibaba"
	CloudProviderAnexia              CloudProvider = "anexia"
	CloudProviderAWS                 CloudProvider = "aws"
	CloudProviderAzure               CloudProvider = "azure"
	CloudProviderBringYourOwn        CloudProvider = "bringyourown"
	CloudProviderDigitalocean        CloudProvider = "digitalocean"
	CloudProviderGCP                 CloudProvider = "gcp"
	CloudProviderHetzner             CloudProvider = "hetzner"
	CloudProviderKubeVirt            CloudProvider = "kubevirt"
	CloudProviderNutanix             CloudProvider = "nutanix"
	CloudProviderOpenStack           CloudProvider = "openstack"
	CloudProviderPacket              CloudProvider = "packet"
	CloudProviderVMwareCloudDirector CloudProvider = "vmwareclouddirector"
	CloudProviderVSphere             CloudProvider = "vsphere"
)

var AllCloudProviders = sets.New(
	CloudProviderFake,
	CloudProviderAlibaba,
	CloudProviderAnexia,
	CloudProviderAWS,
	CloudProviderAzure,
	CloudProviderBringYourOwn,
	CloudProviderDigitalocean,
	CloudProviderGCP,
	CloudProviderHetzner,
	CloudProviderKubeVirt,
	CloudProviderNutanix,
	CloudProviderOpenStack,
	CloudProviderPacket,
	CloudProviderVMwareCloudDirector,
	CloudProviderVSphere,
)

// +kubebuilder:validation:Enum=aks;eks;gke;bringyourown;kubeone

type ExternalClusterProvider string

func (p ExternalClusterProvider) String() string {
	return string(p)
}

const (
	ExternalClusterProviderAKS          ExternalClusterProvider = "aks"
	ExternalClusterProviderEKS          ExternalClusterProvider = "eks"
	ExternalClusterProviderGKE          ExternalClusterProvider = "gke"
	ExternalClusterProviderBringYourOwn ExternalClusterProvider = "bringyourown"
	ExternalClusterProviderKubeOne      ExternalClusterProvider = "kubeone"
)

var AllExternalClusterProviders = sets.New(
	ExternalClusterProviderAKS,
	ExternalClusterProviderEKS,
	ExternalClusterProviderGKE,
	ExternalClusterProviderBringYourOwn,
	ExternalClusterProviderKubeOne,
)

// +kubebuilder:validation:Enum=ubuntu;centos;amzn2;rhel;flatcar;rockylinux

// OperatingSystem defines the a node's operating system. Note that these constants may
// match the machine-controller's constant, but don't have to. Use the functions in
// the helper package to translate between the two.
type OperatingSystem string

func (o OperatingSystem) String() string {
	return string(o)
}

const (
	OperatingSystemUbuntu       OperatingSystem = "ubuntu"
	OperatingSystemCentOS       OperatingSystem = "centos"
	OperatingSystemAmazonLinux2 OperatingSystem = "amzn2"
	OperatingSystemRHEL         OperatingSystem = "rhel"
	OperatingSystemFlatcar      OperatingSystem = "flatcar"
	OperatingSystemRockyLinux   OperatingSystem = "rockylinux"
)

var AllOperatingSystems = sets.New(
	OperatingSystemUbuntu,
	OperatingSystemCentOS,
	OperatingSystemAmazonLinux2,
	OperatingSystemRHEL,
	OperatingSystemFlatcar,
	OperatingSystemRockyLinux,
)

// +kubebuilder:validation:Enum=NodePort;LoadBalancer;Tunneling

// ExposeStrategy is the strategy used to expose a cluster control plane.
// Possible values are `NodePort`, `LoadBalancer` or `Tunneling` (requires a feature gate).
type ExposeStrategy string

func (s ExposeStrategy) String() string {
	return string(s)
}

const (
	// ExposeStrategyNodePort creates a NodePort with a "nodeport-proxy.k8s.io/expose": "true" annotation to expose
	// all clusters on one central Service of type LoadBalancer via the NodePort proxy.
	ExposeStrategyNodePort ExposeStrategy = "NodePort"
	// ExposeStrategyLoadBalancer creates a LoadBalancer service per cluster.
	ExposeStrategyLoadBalancer ExposeStrategy = "LoadBalancer"
	// ExposeStrategyTunneling allows to reach the control plane components by
	// tunneling L4 traffic (TCP only is supported at the moment).
	// The traffic is encapsulated by mean of an agent deployed on the worker
	// nodes.
	// The traffic is decapsulated and forwarded to the recipients by
	// mean of a proxy deployed on the Seed Cluster.
	// The same proxy is also capable of routing TLS traffic without
	// termination, this is to allow the Kubelet to reach the control plane
	// before the agents are running.
	//
	// This strategy has the inconvenience of requiring an agent on worker
	// nodes, but has the notable advantage of requiring one single entry point
	// (e.g. Service of type LoadBalancer) without consuming one or more ports
	// for each user cluster.
	ExposeStrategyTunneling ExposeStrategy = "Tunneling"
)

// AllExposeStrategies is a set containing all the expose strategies.
var AllExposeStrategies = sets.New(
	ExposeStrategyNodePort,
	ExposeStrategyLoadBalancer,
	ExposeStrategyTunneling,
)

// +kubebuilder:validation:Enum=canal;cilium;none

// CNIPluginType defines the type of CNI plugin installed.
type CNIPluginType string

func (c CNIPluginType) String() string {
	return string(c)
}

const (
	// CNIPluginTypeCanal corresponds to Canal CNI plugin (i.e. Flannel +
	// Calico for policy enforcement).
	CNIPluginTypeCanal CNIPluginType = "canal"

	// CNIPluginTypeCilium corresponds to Cilium CNI plugin.
	CNIPluginTypeCilium CNIPluginType = "cilium"

	// CNIPluginTypeNone corresponds to no CNI plugin managed by KKP
	// (cluster users are responsible for managing the CNI in the cluster themselves).
	CNIPluginTypeNone CNIPluginType = "none"
)
