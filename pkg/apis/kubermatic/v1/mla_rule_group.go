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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:validation:Enum=Metrics;Logs

type MLARuleGroupType string

const (
	// RuleGroupTypeMetrics means the rule group defines the rules to generate alerts from metrics.
	MLARuleGroupTypeMetrics MLARuleGroupType = "Metrics"
	// RuleGroupTypeLogs means the rule group defines the rules to generate alerts from logs.
	MLARuleGroupTypeLogs MLARuleGroupType = "Logs"
)

// +genclient
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:object:generate=true
// +kubebuilder:object:root=true
// +kubebuilder:printcolumn:JSONPath=".spec.cluster.name",name="Cluster",type="string"
// +kubebuilder:printcolumn:JSONPath=".metadata.creationTimestamp",name="Age",type="date"

type MLARuleGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec MLARuleGroupSpec `json:"spec,omitempty"`
}

type MLARuleGroupSpec struct {
	// Cluster is the reference to the cluster the ruleGroup should be created in.
	Cluster ClusterReference `json:"cluster"`
	// Type is the type of this ruleGroup applies to.
	Type MLARuleGroupType `json:"ruleGroupType"`
	// IsDefault indicates whether the ruleGroup is the default.
	IsDefault bool `json:"isDefault,omitempty"`
	// Data contains the RuleGroup data. Ref: https://prometheus.io/docs/prometheus/latest/configuration/recording_rules/#rule_group
	Data []byte `json:"data"`
}

// +kubebuilder:object:generate=true
// +kubebuilder:object:root=true

type MLARuleGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []MLARuleGroup `json:"items"`
}
