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
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/sets"
)

// +genclient
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:object:generate=true
// +kubebuilder:object:root=true
// +kubebuilder:printcolumn:JSONPath=".status.clusters",name="Clusters",type="integer"
// +kubebuilder:printcolumn:JSONPath=".metadata.creationTimestamp",name="Age",type="date"

// Datacenter is an allowed cloud provider configuration for user clusters. Each cluster
// must be scheduled to use exactly one of the available datacenters (of the same provider,
// i.e. an AWS cluster cannot use a Hetzner datacenter).
type Datacenter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DatacenterSpec   `json:"spec,omitempty"`
	Status DatacenterStatus `json:"status,omitempty"`
}

/*

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
	ExposeStrategy ExposeStrategy `json:"exposeStrategy,omitempty"`
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
*/

// DatacenterSpec configures a KKP datacenter. Provider configuration is mutually exclusive,
// and as such only a single provider can be configured per datacenter.
type DatacenterSpec struct {
	// Node holds node-specific settings, like e.g. HTTP proxy, Docker
	// registries and the like. Proxy settings are inherited from the seed if
	// not specified here.
	Node *NodeSettings `json:"node,omitempty"`
	// Digitalocean contains settings for Digitalocean (DO).
	Digitalocean *DatacenterSpecDigitalocean `json:"digitalocean,omitempty"`
	// BringYourOwn contains settings for clusters using manually created
	// nodes via kubeadm.
	BringYourOwn *DatacenterSpecBringYourOwn `json:"bringyourown,omitempty"`
	// AWS configures an Amazon Web Services (AWS) datacenter.
	AWS *DatacenterSpecAWS `json:"aws,omitempty"`
	// Azure configures an Azure datacenter.
	Azure *DatacenterSpecAzure `json:"azure,omitempty"`
	// OpenStack configures an OpenStack datacenter.
	OpenStack *DatacenterSpecOpenStack `json:"openstack,omitempty"`
	// Packet configures an Equinix Metal datacenter.
	Packet *DatacenterSpecPacket `json:"packet,omitempty"`
	// Hetzner configures a Hetzner datacenter.
	Hetzner *DatacenterSpecHetzner `json:"hetzner,omitempty"`
	// VSphere configures a VMware vSphere datacenter.
	VSphere *DatacenterSpecVSphere `json:"vsphere,omitempty"`
	// VMwareCloudDirector configures a VMware Cloud Director datacenter.
	VMwareCloudDirector *DatacenterSpecVMwareCloudDirector `json:"vmwareclouddirector,omitempty"`
	// GCP configures a Google Cloud Platform (GCP) datacenter.
	GCP *DatacenterSpecGCP `json:"gcp,omitempty"`
	// KubeVirt configures a KubeVirt datacenter.
	KubeVirt *DatacenterSpecKubeVirt `json:"kubevirt,omitempty"`
	// Alibaba configures an Alibaba Cloud datacenter.
	Alibaba *DatacenterSpecAlibaba `json:"alibaba,omitempty"`
	// Anexia configures an Anexia datacenter.
	Anexia *DatacenterSpecAnexia `json:"anexia,omitempty"`
	// Nutanix configures a Nutanix HCI datacenter.
	Nutanix *DatacenterSpecNutanix `json:"nutanix,omitempty"`

	//nolint:staticcheck
	//lint:ignore SA5008 omitgenyaml is used by the example-yaml-generator
	Fake *DatacenterSpecFake `json:"fake,omitempty,omitgenyaml"`

	// Optional: When defined, only users with an e-mail address on the
	// given domains can make use of this datacenter. You can define multiple
	// domains, e.g. "example.com", one of which must match the email domain
	// exactly (i.e. "example.com" will not match "user@test.example.com").
	RequiredEmails []string `json:"requiredEmails,omitempty"`

	// Optional: EnforceAuditLogging enforces audit logging on every cluster within the DC,
	// ignoring cluster-specific settings.
	EnforceAuditLogging bool `json:"enforceAuditLogging,omitempty"`

	// Optional: EnforcePodSecurityPolicy enforces pod security policy plugin on every clusters within the DC,
	// ignoring cluster-specific settings.
	EnforcePodSecurityPolicy bool `json:"enforcePodSecurityPolicy,omitempty"`

	// Optional: ProviderReconciliationInterval is the time that must have passed since a
	// Cluster's status.lastProviderReconciliation to make the cliuster controller
	// perform an in-depth provider reconciliation, where for example missing security
	// groups will be reconciled.
	// Setting this too low can cause rate limits by the cloud provider, setting this
	// too high means that *if* a resource at a cloud provider is removed/changed outside
	// of KKP, it will take this long to fix it.
	ProviderReconciliationInterval *metav1.Duration `json:"providerReconciliationInterval,omitempty"`

	// Optional: DefaultOperatingSystemProfiles specifies the OperatingSystemProfiles to use for each supported operating system.
	DefaultOperatingSystemProfiles OperatingSystemProfileList `json:"operatingSystemProfiles,omitempty"`

	// Optional: MachineFlavorFilter is used to filter out allowed machine flavors based on the specified resource limits like CPU, Memory, and GPU etc.
	MachineFlavorFilter *MachineFlavorFilter `json:"machineFlavorFilter,omitempty"`
}

