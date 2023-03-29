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
	corev1 "k8s.io/api/core/v1"
)

// EtcdBackupRestore holds the configuration of the automatic backup and restores.
type EtcdBackupRestore struct {
	// Destinations stores all the possible destinations where the backups for the Seed can be stored. If not empty,
	// it enables automatic backup and restore for the seed.
	Destinations map[string]*BackupDestination `json:"destinations,omitempty"`

	// +kubebuilder:validation:Pattern:=`^[a-z0-9]([-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*$`
	// +kubebuilder:validation:MaxLength:=63
	// +kubebuilder:validation:Type=string

	// DefaultDestination marks the default destination that will be used for the default etcd backup config which is
	// created for every user cluster. Has to correspond to a destination in Destinations.
	// If removed, it removes the related default etcd backup configs.
	DefaultDestination string `json:"defaultDestination,omitempty"`
}

// BackupDestination defines the bucket name and endpoint as a backup destination, and holds reference to the credentials secret.
type BackupDestination struct {
	// Endpoint is the API endpoint to use for backup and restore.
	Endpoint string `json:"endpoint"`
	// BucketName is the bucket name to use for backup and restore.
	BucketName string `json:"bucketName"`
	// Credentials hold the ref to the secret with backup credentials
	Credentials *corev1.SecretReference `json:"credentials,omitempty"`
}

type NodeportProxyConfig struct {
	// Disable will prevent the Kubermatic Operator from creating a nodeport-proxy
	// setup on the seed cluster. This should only be used if a suitable replacement
	// is installed (like the nodeport-proxy Helm chart).
	Disable bool `json:"disable,omitempty"`
	// Annotations are used to further tweak the LoadBalancer integration with the
	// cloud provider where the seed cluster is running.
	// Deprecated: Use .envoy.loadBalancerService.annotations instead.
	Annotations map[string]string `json:"annotations,omitempty"`
	// Envoy configures the Envoy application itself.
	Envoy *NodePortProxyComponentEnvoy `json:"envoy,omitempty"`
	// EnvoyManager configures the Kubermatic-internal Envoy manager.
	EnvoyManager *NodeportProxyComponent `json:"envoyManager,omitempty"`
	// Updater configures the component responsible for updating the LoadBalancer
	// service.
	Updater *NodeportProxyComponent `json:"updater,omitempty"`
}

type EnvoyLoadBalancerService struct {
	// Annotations are used to further tweak the LoadBalancer integration with the
	// cloud provider.
	Annotations map[string]string `json:"annotations,omitempty"`
	// SourceRanges will restrict loadbalancer service to IP ranges specified using CIDR notation like 172.25.0.0/16.
	// This field will be ignored if the cloud-provider does not support the feature.
	// More info: https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/
	SourceRanges []CIDR `json:"sourceRanges,omitempty"`
}

type NodePortProxyComponentEnvoy struct {
	NodeportProxyComponent `json:",inline"`
	LoadBalancerService    *EnvoyLoadBalancerService `json:"loadBalancerService,omitempty"`
}

