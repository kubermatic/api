// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"net/http"

	v1 "k8c.io/api/v3/pkg/apis/ee.kubermatic/v1"
	"k8c.io/api/v3/pkg/generated/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type KubermaticEnterpriseV1Interface interface {
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
	ResourceQuotasGetter
	RuleGroupsGetter
	SeedsGetter
	UsersGetter
	UserProjectBindingsGetter
	UserSSHKeysGetter
}

// KubermaticEnterpriseV1Client is used to interact with features provided by the ee.kubermatic.k8c.io group.
type KubermaticEnterpriseV1Client struct {
	restClient rest.Interface
}

func (c *KubermaticEnterpriseV1Client) Addons(namespace string) AddonInterface {
	return newAddons(c, namespace)
}

func (c *KubermaticEnterpriseV1Client) AddonConfigs(namespace string) AddonConfigInterface {
	return newAddonConfigs(c, namespace)
}

func (c *KubermaticEnterpriseV1Client) AdmissionPlugins(namespace string) AdmissionPluginInterface {
	return newAdmissionPlugins(c, namespace)
}

func (c *KubermaticEnterpriseV1Client) Alertmanagers(namespace string) AlertmanagerInterface {
	return newAlertmanagers(c, namespace)
}

func (c *KubermaticEnterpriseV1Client) AllowedRegistries(namespace string) AllowedRegistryInterface {
	return newAllowedRegistries(c, namespace)
}

func (c *KubermaticEnterpriseV1Client) Clusters(namespace string) ClusterInterface {
	return newClusters(c, namespace)
}

func (c *KubermaticEnterpriseV1Client) ClusterTemplates(namespace string) ClusterTemplateInterface {
	return newClusterTemplates(c, namespace)
}

func (c *KubermaticEnterpriseV1Client) ClusterTemplateInstances(namespace string) ClusterTemplateInstanceInterface {
	return newClusterTemplateInstances(c, namespace)
}

func (c *KubermaticEnterpriseV1Client) Constraints(namespace string) ConstraintInterface {
	return newConstraints(c, namespace)
}

func (c *KubermaticEnterpriseV1Client) ConstraintTemplates(namespace string) ConstraintTemplateInterface {
	return newConstraintTemplates(c, namespace)
}

func (c *KubermaticEnterpriseV1Client) EtcdBackupConfigs(namespace string) EtcdBackupConfigInterface {
	return newEtcdBackupConfigs(c, namespace)
}

func (c *KubermaticEnterpriseV1Client) EtcdRestores(namespace string) EtcdRestoreInterface {
	return newEtcdRestores(c, namespace)
}

func (c *KubermaticEnterpriseV1Client) ExternalClusters(namespace string) ExternalClusterInterface {
	return newExternalClusters(c, namespace)
}

func (c *KubermaticEnterpriseV1Client) GroupProjectBindings(namespace string) GroupProjectBindingInterface {
	return newGroupProjectBindings(c, namespace)
}

func (c *KubermaticEnterpriseV1Client) IPAMAllocations(namespace string) IPAMAllocationInterface {
	return newIPAMAllocations(c, namespace)
}

func (c *KubermaticEnterpriseV1Client) IPAMPools(namespace string) IPAMPoolInterface {
	return newIPAMPools(c, namespace)
}

func (c *KubermaticEnterpriseV1Client) KubermaticConfigurations(namespace string) KubermaticConfigurationInterface {
	return newKubermaticConfigurations(c, namespace)
}

func (c *KubermaticEnterpriseV1Client) KubermaticSettings(namespace string) KubermaticSettingInterface {
	return newKubermaticSettings(c, namespace)
}

func (c *KubermaticEnterpriseV1Client) MLAAdminSettings(namespace string) MLAAdminSettingInterface {
	return newMLAAdminSettings(c, namespace)
}

func (c *KubermaticEnterpriseV1Client) Presets(namespace string) PresetInterface {
	return newPresets(c, namespace)
}

func (c *KubermaticEnterpriseV1Client) ResourceQuotas(namespace string) ResourceQuotaInterface {
	return newResourceQuotas(c, namespace)
}

func (c *KubermaticEnterpriseV1Client) RuleGroups(namespace string) RuleGroupInterface {
	return newRuleGroups(c, namespace)
}

func (c *KubermaticEnterpriseV1Client) Seeds(namespace string) SeedInterface {
	return newSeeds(c, namespace)
}

func (c *KubermaticEnterpriseV1Client) Users(namespace string) UserInterface {
	return newUsers(c, namespace)
}

func (c *KubermaticEnterpriseV1Client) UserProjectBindings(namespace string) UserProjectBindingInterface {
	return newUserProjectBindings(c, namespace)
}

func (c *KubermaticEnterpriseV1Client) UserSSHKeys(namespace string) UserSSHKeyInterface {
	return newUserSSHKeys(c, namespace)
}

// NewForConfig creates a new KubermaticEnterpriseV1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*KubermaticEnterpriseV1Client, error) {
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

// NewForConfigAndClient creates a new KubermaticEnterpriseV1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*KubermaticEnterpriseV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &KubermaticEnterpriseV1Client{client}, nil
}

// NewForConfigOrDie creates a new KubermaticEnterpriseV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *KubermaticEnterpriseV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new KubermaticEnterpriseV1Client for the given RESTClient.
func New(c rest.Interface) *KubermaticEnterpriseV1Client {
	return &KubermaticEnterpriseV1Client{c}
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
func (c *KubermaticEnterpriseV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
