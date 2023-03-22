// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	appskubermaticv1 "k8c.io/api/v2/pkg/apis/apps.kubermatic/v1"
	versioned "k8c.io/api/v2/pkg/generated/clientset/versioned"
	internalinterfaces "k8c.io/api/v2/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "k8c.io/api/v2/pkg/generated/listers/apps.kubermatic/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ApplicationDefinitionInformer provides access to a shared informer and lister for
// ApplicationDefinitions.
type ApplicationDefinitionInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ApplicationDefinitionLister
}

type applicationDefinitionInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewApplicationDefinitionInformer constructs a new informer for ApplicationDefinition type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewApplicationDefinitionInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredApplicationDefinitionInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredApplicationDefinitionInformer constructs a new informer for ApplicationDefinition type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredApplicationDefinitionInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppsV1().ApplicationDefinitions(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AppsV1().ApplicationDefinitions(namespace).Watch(context.TODO(), options)
			},
		},
		&appskubermaticv1.ApplicationDefinition{},
		resyncPeriod,
		indexers,
	)
}

func (f *applicationDefinitionInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredApplicationDefinitionInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *applicationDefinitionInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&appskubermaticv1.ApplicationDefinition{}, f.defaultInformer)
}

func (f *applicationDefinitionInformer) Lister() v1.ApplicationDefinitionLister {
	return v1.NewApplicationDefinitionLister(f.Informer().GetIndexer())
}
