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
	"testing"

	kubermaticv1 "k8c.io/apis/v2/pkg/kubermatic/v1"
	"k8c.io/apis/v2/pkg/test/diff"
	"k8s.io/utils/pointer"
)

func TestSetSeedDefaults(t *testing.T) {
	testCases := []struct {
		name     string
		seed     *kubermaticv1.Seed
		expected map[string]kubermaticv1.Datacenter
	}{
		{
			name: "DC settings are being respected",
			seed: &kubermaticv1.Seed{
				Spec: kubermaticv1.SeedSpec{
					ProxySettings: &kubermaticv1.ProxySettings{
						HTTPProxy: pointer.String("seed-proxy"),
						NoProxy:   pointer.String("seed-no-proxy"),
					},
					Datacenters: map[string]kubermaticv1.Datacenter{
						"a": {Node: &kubermaticv1.NodeSettings{ProxySettings: kubermaticv1.ProxySettings{
							HTTPProxy: pointer.String("dc-proxy"),
							NoProxy:   pointer.String("dc-no-proxy"),
						}}},
						"b": {Node: &kubermaticv1.NodeSettings{ProxySettings: kubermaticv1.ProxySettings{
							HTTPProxy: pointer.String("dc-proxy"),
							NoProxy:   pointer.String("dc-no-proxy"),
						}}},
					},
				},
			},
			expected: map[string]kubermaticv1.Datacenter{
				"a": {Node: &kubermaticv1.NodeSettings{ProxySettings: kubermaticv1.ProxySettings{
					HTTPProxy: pointer.String("dc-proxy"),
					NoProxy:   pointer.String("dc-no-proxy"),
				}}},
				"b": {Node: &kubermaticv1.NodeSettings{ProxySettings: kubermaticv1.ProxySettings{
					HTTPProxy: pointer.String("dc-proxy"),
					NoProxy:   pointer.String("dc-no-proxy"),
				}}},
			},
		},
		{
			name: "DC settings are being set",
			seed: &kubermaticv1.Seed{
				Spec: kubermaticv1.SeedSpec{
					ProxySettings: &kubermaticv1.ProxySettings{
						HTTPProxy: pointer.String("seed-proxy"),
						NoProxy:   pointer.String("seed-no-proxy"),
					},
					Datacenters: map[string]kubermaticv1.Datacenter{
						"a": {},
						"b": {},
					},
				},
			},
			expected: map[string]kubermaticv1.Datacenter{
				"a": {Node: &kubermaticv1.NodeSettings{ProxySettings: kubermaticv1.ProxySettings{
					HTTPProxy: pointer.String("seed-proxy"),
					NoProxy:   pointer.String("seed-no-proxy"),
				}}},
				"b": {Node: &kubermaticv1.NodeSettings{ProxySettings: kubermaticv1.ProxySettings{
					HTTPProxy: pointer.String("seed-proxy"),
					NoProxy:   pointer.String("seed-no-proxy"),
				}}},
			},
		},
		{
			name: "Only http_proxy is set",
			seed: &kubermaticv1.Seed{
				Spec: kubermaticv1.SeedSpec{
					ProxySettings: &kubermaticv1.ProxySettings{
						HTTPProxy: pointer.String("seed-proxy"),
						NoProxy:   pointer.String("seed-no-proxy"),
					},
					Datacenters: map[string]kubermaticv1.Datacenter{
						"a": {Node: &kubermaticv1.NodeSettings{ProxySettings: kubermaticv1.ProxySettings{
							NoProxy: pointer.String("dc-no-proxy"),
						}}},
						"b": {Node: &kubermaticv1.NodeSettings{ProxySettings: kubermaticv1.ProxySettings{
							NoProxy: pointer.String("dc-no-proxy"),
						}}},
					},
				},
			},
			expected: map[string]kubermaticv1.Datacenter{
				"a": {Node: &kubermaticv1.NodeSettings{ProxySettings: kubermaticv1.ProxySettings{
					HTTPProxy: pointer.String("seed-proxy"),
					NoProxy:   pointer.String("dc-no-proxy"),
				}}},
				"b": {Node: &kubermaticv1.NodeSettings{ProxySettings: kubermaticv1.ProxySettings{
					HTTPProxy: pointer.String("seed-proxy"),
					NoProxy:   pointer.String("dc-no-proxy"),
				}}},
			},
		},
		{
			name: "Only no_proxy is set",
			seed: &kubermaticv1.Seed{
				Spec: kubermaticv1.SeedSpec{
					ProxySettings: &kubermaticv1.ProxySettings{
						HTTPProxy: pointer.String("seed-proxy"),
						NoProxy:   pointer.String("seed-no-proxy"),
					},
					Datacenters: map[string]kubermaticv1.Datacenter{
						"a": {Node: &kubermaticv1.NodeSettings{ProxySettings: kubermaticv1.ProxySettings{
							HTTPProxy: pointer.String("dc-proxy"),
						}}},
						"b": {Node: &kubermaticv1.NodeSettings{ProxySettings: kubermaticv1.ProxySettings{
							HTTPProxy: pointer.String("dc-proxy"),
						}}},
					},
				},
			},
			expected: map[string]kubermaticv1.Datacenter{
				"a": {Node: &kubermaticv1.NodeSettings{ProxySettings: kubermaticv1.ProxySettings{
					HTTPProxy: pointer.String("dc-proxy"),
					NoProxy:   pointer.String("seed-no-proxy"),
				}}},
				"b": {Node: &kubermaticv1.NodeSettings{ProxySettings: kubermaticv1.ProxySettings{
					HTTPProxy: pointer.String("dc-proxy"),
					NoProxy:   pointer.String("seed-no-proxy"),
				}}},
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			SetSeedDefaults(tc.seed)
			if diff := diff.ObjectDiff(tc.expected, tc.seed.Spec.Datacenters); diff != "" {
				t.Errorf("seed.Spec.Datacenter differs from expected:\n%v", diff)
			}
		})
	}
}
