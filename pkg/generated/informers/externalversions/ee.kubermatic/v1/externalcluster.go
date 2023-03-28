// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	eekubermaticv1 "k8c.io/api/v3/pkg/apis/ee.kubermatic/v1"
	versioned "k8c.io/api/v3/pkg/generated/clientset/versioned"
	internalinterfaces "k8c.io/api/v3/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "k8c.io/api/v3/pkg/generated/listers/ee.kubermatic/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ExternalClusterInformer provides access to a shared informer and lister for
// ExternalClusters.
type ExternalClusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ExternalClusterLister
}

type externalClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewExternalClusterInformer constructs a new informer for ExternalCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewExternalClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredExternalClusterInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredExternalClusterInformer constructs a new informer for ExternalCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredExternalClusterInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EeKubermaticV1().ExternalClusters(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.EeKubermaticV1().ExternalClusters(namespace).Watch(context.TODO(), options)
			},
		},
		&eekubermaticv1.ExternalCluster{},
		resyncPeriod,
		indexers,
	)
}

func (f *externalClusterInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredExternalClusterInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *externalClusterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&eekubermaticv1.ExternalCluster{}, f.defaultInformer)
}

func (f *externalClusterInformer) Lister() v1.ExternalClusterLister {
	return v1.NewExternalClusterLister(f.Informer().GetIndexer())
}
