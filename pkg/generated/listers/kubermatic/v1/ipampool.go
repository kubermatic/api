// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "k8c.io/api/v3/pkg/apis/kubermatic/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// IPAMPoolLister helps list IPAMPools.
// All objects returned here must be treated as read-only.
type IPAMPoolLister interface {
	// List lists all IPAMPools in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.IPAMPool, err error)
	// IPAMPools returns an object that can list and get IPAMPools.
	IPAMPools(namespace string) IPAMPoolNamespaceLister
	IPAMPoolListerExpansion
}

// iPAMPoolLister implements the IPAMPoolLister interface.
type iPAMPoolLister struct {
	indexer cache.Indexer
}

// NewIPAMPoolLister returns a new IPAMPoolLister.
func NewIPAMPoolLister(indexer cache.Indexer) IPAMPoolLister {
	return &iPAMPoolLister{indexer: indexer}
}

// List lists all IPAMPools in the indexer.
func (s *iPAMPoolLister) List(selector labels.Selector) (ret []*v1.IPAMPool, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.IPAMPool))
	})
	return ret, err
}

// IPAMPools returns an object that can list and get IPAMPools.
func (s *iPAMPoolLister) IPAMPools(namespace string) IPAMPoolNamespaceLister {
	return iPAMPoolNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// IPAMPoolNamespaceLister helps list and get IPAMPools.
// All objects returned here must be treated as read-only.
type IPAMPoolNamespaceLister interface {
	// List lists all IPAMPools in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.IPAMPool, err error)
	// Get retrieves the IPAMPool from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.IPAMPool, error)
	IPAMPoolNamespaceListerExpansion
}

// iPAMPoolNamespaceLister implements the IPAMPoolNamespaceLister
// interface.
type iPAMPoolNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all IPAMPools in the indexer for a given namespace.
func (s iPAMPoolNamespaceLister) List(selector labels.Selector) (ret []*v1.IPAMPool, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.IPAMPool))
	})
	return ret, err
}

// Get retrieves the IPAMPool from the indexer for a given namespace and name.
func (s iPAMPoolNamespaceLister) Get(name string) (*v1.IPAMPool, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("ipampool"), name)
	}
	return obj.(*v1.IPAMPool), nil
}
