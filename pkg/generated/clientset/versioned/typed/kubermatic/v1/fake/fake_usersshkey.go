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

// FakeUserSSHKeys implements UserSSHKeyInterface
type FakeUserSSHKeys struct {
	Fake *FakeKubermaticV1
	ns   string
}

var usersshkeysResource = v1.SchemeGroupVersion.WithResource("usersshkeys")

var usersshkeysKind = v1.SchemeGroupVersion.WithKind("UserSSHKey")

// Get takes name of the userSSHKey, and returns the corresponding userSSHKey object, and an error if there is any.
func (c *FakeUserSSHKeys) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.UserSSHKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(usersshkeysResource, c.ns, name), &v1.UserSSHKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.UserSSHKey), err
}

// List takes label and field selectors, and returns the list of UserSSHKeys that match those selectors.
func (c *FakeUserSSHKeys) List(ctx context.Context, opts metav1.ListOptions) (result *v1.UserSSHKeyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(usersshkeysResource, usersshkeysKind, c.ns, opts), &v1.UserSSHKeyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.UserSSHKeyList{ListMeta: obj.(*v1.UserSSHKeyList).ListMeta}
	for _, item := range obj.(*v1.UserSSHKeyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested userSSHKeys.
func (c *FakeUserSSHKeys) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(usersshkeysResource, c.ns, opts))

}

// Create takes the representation of a userSSHKey and creates it.  Returns the server's representation of the userSSHKey, and an error, if there is any.
func (c *FakeUserSSHKeys) Create(ctx context.Context, userSSHKey *v1.UserSSHKey, opts metav1.CreateOptions) (result *v1.UserSSHKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(usersshkeysResource, c.ns, userSSHKey), &v1.UserSSHKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.UserSSHKey), err
}

// Update takes the representation of a userSSHKey and updates it. Returns the server's representation of the userSSHKey, and an error, if there is any.
func (c *FakeUserSSHKeys) Update(ctx context.Context, userSSHKey *v1.UserSSHKey, opts metav1.UpdateOptions) (result *v1.UserSSHKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(usersshkeysResource, c.ns, userSSHKey), &v1.UserSSHKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.UserSSHKey), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeUserSSHKeys) UpdateStatus(ctx context.Context, userSSHKey *v1.UserSSHKey, opts metav1.UpdateOptions) (*v1.UserSSHKey, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(usersshkeysResource, "status", c.ns, userSSHKey), &v1.UserSSHKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.UserSSHKey), err
}

// Delete takes name of the userSSHKey and deletes it. Returns an error if one occurs.
func (c *FakeUserSSHKeys) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(usersshkeysResource, c.ns, name, opts), &v1.UserSSHKey{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeUserSSHKeys) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(usersshkeysResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.UserSSHKeyList{})
	return err
}

// Patch applies the patch and returns the patched userSSHKey.
func (c *FakeUserSSHKeys) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.UserSSHKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(usersshkeysResource, c.ns, name, pt, data, subresources...), &v1.UserSSHKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.UserSSHKey), err
}
