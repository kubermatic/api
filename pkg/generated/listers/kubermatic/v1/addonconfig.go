/*
Copyright 2023 The Kubermatic Kubernetes Platform contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "k8c.io/api/v2/pkg/apis/kubermatic/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AddonConfigLister helps list AddonConfigs.
// All objects returned here must be treated as read-only.
type AddonConfigLister interface {
	// List lists all AddonConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.AddonConfig, err error)
	// AddonConfigs returns an object that can list and get AddonConfigs.
	AddonConfigs(namespace string) AddonConfigNamespaceLister
	AddonConfigListerExpansion
}

// addonConfigLister implements the AddonConfigLister interface.
type addonConfigLister struct {
	indexer cache.Indexer
}

// NewAddonConfigLister returns a new AddonConfigLister.
func NewAddonConfigLister(indexer cache.Indexer) AddonConfigLister {
	return &addonConfigLister{indexer: indexer}
}

// List lists all AddonConfigs in the indexer.
func (s *addonConfigLister) List(selector labels.Selector) (ret []*v1.AddonConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AddonConfig))
	})
	return ret, err
}

// AddonConfigs returns an object that can list and get AddonConfigs.
func (s *addonConfigLister) AddonConfigs(namespace string) AddonConfigNamespaceLister {
	return addonConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AddonConfigNamespaceLister helps list and get AddonConfigs.
// All objects returned here must be treated as read-only.
type AddonConfigNamespaceLister interface {
	// List lists all AddonConfigs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.AddonConfig, err error)
	// Get retrieves the AddonConfig from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.AddonConfig, error)
	AddonConfigNamespaceListerExpansion
}

// addonConfigNamespaceLister implements the AddonConfigNamespaceLister
// interface.
type addonConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AddonConfigs in the indexer for a given namespace.
func (s addonConfigNamespaceLister) List(selector labels.Selector) (ret []*v1.AddonConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AddonConfig))
	})
	return ret, err
}

// Get retrieves the AddonConfig from the indexer for a given namespace and name.
func (s addonConfigNamespaceLister) Get(name string) (*v1.AddonConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("addonconfig"), name)
	}
	return obj.(*v1.AddonConfig), nil
}
