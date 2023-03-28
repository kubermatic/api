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

	kubermaticeev1 "k8c.io/api/v3/pkg/apis/ee.kubermatic/v1"
	machinecontroller "k8c.io/api/v3/pkg/apis/machine-controller"
)

var (
	cloudProviderMap = map[kubermaticeev1.CloudProvider]machinecontroller.CloudProvider{
		kubermaticeev1.CloudProviderFake:                machinecontroller.CloudProviderFake,
		kubermaticeev1.CloudProviderAlibaba:             machinecontroller.CloudProviderAlibaba,
		kubermaticeev1.CloudProviderAnexia:              machinecontroller.CloudProviderAnexia,
		kubermaticeev1.CloudProviderAWS:                 machinecontroller.CloudProviderAWS,
		kubermaticeev1.CloudProviderAzure:               machinecontroller.CloudProviderAzure,
		kubermaticeev1.CloudProviderBringYourOwn:        machinecontroller.CloudProvider(""),
		kubermaticeev1.CloudProviderDigitalocean:        machinecontroller.CloudProviderDigitalocean,
		kubermaticeev1.CloudProviderGCP:                 machinecontroller.CloudProviderGoogle,
		kubermaticeev1.CloudProviderHetzner:             machinecontroller.CloudProviderHetzner,
		kubermaticeev1.CloudProviderKubeVirt:            machinecontroller.CloudProviderKubeVirt,
		kubermaticeev1.CloudProviderNutanix:             machinecontroller.CloudProviderNutanix,
		kubermaticeev1.CloudProviderOpenStack:           machinecontroller.CloudProviderOpenStack,
		kubermaticeev1.CloudProviderPacket:              machinecontroller.CloudProviderPacket,
		kubermaticeev1.CloudProviderVMwareCloudDirector: machinecontroller.CloudProviderVMwareCloudDirector,
		kubermaticeev1.CloudProviderVSphere:             machinecontroller.CloudProviderVSphere,
	}
)

func CloudProviderToMachineController(provider kubermaticeev1.CloudProvider) (machinecontroller.CloudProvider, error) {
	mapped, exists := cloudProviderMap[provider]
	if !exists {
		return machinecontroller.CloudProvider(""), fmt.Errorf("unknown cloud provider %q", provider)
	}

	if mapped == "" {
		return machinecontroller.CloudProvider(""), fmt.Errorf("KKP provider %q has no equivalent in the machine-controller", provider)
	}

	return mapped, nil
}

func CloudProviderToKKP(provider machinecontroller.CloudProvider) (kubermaticeev1.CloudProvider, error) {
	// make sure non-existing mappings in one direction (kkp=>mc) do not accidentally create
	// a fake mapping when going the other direction (mc=>kkp)
	if provider == "" {
		return kubermaticeev1.CloudProvider(""), errors.New("no cloud provider given")
	}

	for kkp, mc := range cloudProviderMap {
		if provider == mc {
			return kkp, nil
		}
	}

	return kubermaticeev1.CloudProvider(""), fmt.Errorf("unknown cloud provider %q", provider)
}

var (
	operatingSystemMap = map[kubermaticeev1.OperatingSystem]machinecontroller.OperatingSystem{
		kubermaticeev1.OperatingSystemUbuntu:       machinecontroller.OperatingSystemUbuntu,
		kubermaticeev1.OperatingSystemCentOS:       machinecontroller.OperatingSystemCentOS,
		kubermaticeev1.OperatingSystemAmazonLinux2: machinecontroller.OperatingSystemAmazonLinux2,
		kubermaticeev1.OperatingSystemRHEL:         machinecontroller.OperatingSystemRHEL,
		kubermaticeev1.OperatingSystemFlatcar:      machinecontroller.OperatingSystemFlatcar,
		kubermaticeev1.OperatingSystemRockyLinux:   machinecontroller.OperatingSystemRockyLinux,
	}
)

func OperatingSystemToMachineController(os kubermaticeev1.OperatingSystem) (machinecontroller.OperatingSystem, error) {
	mapped, exists := operatingSystemMap[os]
	if !exists {
		return machinecontroller.OperatingSystem(""), fmt.Errorf("unknown operating system %q", os)
	}

	if mapped == "" {
		return machinecontroller.OperatingSystem(""), fmt.Errorf("KKP cloud provider %q has no equivalent in the machine-controller", os)
	}

	return mapped, nil
}

func OperatingSystemToKKP(os machinecontroller.OperatingSystem) (kubermaticeev1.OperatingSystem, error) {
	// make sure non-existing mappings in one direction (kkp=>mc) do not accidentally create
	// a fake mapping when going the other direction (mc=>kkp)
	if os == "" {
		return kubermaticeev1.OperatingSystem(""), errors.New("no operating system given")
	}

	for kkp, mc := range operatingSystemMap {
		if os == mc {
			return kkp, nil
		}
	}

	return kubermaticeev1.OperatingSystem(""), fmt.Errorf("unknown operating system %q", os)
}
