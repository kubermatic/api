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
	"k8c.io/api/v2/pkg/types"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:validation:Pattern:=`^((\d{1,3}\.){3}\d{1,3}\/([0-9]|[1-2][0-9]|3[0-2]))$`
type CIDR string

// +kubebuilder:validation:Enum="";Healthy;Unhealthy;Invalid;Terminating;Paused

type SeedPhase string

// These are the valid phases of a seed.
const (
	// SeedPhaseHealthy means the seed is reachable and was successfully reconciled.
	SeedPhaseHealthy SeedPhase = "Healthy"

	// SeedPhaseUnhealthy means the KKP resources on the seed cluster could not be
	// successfully reconciled.
	SeedPhaseUnhealthy SeedPhase = "Unhealthy"

	// SeedPhaseInvalid means the seed kubeconfig is defunct.
	SeedPhaseInvalid SeedPhase = "Invalid"

	// SeedPhaseTerminating means the seed is currently being deleted.
	SeedPhaseTerminating SeedPhase = "Terminating"

	// SeedPhasePaused means the seed is not being reconciled because the SkipReconciling
	// annotation is set.
	SeedPhasePaused SeedPhase = "Paused"
)

// +kubebuilder:object:generate=true
// +kubebuilder:object:root=true

// SeedDatacenterList is the type representing a SeedDatacenterList.
type SeedList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	// List of seeds
	Items []Seed `json:"items"`
}

// +kubebuilder:object:generate=true
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:JSONPath=".status.clusters",name="Clusters",type="integer"
// +kubebuilder:printcolumn:JSONPath=".spec.location",name="Location",type="string"
// +kubebuilder:printcolumn:JSONPath=".status.versions.kubermatic",name="KKP Version",type="string"
// +kubebuilder:printcolumn:JSONPath=".status.versions.cluster",name="Cluster Version",type="string"
// +kubebuilder:printcolumn:JSONPath=".status.phase",name="Phase",type="string"
// +kubebuilder:printcolumn:JSONPath=".metadata.creationTimestamp",name="Age",type="date"

// Seed is the type representing a Seed cluster. Seed clusters host the the control planes
// for KKP user clusters.
type Seed struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec SeedSpec `json:"spec"`
	//nolint:staticcheck
	//lint:ignore SA5008 omitgenyaml is used by the example-yaml-generator
	Status SeedStatus `json:"status,omitempty,omitgenyaml"`
}

// SeedStatus contains runtime information regarding the seed.
type SeedStatus struct {
	// Phase contains a human readable text to indicate the seed cluster status. No logic should be tied
	// to this field, as its content can change in between KKP releases.
	Phase SeedPhase `json:"phase,omitempty"`

	// +kubebuilder:default=0
	// +kubebuilder:validation:Minimum:=0

	// Clusters is the total number of user clusters that exist on this seed.
	Clusters int `json:"clusters"`

	// Versions contains information regarding versions of components in the cluster and the cluster
	// itself.
	// +optional
	Versions SeedVersionsStatus `json:"versions,omitempty"`

	// Conditions contains conditions the seed is in, its primary use case is status signaling
	// between controllers or between controllers and the API.
	// +optional
	Conditions map[SeedConditionType]SeedCondition `json:"conditions,omitempty"`
}

// +kubebuilder:validation:Enum="";KubeconfigValid;ResourcesReconciled;ClusterInitialized

// SeedConditionType is used to indicate the type of a seed condition. For all condition
// types, the `true` value must indicate success.
type SeedConditionType string

const (
	// SeedConditionKubeconfigValid indicates that the configured kubeconfig for the seed is valid.
	// The seed-sync controller manages this condition.
	SeedConditionKubeconfigValid SeedConditionType = "KubeconfigValid"
	// SeedConditionResourcesReconciled indicates that the KKP operator has finished setting up the
	// resources inside the seed cluster.
	SeedConditionResourcesReconciled SeedConditionType = "ResourcesReconciled"
	// SeedConditionClusterInitialized indicates that the KKP operator has finished setting up the
	// CRDs and other prerequisites on the Seed cluster. After this condition is true, other
	// controllers can begin to create watches and reconcile resources (i.e. this condition is
	// a precondition to ResourcesReconciled). Once this condition is true, it is never set to false
	// again.
	SeedConditionClusterInitialized SeedConditionType = "ClusterInitialized"
)

