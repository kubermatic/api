// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "k8c.io/api/v3/pkg/generated/clientset/versioned/typed/ee.kubermatic/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeKubermaticEnterpriseV1 struct {
	*testing.Fake
}

func (c *FakeKubermaticEnterpriseV1) Addons(namespace string) v1.AddonInterface {
	return &FakeAddons{c, namespace}
}

func (c *FakeKubermaticEnterpriseV1) AddonConfigs(namespace string) v1.AddonConfigInterface {
	return &FakeAddonConfigs{c, namespace}
}

func (c *FakeKubermaticEnterpriseV1) AdmissionPlugins(namespace string) v1.AdmissionPluginInterface {
	return &FakeAdmissionPlugins{c, namespace}
}

func (c *FakeKubermaticEnterpriseV1) Alertmanagers(namespace string) v1.AlertmanagerInterface {
	return &FakeAlertmanagers{c, namespace}
}

func (c *FakeKubermaticEnterpriseV1) AllowedRegistries(namespace string) v1.AllowedRegistryInterface {
	return &FakeAllowedRegistries{c, namespace}
}

func (c *FakeKubermaticEnterpriseV1) Clusters(namespace string) v1.ClusterInterface {
	return &FakeClusters{c, namespace}
}

func (c *FakeKubermaticEnterpriseV1) ClusterTemplates(namespace string) v1.ClusterTemplateInterface {
	return &FakeClusterTemplates{c, namespace}
}

func (c *FakeKubermaticEnterpriseV1) ClusterTemplateInstances(namespace string) v1.ClusterTemplateInstanceInterface {
	return &FakeClusterTemplateInstances{c, namespace}
}

func (c *FakeKubermaticEnterpriseV1) Constraints(namespace string) v1.ConstraintInterface {
	return &FakeConstraints{c, namespace}
}

func (c *FakeKubermaticEnterpriseV1) ConstraintTemplates(namespace string) v1.ConstraintTemplateInterface {
	return &FakeConstraintTemplates{c, namespace}
}

func (c *FakeKubermaticEnterpriseV1) EtcdBackupConfigs(namespace string) v1.EtcdBackupConfigInterface {
	return &FakeEtcdBackupConfigs{c, namespace}
}

func (c *FakeKubermaticEnterpriseV1) EtcdRestores(namespace string) v1.EtcdRestoreInterface {
	return &FakeEtcdRestores{c, namespace}
}

func (c *FakeKubermaticEnterpriseV1) ExternalClusters(namespace string) v1.ExternalClusterInterface {
	return &FakeExternalClusters{c, namespace}
}

func (c *FakeKubermaticEnterpriseV1) GroupProjectBindings(namespace string) v1.GroupProjectBindingInterface {
	return &FakeGroupProjectBindings{c, namespace}
}

func (c *FakeKubermaticEnterpriseV1) IPAMAllocations(namespace string) v1.IPAMAllocationInterface {
	return &FakeIPAMAllocations{c, namespace}
}

func (c *FakeKubermaticEnterpriseV1) IPAMPools(namespace string) v1.IPAMPoolInterface {
	return &FakeIPAMPools{c, namespace}
}

func (c *FakeKubermaticEnterpriseV1) KubermaticConfigurations(namespace string) v1.KubermaticConfigurationInterface {
	return &FakeKubermaticConfigurations{c, namespace}
}

func (c *FakeKubermaticEnterpriseV1) KubermaticSettings(namespace string) v1.KubermaticSettingInterface {
	return &FakeKubermaticSettings{c, namespace}
}

func (c *FakeKubermaticEnterpriseV1) MLAAdminSettings(namespace string) v1.MLAAdminSettingInterface {
	return &FakeMLAAdminSettings{c, namespace}
}

func (c *FakeKubermaticEnterpriseV1) Presets(namespace string) v1.PresetInterface {
	return &FakePresets{c, namespace}
}

func (c *FakeKubermaticEnterpriseV1) ResourceQuotas(namespace string) v1.ResourceQuotaInterface {
	return &FakeResourceQuotas{c, namespace}
}

func (c *FakeKubermaticEnterpriseV1) RuleGroups(namespace string) v1.RuleGroupInterface {
	return &FakeRuleGroups{c, namespace}
}

func (c *FakeKubermaticEnterpriseV1) Seeds(namespace string) v1.SeedInterface {
	return &FakeSeeds{c, namespace}
}

func (c *FakeKubermaticEnterpriseV1) Users(namespace string) v1.UserInterface {
	return &FakeUsers{c, namespace}
}

func (c *FakeKubermaticEnterpriseV1) UserProjectBindings(namespace string) v1.UserProjectBindingInterface {
	return &FakeUserProjectBindings{c, namespace}
}

func (c *FakeKubermaticEnterpriseV1) UserSSHKeys(namespace string) v1.UserSSHKeyInterface {
	return &FakeUserSSHKeys{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeKubermaticEnterpriseV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}