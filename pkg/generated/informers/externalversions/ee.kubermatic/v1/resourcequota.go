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

// ResourceQuotaInformer provides access to a shared informer and lister for
// ResourceQuotas.
type ResourceQuotaInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ResourceQuotaLister
}

type resourceQuotaInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewResourceQuotaInformer constructs a new informer for ResourceQuota type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewResourceQuotaInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredResourceQuotaInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredResourceQuotaInformer constructs a new informer for ResourceQuota type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredResourceQuotaInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubermaticEnterpriseV1().ResourceQuotas(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubermaticEnterpriseV1().ResourceQuotas(namespace).Watch(context.TODO(), options)
			},
		},
		&eekubermaticv1.ResourceQuota{},
		resyncPeriod,
		indexers,
	)
}

func (f *resourceQuotaInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredResourceQuotaInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *resourceQuotaInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&eekubermaticv1.ResourceQuota{}, f.defaultInformer)
}

func (f *resourceQuotaInformer) Lister() v1.ResourceQuotaLister {
	return v1.NewResourceQuotaLister(f.Informer().GetIndexer())
}
