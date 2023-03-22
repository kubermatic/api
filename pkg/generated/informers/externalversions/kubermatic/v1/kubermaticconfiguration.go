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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	kubermaticv1 "k8c.io/api/v2/pkg/apis/kubermatic/v1"
	versioned "k8c.io/api/v2/pkg/generated/clientset/versioned"
	internalinterfaces "k8c.io/api/v2/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "k8c.io/api/v2/pkg/generated/listers/kubermatic/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// KubermaticConfigurationInformer provides access to a shared informer and lister for
// KubermaticConfigurations.
type KubermaticConfigurationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.KubermaticConfigurationLister
}

type kubermaticConfigurationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewKubermaticConfigurationInformer constructs a new informer for KubermaticConfiguration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewKubermaticConfigurationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredKubermaticConfigurationInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredKubermaticConfigurationInformer constructs a new informer for KubermaticConfiguration type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredKubermaticConfigurationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubermaticV1().KubermaticConfigurations(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubermaticV1().KubermaticConfigurations(namespace).Watch(context.TODO(), options)
			},
		},
		&kubermaticv1.KubermaticConfiguration{},
		resyncPeriod,
		indexers,
	)
}

func (f *kubermaticConfigurationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredKubermaticConfigurationInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *kubermaticConfigurationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kubermaticv1.KubermaticConfiguration{}, f.defaultInformer)
}

func (f *kubermaticConfigurationInformer) Lister() v1.KubermaticConfigurationLister {
	return v1.NewKubermaticConfigurationLister(f.Informer().GetIndexer())
}
