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

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"net/http"

	v1 "k8c.io/api/v2/pkg/apis/kubermatic/v1"
	"k8c.io/api/v2/pkg/generated/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type KubermaticV1Interface interface {
	RESTClient() rest.Interface
	AddonsGetter
	AddonConfigsGetter
	AdmissionPluginsGetter
	AlertmanagersGetter
	AllowedRegistriesGetter
	ClustersGetter
	ClusterTemplatesGetter
	ClusterTemplateInstancesGetter
	ConstraintsGetter
	ConstraintTemplatesGetter
	EtcdBackupConfigsGetter
	EtcdRestoresGetter
	ExternalClustersGetter
	GroupProjectBindingsGetter
	IPAMAllocationsGetter
	IPAMPoolsGetter
	KubermaticConfigurationsGetter
	KubermaticSettingsGetter
	MLAAdminSettingsGetter
	PresetsGetter
	ProjectsGetter
	ResourceQuotasGetter
	RuleGroupsGetter
	SeedsGetter
	UsersGetter
	UserProjectBindingsGetter
	UserSSHKeysGetter
}

// KubermaticV1Client is used to interact with features provided by the kubermatic.k8c.io group.
type KubermaticV1Client struct {
	restClient rest.Interface
}

func (c *KubermaticV1Client) Addons(namespace string) AddonInterface {
	return newAddons(c, namespace)
}

func (c *KubermaticV1Client) AddonConfigs(namespace string) AddonConfigInterface {
	return newAddonConfigs(c, namespace)
}

func (c *KubermaticV1Client) AdmissionPlugins(namespace string) AdmissionPluginInterface {
	return newAdmissionPlugins(c, namespace)
}

func (c *KubermaticV1Client) Alertmanagers(namespace string) AlertmanagerInterface {
	return newAlertmanagers(c, namespace)
}

func (c *KubermaticV1Client) AllowedRegistries(namespace string) AllowedRegistryInterface {
	return newAllowedRegistries(c, namespace)
}

func (c *KubermaticV1Client) Clusters(namespace string) ClusterInterface {
	return newClusters(c, namespace)
}

func (c *KubermaticV1Client) ClusterTemplates(namespace string) ClusterTemplateInterface {
	return newClusterTemplates(c, namespace)
}

func (c *KubermaticV1Client) ClusterTemplateInstances(namespace string) ClusterTemplateInstanceInterface {
	return newClusterTemplateInstances(c, namespace)
}

func (c *KubermaticV1Client) Constraints(namespace string) ConstraintInterface {
	return newConstraints(c, namespace)
}

func (c *KubermaticV1Client) ConstraintTemplates(namespace string) ConstraintTemplateInterface {
	return newConstraintTemplates(c, namespace)
}

func (c *KubermaticV1Client) EtcdBackupConfigs(namespace string) EtcdBackupConfigInterface {
	return newEtcdBackupConfigs(c, namespace)
}

func (c *KubermaticV1Client) EtcdRestores(namespace string) EtcdRestoreInterface {
	return newEtcdRestores(c, namespace)
}

func (c *KubermaticV1Client) ExternalClusters(namespace string) ExternalClusterInterface {
	return newExternalClusters(c, namespace)
}

func (c *KubermaticV1Client) GroupProjectBindings(namespace string) GroupProjectBindingInterface {
	return newGroupProjectBindings(c, namespace)
}

func (c *KubermaticV1Client) IPAMAllocations(namespace string) IPAMAllocationInterface {
	return newIPAMAllocations(c, namespace)
}

func (c *KubermaticV1Client) IPAMPools(namespace string) IPAMPoolInterface {
	return newIPAMPools(c, namespace)
}

func (c *KubermaticV1Client) KubermaticConfigurations(namespace string) KubermaticConfigurationInterface {
	return newKubermaticConfigurations(c, namespace)
}

func (c *KubermaticV1Client) KubermaticSettings(namespace string) KubermaticSettingInterface {
	return newKubermaticSettings(c, namespace)
}

func (c *KubermaticV1Client) MLAAdminSettings(namespace string) MLAAdminSettingInterface {
	return newMLAAdminSettings(c, namespace)
}

func (c *KubermaticV1Client) Presets(namespace string) PresetInterface {
	return newPresets(c, namespace)
}

func (c *KubermaticV1Client) Projects(namespace string) ProjectInterface {
	return newProjects(c, namespace)
}

func (c *KubermaticV1Client) ResourceQuotas(namespace string) ResourceQuotaInterface {
	return newResourceQuotas(c, namespace)
}

func (c *KubermaticV1Client) RuleGroups(namespace string) RuleGroupInterface {
	return newRuleGroups(c, namespace)
}

func (c *KubermaticV1Client) Seeds(namespace string) SeedInterface {
	return newSeeds(c, namespace)
}

func (c *KubermaticV1Client) Users(namespace string) UserInterface {
	return newUsers(c, namespace)
}

func (c *KubermaticV1Client) UserProjectBindings(namespace string) UserProjectBindingInterface {
	return newUserProjectBindings(c, namespace)
}

func (c *KubermaticV1Client) UserSSHKeys(namespace string) UserSSHKeyInterface {
	return newUserSSHKeys(c, namespace)
}

// NewForConfig creates a new KubermaticV1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*KubermaticV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new KubermaticV1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*KubermaticV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &KubermaticV1Client{client}, nil
}

// NewForConfigOrDie creates a new KubermaticV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *KubermaticV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new KubermaticV1Client for the given RESTClient.
func New(c rest.Interface) *KubermaticV1Client {
	return &KubermaticV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *KubermaticV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
