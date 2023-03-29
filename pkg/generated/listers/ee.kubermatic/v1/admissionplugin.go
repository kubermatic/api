// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "k8c.io/api/v3/pkg/apis/ee.kubermatic/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AdmissionPluginLister helps list AdmissionPlugins.
// All objects returned here must be treated as read-only.
type AdmissionPluginLister interface {
	// List lists all AdmissionPlugins in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.AdmissionPlugin, err error)
	// AdmissionPlugins returns an object that can list and get AdmissionPlugins.
	AdmissionPlugins(namespace string) AdmissionPluginNamespaceLister
	AdmissionPluginListerExpansion
}

// admissionPluginLister implements the AdmissionPluginLister interface.
type admissionPluginLister struct {
	indexer cache.Indexer
}

// NewAdmissionPluginLister returns a new AdmissionPluginLister.
func NewAdmissionPluginLister(indexer cache.Indexer) AdmissionPluginLister {
	return &admissionPluginLister{indexer: indexer}
}

// List lists all AdmissionPlugins in the indexer.
func (s *admissionPluginLister) List(selector labels.Selector) (ret []*v1.AdmissionPlugin, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AdmissionPlugin))
	})
	return ret, err
}

// AdmissionPlugins returns an object that can list and get AdmissionPlugins.
func (s *admissionPluginLister) AdmissionPlugins(namespace string) AdmissionPluginNamespaceLister {
	return admissionPluginNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AdmissionPluginNamespaceLister helps list and get AdmissionPlugins.
// All objects returned here must be treated as read-only.
type AdmissionPluginNamespaceLister interface {
	// List lists all AdmissionPlugins in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.AdmissionPlugin, err error)
	// Get retrieves the AdmissionPlugin from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.AdmissionPlugin, error)
	AdmissionPluginNamespaceListerExpansion
}

// admissionPluginNamespaceLister implements the AdmissionPluginNamespaceLister
// interface.
type admissionPluginNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AdmissionPlugins in the indexer for a given namespace.
func (s admissionPluginNamespaceLister) List(selector labels.Selector) (ret []*v1.AdmissionPlugin, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AdmissionPlugin))
	})
	return ret, err
}

// Get retrieves the AdmissionPlugin from the indexer for a given namespace and name.
func (s admissionPluginNamespaceLister) Get(name string) (*v1.AdmissionPlugin, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("admissionplugin"), name)
	}
	return obj.(*v1.AdmissionPlugin), nil
}