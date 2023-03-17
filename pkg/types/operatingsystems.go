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

package types

import (
	"strings"
)

// OperatingSystem defines the a node's operating system.
type OperatingSystem string

const (
	OperatingSystemUbuntu       OperatingSystem = "ubuntu"
	OperatingSystemCentOS       OperatingSystem = "centos"
	OperatingSystemAmazonLinux2 OperatingSystem = "amzn2"
	OperatingSystemRHEL         OperatingSystem = "rhel"
	OperatingSystemFlatcar      OperatingSystem = "flatcar"
	OperatingSystemRockyLinux   OperatingSystem = "rockylinux"
)

var (
	SupportedOperatingSystems = []OperatingSystem{
		OperatingSystemUbuntu,
		OperatingSystemCentOS,
		OperatingSystemAmazonLinux2,
		OperatingSystemRHEL,
		OperatingSystemFlatcar,
		OperatingSystemRockyLinux,
	}
)

func IsOperatingSystemSupported(name string) bool {
	for _, os := range SupportedOperatingSystems {
		if strings.EqualFold(name, string(os)) {
			return true
		}
	}

	return false
}