type SeedCondition struct {
	// Status of the condition, one of True, False, Unknown.
	Status corev1.ConditionStatus `json:"status"`
	// Last time we got an update on a given condition.
	LastHeartbeatTime metav1.Time `json:"lastHeartbeatTime"`
	// Last time the condition transit from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
	// (brief) reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty"`
	// Human readable message indicating details about last transition.
	// +optional
	Message string `json:"message,omitempty"`
}

// SeedVersionsStatus contains information regarding versions of components in the cluster
// and the cluster itself.
type SeedVersionsStatus struct {
	// Kubermatic is the version of the currently deployed KKP components. Note that a permanent
	// version skew between master and seed is not supported and KKP setups should never run for
	// longer times with a skew between the clusters.
	Kubermatic string `json:"kubermatic,omitempty"`
	// Cluster is the Kubernetes version of the cluster's control plane.
	Cluster string `json:"cluster,omitempty"`
}

// The spec for a seed cluster.
type SeedSpec struct {
	// Optional: Country of the seed as ISO-3166 two-letter code, e.g. DE or UK.
	// For informational purposes in the Kubermatic dashboard only.
	Country string `json:"country,omitempty"`
	// Optional: Detailed location of the cluster, like "Hamburg" or "Datacenter 7".
	// For informational purposes in the Kubermatic dashboard only.
	Location string `json:"location,omitempty"`
	// A reference to the Kubeconfig of this cluster. The Kubeconfig must
	// have cluster-admin privileges. This field is mandatory for every
	// seed, even if there are no datacenters defined yet.
	Kubeconfig corev1.ObjectReference `json:"kubeconfig"`
	// Datacenters contains a map of the possible datacenters (DCs) in this seed.
	// Each DC must have a globally unique identifier (i.e. names must be unique
	// across all seeds).
	Datacenters map[string]Datacenter `json:"datacenters,omitempty"`
	// Optional: This can be used to override the DNS name used for this seed.
	// By default the seed name is used.
	SeedDNSOverwrite string `json:"seedDNSOverwrite,omitempty"`
	// NodeportProxy can be used to configure the NodePort proxy service that is
	// responsible for making user-cluster control planes accessible from the outside.
	NodeportProxy *NodeportProxyConfig `json:"nodeportProxy,omitempty"`
	// Optional: ProxySettings can be used to configure HTTP proxy settings on the
	// worker nodes in user clusters. However, proxy settings on nodes take precedence.
	ProxySettings *ProxySettings `json:"proxySettings,omitempty"`
	// Optional: ExposeStrategy explicitly sets the expose strategy for this seed cluster, if not set, the default provided by the master is used.
	ExposeStrategy types.ExposeStrategy `json:"exposeStrategy,omitempty"`
	// Optional: MLA allows configuring seed level MLA (Monitoring, Logging & Alerting) stack settings.
	MLA *SeedMLASettings `json:"mla,omitempty"`
	// DefaultComponentSettings are default values to set for newly created clusters.
	// Deprecated: Use DefaultClusterTemplate instead.
	DefaultComponentSettings *ComponentSettings `json:"defaultComponentSettings,omitempty"`
	// DefaultClusterTemplate is the name of a cluster template of scope "seed" that is used
	// to default all new created clusters
	DefaultClusterTemplate string `json:"defaultClusterTemplate,omitempty"`
	// Metering configures the metering tool on user clusters across the seed.
	Metering *MeteringConfiguration `json:"metering,omitempty"`
	// EtcdBackupRestore holds the configuration of the automatic etcd backup restores for the Seed;
	// if this is set, the new backup/restore controllers are enabled for this Seed.
	EtcdBackupRestore *EtcdBackupRestore `json:"etcdBackupRestore,omitempty"`
	// OIDCProviderConfiguration allows to configure OIDC provider at the Seed level.
	OIDCProviderConfiguration *OIDCProviderConfiguration `json:"oidcProviderConfiguration,omitempty"`
}

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
