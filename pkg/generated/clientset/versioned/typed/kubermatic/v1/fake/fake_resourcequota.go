// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	kubermaticv1 "k8c.io/api/v2/pkg/apis/kubermatic/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeResourceQuotas implements ResourceQuotaInterface
type FakeResourceQuotas struct {
	Fake *FakeKubermaticV1
	ns   string
}

var resourcequotasResource = schema.GroupVersionResource{Group: "kubermatic.k8c.io", Version: "v1", Resource: "resourcequotas"}

var resourcequotasKind = schema.GroupVersionKind{Group: "kubermatic.k8c.io", Version: "v1", Kind: "ResourceQuota"}

// Get takes name of the resourceQuota, and returns the corresponding resourceQuota object, and an error if there is any.
func (c *FakeResourceQuotas) Get(ctx context.Context, name string, options v1.GetOptions) (result *kubermaticv1.ResourceQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(resourcequotasResource, c.ns, name), &kubermaticv1.ResourceQuota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.ResourceQuota), err
}

// List takes label and field selectors, and returns the list of ResourceQuotas that match those selectors.
func (c *FakeResourceQuotas) List(ctx context.Context, opts v1.ListOptions) (result *kubermaticv1.ResourceQuotaList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(resourcequotasResource, resourcequotasKind, c.ns, opts), &kubermaticv1.ResourceQuotaList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &kubermaticv1.ResourceQuotaList{ListMeta: obj.(*kubermaticv1.ResourceQuotaList).ListMeta}
	for _, item := range obj.(*kubermaticv1.ResourceQuotaList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested resourceQuotas.
func (c *FakeResourceQuotas) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(resourcequotasResource, c.ns, opts))

}

// Create takes the representation of a resourceQuota and creates it.  Returns the server's representation of the resourceQuota, and an error, if there is any.
func (c *FakeResourceQuotas) Create(ctx context.Context, resourceQuota *kubermaticv1.ResourceQuota, opts v1.CreateOptions) (result *kubermaticv1.ResourceQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(resourcequotasResource, c.ns, resourceQuota), &kubermaticv1.ResourceQuota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.ResourceQuota), err
}

// Update takes the representation of a resourceQuota and updates it. Returns the server's representation of the resourceQuota, and an error, if there is any.
func (c *FakeResourceQuotas) Update(ctx context.Context, resourceQuota *kubermaticv1.ResourceQuota, opts v1.UpdateOptions) (result *kubermaticv1.ResourceQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(resourcequotasResource, c.ns, resourceQuota), &kubermaticv1.ResourceQuota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.ResourceQuota), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeResourceQuotas) UpdateStatus(ctx context.Context, resourceQuota *kubermaticv1.ResourceQuota, opts v1.UpdateOptions) (*kubermaticv1.ResourceQuota, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(resourcequotasResource, "status", c.ns, resourceQuota), &kubermaticv1.ResourceQuota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.ResourceQuota), err
}

// Delete takes name of the resourceQuota and deletes it. Returns an error if one occurs.
func (c *FakeResourceQuotas) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(resourcequotasResource, c.ns, name, opts), &kubermaticv1.ResourceQuota{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeResourceQuotas) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(resourcequotasResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &kubermaticv1.ResourceQuotaList{})
	return err
}

// Patch applies the patch and returns the patched resourceQuota.
func (c *FakeResourceQuotas) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *kubermaticv1.ResourceQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(resourcequotasResource, c.ns, name, pt, data, subresources...), &kubermaticv1.ResourceQuota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.ResourceQuota), err
}
