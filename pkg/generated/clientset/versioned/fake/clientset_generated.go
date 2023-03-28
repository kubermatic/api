// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	clientset "k8c.io/api/v3/pkg/generated/clientset/versioned"
	appskubermaticv1 "k8c.io/api/v3/pkg/generated/clientset/versioned/typed/apps.kubermatic/v1"
	fakeappskubermaticv1 "k8c.io/api/v3/pkg/generated/clientset/versioned/typed/apps.kubermatic/v1/fake"
	eekubermaticv1 "k8c.io/api/v3/pkg/generated/clientset/versioned/typed/ee.kubermatic/v1"
	fakeeekubermaticv1 "k8c.io/api/v3/pkg/generated/clientset/versioned/typed/ee.kubermatic/v1/fake"
	kubermaticv1 "k8c.io/api/v3/pkg/generated/clientset/versioned/typed/kubermatic/v1"
	fakekubermaticv1 "k8c.io/api/v3/pkg/generated/clientset/versioned/typed/kubermatic/v1/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var (
	_ clientset.Interface = &Clientset{}
	_ testing.FakeClient  = &Clientset{}
)

// AppsKubermaticV1 retrieves the AppsKubermaticV1Client
func (c *Clientset) AppsKubermaticV1() appskubermaticv1.AppsKubermaticV1Interface {
	return &fakeappskubermaticv1.FakeAppsKubermaticV1{Fake: &c.Fake}
}

// EeKubermaticV1 retrieves the EeKubermaticV1Client
func (c *Clientset) EeKubermaticV1() eekubermaticv1.EeKubermaticV1Interface {
	return &fakeeekubermaticv1.FakeEeKubermaticV1{Fake: &c.Fake}
}

// KubermaticV1 retrieves the KubermaticV1Client
func (c *Clientset) KubermaticV1() kubermaticv1.KubermaticV1Interface {
	return &fakekubermaticv1.FakeKubermaticV1{Fake: &c.Fake}
}