// ImageList defines a map of operating system and the image to use.
type ImageList map[OperatingSystem]string

// ImageListWithVersions defines a map of operating system with their versions to use.
type ImageListWithVersions map[OperatingSystem]OSVersions

// OSVersions defines a map of OS version and the source to download the image.
type OSVersions map[string]string

// OperatingSystemProfileList defines a map of operating system and the OperatingSystemProfile to use.
type OperatingSystemProfileList map[OperatingSystem]string

// DatacenterSpecHetzner describes a Hetzner cloud datacenter.
type DatacenterSpecHetzner struct {
	// Datacenter location, e.g. "nbg1-dc3". A list of existing datacenters can be found
	// at https://docs.hetzner.com/general/others/data-centers-and-connection/
	Datacenter string `json:"datacenter"`
	// Network is the pre-existing Hetzner network in which the machines are running.
	// While machines can be in multiple networks, a single one must be chosen for the
	// HCloud CCM to work.
	Network string `json:"network"`
	// Optional: Detailed location of the datacenter, like "Hamburg" or "Datacenter 7".
	// For informational purposes only.
	Location string `json:"location,omitempty"`
}

// DatacenterSpecDigitalocean describes a DigitalOcean datacenter.
type DatacenterSpecDigitalocean struct {
	// Datacenter location, e.g. "ams3". A list of existing datacenters can be found
	// at https://www.digitalocean.com/docs/platform/availability-matrix/
	Region string `json:"region"`
}

// DatacenterSpecOpenStack describes an OpenStack datacenter.
type DatacenterSpecOpenStack struct {
	AuthURL          string `json:"authURL"`
	AvailabilityZone string `json:"availabilityZone,omitempty"`
	Region           string `json:"region"`
	// Optional
	IgnoreVolumeAZ bool `json:"ignoreVolumeAZ,omitempty"` //nolint:tagliatelle
	// Optional
	EnforceFloatingIP bool `json:"enforceFloatingIP,omitempty"`
	// Used for automatic network creation
	DNSServers []string `json:"dnsServers,omitempty"`
	// Images to use for each supported operating system.
	Images ImageList `json:"images"`
	// Optional: Gets mapped to the "manage-security-groups" setting in the cloud config.
	// This setting defaults to true.
	ManageSecurityGroups *bool `json:"manageSecurityGroups,omitempty"`
	// Optional: Gets mapped to the "use-octavia" setting in the cloud config.
	// use-octavia is enabled by default in CCM since v1.17.0, and disabled by
	// default with the in-tree cloud provider.
	UseOctavia *bool `json:"useOctavia,omitempty"`
	// Optional: Gets mapped to the "trust-device-path" setting in the cloud config.
	// This setting defaults to false.
	TrustDevicePath *bool `json:"trustDevicePath,omitempty"`
	// Optional: Restrict the allowed VM configurations that can be chosen in
	// the KKP dashboard. This setting does not affect the validation webhook for
	// MachineDeployments.
	NodeSizeRequirements *OpenStackNodeSizeRequirements `json:"nodeSizeRequirements,omitempty"`
	// Optional: List of enabled flavors for the given datacenter
	EnabledFlavors []string `json:"enabledFlavors,omitempty"`
	// Optional: defines if the IPv6 is enabled for the datacenter
	IPv6Enabled *bool `json:"ipv6Enabled,omitempty"`
}

