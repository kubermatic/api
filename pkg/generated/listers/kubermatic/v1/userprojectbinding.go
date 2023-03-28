// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "k8c.io/api/v3/pkg/apis/kubermatic/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// UserProjectBindingLister helps list UserProjectBindings.
// All objects returned here must be treated as read-only.
type UserProjectBindingLister interface {
	// List lists all UserProjectBindings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.UserProjectBinding, err error)
	// UserProjectBindings returns an object that can list and get UserProjectBindings.
	UserProjectBindings(namespace string) UserProjectBindingNamespaceLister
	UserProjectBindingListerExpansion
}

// userProjectBindingLister implements the UserProjectBindingLister interface.
type userProjectBindingLister struct {
	indexer cache.Indexer
}

// NewUserProjectBindingLister returns a new UserProjectBindingLister.
func NewUserProjectBindingLister(indexer cache.Indexer) UserProjectBindingLister {
	return &userProjectBindingLister{indexer: indexer}
}

// List lists all UserProjectBindings in the indexer.
func (s *userProjectBindingLister) List(selector labels.Selector) (ret []*v1.UserProjectBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.UserProjectBinding))
	})
	return ret, err
}

// UserProjectBindings returns an object that can list and get UserProjectBindings.
func (s *userProjectBindingLister) UserProjectBindings(namespace string) UserProjectBindingNamespaceLister {
	return userProjectBindingNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// UserProjectBindingNamespaceLister helps list and get UserProjectBindings.
// All objects returned here must be treated as read-only.
type UserProjectBindingNamespaceLister interface {
	// List lists all UserProjectBindings in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.UserProjectBinding, err error)
	// Get retrieves the UserProjectBinding from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.UserProjectBinding, error)
	UserProjectBindingNamespaceListerExpansion
}

// userProjectBindingNamespaceLister implements the UserProjectBindingNamespaceLister
// interface.
type userProjectBindingNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all UserProjectBindings in the indexer for a given namespace.
func (s userProjectBindingNamespaceLister) List(selector labels.Selector) (ret []*v1.UserProjectBinding, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.UserProjectBinding))
	})
	return ret, err
}

// Get retrieves the UserProjectBinding from the indexer for a given namespace and name.
func (s userProjectBindingNamespaceLister) Get(name string) (*v1.UserProjectBinding, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("userprojectbinding"), name)
	}
	return obj.(*v1.UserProjectBinding), nil
}
