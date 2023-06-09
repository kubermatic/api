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

	kubermaticv1 "k8c.io/api/v3/pkg/apis/kubermatic/v1"

	"k8s.io/apimachinery/pkg/util/sets"
)

func TestMapCloudProvider(t *testing.T) {
	for _, kkpProvider := range sets.List(kubermaticv1.AllCloudProviders) {
		t.Run(string(kkpProvider), func(t *testing.T) {
			// provider has purposefully no mapping
			if kkpProvider == kubermaticv1.CloudProviderBringYourOwn {
				return
			}

			mcProvider, err := CloudProviderToMachineController(kkpProvider)
			if err != nil {
				t.Fatalf("Cannot map cloud provider %q to machine-controller: %v", kkpProvider, err)
			}

			mapped, err := CloudProviderToKKP(mcProvider)
			if err != nil {
				t.Fatalf("Cannot map cloud provider %q (originally %q) back to KKP: %v", mcProvider, kkpProvider, err)
			}

			if mapped != kkpProvider {
				t.Fatalf("%q in KKP maps to %q in machine-controller, but in reverse maps to %q", kkpProvider, mcProvider, mapped)
			}
		})
	}
}

func TestMapOperatingSystem(t *testing.T) {
	for _, kkpOperatingSystem := range sets.List(kubermaticv1.AllOperatingSystems) {
		t.Run(string(kkpOperatingSystem), func(t *testing.T) {
			mcOperatingSystem, err := OperatingSystemToMachineController(kkpOperatingSystem)
			if err != nil {
				t.Fatalf("Cannot map operating system %q to machine-controller: %v", kkpOperatingSystem, err)
			}

			mapped, err := OperatingSystemToKKP(mcOperatingSystem)
			if err != nil {
				t.Fatalf("Cannot map operating system %q (originally %q) back to KKP: %v", mcOperatingSystem, kkpOperatingSystem, err)
			}

			if mapped != kkpOperatingSystem {
				t.Fatalf("%q in KKP maps to %q in machine-controller, but in reverse maps to %q", kkpOperatingSystem, mcOperatingSystem, mapped)
			}
		})
	}
}
