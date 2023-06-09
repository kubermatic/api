// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "k8c.io/api/v3/pkg/apis/kubermatic/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// EtcdBackupConfigLister helps list EtcdBackupConfigs.
// All objects returned here must be treated as read-only.
type EtcdBackupConfigLister interface {
	// List lists all EtcdBackupConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.EtcdBackupConfig, err error)
	// EtcdBackupConfigs returns an object that can list and get EtcdBackupConfigs.
	EtcdBackupConfigs(namespace string) EtcdBackupConfigNamespaceLister
	EtcdBackupConfigListerExpansion
}

// etcdBackupConfigLister implements the EtcdBackupConfigLister interface.
type etcdBackupConfigLister struct {
	indexer cache.Indexer
}

// NewEtcdBackupConfigLister returns a new EtcdBackupConfigLister.
func NewEtcdBackupConfigLister(indexer cache.Indexer) EtcdBackupConfigLister {
	return &etcdBackupConfigLister{indexer: indexer}
}

// List lists all EtcdBackupConfigs in the indexer.
func (s *etcdBackupConfigLister) List(selector labels.Selector) (ret []*v1.EtcdBackupConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.EtcdBackupConfig))
	})
	return ret, err
}

// EtcdBackupConfigs returns an object that can list and get EtcdBackupConfigs.
func (s *etcdBackupConfigLister) EtcdBackupConfigs(namespace string) EtcdBackupConfigNamespaceLister {
	return etcdBackupConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EtcdBackupConfigNamespaceLister helps list and get EtcdBackupConfigs.
// All objects returned here must be treated as read-only.
type EtcdBackupConfigNamespaceLister interface {
	// List lists all EtcdBackupConfigs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.EtcdBackupConfig, err error)
	// Get retrieves the EtcdBackupConfig from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.EtcdBackupConfig, error)
	EtcdBackupConfigNamespaceListerExpansion
}

// etcdBackupConfigNamespaceLister implements the EtcdBackupConfigNamespaceLister
// interface.
type etcdBackupConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all EtcdBackupConfigs in the indexer for a given namespace.
func (s etcdBackupConfigNamespaceLister) List(selector labels.Selector) (ret []*v1.EtcdBackupConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.EtcdBackupConfig))
	})
	return ret, err
}

// Get retrieves the EtcdBackupConfig from the indexer for a given namespace and name.
func (s etcdBackupConfigNamespaceLister) Get(name string) (*v1.EtcdBackupConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("etcdbackupconfig"), name)
	}
	return obj.(*v1.EtcdBackupConfig), nil
}
