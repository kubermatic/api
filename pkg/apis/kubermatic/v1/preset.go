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

// +kubebuilder:object:generate=true
// +kubebuilder:object:root=true
// +kubebuilder:printcolumn:JSONPath=".metadata.creationTimestamp",name="Age",type="date"

// PresetList is the type representing a PresetList.
type PresetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	// List of presets
	Items []Preset `json:"items"`
}

// +kubebuilder:resource:scope=Cluster
// +kubebuilder:object:generate=true
// +kubebuilder:object:root=true

// Presets are preconfigured cloud provider credentials that can be applied
// to new clusters. This frees end users from having to know the actual
// credentials used for their clusters.
type Preset struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec PresetSpec `json:"spec"`
}

// Presets specifies default presets for supported providers.
type PresetSpec struct {
	Digitalocean        *DigitaloceanPreset        `json:"digitalocean,omitempty"`
	Hetzner             *HetznerPreset             `json:"hetzner,omitempty"`
	Azure               *AzurePreset               `json:"azure,omitempty"`
	VSphere             *VSpherePreset             `json:"vsphere,omitempty"`
	AWS                 *AWSPreset                 `json:"aws,omitempty"`
	OpenStack           *OpenStackPreset           `json:"openstack,omitempty"`
	Packet              *PacketPreset              `json:"packet,omitempty"`
	GCP                 *GCPPreset                 `json:"gcp,omitempty"`
	KubeVirt            *KubeVirtPreset            `json:"kubevirt,omitempty"`
	Alibaba             *AlibabaPreset             `json:"alibaba,omitempty"`
	Anexia              *AnexiaPreset              `json:"anexia,omitempty"`
	Nutanix             *NutanixPreset             `json:"nutanix,omitempty"`
	VMwareCloudDirector *VMwareCloudDirectorPreset `json:"vmwareclouddirector,omitempty"`
	GKE                 *GKEPreset                 `json:"gke,omitempty"`
	EKS                 *EKSPreset                 `json:"eks,omitempty"`
	AKS                 *AKSPreset                 `json:"aks,omitempty"`

	Fake *FakePreset `json:"fake,omitempty"`

	// RequiredEmails is a list of e-mail addresses that this presets should
	// be restricted to. Each item in the list can be either a full e-mail
	// address or just a domain name. This restriction is only enforced in the
	// KKP API.
	RequiredEmails []string `json:"requiredEmails,omitempty"`

	// Projects is a list of project IDs that this preset is limited to.
	Projects []string `json:"projects,omitempty"`

	// Only enabled presets will be available in the KKP dashboard.
	Enabled *bool `json:"enabled,omitempty"`
}

type ProviderPreset struct {
	// Only enabled presets will be available in the KKP dashboard.
	Enabled *bool `json:"enabled,omitempty"`
	// If datacenter is set, this preset is only applicable to the
	// configured datacenter.
	Datacenter string `json:"datacenter,omitempty"`
}

type DigitaloceanPreset struct {
	ProviderPreset `json:",inline"`

	// Token is used to authenticate with the DigitalOcean API.
	Token string `json:"token"`
}

func (s DigitaloceanPreset) IsValid() bool {
	return len(s.Token) > 0
}

type HetznerPreset struct {
	ProviderPreset `json:",inline"`

	// Token is used to authenticate with the Hetzner API.
	Token string `json:"token"`

	// Network is the pre-existing Hetzner network in which the machines are running.
	// While machines can be in multiple networks, a single one must be chosen for the
	// HCloud CCM to work.
	// If this is empty, the network configured on the datacenter will be used.
	Network string `json:"network,omitempty"`
}

func (s HetznerPreset) IsValid() bool {
	return len(s.Token) > 0
}

type AzurePreset struct {
	ProviderPreset `json:",inline"`

	TenantID       string `json:"tenantID"`
	SubscriptionID string `json:"subscriptionID"`
	ClientID       string `json:"clientID"`
	ClientSecret   string `json:"clientSecret"`

	ResourceGroup     string `json:"resourceGroup,omitempty"`
	VNetResourceGroup string `json:"vnetResourceGroup,omitempty"`
	VNetName          string `json:"vnet,omitempty"`
	SubnetName        string `json:"subnet,omitempty"`
	RouteTableName    string `json:"routeTable,omitempty"`
	SecurityGroup     string `json:"securityGroup,omitempty"`
	// LoadBalancerSKU sets the LB type that will be used for the Azure cluster, possible values are "basic" and "standard", if empty, "basic" will be used
	LoadBalancerSKU AzureLBSKU `json:"loadBalancerSKU"` //nolint:tagliatelle
}

