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

// +genclient
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:object:generate=true
// +kubebuilder:object:root=true
// +kubebuilder:printcolumn:JSONPath=".spec.cluster.name",name="Cluster",type="string"
// +kubebuilder:printcolumn:JSONPath=".metadata.creationTimestamp",name="Age",type="date"

// MLAClusterConfiguration is the object representing cluster-specific settings for user cluster
// MLA (monitoring, logging & alerting) stack.
type MLAClusterConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec MLAClusterConfigurationSpec `json:"spec,omitempty"`
}

// MLAClusterConfigurationSpec specifies the cluster-specific administrator settings
// for KKP user cluster MLA (monitoring, logging & alerting) stack.
type MLAClusterConfigurationSpec struct {
	// Cluster is the name of the user cluster whose MLA settings are defined in this object.
	Cluster ClusterReference `json:"cluster"`
	// Monitoring contains configuration for monitoring in the user cluster.
	Monitoring *MonitoringConfiguration `json:"monitoring,omitempty"`
	// Logging contains configuration logging in the user cluster.
	Logging *LoggingConfiguration `json:"logging,omitempty"`
}

type MonitoringConfiguration struct {
	// RateLimits contains rate-limiting configuration for monitoring in the user cluster.
	RateLimits *MonitoringRateLimitSettings `json:"rateLimits,omitempty"`
}

type LoggingConfiguration struct {
	// RateLimits contains rate-limiting configuration for monitoring in the user cluster.
	RateLimits *LoggingRateLimitSettings `json:"rateLimits,omitempty"`
}

// MonitoringRateLimitSettings contains rate-limiting configuration for monitoring in the user cluster.
type MonitoringRateLimitSettings struct {
	// IngestionRate represents the ingestion rate limit in samples per second (Cortex `ingestion_rate`).
	IngestionRate int32 `json:"ingestionRate,omitempty"`
	// IngestionBurstSize represents ingestion burst size in samples per second (Cortex `ingestion_burst_size`).
	IngestionBurstSize int32 `json:"ingestionBurstSize,omitempty"`
	// MaxSeriesPerMetric represents maximum number of series per metric (Cortex `max_series_per_metric`).
	MaxSeriesPerMetric int32 `json:"maxSeriesPerMetric,omitempty"`
	// MaxSeriesTotal represents maximum number of series per this user cluster (Cortex `max_series_per_user`).
	MaxSeriesTotal int32 `json:"maxSeriesTotal,omitempty"`

	// QueryRate represents  query request rate limit per second (nginx `rate` in `r/s`).
	QueryRate int32 `json:"queryRate,omitempty"`
	// QueryBurstSize represents query burst size in number of requests (nginx `burst`).
	QueryBurstSize int32 `json:"queryBurstSize,omitempty"`
	// MaxSamplesPerQuery represents maximum number of samples during a query (Cortex `max_samples_per_query`).
	MaxSamplesPerQuery int32 `json:"maxSamplesPerQuery,omitempty"`
	// MaxSeriesPerQuery represents maximum number of timeseries during a query (Cortex `max_series_per_query`).
	MaxSeriesPerQuery int32 `json:"maxSeriesPerQuery,omitempty"`
}

// LoggingRateLimitSettings contains rate-limiting configuration for logging in the user cluster.
type LoggingRateLimitSettings struct {
	// IngestionRate represents ingestion rate limit in requests per second (nginx `rate` in `r/s`).
	IngestionRate int32 `json:"ingestionRate,omitempty"`
	// IngestionBurstSize represents ingestion burst size in number of requests (nginx `burst`).
	IngestionBurstSize int32 `json:"ingestionBurstSize,omitempty"`

	// QueryRate represents query request rate limit per second (nginx `rate` in `r/s`).
	QueryRate int32 `json:"queryRate,omitempty"`
	// QueryBurstSize represents query burst size in number of requests (nginx `burst`).
	QueryBurstSize int32 `json:"queryBurstSize,omitempty"`
}

// +kubebuilder:object:generate=true
// +kubebuilder:object:root=true

// MLAClusterConfigurationList specifies a list of administrator settings for KKP
// user cluster MLA (monitoring, logging & alerting) stack.
type MLAClusterConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []MLAClusterConfiguration `json:"items"`
}
