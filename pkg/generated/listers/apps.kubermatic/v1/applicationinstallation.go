// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "k8c.io/api/v2/pkg/apis/apps.kubermatic/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ApplicationInstallationLister helps list ApplicationInstallations.
// All objects returned here must be treated as read-only.
type ApplicationInstallationLister interface {
	// List lists all ApplicationInstallations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ApplicationInstallation, err error)
	// ApplicationInstallations returns an object that can list and get ApplicationInstallations.
	ApplicationInstallations(namespace string) ApplicationInstallationNamespaceLister
	ApplicationInstallationListerExpansion
}

// applicationInstallationLister implements the ApplicationInstallationLister interface.
type applicationInstallationLister struct {
	indexer cache.Indexer
}

// NewApplicationInstallationLister returns a new ApplicationInstallationLister.
func NewApplicationInstallationLister(indexer cache.Indexer) ApplicationInstallationLister {
	return &applicationInstallationLister{indexer: indexer}
}

// List lists all ApplicationInstallations in the indexer.
func (s *applicationInstallationLister) List(selector labels.Selector) (ret []*v1.ApplicationInstallation, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ApplicationInstallation))
	})
	return ret, err
}

// ApplicationInstallations returns an object that can list and get ApplicationInstallations.
func (s *applicationInstallationLister) ApplicationInstallations(namespace string) ApplicationInstallationNamespaceLister {
	return applicationInstallationNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ApplicationInstallationNamespaceLister helps list and get ApplicationInstallations.
// All objects returned here must be treated as read-only.
type ApplicationInstallationNamespaceLister interface {
	// List lists all ApplicationInstallations in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ApplicationInstallation, err error)
	// Get retrieves the ApplicationInstallation from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ApplicationInstallation, error)
	ApplicationInstallationNamespaceListerExpansion
}

// applicationInstallationNamespaceLister implements the ApplicationInstallationNamespaceLister
// interface.
type applicationInstallationNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ApplicationInstallations in the indexer for a given namespace.
func (s applicationInstallationNamespaceLister) List(selector labels.Selector) (ret []*v1.ApplicationInstallation, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ApplicationInstallation))
	})
	return ret, err
}

// Get retrieves the ApplicationInstallation from the indexer for a given namespace and name.
func (s applicationInstallationNamespaceLister) Get(name string) (*v1.ApplicationInstallation, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("applicationinstallation"), name)
	}
	return obj.(*v1.ApplicationInstallation), nil
}
