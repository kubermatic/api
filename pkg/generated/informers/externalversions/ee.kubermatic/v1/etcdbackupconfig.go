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

// EtcdBackupConfigInformer provides access to a shared informer and lister for
// EtcdBackupConfigs.
type EtcdBackupConfigInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.EtcdBackupConfigLister
}

type etcdBackupConfigInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewEtcdBackupConfigInformer constructs a new informer for EtcdBackupConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewEtcdBackupConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredEtcdBackupConfigInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredEtcdBackupConfigInformer constructs a new informer for EtcdBackupConfig type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredEtcdBackupConfigInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubermaticEnterpriseV1().EtcdBackupConfigs(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubermaticEnterpriseV1().EtcdBackupConfigs(namespace).Watch(context.TODO(), options)
			},
		},
		&eekubermaticv1.EtcdBackupConfig{},
		resyncPeriod,
		indexers,
	)
}

func (f *etcdBackupConfigInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredEtcdBackupConfigInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *etcdBackupConfigInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&eekubermaticv1.EtcdBackupConfig{}, f.defaultInformer)
}

func (f *etcdBackupConfigInformer) Lister() v1.EtcdBackupConfigLister {
	return v1.NewEtcdBackupConfigLister(f.Informer().GetIndexer())
}
