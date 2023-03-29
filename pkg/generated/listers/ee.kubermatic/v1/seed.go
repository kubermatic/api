// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "k8c.io/api/v3/pkg/apis/ee.kubermatic/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// SeedLister helps list Seeds.
// All objects returned here must be treated as read-only.
type SeedLister interface {
	// List lists all Seeds in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Seed, err error)
	// Seeds returns an object that can list and get Seeds.
	Seeds(namespace string) SeedNamespaceLister
	SeedListerExpansion
}

// seedLister implements the SeedLister interface.
type seedLister struct {
	indexer cache.Indexer
}

// NewSeedLister returns a new SeedLister.
func NewSeedLister(indexer cache.Indexer) SeedLister {
	return &seedLister{indexer: indexer}
}

// List lists all Seeds in the indexer.
func (s *seedLister) List(selector labels.Selector) (ret []*v1.Seed, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Seed))
	})
	return ret, err
}

// Seeds returns an object that can list and get Seeds.
func (s *seedLister) Seeds(namespace string) SeedNamespaceLister {
	return seedNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SeedNamespaceLister helps list and get Seeds.
// All objects returned here must be treated as read-only.
type SeedNamespaceLister interface {
	// List lists all Seeds in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Seed, err error)
	// Get retrieves the Seed from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Seed, error)
	SeedNamespaceListerExpansion
}

// seedNamespaceLister implements the SeedNamespaceLister
// interface.
type seedNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Seeds in the indexer for a given namespace.
func (s seedNamespaceLister) List(selector labels.Selector) (ret []*v1.Seed, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Seed))
	})
	return ret, err
}

// Get retrieves the Seed from the indexer for a given namespace and name.
func (s seedNamespaceLister) Get(name string) (*v1.Seed, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("seed"), name)
	}
	return obj.(*v1.Seed), nil
}
