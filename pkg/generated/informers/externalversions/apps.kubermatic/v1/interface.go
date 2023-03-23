// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	internalinterfaces "k8c.io/api/v2/pkg/generated/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// ApplicationDefinitions returns a ApplicationDefinitionInformer.
	ApplicationDefinitions() ApplicationDefinitionInformer
	// ApplicationInstallations returns a ApplicationInstallationInformer.
	ApplicationInstallations() ApplicationInstallationInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// ApplicationDefinitions returns a ApplicationDefinitionInformer.
func (v *version) ApplicationDefinitions() ApplicationDefinitionInformer {
	return &applicationDefinitionInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ApplicationInstallations returns a ApplicationInstallationInformer.
func (v *version) ApplicationInstallations() ApplicationInstallationInformer {
	return &applicationInstallationInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
