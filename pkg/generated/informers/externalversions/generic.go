// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	v1 "k8c.io/api/v3/pkg/apis/apps.kubermatic/v1"
	eeappskubermaticv1 "k8c.io/api/v3/pkg/apis/ee.apps.kubermatic/v1"
	eekubermaticv1 "k8c.io/api/v3/pkg/apis/ee.kubermatic/v1"
	kubermaticv1 "k8c.io/api/v3/pkg/apis/kubermatic/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=apps.kubermatic.k8c.io, Version=v1
	case v1.SchemeGroupVersion.WithResource("applicationdefinitions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.KubermaticApps().V1().ApplicationDefinitions().Informer()}, nil
	case v1.SchemeGroupVersion.WithResource("applicationinstallations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.KubermaticApps().V1().ApplicationInstallations().Informer()}, nil

		// Group=ee.apps.kubermatic.k8c.io, Version=v1
	case eeappskubermaticv1.SchemeGroupVersion.WithResource("applicationdefinitions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.KubermaticEnterpriseApps().V1().ApplicationDefinitions().Informer()}, nil
	case eeappskubermaticv1.SchemeGroupVersion.WithResource("applicationinstallations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.KubermaticEnterpriseApps().V1().ApplicationInstallations().Informer()}, nil

		// Group=ee.kubermatic.k8c.io, Version=v1
	case eekubermaticv1.SchemeGroupVersion.WithResource("addons"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.KubermaticEnterprise().V1().Addons().Informer()}, nil
	case eekubermaticv1.SchemeGroupVersion.WithResource("addonconfigs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.KubermaticEnterprise().V1().AddonConfigs().Informer()}, nil
	case eekubermaticv1.SchemeGroupVersion.WithResource("admissionplugins"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.KubermaticEnterprise().V1().AdmissionPlugins().Informer()}, nil
	case eekubermaticv1.SchemeGroupVersion.WithResource("alertmanagers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.KubermaticEnterprise().V1().Alertmanagers().Informer()}, nil
	case eekubermaticv1.SchemeGroupVersion.WithResource("allowedregistries"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.KubermaticEnterprise().V1().AllowedRegistries().Informer()}, nil
	case eekubermaticv1.SchemeGroupVersion.WithResource("clusters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.KubermaticEnterprise().V1().Clusters().Informer()}, nil
	case eekubermaticv1.SchemeGroupVersion.WithResource("clustertemplates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.KubermaticEnterprise().V1().ClusterTemplates().Informer()}, nil
	case eekubermaticv1.SchemeGroupVersion.WithResource("clustertemplateinstances"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.KubermaticEnterprise().V1().ClusterTemplateInstances().Informer()}, nil
	case eekubermaticv1.SchemeGroupVersion.WithResource("constraints"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.KubermaticEnterprise().V1().Constraints().Informer()}, nil
	case eekubermaticv1.SchemeGroupVersion.WithResource("constrainttemplates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.KubermaticEnterprise().V1().ConstraintTemplates().Informer()}, nil
	case eekubermaticv1.SchemeGroupVersion.WithResource("etcdbackupconfigs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.KubermaticEnterprise().V1().EtcdBackupConfigs().Informer()}, nil
	case eekubermaticv1.SchemeGroupVersion.WithResource("etcdrestores"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.KubermaticEnterprise().V1().EtcdRestores().Informer()}, nil
	case eekubermaticv1.SchemeGroupVersion.WithResource("externalclusters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.KubermaticEnterprise().V1().ExternalClusters().Informer()}, nil
	case eekubermaticv1.SchemeGroupVersion.WithResource("groupprojectbindings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.KubermaticEnterprise().V1().GroupProjectBindings().Informer()}, nil
	case eekubermaticv1.SchemeGroupVersion.WithResource("ipamallocations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.KubermaticEnterprise().V1().IPAMAllocations().Informer()}, nil
	case eekubermaticv1.SchemeGroupVersion.WithResource("ipampools"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.KubermaticEnterprise().V1().IPAMPools().Informer()}, nil
	case eekubermaticv1.SchemeGroupVersion.WithResource("kubermaticconfigurations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.KubermaticEnterprise().V1().KubermaticConfigurations().Informer()}, nil
	case eekubermaticv1.SchemeGroupVersion.WithResource("kubermaticsettings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.KubermaticEnterprise().V1().KubermaticSettings().Informer()}, nil
	case eekubermaticv1.SchemeGroupVersion.WithResource("mlaadminsettings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.KubermaticEnterprise().V1().MLAAdminSettings().Informer()}, nil
	case eekubermaticv1.SchemeGroupVersion.WithResource("presets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.KubermaticEnterprise().V1().Presets().Informer()}, nil
	case eekubermaticv1.SchemeGroupVersion.WithResource("resourcequotas"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.KubermaticEnterprise().V1().ResourceQuotas().Informer()}, nil
	case eekubermaticv1.SchemeGroupVersion.WithResource("rulegroups"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.KubermaticEnterprise().V1().RuleGroups().Informer()}, nil
	case eekubermaticv1.SchemeGroupVersion.WithResource("sshkeys"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.KubermaticEnterprise().V1().SSHKeys().Informer()}, nil
	case eekubermaticv1.SchemeGroupVersion.WithResource("seeds"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.KubermaticEnterprise().V1().Seeds().Informer()}, nil
	case eekubermaticv1.SchemeGroupVersion.WithResource("users"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.KubermaticEnterprise().V1().Users().Informer()}, nil
	case eekubermaticv1.SchemeGroupVersion.WithResource("userprojectbindings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.KubermaticEnterprise().V1().UserProjectBindings().Informer()}, nil

		// Group=kubermatic.k8c.io, Version=v1
	case kubermaticv1.SchemeGroupVersion.WithResource("addons"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubermatic().V1().Addons().Informer()}, nil
	case kubermaticv1.SchemeGroupVersion.WithResource("addonconfigs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubermatic().V1().AddonConfigs().Informer()}, nil
	case kubermaticv1.SchemeGroupVersion.WithResource("admissionplugins"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubermatic().V1().AdmissionPlugins().Informer()}, nil
	case kubermaticv1.SchemeGroupVersion.WithResource("alertmanagers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubermatic().V1().Alertmanagers().Informer()}, nil
	case kubermaticv1.SchemeGroupVersion.WithResource("allowedregistries"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubermatic().V1().AllowedRegistries().Informer()}, nil
	case kubermaticv1.SchemeGroupVersion.WithResource("clusters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubermatic().V1().Clusters().Informer()}, nil
	case kubermaticv1.SchemeGroupVersion.WithResource("clustertemplates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubermatic().V1().ClusterTemplates().Informer()}, nil
	case kubermaticv1.SchemeGroupVersion.WithResource("clustertemplateinstances"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubermatic().V1().ClusterTemplateInstances().Informer()}, nil
	case kubermaticv1.SchemeGroupVersion.WithResource("constraints"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubermatic().V1().Constraints().Informer()}, nil
	case kubermaticv1.SchemeGroupVersion.WithResource("constrainttemplates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubermatic().V1().ConstraintTemplates().Informer()}, nil
	case kubermaticv1.SchemeGroupVersion.WithResource("datacenters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubermatic().V1().Datacenters().Informer()}, nil
	case kubermaticv1.SchemeGroupVersion.WithResource("etcdbackupconfigs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubermatic().V1().EtcdBackupConfigs().Informer()}, nil
	case kubermaticv1.SchemeGroupVersion.WithResource("etcdrestores"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubermatic().V1().EtcdRestores().Informer()}, nil
	case kubermaticv1.SchemeGroupVersion.WithResource("externalclusters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubermatic().V1().ExternalClusters().Informer()}, nil
	case kubermaticv1.SchemeGroupVersion.WithResource("ipamallocations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubermatic().V1().IPAMAllocations().Informer()}, nil
	case kubermaticv1.SchemeGroupVersion.WithResource("ipampools"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubermatic().V1().IPAMPools().Informer()}, nil
	case kubermaticv1.SchemeGroupVersion.WithResource("kubermaticconfigurations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubermatic().V1().KubermaticConfigurations().Informer()}, nil
	case kubermaticv1.SchemeGroupVersion.WithResource("kubermaticsettings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubermatic().V1().KubermaticSettings().Informer()}, nil
	case kubermaticv1.SchemeGroupVersion.WithResource("mlaadminsettings"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubermatic().V1().MLAAdminSettings().Informer()}, nil
	case kubermaticv1.SchemeGroupVersion.WithResource("presets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubermatic().V1().Presets().Informer()}, nil
	case kubermaticv1.SchemeGroupVersion.WithResource("resourcequotas"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubermatic().V1().ResourceQuotas().Informer()}, nil
	case kubermaticv1.SchemeGroupVersion.WithResource("rulegroups"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubermatic().V1().RuleGroups().Informer()}, nil
	case kubermaticv1.SchemeGroupVersion.WithResource("sshkeys"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubermatic().V1().SSHKeys().Informer()}, nil
	case kubermaticv1.SchemeGroupVersion.WithResource("users"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Kubermatic().V1().Users().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
