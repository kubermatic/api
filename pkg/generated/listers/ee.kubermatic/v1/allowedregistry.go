// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "k8c.io/api/v3/pkg/apis/ee.kubermatic/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AllowedRegistryLister helps list AllowedRegistries.
// All objects returned here must be treated as read-only.
type AllowedRegistryLister interface {
	// List lists all AllowedRegistries in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.AllowedRegistry, err error)
	// AllowedRegistries returns an object that can list and get AllowedRegistries.
	AllowedRegistries(namespace string) AllowedRegistryNamespaceLister
	AllowedRegistryListerExpansion
}

// allowedRegistryLister implements the AllowedRegistryLister interface.
type allowedRegistryLister struct {
	indexer cache.Indexer
}

// NewAllowedRegistryLister returns a new AllowedRegistryLister.
func NewAllowedRegistryLister(indexer cache.Indexer) AllowedRegistryLister {
	return &allowedRegistryLister{indexer: indexer}
}

// List lists all AllowedRegistries in the indexer.
func (s *allowedRegistryLister) List(selector labels.Selector) (ret []*v1.AllowedRegistry, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AllowedRegistry))
	})
	return ret, err
}

// AllowedRegistries returns an object that can list and get AllowedRegistries.
func (s *allowedRegistryLister) AllowedRegistries(namespace string) AllowedRegistryNamespaceLister {
	return allowedRegistryNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AllowedRegistryNamespaceLister helps list and get AllowedRegistries.
// All objects returned here must be treated as read-only.
type AllowedRegistryNamespaceLister interface {
	// List lists all AllowedRegistries in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.AllowedRegistry, err error)
	// Get retrieves the AllowedRegistry from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.AllowedRegistry, error)
	AllowedRegistryNamespaceListerExpansion
}

// allowedRegistryNamespaceLister implements the AllowedRegistryNamespaceLister
// interface.
type allowedRegistryNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AllowedRegistries in the indexer for a given namespace.
func (s allowedRegistryNamespaceLister) List(selector labels.Selector) (ret []*v1.AllowedRegistry, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AllowedRegistry))
	})
	return ret, err
}

// Get retrieves the AllowedRegistry from the indexer for a given namespace and name.
func (s allowedRegistryNamespaceLister) Get(name string) (*v1.AllowedRegistry, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("allowedregistry"), name)
	}
	return obj.(*v1.AllowedRegistry), nil
}
