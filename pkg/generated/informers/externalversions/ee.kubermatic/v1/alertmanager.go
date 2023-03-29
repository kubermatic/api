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

// AlertmanagerInformer provides access to a shared informer and lister for
// Alertmanagers.
type AlertmanagerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.AlertmanagerLister
}

type alertmanagerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewAlertmanagerInformer constructs a new informer for Alertmanager type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAlertmanagerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAlertmanagerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredAlertmanagerInformer constructs a new informer for Alertmanager type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAlertmanagerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubermaticEnterpriseV1().Alertmanagers(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubermaticEnterpriseV1().Alertmanagers(namespace).Watch(context.TODO(), options)
			},
		},
		&eekubermaticv1.Alertmanager{},
		resyncPeriod,
		indexers,
	)
}

func (f *alertmanagerInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAlertmanagerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *alertmanagerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&eekubermaticv1.Alertmanager{}, f.defaultInformer)
}

func (f *alertmanagerInformer) Lister() v1.AlertmanagerLister {
	return v1.NewAlertmanagerLister(f.Informer().GetIndexer())
}
