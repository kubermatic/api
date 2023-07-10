// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8c.io/api/v3/pkg/apis/kubermatic/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEtcdRestores implements EtcdRestoreInterface
type FakeEtcdRestores struct {
	Fake *FakeKubermaticV1
	ns   string
}

var etcdrestoresResource = v1.SchemeGroupVersion.WithResource("etcdrestores")

var etcdrestoresKind = v1.SchemeGroupVersion.WithKind("EtcdRestore")

// Get takes name of the etcdRestore, and returns the corresponding etcdRestore object, and an error if there is any.
func (c *FakeEtcdRestores) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.EtcdRestore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(etcdrestoresResource, c.ns, name), &v1.EtcdRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.EtcdRestore), err
}

// List takes label and field selectors, and returns the list of EtcdRestores that match those selectors.
func (c *FakeEtcdRestores) List(ctx context.Context, opts metav1.ListOptions) (result *v1.EtcdRestoreList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(etcdrestoresResource, etcdrestoresKind, c.ns, opts), &v1.EtcdRestoreList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.EtcdRestoreList{ListMeta: obj.(*v1.EtcdRestoreList).ListMeta}
	for _, item := range obj.(*v1.EtcdRestoreList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested etcdRestores.
func (c *FakeEtcdRestores) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(etcdrestoresResource, c.ns, opts))

}

// Create takes the representation of a etcdRestore and creates it.  Returns the server's representation of the etcdRestore, and an error, if there is any.
func (c *FakeEtcdRestores) Create(ctx context.Context, etcdRestore *v1.EtcdRestore, opts metav1.CreateOptions) (result *v1.EtcdRestore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(etcdrestoresResource, c.ns, etcdRestore), &v1.EtcdRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.EtcdRestore), err
}

// Update takes the representation of a etcdRestore and updates it. Returns the server's representation of the etcdRestore, and an error, if there is any.
func (c *FakeEtcdRestores) Update(ctx context.Context, etcdRestore *v1.EtcdRestore, opts metav1.UpdateOptions) (result *v1.EtcdRestore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(etcdrestoresResource, c.ns, etcdRestore), &v1.EtcdRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.EtcdRestore), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEtcdRestores) UpdateStatus(ctx context.Context, etcdRestore *v1.EtcdRestore, opts metav1.UpdateOptions) (*v1.EtcdRestore, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(etcdrestoresResource, "status", c.ns, etcdRestore), &v1.EtcdRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.EtcdRestore), err
}

// Delete takes name of the etcdRestore and deletes it. Returns an error if one occurs.
func (c *FakeEtcdRestores) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(etcdrestoresResource, c.ns, name, opts), &v1.EtcdRestore{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEtcdRestores) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(etcdrestoresResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.EtcdRestoreList{})
	return err
}

// Patch applies the patch and returns the patched etcdRestore.
func (c *FakeEtcdRestores) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.EtcdRestore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(etcdrestoresResource, c.ns, name, pt, data, subresources...), &v1.EtcdRestore{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.EtcdRestore), err
}
