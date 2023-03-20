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
	"errors"
	"fmt"

	kubermaticv1 "k8c.io/api/v2/pkg/apis/kubermatic/v1"
	machinecontroller "k8c.io/api/v2/pkg/apis/machine-controller"
)

var (
	cloudProviderMap = map[kubermaticv1.CloudProvider]machinecontroller.CloudProvider{
		kubermaticv1.CloudProviderFake:                machinecontroller.CloudProviderFake,
		kubermaticv1.CloudProviderAlibaba:             machinecontroller.CloudProviderAlibaba,
		kubermaticv1.CloudProviderAnexia:              machinecontroller.CloudProviderAnexia,
		kubermaticv1.CloudProviderAWS:                 machinecontroller.CloudProviderAWS,
		kubermaticv1.CloudProviderAzure:               machinecontroller.CloudProviderAzure,
		kubermaticv1.CloudProviderBringYourOwn:        machinecontroller.CloudProvider(""),
		kubermaticv1.CloudProviderDigitalocean:        machinecontroller.CloudProviderDigitalocean,
		kubermaticv1.CloudProviderGCP:                 machinecontroller.CloudProviderGoogle,
		kubermaticv1.CloudProviderHetzner:             machinecontroller.CloudProviderHetzner,
		kubermaticv1.CloudProviderKubeVirt:            machinecontroller.CloudProviderKubeVirt,
		kubermaticv1.CloudProviderNutanix:             machinecontroller.CloudProviderNutanix,
		kubermaticv1.CloudProviderOpenStack:           machinecontroller.CloudProviderOpenStack,
		kubermaticv1.CloudProviderPacket:              machinecontroller.CloudProviderPacket,
		kubermaticv1.CloudProviderVMwareCloudDirector: machinecontroller.CloudProviderVMwareCloudDirector,
		kubermaticv1.CloudProviderVSphere:             machinecontroller.CloudProviderVSphere,
	}
)

func CloudProviderToMachineController(provider kubermaticv1.CloudProvider) (machinecontroller.CloudProvider, error) {
	mapped, exists := cloudProviderMap[provider]
	if !exists {
		return machinecontroller.CloudProvider(""), fmt.Errorf("unknown cloud provider %q", provider)
	}

	if mapped == "" {
		return machinecontroller.CloudProvider(""), fmt.Errorf("KKP provider %q has no equivalent in the machine-controller", provider)
	}

	return mapped, nil
}

func CloudProviderToKKP(provider machinecontroller.CloudProvider) (kubermaticv1.CloudProvider, error) {
	// make sure non-existing mappings in one direction (kkp=>mc) do not accidentally create
	// a fake mapping when going the other direction (mc=>kkp)
	if provider == "" {
		return kubermaticv1.CloudProvider(""), errors.New("no cloud provider given")
	}

	for kkp, mc := range cloudProviderMap {
		if provider == mc {
			return kkp, nil
		}
	}

	return kubermaticv1.CloudProvider(""), fmt.Errorf("unknown cloud provider %q", provider)
}

var (
	operatingSystemMap = map[kubermaticv1.OperatingSystem]machinecontroller.OperatingSystem{
		kubermaticv1.OperatingSystemUbuntu:       machinecontroller.OperatingSystemUbuntu,
		kubermaticv1.OperatingSystemCentOS:       machinecontroller.OperatingSystemCentOS,
		kubermaticv1.OperatingSystemAmazonLinux2: machinecontroller.OperatingSystemAmazonLinux2,
		kubermaticv1.OperatingSystemRHEL:         machinecontroller.OperatingSystemRHEL,
		kubermaticv1.OperatingSystemFlatcar:      machinecontroller.OperatingSystemFlatcar,
		kubermaticv1.OperatingSystemRockyLinux:   machinecontroller.OperatingSystemRockyLinux,
	}
)

func OperatingSystemToMachineController(os kubermaticv1.OperatingSystem) (machinecontroller.OperatingSystem, error) {
	mapped, exists := operatingSystemMap[os]
	if !exists {
		return machinecontroller.OperatingSystem(""), fmt.Errorf("unknown operating system %q", os)
	}

	if mapped == "" {
		return machinecontroller.OperatingSystem(""), fmt.Errorf("KKP cloud provider %q has no equivalent in the machine-controller", os)
	}

	return mapped, nil
}

func OperatingSystemToKKP(os machinecontroller.OperatingSystem) (kubermaticv1.OperatingSystem, error) {
	// make sure non-existing mappings in one direction (kkp=>mc) do not accidentally create
	// a fake mapping when going the other direction (mc=>kkp)
	if os == "" {
		return kubermaticv1.OperatingSystem(""), errors.New("no operating system given")
	}

	for kkp, mc := range operatingSystemMap {
		if os == mc {
			return kkp, nil
		}
	}

	return kubermaticv1.OperatingSystem(""), fmt.Errorf("unknown operating system %q", os)
}
