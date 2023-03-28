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

	kubermaticv1 "k8c.io/api/v3/pkg/apis/kubermatic/v1/ee"
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