type OpenStackNodeSizeRequirements struct {
	// VCPUs is the minimum required amount of (virtual) CPUs
	MinimumVCPUs int `json:"minimumVCPUs,omitempty"` //nolint:tagliatelle
	// MinimumMemory is the minimum required amount of memory, measured in MB
	MinimumMemory int `json:"minimumMemory,omitempty"`
}

// DatacenterSpecAzure describes an Azure cloud datacenter.
type DatacenterSpecAzure struct {
	// Region to use, for example "westeurope". A list of available regions can be
	// found at https://azure.microsoft.com/en-us/global-infrastructure/locations/
	Location string `json:"location"`
}

// DatacenterSpecVSphere describes a vSphere datacenter.
type DatacenterSpecVSphere struct {
	// Endpoint URL to use, including protocol, for example "https://vcenter.example.com".
	Endpoint string `json:"endpoint"`
	// If set to true, disables the TLS certificate check against the endpoint.
	AllowInsecure bool `json:"allowInsecure,omitempty"`
	// The default Datastore to be used for provisioning volumes using storage
	// classes/dynamic provisioning and for storing virtual machine files in
	// case no `Datastore` or `DatastoreCluster` is provided at Cluster level.
	DefaultDatastore string `json:"datastore"`
	// The name of the datacenter to use.
	Datacenter string `json:"datacenter"`
	// The name of the vSphere cluster to use. Used for out-of-tree CSI Driver.
	Cluster string `json:"cluster"`
	// The name of the storage policy to use for the storage class created in the user cluster.
	DefaultStoragePolicy string `json:"storagePolicy,omitempty"`
	// Optional: The root path for cluster specific VM folders. Each cluster gets its own
	// folder below the root folder. Must be the FQDN (for example
	// "/datacenter-1/vm/all-kubermatic-vms-in-here") and defaults to the root VM
	// folder: "/datacenter-1/vm"
	RootPath string `json:"rootPath,omitempty"`
	// A list of VM templates to use for a given operating system. You must
	// define at least one template.
	// See: https://github.com/kubermatic/machine-controller/blob/master/docs/vsphere.md#template-vms-preparation
	Templates ImageList `json:"templates"`
	// Optional: Infra management user is the user that will be used for everything
	// except the cloud provider functionality, which will still use the credentials
	// passed in via the Kubermatic dashboard/API.
	InfraManagementUser *VSphereCredentials `json:"infraManagementUser,omitempty"`
	// Optional: defines if the IPv6 is enabled for the datacenter
	IPv6Enabled *bool `json:"ipv6Enabled,omitempty"`
	// DefaultTagCategoryID is the tag category id that will be used as default, if users don't specify it on a cluster level,
	// and they don't wish KKP to create default generated tag category, upon cluster creation.
	DefaultTagCategoryID string `json:"defaultTagCategoryID,omitempty"`
}

type DatacenterSpecVMwareCloudDirector struct {
	// Endpoint URL to use, including protocol, for example "https://vclouddirector.example.com".
	URL string `json:"url"`
	// If set to true, disables the TLS certificate check against the endpoint.
	AllowInsecure bool `json:"allowInsecure,omitempty"`
	// The default catalog which contains the VM templates.
	DefaultCatalog string `json:"catalog,omitempty"`
	// The name of the storage profile to use for disks attached to the VMs.
	DefaultStorageProfile string `json:"storageProfile,omitempty"`
	// A list of VM templates to use for a given operating system. You must
	// define at least one template.
	Templates ImageList `json:"templates"`
}

// DatacenterSpecAWS describes an AWS datacenter.
type DatacenterSpecAWS struct {
	// The AWS region to use, e.g. "us-east-1". For a list of available regions, see
	// https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html
	Region string `json:"region"`

	// List of AMIs to use for a given operating system.
	// This gets defaulted by querying for the latest AMI for the given distribution
	// when machines are created, so under normal circumstances it is not necessary
	// to define the AMIs statically.
	Images ImageList `json:"images,omitempty"`
}

// DatacenterSpecBringYourOwn describes a datacenter our of bring your own nodes.
type DatacenterSpecBringYourOwn struct {
}

