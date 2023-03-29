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
	kubermaticeev1 "k8c.io/api/v3/pkg/apis/ee.kubermatic/v1"
)

var (
	// knownIPv6CloudProviders configures which providers have IPv6 and if it's enabled for all datacenters.
	knownIPv6CloudProviders = map[kubermaticeev1.CloudProvider]struct {
		ipv6EnabledForAllDatacenters bool
	}{
		kubermaticeev1.CloudProviderAWS: {
			ipv6EnabledForAllDatacenters: true,
		},
		kubermaticeev1.CloudProviderAzure: {
			ipv6EnabledForAllDatacenters: true,
		},
		kubermaticeev1.CloudProviderBringYourOwn: {
			ipv6EnabledForAllDatacenters: true,
		},
		kubermaticeev1.CloudProviderDigitalocean: {
			ipv6EnabledForAllDatacenters: true,
		},
		kubermaticeev1.CloudProviderGCP: {
			ipv6EnabledForAllDatacenters: true,
		},
		kubermaticeev1.CloudProviderHetzner: {
			ipv6EnabledForAllDatacenters: true,
		},
		kubermaticeev1.CloudProviderOpenStack: {
			ipv6EnabledForAllDatacenters: false,
		},
		kubermaticeev1.CloudProviderPacket: {
			ipv6EnabledForAllDatacenters: true,
		},
		kubermaticeev1.CloudProviderVSphere: {
			ipv6EnabledForAllDatacenters: false,
		},
	}
)

func IsIPv6KnownProvider(provider kubermaticeev1.CloudProvider) bool {
	_, isIPv6KnownProvider := knownIPv6CloudProviders[provider]
	return isIPv6KnownProvider
}

// IsIPv6EnabledDatacenter returns true if IPv6 is enabled for the datacenter.
func IsIPv6EnabledDatacenter(dc *kubermaticeev1.Datacenter) bool {
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
	case kubermaticeev1.CloudProviderOpenStack:
		flag = dc.Spec.OpenStack.IPv6Enabled
	case kubermaticeev1.CloudProviderVSphere:
		flag = dc.Spec.VSphere.IPv6Enabled
	}

	return flag != nil && *flag
}