func (s AzurePreset) IsValid() bool {
	return len(s.TenantID) > 0 &&
		len(s.SubscriptionID) > 0 &&
		len(s.ClientID) > 0 &&
		len(s.ClientSecret) > 0
}

type VSpherePreset struct {
	ProviderPreset `json:",inline"`

	Username string `json:"username"`
	Password string `json:"password"`

	VMNetName        string `json:"vmNetName,omitempty"`
	Datastore        string `json:"datastore,omitempty"`
	DatastoreCluster string `json:"datastoreCluster,omitempty"`
	ResourcePool     string `json:"resourcePool,omitempty"`
}

func (s VSpherePreset) IsValid() bool {
	return len(s.Username) > 0 && len(s.Password) > 0
}

type VMwareCloudDirectorPreset struct {
	ProviderPreset `json:",inline"`

	Username     string `json:"username"`
	Password     string `json:"password"`
	VDC          string `json:"vdc"`
	Organization string `json:"organization"`
	OVDCNetwork  string `json:"ovdcNetwork"`
}

func (s VMwareCloudDirectorPreset) IsValid() bool {
	return len(s.Username) > 0 &&
		len(s.Password) > 0 &&
		len(s.VDC) > 0 &&
		len(s.Organization) > 0 &&
		len(s.OVDCNetwork) > 0
}

type AWSPreset struct {
	ProviderPreset `json:",inline"`

	// Access Key ID to authenticate against AWS.
	AccessKeyID string `json:"accessKeyID"`
	// Secret Access Key to authenticate against AWS.
	SecretAccessKey string `json:"secretAccessKey"`

	AssumeRoleARN        string `json:"assumeRoleARN,omitempty"` //nolint:tagliatelle
	AssumeRoleExternalID string `json:"assumeRoleExternalID,omitempty"`

	// AWS VPC to use. Must be configured.
	VPCID string `json:"vpcID,omitempty"`
	// Route table to use. This can be configured, but if left empty will be
	// automatically filled in during reconciliation.
	RouteTableID string `json:"routeTableID,omitempty"`
	// Instance profile to use. This can be configured, but if left empty will be
	// automatically filled in during reconciliation.
	InstanceProfileName string `json:"instanceProfileName,omitempty"`
	// Security group to use. This can be configured, but if left empty will be
	// automatically filled in during reconciliation.
	SecurityGroupID string `json:"securityGroupID,omitempty"`
	// ARN to use. This can be configured, but if left empty will be
	// automatically filled in during reconciliation.
	ControlPlaneRoleARN string `json:"roleARN,omitempty"` //nolint:tagliatelle
}

func (s AWSPreset) IsValid() bool {
	return len(s.AccessKeyID) > 0 && len(s.SecretAccessKey) > 0
}

type OpenStackPreset struct {
	ProviderPreset `json:",inline"`

	UseToken bool `json:"useToken,omitempty"`

	ApplicationCredentialID     string `json:"applicationCredentialID,omitempty"`
	ApplicationCredentialSecret string `json:"applicationCredentialSecret,omitempty"`

	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
	Project   string `json:"project,omitempty"`
	ProjectID string `json:"projectID,omitempty"`
	Domain    string `json:"domain"`

	Network        string `json:"network,omitempty"`
	SecurityGroups string `json:"securityGroups,omitempty"`
	FloatingIPPool string `json:"floatingIPPool,omitempty"`
	RouterID       string `json:"routerID,omitempty"`
	SubnetID       string `json:"subnetID,omitempty"`
}

func (s OpenStackPreset) IsValid() bool {
	if s.UseToken {
		return true
	}

	if len(s.ApplicationCredentialID) > 0 {
		return len(s.ApplicationCredentialSecret) > 0
	}

	return len(s.Username) > 0 &&
		len(s.Password) > 0 &&
		(len(s.Project) > 0 || len(s.ProjectID) > 0) &&
		len(s.Domain) > 0
}

