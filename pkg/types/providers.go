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

package types

import (
	"strings"
)

type CloudProvider string

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
	CloudProviderKubevirt            CloudProvider = "kubevirt"
	CloudProviderNutanix             CloudProvider = "nutanix"
	CloudProviderOpenStack           CloudProvider = "openstack"
	CloudProviderPacket              CloudProvider = "packet"
	CloudProviderVMwareCloudDirector CloudProvider = "vmwareclouddirector"
	CloudProviderVSphere             CloudProvider = "vsphere"
)

var (
	SupportedProviders = []CloudProvider{
		CloudProviderFake,
		CloudProviderAlibaba,
		CloudProviderAnexia,
		CloudProviderAWS,
		CloudProviderAzure,
		CloudProviderBringYourOwn,
		CloudProviderDigitalocean,
		CloudProviderGCP,
		CloudProviderHetzner,
		CloudProviderKubevirt,
		CloudProviderNutanix,
		CloudProviderOpenStack,
		CloudProviderPacket,
		CloudProviderVMwareCloudDirector,
		CloudProviderVSphere,
	}
)

func IsProviderSupported(name string) bool {
	for _, provider := range SupportedProviders {
		if strings.EqualFold(name, string(provider)) {
			return true
		}
	}

	return false
}

type ExternalClusterProvider string

const (
	ExternalClusterProviderAKS          ExternalClusterProvider = "aks"
	ExternalClusterProviderEKS          ExternalClusterProvider = "eks"
	ExternalClusterProviderGKE          ExternalClusterProvider = "gke"
	ExternalClusterProviderBringYourOwn ExternalClusterProvider = "bringyourown"
	ExternalClusterProviderKubeOne      ExternalClusterProvider = "kubeone"
)

var (
	SupportedExternalProviders = []ExternalClusterProvider{
		ExternalClusterProviderAKS,
		ExternalClusterProviderEKS,
		ExternalClusterProviderGKE,
		ExternalClusterProviderBringYourOwn,
		ExternalClusterProviderKubeOne,
	}
)

func IsExternalProviderSupported(name string) bool {
	for _, provider := range SupportedExternalProviders {
		if strings.EqualFold(name, string(provider)) {
			return true
		}
	}

	return false
}
