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

package machinecontroller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	corev1 "k8s.io/api/core/v1"
)

// OperatingSystem defines the host operating system.
type OperatingSystem string

const (
	OperatingSystemUbuntu       OperatingSystem = "ubuntu"
	OperatingSystemCentOS       OperatingSystem = "centos"
	OperatingSystemAmazonLinux2 OperatingSystem = "amzn2"
	OperatingSystemRHEL         OperatingSystem = "rhel"
	OperatingSystemFlatcar      OperatingSystem = "flatcar"
	OperatingSystemRockyLinux   OperatingSystem = "rockylinux"
)

// GlobalObjectKeySelector is needed as we can not use v1.SecretKeySelector
// because it is not cross namespace.
type GlobalObjectKeySelector struct {
	corev1.ObjectReference `json:",inline"`
	Key                    string `json:"key,omitempty"`
}

type GlobalSecretKeySelector GlobalObjectKeySelector
type GlobalConfigMapKeySelector GlobalObjectKeySelector

type ConfigVarString struct {
	Value           string                     `json:"value,omitempty"`
	SecretKeyRef    GlobalSecretKeySelector    `json:"secretKeyRef,omitempty"`
	ConfigMapKeyRef GlobalConfigMapKeySelector `json:"configMapKeyRef,omitempty"`
}

// This type only exists to have the same fields as ConfigVarString but
// not its funcs, so it can be used as target for json.Unmarshal without
// causing a recursion.
type configVarStringWithoutUnmarshaller ConfigVarString

// MarshalJSON converts a configVarString to its JSON form, omitting empty strings.
// This is done to not have the json object cluttered with empty strings
// This will eventually hopefully be resolved within golang itself
// https://github.com/golang/go/issues/11939.
func (configVarString ConfigVarString) MarshalJSON() ([]byte, error) {
	var secretKeyRefEmpty, configMapKeyRefEmpty bool
	if configVarString.SecretKeyRef.ObjectReference.Namespace == "" &&
		configVarString.SecretKeyRef.ObjectReference.Name == "" &&
		configVarString.SecretKeyRef.Key == "" {
		secretKeyRefEmpty = true
	}

	if configVarString.ConfigMapKeyRef.ObjectReference.Namespace == "" &&
		configVarString.ConfigMapKeyRef.ObjectReference.Name == "" &&
		configVarString.ConfigMapKeyRef.Key == "" {
		configMapKeyRefEmpty = true
	}

	if secretKeyRefEmpty && configMapKeyRefEmpty {
		return []byte(fmt.Sprintf(`"%s"`, configVarString.Value)), nil
	}

	buffer := bytes.NewBufferString("{")
	if !secretKeyRefEmpty {
		jsonVal, err := json.Marshal(configVarString.SecretKeyRef)
		if err != nil {
			return nil, err
		}
		buffer.WriteString(fmt.Sprintf(`"secretKeyRef":%s`, string(jsonVal)))
	}

	if !configMapKeyRefEmpty {
		var leadingComma string
		if !secretKeyRefEmpty {
			leadingComma = ","
		}
		jsonVal, err := json.Marshal(configVarString.ConfigMapKeyRef)
		if err != nil {
			return nil, err
		}
		buffer.WriteString(fmt.Sprintf(`%s"configMapKeyRef":%s`, leadingComma, jsonVal))
	}

	if configVarString.Value != "" {
		buffer.WriteString(fmt.Sprintf(`,"value":"%s"`, configVarString.Value))
	}

	buffer.WriteString("}")
	return buffer.Bytes(), nil
}