type PacketPreset struct {
	ProviderPreset `json:",inline"`

	APIKey    string `json:"apiKey"`
	ProjectID string `json:"projectID"`

	BillingCycle string `json:"billingCycle,omitempty"`
}

func (s PacketPreset) IsValid() bool {
	return len(s.APIKey) > 0 && len(s.ProjectID) > 0
}

type GCPPreset struct {
	ProviderPreset `json:",inline"`

	ServiceAccount string `json:"serviceAccount"`

	Network    string `json:"network,omitempty"`
	Subnetwork string `json:"subnetwork,omitempty"`
}

func (s GCPPreset) IsValid() bool {
	return len(s.ServiceAccount) > 0
}

type FakePreset struct {
	ProviderPreset `json:",inline"`

	Token string `json:"token"`
}

func (s FakePreset) IsValid() bool {
	return len(s.Token) > 0
}

type KubeVirtPreset struct {
	ProviderPreset `json:",inline"`

	Kubeconfig string `json:"kubeconfig"`
}

func (s KubeVirtPreset) IsValid() bool {
	return len(s.Kubeconfig) > 0
}

type AlibabaPreset struct {
	ProviderPreset `json:",inline"`

	// Access Key ID to authenticate against Alibaba.
	AccessKeyID string `json:"accessKeyID"`
	// Access Key Secret to authenticate against Alibaba.
	AccessKeySecret string `json:"accessKeySecret"`
}

func (s AlibabaPreset) IsValid() bool {
	return len(s.AccessKeyID) > 0 &&
		len(s.AccessKeySecret) > 0
}

type AnexiaPreset struct {
	ProviderPreset `json:",inline"`

	// Token is used to authenticate with the Anexia API.
	Token string `json:"token"`
}

func (s AnexiaPreset) IsValid() bool {
	return len(s.Token) > 0
}

type NutanixPreset struct {
	ProviderPreset `json:",inline"`

	// ProxyURL is used to optionally configure a HTTP proxy to access Nutanix Prism Central.
	ProxyURL string `json:"proxyURL,omitempty"`
	// Username is the username to access the Nutanix Prism Central API.
	Username string `json:"username"`
	// Password is the password corresponding to the provided user.
	Password string `json:"password"`

	// ClusterName is the Nutanix cluster to deploy resources and nodes to.
	ClusterName string `json:"clusterName"`
	// ProjectName is the optional Nutanix project to use. If none is given,
	// no project will be used.
	ProjectName string `json:"projectName,omitempty"`

	// Prism Element Username for csi driver
	CSIUsername string `json:"csiUsername,omitempty"`

	// Prism Element Password for csi driver
	CSIPassword string `json:"csiPassword,omitempty"`

	// CSIEndpoint to access Nutanix Prism Element for csi driver
	CSIEndpoint string `json:"csiEndpoint,omitempty"`

	// CSIPort to use when connecting to the Nutanix Prism Element endpoint (defaults to 9440)
	CSIPort *int32 `json:"csiPort,omitempty"`
}

func (s NutanixPreset) IsValid() bool {
	return len(s.Username) > 0 && len(s.Password) > 0
}

type GKEPreset struct {
	ProviderPreset `json:",inline"`

	ServiceAccount string `json:"serviceAccount"`
}

func (s GKEPreset) IsValid() bool {
	return len(s.ServiceAccount) > 0
}

type EKSPreset struct {
	ProviderPreset `json:",inline"`

	AccessKeyID          string `json:"accessKeyID"`
	SecretAccessKey      string `json:"secretAccessKey"`
	AssumeRoleARN        string `json:"assumeRoleARN,omitempty"` //nolint:tagliatelle
	AssumeRoleExternalID string `json:"assumeRoleExternalID,omitempty"`
}

func (s EKSPreset) IsValid() bool {
	return len(s.AccessKeyID) > 0 &&
		len(s.SecretAccessKey) > 0
}

type AKSPreset struct {
	ProviderPreset `json:",inline"`

	TenantID       string `json:"tenantID"`
	SubscriptionID string `json:"subscriptionID"`
	ClientID       string `json:"clientID"`
	ClientSecret   string `json:"clientSecret"`
}

func (s AKSPreset) IsValid() bool {
	return len(s.TenantID) > 0 &&
		len(s.SubscriptionID) > 0 &&
		len(s.ClientID) > 0 &&
		len(s.ClientSecret) > 0
}
