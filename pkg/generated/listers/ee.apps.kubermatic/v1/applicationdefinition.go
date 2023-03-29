// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "k8c.io/api/v3/pkg/apis/ee.apps.kubermatic/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ApplicationDefinitionLister helps list ApplicationDefinitions.
// All objects returned here must be treated as read-only.
type ApplicationDefinitionLister interface {
	// List lists all ApplicationDefinitions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ApplicationDefinition, err error)
	// ApplicationDefinitions returns an object that can list and get ApplicationDefinitions.
	ApplicationDefinitions(namespace string) ApplicationDefinitionNamespaceLister
	ApplicationDefinitionListerExpansion
}

// applicationDefinitionLister implements the ApplicationDefinitionLister interface.
type applicationDefinitionLister struct {
	indexer cache.Indexer
}

// NewApplicationDefinitionLister returns a new ApplicationDefinitionLister.
func NewApplicationDefinitionLister(indexer cache.Indexer) ApplicationDefinitionLister {
	return &applicationDefinitionLister{indexer: indexer}
}

// List lists all ApplicationDefinitions in the indexer.
func (s *applicationDefinitionLister) List(selector labels.Selector) (ret []*v1.ApplicationDefinition, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ApplicationDefinition))
	})
	return ret, err
}

// ApplicationDefinitions returns an object that can list and get ApplicationDefinitions.
func (s *applicationDefinitionLister) ApplicationDefinitions(namespace string) ApplicationDefinitionNamespaceLister {
	return applicationDefinitionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ApplicationDefinitionNamespaceLister helps list and get ApplicationDefinitions.
// All objects returned here must be treated as read-only.
type ApplicationDefinitionNamespaceLister interface {
	// List lists all ApplicationDefinitions in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ApplicationDefinition, err error)
	// Get retrieves the ApplicationDefinition from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ApplicationDefinition, error)
	ApplicationDefinitionNamespaceListerExpansion
}

// applicationDefinitionNamespaceLister implements the ApplicationDefinitionNamespaceLister
// interface.
type applicationDefinitionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ApplicationDefinitions in the indexer for a given namespace.
func (s applicationDefinitionNamespaceLister) List(selector labels.Selector) (ret []*v1.ApplicationDefinition, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ApplicationDefinition))
	})
	return ret, err
}

// Get retrieves the ApplicationDefinition from the indexer for a given namespace and name.
func (s applicationDefinitionNamespaceLister) Get(name string) (*v1.ApplicationDefinition, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("applicationdefinition"), name)
	}
	return obj.(*v1.ApplicationDefinition), nil
}
