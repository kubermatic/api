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
