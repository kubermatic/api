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

	kubermaticeev1 "k8c.io/api/v3/pkg/apis/ee.kubermatic/v1"
	kubermaticv1 "k8c.io/api/v3/pkg/apis/kubermatic/v1"
)

// ClusterCloudProviderName returns the provider name for the given CloudSpec.
func ClusterCloudProviderName(spec kubermaticv1.CloudSpec) (kubermaticeev1.CloudProvider, error) {
	var providers []kubermaticeev1.CloudProvider
	if spec.AWS != nil {
		providers = append(providers, kubermaticeev1.CloudProviderAWS)
	}
	if spec.Alibaba != nil {
		providers = append(providers, kubermaticeev1.CloudProviderAlibaba)
	}
	if spec.Anexia != nil {
		providers = append(providers, kubermaticeev1.CloudProviderAnexia)
	}
	if spec.Azure != nil {
		providers = append(providers, kubermaticeev1.CloudProviderAzure)
	}
	if spec.BringYourOwn != nil {
		providers = append(providers, kubermaticeev1.CloudProviderBringYourOwn)
	}
	if spec.Digitalocean != nil {
		providers = append(providers, kubermaticeev1.CloudProviderDigitalocean)
	}
	if spec.Fake != nil {
		providers = append(providers, kubermaticeev1.CloudProviderFake)
	}
	if spec.GCP != nil {
		providers = append(providers, kubermaticeev1.CloudProviderGCP)
	}
	if spec.Hetzner != nil {
		providers = append(providers, kubermaticeev1.CloudProviderHetzner)
	}
	if spec.KubeVirt != nil {
		providers = append(providers, kubermaticeev1.CloudProviderKubeVirt)
	}
	if spec.OpenStack != nil {
		providers = append(providers, kubermaticeev1.CloudProviderOpenStack)
	}
	if spec.Packet != nil {
		providers = append(providers, kubermaticeev1.CloudProviderPacket)
	}
	if spec.VSphere != nil {
		providers = append(providers, kubermaticeev1.CloudProviderVSphere)
	}
	if spec.Nutanix != nil {
		providers = append(providers, kubermaticeev1.CloudProviderNutanix)
	}
	if spec.VMwareCloudDirector != nil {
		providers = append(providers, kubermaticeev1.CloudProviderVMwareCloudDirector)
	}
	if len(providers) == 0 {
		return "", nil
	}
	if len(providers) != 1 {
		return "", fmt.Errorf("only one cloud provider can be set in CloudSpec, but found the following providers: %v", providers)
	}
	return providers[0], nil
}
