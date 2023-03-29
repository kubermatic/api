// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	kubermaticv1 "k8c.io/api/v3/pkg/apis/kubermatic/v1"
	versioned "k8c.io/api/v3/pkg/generated/clientset/versioned"
	internalinterfaces "k8c.io/api/v3/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "k8c.io/api/v3/pkg/generated/listers/kubermatic/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// UserSSHKeyInformer provides access to a shared informer and lister for
// UserSSHKeys.
type UserSSHKeyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.UserSSHKeyLister
}

type userSSHKeyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewUserSSHKeyInformer constructs a new informer for UserSSHKey type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewUserSSHKeyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredUserSSHKeyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredUserSSHKeyInformer constructs a new informer for UserSSHKey type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredUserSSHKeyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubermaticV1().UserSSHKeys(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubermaticV1().UserSSHKeys(namespace).Watch(context.TODO(), options)
			},
		},
		&kubermaticv1.UserSSHKey{},
		resyncPeriod,
		indexers,
	)
}

func (f *userSSHKeyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredUserSSHKeyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *userSSHKeyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kubermaticv1.UserSSHKey{}, f.defaultInformer)
}

func (f *userSSHKeyInformer) Lister() v1.UserSSHKeyLister {
	return v1.NewUserSSHKeyLister(f.Informer().GetIndexer())
}