// DatacenterSpecPacket describes a Packet datacenter.
type DatacenterSpecPacket struct {
	// The list of enabled facilities, for example "ams1", for a full list of available
	// facilities see https://metal.equinix.com/developers/docs/locations/facilities/
	Facilities []string `json:"facilities,omitempty"`
	// Metros are facilities that are grouped together geographically and share capacity
	// and networking features, see https://metal.equinix.com/developers/docs/locations/metros/
	Metro string `json:"metro,omitempty"`
}

// DatacenterSpecGCP describes a GCP datacenter.
type DatacenterSpecGCP struct {
	// Region to use, for example "europe-west3", for a full list of regions see
	// https://cloud.google.com/compute/docs/regions-zones/
	Region string `json:"region"`
	// List of enabled zones, for example [a, c]. See the link above for the available
	// zones in your chosen region.
	ZoneSuffixes []string `json:"zoneSuffixes"`

	// Optional: Regional clusters spread their resources across multiple availability zones.
	// Refer to the official documentation for more details on this:
	// https://cloud.google.com/kubernetes-engine/docs/concepts/regional-clusters
	Regional bool `json:"regional,omitempty"`
}

// DatacenterSpecFake describes a fake datacenter.
type DatacenterSpecFake struct {
	FakeProperty string `json:"fakeProperty,omitempty"`
}

// DatacenterSpecKubeVirt describes a kubevirt datacenter.
type DatacenterSpecKubeVirt struct {
	// +kubebuilder:validation:Enum=ClusterFirstWithHostNet;ClusterFirst;Default;None
	// +kubebuilder:default=ClusterFirst

	// DNSPolicy represents the dns policy for the pod. Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst',
	// 'Default' or 'None'. Defaults to "ClusterFirst". DNS parameters given in DNSConfig will be merged with the
	// policy selected with DNSPolicy.
	DNSPolicy string `json:"dnsPolicy,omitempty"`

	// DNSConfig represents the DNS parameters of a pod. Parameters specified here will be merged to the generated DNS
	// configuration based on DNSPolicy.
	DNSConfig *corev1.PodDNSConfig `json:"dnsConfig,omitempty"`

	// CustomNetworkPolicies (optional) allows to add some extra custom NetworkPolicies, that are deployed
	// in the dedicated infra KubeVirt cluster. They are added to the defaults.
	CustomNetworkPolicies []*CustomNetworkPolicy `json:"customNetworkPolicies,omitempty"`

	// Images represents standard VM Image sources.
	Images *KubeVirtImageSources `json:"images,omitempty"`

	// InfraStorageClasses contains a list of KubeVirt infra cluster StorageClasses names
	// that will be used to initialise StorageClasses in the tenant cluster.
	// In the tenant cluster, the created StorageClass name will have as name:
	// kubevirt-<infra-storageClass-name>
	InfraStorageClasses []KubeVirtInfraStorageClass `json:"infraStorageClasses,omitempty"`
}

type KubeVirtInfraStorageClass struct {
	Name string `json:"name"`
	// Optional: IsDefaultClass. If true, the created StorageClass in the tenant cluster will be annotated with:
	// storageclass.kubernetes.io/is-default-class : true
	// If missing or false, annotation will be:
	// storageclass.kubernetes.io/is-default-class : false
	IsDefaultClass *bool `json:"isDefaultClass,omitempty"`
}

// CustomNetworkPolicy contains a name and the Spec of a NetworkPolicy.
type CustomNetworkPolicy struct {
	// Name is the name of the Custom Network Policy.
	Name string `json:"name"`
	// Spec is the Spec of the NetworkPolicy, using the standard type.
	Spec networkingv1.NetworkPolicySpec `json:"spec"`
}

var (
	SupportedKubeVirtOS = sets.New(
		OperatingSystemCentOS,
		OperatingSystemUbuntu,
		OperatingSystemRHEL,
		OperatingSystemFlatcar,
		OperatingSystemRockyLinux,
	)
)

// KubeVirtImageSources represents KubeVirt image sources.
type KubeVirtImageSources struct {
	// HTTP represents a http source.
	HTTP *KubeVirtHTTPSource `json:"http,omitempty"`
}

// KubeVirtHTTPSource represents list of images and their versions that can be downloaded over HTTP.
type KubeVirtHTTPSource struct {
	// OperatingSystems represents list of supported operating-systems with their URLs.
	OperatingSystems map[OperatingSystem]OSVersions `json:"operatingSystems"`
}

