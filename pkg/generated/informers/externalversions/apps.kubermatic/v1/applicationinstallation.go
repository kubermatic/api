// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	appskubermaticv1 "k8c.io/api/v3/pkg/apis/apps.kubermatic/v1"
	versioned "k8c.io/api/v3/pkg/generated/clientset/versioned"
	internalinterfaces "k8c.io/api/v3/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "k8c.io/api/v3/pkg/generated/listers/apps.kubermatic/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ApplicationInstallationInformer provides access to a shared informer and lister for
// ApplicationInstallations.
type ApplicationInstallationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ApplicationInstallationLister
}

type applicationInstallationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewApplicationInstallationInformer constructs a new informer for ApplicationInstallation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewApplicationInstallationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredApplicationInstallationInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredApplicationInstallationInformer constructs a new informer for ApplicationInstallation type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredApplicationInstallationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubermaticAppsV1().ApplicationInstallations(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubermaticAppsV1().ApplicationInstallations(namespace).Watch(context.TODO(), options)
			},
		},
		&appskubermaticv1.ApplicationInstallation{},
		resyncPeriod,
		indexers,
	)
}

func (f *applicationInstallationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredApplicationInstallationInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *applicationInstallationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&appskubermaticv1.ApplicationInstallation{}, f.defaultInformer)
}

func (f *applicationInstallationInformer) Lister() v1.ApplicationInstallationLister {
	return v1.NewApplicationInstallationLister(f.Informer().GetIndexer())
}
