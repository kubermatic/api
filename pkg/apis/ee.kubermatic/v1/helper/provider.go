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

	kubermaticv1 "k8c.io/api/v3/pkg/apis/ee.kubermatic/v1"
)

// ExternalClusterCloudProviderName returns the provider name for the given ExternalClusterCloudSpec.
func ExternalClusterCloudProviderName(spec kubermaticv1.ExternalClusterCloudSpec) (kubermaticv1.ExternalClusterProvider, error) {
	var providers []kubermaticv1.ExternalClusterProvider
	if spec.AKS != nil {
		providers = append(providers, kubermaticv1.ExternalClusterProviderAKS)
	}
	if spec.EKS != nil {
		providers = append(providers, kubermaticv1.ExternalClusterProviderEKS)
	}
	if spec.GKE != nil {
		providers = append(providers, kubermaticv1.ExternalClusterProviderGKE)
	}
	if spec.KubeOne != nil {
		providers = append(providers, kubermaticv1.ExternalClusterProviderKubeOne)
	}
	if spec.BringYourOwn != nil {
		providers = append(providers, kubermaticv1.ExternalClusterProviderBringYourOwn)
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
func ClusterCloudProviderName(spec kubermaticv1.CloudSpec) (kubermaticv1.CloudProvider, error) {
	var providers []kubermaticv1.CloudProvider
	if spec.AWS != nil {
		providers = append(providers, kubermaticv1.CloudProviderAWS)
	}
	if spec.Alibaba != nil {
		providers = append(providers, kubermaticv1.CloudProviderAlibaba)
	}
	if spec.Anexia != nil {
		providers = append(providers, kubermaticv1.CloudProviderAnexia)
	}
	if spec.Azure != nil {
		providers = append(providers, kubermaticv1.CloudProviderAzure)
	}
	if spec.BringYourOwn != nil {
		providers = append(providers, kubermaticv1.CloudProviderBringYourOwn)
	}
	if spec.Digitalocean != nil {
		providers = append(providers, kubermaticv1.CloudProviderDigitalocean)
	}
	if spec.Fake != nil {
		providers = append(providers, kubermaticv1.CloudProviderFake)
	}
	if spec.GCP != nil {
		providers = append(providers, kubermaticv1.CloudProviderGCP)
	}
	if spec.Hetzner != nil {
		providers = append(providers, kubermaticv1.CloudProviderHetzner)
	}
	if spec.KubeVirt != nil {
		providers = append(providers, kubermaticv1.CloudProviderKubeVirt)
	}
	if spec.OpenStack != nil {
		providers = append(providers, kubermaticv1.CloudProviderOpenStack)
	}
	if spec.Packet != nil {
		providers = append(providers, kubermaticv1.CloudProviderPacket)
	}
	if spec.VSphere != nil {
		providers = append(providers, kubermaticv1.CloudProviderVSphere)
	}
	if spec.Nutanix != nil {
		providers = append(providers, kubermaticv1.CloudProviderNutanix)
	}
	if spec.VMwareCloudDirector != nil {
		providers = append(providers, kubermaticv1.CloudProviderVMwareCloudDirector)
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
func DatacenterCloudProviderName(spec *kubermaticv1.DatacenterSpec) (kubermaticv1.CloudProvider, error) {
	if spec == nil {
		return "", nil
	}
	var providers []kubermaticv1.CloudProvider
	if spec.BringYourOwn != nil {
		providers = append(providers, kubermaticv1.CloudProviderBringYourOwn)
	}
	if spec.Digitalocean != nil {
		providers = append(providers, kubermaticv1.CloudProviderDigitalocean)
	}
	if spec.AWS != nil {
		providers = append(providers, kubermaticv1.CloudProviderAWS)
	}
	if spec.OpenStack != nil {
		providers = append(providers, kubermaticv1.CloudProviderOpenStack)
	}
	if spec.Packet != nil {
		providers = append(providers, kubermaticv1.CloudProviderPacket)
	}
	if spec.Hetzner != nil {
		providers = append(providers, kubermaticv1.CloudProviderHetzner)
	}
	if spec.VSphere != nil {
		providers = append(providers, kubermaticv1.CloudProviderVSphere)
	}
	if spec.Azure != nil {
		providers = append(providers, kubermaticv1.CloudProviderAzure)
	}
	if spec.GCP != nil {
		providers = append(providers, kubermaticv1.CloudProviderGCP)
	}
	if spec.Fake != nil {
		providers = append(providers, kubermaticv1.CloudProviderFake)
	}
	if spec.KubeVirt != nil {
		providers = append(providers, kubermaticv1.CloudProviderKubeVirt)
	}
	if spec.Alibaba != nil {
		providers = append(providers, kubermaticv1.CloudProviderAlibaba)
	}
	if spec.Anexia != nil {
		providers = append(providers, kubermaticv1.CloudProviderAnexia)
	}
	if spec.Nutanix != nil {
		providers = append(providers, kubermaticv1.CloudProviderNutanix)
	}
	if spec.VMwareCloudDirector != nil {
		providers = append(providers, kubermaticv1.CloudProviderVMwareCloudDirector)
	}
	if len(providers) == 0 {
		return "", nil
	}
	if len(providers) != 1 {
		return "", fmt.Errorf("only one cloud provider can be set in DatacenterSpec: %+v", spec)
	}
	return providers[0], nil
}
