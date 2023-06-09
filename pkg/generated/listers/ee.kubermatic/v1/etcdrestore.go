// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "k8c.io/api/v3/pkg/apis/ee.kubermatic/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// EtcdRestoreLister helps list EtcdRestores.
// All objects returned here must be treated as read-only.
type EtcdRestoreLister interface {
	// List lists all EtcdRestores in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.EtcdRestore, err error)
	// EtcdRestores returns an object that can list and get EtcdRestores.
	EtcdRestores(namespace string) EtcdRestoreNamespaceLister
	EtcdRestoreListerExpansion
}

// etcdRestoreLister implements the EtcdRestoreLister interface.
type etcdRestoreLister struct {
	indexer cache.Indexer
}

// NewEtcdRestoreLister returns a new EtcdRestoreLister.
func NewEtcdRestoreLister(indexer cache.Indexer) EtcdRestoreLister {
	return &etcdRestoreLister{indexer: indexer}
}

// List lists all EtcdRestores in the indexer.
func (s *etcdRestoreLister) List(selector labels.Selector) (ret []*v1.EtcdRestore, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.EtcdRestore))
	})
	return ret, err
}

// EtcdRestores returns an object that can list and get EtcdRestores.
func (s *etcdRestoreLister) EtcdRestores(namespace string) EtcdRestoreNamespaceLister {
	return etcdRestoreNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EtcdRestoreNamespaceLister helps list and get EtcdRestores.
// All objects returned here must be treated as read-only.
type EtcdRestoreNamespaceLister interface {
	// List lists all EtcdRestores in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.EtcdRestore, err error)
	// Get retrieves the EtcdRestore from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.EtcdRestore, error)
	EtcdRestoreNamespaceListerExpansion
}

// etcdRestoreNamespaceLister implements the EtcdRestoreNamespaceLister
// interface.
type etcdRestoreNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EtcdRestores in the indexer for a given namespace.
func (s etcdRestoreNamespaceLister) List(selector labels.Selector) (ret []*v1.EtcdRestore, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.EtcdRestore))
	})
	return ret, err
}

// Get retrieves the EtcdRestore from the indexer for a given namespace and name.
func (s etcdRestoreNamespaceLister) Get(name string) (*v1.EtcdRestore, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("etcdrestore"), name)
	}
	return obj.(*v1.EtcdRestore), nil
}