// DatacenterSpecNutanix describes a Nutanix datacenter.
type DatacenterSpecNutanix struct {
	// Endpoint to use for accessing Nutanix Prism Central. No protocol or port should be passed,
	// for example "nutanix.example.com" or "10.0.0.1"
	Endpoint string `json:"endpoint"`
	// Optional: Port to use when connecting to the Nutanix Prism Central endpoint (defaults to 9440)
	Port *int32 `json:"port,omitempty"`

	// Optional: AllowInsecure allows to disable the TLS certificate check against the endpoint (defaults to false)
	AllowInsecure bool `json:"allowInsecure,omitempty"`
	// Images to use for each supported operating system
	Images ImageList `json:"images"`
}

// DatacenterSpecAlibaba describes a alibaba datacenter.
type DatacenterSpecAlibaba struct {
	// Region to use, for a full list of regions see
	// https://www.alibabacloud.com/help/doc-detail/40654.htm
	Region string `json:"region"`
}

// DatacenterSpecAnexia describes a anexia datacenter.
type DatacenterSpecAnexia struct {
	// LocationID the location of the region
	LocationID string `json:"locationID"`
}

// NodeSettings are node specific flags which can be configured on datacenter level.
type NodeSettings struct {
	// Optional: Proxy settings for the Nodes in this datacenter.
	// Defaults to the Proxy settings of the seed.
	ProxySettings `json:",inline"`
	// Optional: These image registries will be configured as insecure
	// on the container runtime.
	InsecureRegistries []string `json:"insecureRegistries,omitempty"`
	// Optional: These image registries will be configured as registry mirrors
	// on the container runtime.
	RegistryMirrors []string `json:"registryMirrors,omitempty"`
	// Optional: Translates to --pod-infra-container-image on the kubelet.
	// If not set, the kubelet will default it.
	PauseImage string `json:"pauseImage,omitempty"`
	// Optional: ContainerdRegistryMirrors configure registry mirrors endpoints. Can be used multiple times to specify multiple mirrors.
	ContainerdRegistryMirrors *ContainerRuntimeContainerd `json:"containerdRegistryMirrors,omitempty"`
}

// ContainerRuntimeContainerd defines containerd container runtime registries configs.
type ContainerRuntimeContainerd struct {
	// A map of registries to use to render configs and mirrors for containerd registries
	Registries map[string]ContainerdRegistry `json:"registries,omitempty"`
}

// ContainerdRegistry defines endpoints and security for given container registry.
type ContainerdRegistry struct {
	// List of registry mirrors to use
	Mirrors []string `json:"mirrors,omitempty"`
}

// ProxySettings allow configuring a HTTP proxy for the control planes and nodes.
type ProxySettings struct {
	// Optional: If set, this proxy will be configured for both HTTP and HTTPS.
	HTTPProxy *string `json:"httpProxy,omitempty"`
	// Optional: If set this will be set as NO_PROXY environment variable on the node;
	// The value must be a comma-separated list of domains for which no proxy
	// should be used, e.g. "*.example.com,internal.dev".
	// Note that the in-cluster apiserver URL will be automatically prepended
	// to this value.
	NoProxy *string `json:"noProxy,omitempty"`
}

func emptyStrPtr(s *string) bool {
	return s == nil || *s == ""
}

// Empty returns true if p or all of its children are nil or empty strings.
func (p *ProxySettings) Empty() bool {
	return p == nil || (emptyStrPtr(p.HTTPProxy) && emptyStrPtr(p.NoProxy))
}

// Merge applies the settings from p into dst if the corresponding setting
// in dst is nil or an empty string.
func (p *ProxySettings) Merge(dst *ProxySettings) {
	if emptyStrPtr(dst.HTTPProxy) {
		dst.HTTPProxy = p.HTTPProxy
	}
	if emptyStrPtr(dst.NoProxy) {
		dst.NoProxy = p.NoProxy
	}
}

// DatacenterStatus contains runtime information regarding the datacenter.
type DatacenterStatus struct {
	// +kubebuilder:default=0
	// +kubebuilder:validation:Minimum:=0

	// Clusters is the total number of user clusters that exist on this seed.
	Clusters int `json:"clusters"`
}

// +kubebuilder:object:generate=true
// +kubebuilder:object:root=true

// DatacenterList is a list of datacenters.
type DatacenterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Datacenter `json:"items"`
}