func (configVarString *ConfigVarString) UnmarshalJSON(b []byte) error {
	if !bytes.HasPrefix(b, []byte("{")) {
		b = bytes.TrimPrefix(b, []byte(`"`))
		b = bytes.TrimSuffix(b, []byte(`"`))

		// `Unquote` expects the input string to be inside quotation marks.
		//  Since we can have a string without any quotations, in which case `TrimPrefix` and
		// `TrimSuffix` will be noop. We explicitly add quotation marks to the input string
		// to make sure that `Unquote` never fails.
		s, err := strconv.Unquote("\"" + string(b) + "\"")
		if err != nil {
			return err
		}
		configVarString.Value = s
		return nil
	}
	// This type must have the same fields as ConfigVarString but not
	// its UnmarshalJSON, otherwise we cause a recursion
	var cvsDummy configVarStringWithoutUnmarshaller
	err := json.Unmarshal(b, &cvsDummy)
	if err != nil {
		return err
	}
	configVarString.Value = cvsDummy.Value
	configVarString.SecretKeyRef = cvsDummy.SecretKeyRef
	configVarString.ConfigMapKeyRef = cvsDummy.ConfigMapKeyRef
	return nil
}

type ConfigVarBool struct {
	Value           *bool                      `json:"value,omitempty"`
	SecretKeyRef    GlobalSecretKeySelector    `json:"secretKeyRef,omitempty"`
	ConfigMapKeyRef GlobalConfigMapKeySelector `json:"configMapKeyRef,omitempty"`
}

type configVarBoolWithoutUnmarshaller ConfigVarBool

// MarshalJSON encodes the configVarBool, omitting empty strings
// This is done to not have the json object cluttered with empty strings
// This will eventually hopefully be resolved within golang itself
// https://github.com/golang/go/issues/11939
func (configVarBool ConfigVarBool) MarshalJSON() ([]byte, error) {
	var secretKeyRefEmpty, configMapKeyRefEmpty bool
	if configVarBool.SecretKeyRef.ObjectReference.Namespace == "" &&
		configVarBool.SecretKeyRef.ObjectReference.Name == "" &&
		configVarBool.SecretKeyRef.Key == "" {
		secretKeyRefEmpty = true
	}

	if configVarBool.ConfigMapKeyRef.ObjectReference.Namespace == "" &&
		configVarBool.ConfigMapKeyRef.ObjectReference.Name == "" &&
		configVarBool.ConfigMapKeyRef.Key == "" {
		configMapKeyRefEmpty = true
	}

	if secretKeyRefEmpty && configMapKeyRefEmpty {
		jsonVal, err := json.Marshal(configVarBool.Value)
		if err != nil {
			return []byte{}, err
		}
		return jsonVal, nil
	}

	buffer := bytes.NewBufferString("{")
	if !secretKeyRefEmpty {
		jsonVal, err := json.Marshal(configVarBool.SecretKeyRef)
		if err != nil {
			return nil, err
		}
		buffer.WriteString(fmt.Sprintf(`"secretKeyRef":%s`, string(jsonVal)))
	}

	if !configMapKeyRefEmpty {
		var leadingComma string
		if !secretKeyRefEmpty {
			leadingComma = ","
		}
		jsonVal, err := json.Marshal(configVarBool.ConfigMapKeyRef)
		if err != nil {
			return nil, err
		}
		buffer.WriteString(fmt.Sprintf(`%s"configMapKeyRef":%s`, leadingComma, jsonVal))
	}

	if configVarBool.Value != nil {
		jsonVal, err := json.Marshal(configVarBool.Value)
		if err != nil {
			return []byte{}, err
		}

		buffer.WriteString(fmt.Sprintf(`,"value":%v`, string(jsonVal)))
	}

	buffer.WriteString("}")

	return buffer.Bytes(), nil
}

func (configVarBool *ConfigVarBool) UnmarshalJSON(b []byte) error {
	if !bytes.HasPrefix(b, []byte("{")) {
		var val *bool
		if err := json.Unmarshal(b, &val); err != nil {
			return fmt.Errorf("Error parsing value: '%w'", err)
		}
		configVarBool.Value = val

		return nil
	}

	var cvbDummy configVarBoolWithoutUnmarshaller

	err := json.Unmarshal(b, &cvbDummy)
	if err != nil {
		return err
	}

	configVarBool.Value = cvbDummy.Value
	configVarBool.SecretKeyRef = cvbDummy.SecretKeyRef
	configVarBool.ConfigMapKeyRef = cvbDummy.ConfigMapKeyRef

	return nil
}