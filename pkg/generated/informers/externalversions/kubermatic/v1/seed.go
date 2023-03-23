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

// SeedInformer provides access to a shared informer and lister for
// Seeds.
type SeedInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.SeedLister
}

type seedInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSeedInformer constructs a new informer for Seed type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSeedInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSeedInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSeedInformer constructs a new informer for Seed type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSeedInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubermaticV1().Seeds(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubermaticV1().Seeds(namespace).Watch(context.TODO(), options)
			},
		},
		&kubermaticv1.Seed{},
		resyncPeriod,
		indexers,
	)
}

func (f *seedInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSeedInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *seedInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kubermaticv1.Seed{}, f.defaultInformer)
}

func (f *seedInformer) Lister() v1.SeedLister {
	return v1.NewSeedLister(f.Informer().GetIndexer())
}
