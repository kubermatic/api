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

// SSHKeyBindingInformer provides access to a shared informer and lister for
// SSHKeyBindings.
type SSHKeyBindingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.SSHKeyBindingLister
}

type sSHKeyBindingInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSSHKeyBindingInformer constructs a new informer for SSHKeyBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSSHKeyBindingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSSHKeyBindingInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSSHKeyBindingInformer constructs a new informer for SSHKeyBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSSHKeyBindingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubermaticV1().SSHKeyBindings(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.KubermaticV1().SSHKeyBindings(namespace).Watch(context.TODO(), options)
			},
		},
		&kubermaticv1.SSHKeyBinding{},
		resyncPeriod,
		indexers,
	)
}

func (f *sSHKeyBindingInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSSHKeyBindingInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *sSHKeyBindingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&kubermaticv1.SSHKeyBinding{}, f.defaultInformer)
}

func (f *sSHKeyBindingInformer) Lister() v1.SSHKeyBindingLister {
	return v1.NewSSHKeyBindingLister(f.Informer().GetIndexer())
}