type NodeportProxyComponent struct {
	// DockerRepository is the repository containing the component's image.
	DockerRepository string `json:"dockerRepository,omitempty"`
	// Resources describes the requested and maximum allowed CPU/memory usage.
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`
}

// SeedMLASettings allow configuring seed level MLA (Monitoring, Logging & Alerting) stack settings.
type SeedMLASettings struct {
	// Optional: controls whether the user cluster MLA (Monitoring, Logging & Alerting) stack is enabled in the seed.
	UserClusterMLAEnabled bool `json:"userClusterMLAEnabled,omitempty"` //nolint:tagliatelle
}

// MeteringConfiguration contains all the configuration for the metering tool.
type MeteringConfiguration struct {
	Enabled bool `json:"enabled"`

	// StorageClassName is the name of the storage class that the metering prometheus instance uses to store metric data for reporting.
	StorageClassName string `json:"storageClassName"`
	// StorageSize is the size of the storage class. Default value is 100Gi.
	StorageSize string `json:"storageSize"`

	// +kubebuilder:default:={weekly: {schedule: "0 1 * * 6", interval: 7}}

	// ReportConfigurations is a map of report configuration definitions.
	ReportConfigurations map[string]*MeteringReportConfiguration `json:"reports,omitempty"`
}

type MeteringReportConfiguration struct {
	// +kubebuilder:default:=`0 1 * * 6`

	// Schedule in Cron format, see https://en.wikipedia.org/wiki/Cron. Please take a note that Schedule is responsible
	// only for setting the time when a report generation mechanism kicks off. The Interval MUST be set independently.
	Schedule string `json:"schedule,omitempty"`

	// +kubebuilder:default=7
	// +kubebuilder:validation:Minimum:=1

	// Interval defines the number of days consulted in the metering report.
	Interval uint32 `json:"interval,omitempty"`

	// +optional
	// +kubebuilder:validation:Minimum:=1

	// Retention defines a number of days after which reports are queued for removal. If not set, reports are kept forever.
	// Please note that this functionality works only for object storage that supports an object lifecycle management mechanism.
	Retention *uint32 `json:"retention,omitempty"`

	// +optional
	// +kubebuilder:default:={"cluster","namespace"}

	// Types of reports to generate. Available report types are cluster and namespace. By default, all types of reports are generated.
	Types []string `json:"type,omitempty"`
}

// OIDCProviderConfiguration allows to configure OIDC provider at the Seed level. If set, it overwrites the OIDC configuration from the KubermaticConfiguration.
// OIDC is later used to configure:
// - access to User Cluster API-Servers (via user kubeconfigs) - https://kubernetes.io/docs/reference/access-authn-authz/authentication/#openid-connect-tokens,
// - access to User Cluster's Kubernetes Dashboards.
type OIDCProviderConfiguration struct {
	// URL of the provider which allows the API server to discover public signing keys.
	IssuerURL string `json:"issuerURL"`

	// IssuerClientID is the application's ID.
	IssuerClientID string `json:"issuerClientID"`

	// IssuerClientSecret is the application's secret.
	IssuerClientSecret string `json:"issuerClientSecret"`

	// Optional: CookieHashKey is required, used to authenticate the cookie value using HMAC.
	// It is recommended to use a key with 32 or 64 bytes.
	// If not set, configuration is inherited from the default OIDC provider.
	CookieHashKey *string `json:"cookieHashKey,omitempty"`

	// Optional: CookieSecureMode if true then cookie received only with HTTPS otherwise with HTTP.
	// If not set, configuration is inherited from the default OIDC provider.
	CookieSecureMode *bool `json:"cookieSecureMode,omitempty"`

	// Optional:  OfflineAccessAsScope if true then "offline_access" scope will be used
	// otherwise 'access_type=offline" query param will be passed.
	// If not set, configuration is inherited from the default OIDC provider.
	OfflineAccessAsScope *bool `json:"offlineAccessAsScope,omitempty"`

	// Optional: SkipTLSVerify skip TLS verification for the token issuer.
	// If not set, configuration is inherited from the default OIDC provider.
	SkipTLSVerify *bool `json:"skipTLSVerify,omitempty"`
}

// // IsEtcdAutomaticBackupEnabled returns true if etcd automatic backup is configured for the seed.
// func (s *Seed) IsEtcdAutomaticBackupEnabled() bool {
// 	if cfg := s.Spec.EtcdBackupRestore; cfg != nil {
// 		return len(cfg.Destinations) > 0
// 	}
// 	return false
// }

// // IsDefaultEtcdAutomaticBackupEnabled returns true if etcd automatic backup with default destination is configured for the seed.
// func (s *Seed) IsDefaultEtcdAutomaticBackupEnabled() bool {
// 	return s.IsEtcdAutomaticBackupEnabled() && s.Spec.EtcdBackupRestore.DefaultDestination != ""
// }

// func (s *Seed) GetEtcdBackupDestination(destinationName string) *BackupDestination {
// 	if s.Spec.EtcdBackupRestore == nil {
// 		return nil
// 	}

// 	return s.Spec.EtcdBackupRestore.Destinations[destinationName]
// }
