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

package machinecontroller

import (
	"k8s.io/apimachinery/pkg/util/sets"
)

// CloudProvider defines the cloud provider where the a cluster's nodes are running.
// Note that these constants may match KKP's constant, but don't have to.
// Use the functions in the helper package to translate between the two.
type CloudProvider string

const (
	CloudProviderFake                CloudProvider = "fake"
	CloudProviderAlibaba             CloudProvider = "alibaba"
	CloudProviderAnexia              CloudProvider = "anexia"
	CloudProviderAWS                 CloudProvider = "aws"
	CloudProviderAzure               CloudProvider = "azure"
	CloudProviderBaremetal           CloudProvider = "baremetal"
	CloudProviderDigitalocean        CloudProvider = "digitalocean"
	CloudProviderEquinixMetal        CloudProvider = "equinixmetal"
	CloudProviderExternal            CloudProvider = "external"
	CloudProviderGoogle              CloudProvider = "gce"
	CloudProviderHetzner             CloudProvider = "hetzner"
	CloudProviderKubeVirt            CloudProvider = "kubevirt"
	CloudProviderLinode              CloudProvider = "linode"
	CloudProviderNutanix             CloudProvider = "nutanix"
	CloudProviderOpenStack           CloudProvider = "openstack"
	CloudProviderPacket              CloudProvider = "packet"
	CloudProviderScaleway            CloudProvider = "scaleway"
	CloudProviderVMwareCloudDirector CloudProvider = "vmware-cloud-director"
	CloudProviderVSphere             CloudProvider = "vsphere"
	CloudProviderVultr               CloudProvider = "vultr"
)

var AllCloudProviders = sets.New(
	CloudProviderFake,
	CloudProviderAlibaba,
	CloudProviderAnexia,
	CloudProviderAWS,
	CloudProviderAzure,
	CloudProviderBaremetal,
	CloudProviderDigitalocean,
	CloudProviderEquinixMetal,
	CloudProviderExternal,
	CloudProviderGoogle,
	CloudProviderHetzner,
	CloudProviderKubeVirt,
	CloudProviderLinode,
	CloudProviderNutanix,
	CloudProviderOpenStack,
	CloudProviderPacket,
	CloudProviderScaleway,
	CloudProviderVMwareCloudDirector,
	CloudProviderVSphere,
	CloudProviderVultr,
)

// OperatingSystem defines the a node's operating system. Note that these constants may
// match KKP's constant, but don't have to. Use the functions in the helper package to
// translate between the two.
type OperatingSystem string

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
