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
	kubermaticv1 "k8c.io/api/v2/pkg/kubermatic/v1"
	"k8c.io/api/v2/pkg/types"
)

var (
	// knownIPv6CloudProviders configures which providers have IPv6 and if it's enabled for all datacenters.
	knownIPv6CloudProviders = map[types.CloudProvider]struct {
		ipv6EnabledForAllDatacenters bool
	}{
		types.CloudProviderAWS: {
			ipv6EnabledForAllDatacenters: true,
		},
		types.CloudProviderAzure: {
			ipv6EnabledForAllDatacenters: true,
		},
		types.CloudProviderBringYourOwn: {
			ipv6EnabledForAllDatacenters: true,
		},
		types.CloudProviderDigitalocean: {
			ipv6EnabledForAllDatacenters: true,
		},
		types.CloudProviderGCP: {
			ipv6EnabledForAllDatacenters: true,
		},
		types.CloudProviderHetzner: {
			ipv6EnabledForAllDatacenters: true,
		},
		types.CloudProviderOpenStack: {
			ipv6EnabledForAllDatacenters: false,
		},
		types.CloudProviderPacket: {
			ipv6EnabledForAllDatacenters: true,
		},
		types.CloudProviderVSphere: {
			ipv6EnabledForAllDatacenters: false,
		},
	}
)

// IsIPv6EnabledDatacenter returns true if IPv6 is enabled for the datacenter.
func IsIPv6EnabledDatacenter(dc *kubermaticv1.Datacenter) bool {
	provider, err := DatacenterCloudProviderName(&dc.Spec)
	if err != nil {
		return false
	}

	cloudProviderCfg, exists := knownIPv6CloudProviders[provider]
	if !exists {
		return false
	}

	if cloudProviderCfg.ipv6EnabledForAllDatacenters {
		return true
	}

	var flag *bool

	switch provider {
	case types.CloudProviderOpenStack:
		flag = dc.Spec.OpenStack.IPv6Enabled
	case types.CloudProviderVSphere:
		flag = dc.Spec.VSphere.IPv6Enabled
	}

	return flag != nil && *flag
}
