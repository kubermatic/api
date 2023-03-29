// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	eekubermaticv1 "k8c.io/api/v3/pkg/apis/ee.kubermatic/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAllowedRegistries implements AllowedRegistryInterface
type FakeAllowedRegistries struct {
	Fake *FakeKubermaticEnterpriseV1
	ns   string
}

var allowedregistriesResource = schema.GroupVersionResource{Group: "ee.kubermatic.k8c.io", Version: "v1", Resource: "allowedregistries"}

var allowedregistriesKind = schema.GroupVersionKind{Group: "ee.kubermatic.k8c.io", Version: "v1", Kind: "AllowedRegistry"}

// Get takes name of the allowedRegistry, and returns the corresponding allowedRegistry object, and an error if there is any.
func (c *FakeAllowedRegistries) Get(ctx context.Context, name string, options v1.GetOptions) (result *eekubermaticv1.AllowedRegistry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(allowedregistriesResource, c.ns, name), &eekubermaticv1.AllowedRegistry{})

	if obj == nil {
		return nil, err
	}
	return obj.(*eekubermaticv1.AllowedRegistry), err
}

// List takes label and field selectors, and returns the list of AllowedRegistries that match those selectors.
func (c *FakeAllowedRegistries) List(ctx context.Context, opts v1.ListOptions) (result *eekubermaticv1.AllowedRegistryList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(allowedregistriesResource, allowedregistriesKind, c.ns, opts), &eekubermaticv1.AllowedRegistryList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &eekubermaticv1.AllowedRegistryList{ListMeta: obj.(*eekubermaticv1.AllowedRegistryList).ListMeta}
	for _, item := range obj.(*eekubermaticv1.AllowedRegistryList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested allowedRegistries.
func (c *FakeAllowedRegistries) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(allowedregistriesResource, c.ns, opts))

}

// Create takes the representation of a allowedRegistry and creates it.  Returns the server's representation of the allowedRegistry, and an error, if there is any.
func (c *FakeAllowedRegistries) Create(ctx context.Context, allowedRegistry *eekubermaticv1.AllowedRegistry, opts v1.CreateOptions) (result *eekubermaticv1.AllowedRegistry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(allowedregistriesResource, c.ns, allowedRegistry), &eekubermaticv1.AllowedRegistry{})

	if obj == nil {
		return nil, err
	}
	return obj.(*eekubermaticv1.AllowedRegistry), err
}

// Update takes the representation of a allowedRegistry and updates it. Returns the server's representation of the allowedRegistry, and an error, if there is any.
func (c *FakeAllowedRegistries) Update(ctx context.Context, allowedRegistry *eekubermaticv1.AllowedRegistry, opts v1.UpdateOptions) (result *eekubermaticv1.AllowedRegistry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(allowedregistriesResource, c.ns, allowedRegistry), &eekubermaticv1.AllowedRegistry{})

	if obj == nil {
		return nil, err
	}
	return obj.(*eekubermaticv1.AllowedRegistry), err
}

// Delete takes name of the allowedRegistry and deletes it. Returns an error if one occurs.
func (c *FakeAllowedRegistries) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(allowedregistriesResource, c.ns, name, opts), &eekubermaticv1.AllowedRegistry{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAllowedRegistries) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(allowedregistriesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &eekubermaticv1.AllowedRegistryList{})
	return err
}

// Patch applies the patch and returns the patched allowedRegistry.
func (c *FakeAllowedRegistries) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *eekubermaticv1.AllowedRegistry, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(allowedregistriesResource, c.ns, name, pt, data, subresources...), &eekubermaticv1.AllowedRegistry{})

	if obj == nil {
		return nil, err
	}
	return obj.(*eekubermaticv1.AllowedRegistry), err
}