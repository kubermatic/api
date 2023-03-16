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

package helper

import (
	"fmt"

	kubermaticv1 "k8c.io/apis/v2/pkg/kubermatic/v1"
	"k8c.io/apis/v2/pkg/types"
)

// ExternalClusterCloudProviderName returns the provider name for the given ExternalClusterCloudSpec.
func ExternalClusterCloudProviderName(spec kubermaticv1.ExternalClusterCloudSpec) (types.ExternalClusterProvider, error) {
	var providers []types.ExternalClusterProvider
	if spec.AKS != nil {
		providers = append(providers, types.ExternalClusterProviderAKS)
	}
	if spec.EKS != nil {
		providers = append(providers, types.ExternalClusterProviderEKS)
	}
	if spec.GKE != nil {
		providers = append(providers, types.ExternalClusterProviderGKE)
	}
	if spec.KubeOne != nil {
		providers = append(providers, types.ExternalClusterProviderKubeOne)
	}
	if spec.BringYourOwn != nil {
		providers = append(providers, types.ExternalClusterProviderBringYourOwn)
	}
	if len(providers) == 0 {
		return "", nil
	}
	if len(providers) != 1 {
		return "", fmt.Errorf("only one cloud provider can be set in ExternalClusterCloudSpec, but found the following providers: %v", providers)
	}
	return providers[0], nil
}

// ClusterCloudProviderName returns the provider name for the given CloudSpec.
func ClusterCloudProviderName(spec kubermaticv1.CloudSpec) (types.CloudProvider, error) {
	var providers []types.CloudProvider
	if spec.AWS != nil {
		providers = append(providers, types.CloudProviderAWS)
	}
	if spec.Alibaba != nil {
		providers = append(providers, types.CloudProviderAlibaba)
	}
	if spec.Anexia != nil {
		providers = append(providers, types.CloudProviderAnexia)
	}
	if spec.Azure != nil {
		providers = append(providers, types.CloudProviderAzure)
	}
	if spec.BringYourOwn != nil {
		providers = append(providers, types.CloudProviderBringYourOwn)
	}
	if spec.Digitalocean != nil {
		providers = append(providers, types.CloudProviderDigitalocean)
	}
	if spec.Fake != nil {
		providers = append(providers, types.CloudProviderFake)
	}
	if spec.GCP != nil {
		providers = append(providers, types.CloudProviderGCP)
	}
	if spec.Hetzner != nil {
		providers = append(providers, types.CloudProviderHetzner)
	}
	if spec.Kubevirt != nil {
		providers = append(providers, types.CloudProviderKubevirt)
	}
	if spec.OpenStack != nil {
		providers = append(providers, types.CloudProviderOpenStack)
	}
	if spec.Packet != nil {
		providers = append(providers, types.CloudProviderPacket)
	}
	if spec.VSphere != nil {
		providers = append(providers, types.CloudProviderVSphere)
	}
	if spec.Nutanix != nil {
		providers = append(providers, types.CloudProviderNutanix)
	}
	if spec.VMwareCloudDirector != nil {
		providers = append(providers, types.CloudProviderVMwareCloudDirector)
	}
	if len(providers) == 0 {
		return "", nil
	}
	if len(providers) != 1 {
		return "", fmt.Errorf("only one cloud provider can be set in CloudSpec, but found the following providers: %v", providers)
	}
	return providers[0], nil
}

// DatacenterCloudProviderName returns the provider name for the given Datacenter.
func DatacenterCloudProviderName(spec *kubermaticv1.DatacenterSpec) (types.CloudProvider, error) {
	if spec == nil {
		return "", nil
	}
	var providers []types.CloudProvider
	if spec.BringYourOwn != nil {
		providers = append(providers, types.CloudProviderBringYourOwn)
	}
	if spec.Digitalocean != nil {
		providers = append(providers, types.CloudProviderDigitalocean)
	}
	if spec.AWS != nil {
		providers = append(providers, types.CloudProviderAWS)
	}
	if spec.OpenStack != nil {
		providers = append(providers, types.CloudProviderOpenStack)
	}
	if spec.Packet != nil {
		providers = append(providers, types.CloudProviderPacket)
	}
	if spec.Hetzner != nil {
		providers = append(providers, types.CloudProviderHetzner)
	}
	if spec.VSphere != nil {
		providers = append(providers, types.CloudProviderVSphere)
	}
	if spec.Azure != nil {
		providers = append(providers, types.CloudProviderAzure)
	}
	if spec.GCP != nil {
		providers = append(providers, types.CloudProviderGCP)
	}
	if spec.Fake != nil {
		providers = append(providers, types.CloudProviderFake)
	}
	if spec.Kubevirt != nil {
		providers = append(providers, types.CloudProviderKubevirt)
	}
	if spec.Alibaba != nil {
		providers = append(providers, types.CloudProviderAlibaba)
	}
	if spec.Anexia != nil {
		providers = append(providers, types.CloudProviderAnexia)
	}
	if spec.Nutanix != nil {
		providers = append(providers, types.CloudProviderNutanix)
	}
	if spec.VMwareCloudDirector != nil {
		providers = append(providers, types.CloudProviderVMwareCloudDirector)
	}
	if len(providers) == 0 {
		return "", nil
	}
	if len(providers) != 1 {
		return "", fmt.Errorf("only one cloud provider can be set in DatacenterSpec: %+v", spec)
	}
	return providers[0], nil
}
