// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8c.io/api/v3/pkg/apis/ee.kubermatic/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeResourceQuotas implements ResourceQuotaInterface
type FakeResourceQuotas struct {
	Fake *FakeKubermaticEnterpriseV1
	ns   string
}

var resourcequotasResource = v1.SchemeGroupVersion.WithResource("resourcequotas")

var resourcequotasKind = v1.SchemeGroupVersion.WithKind("ResourceQuota")

// Get takes name of the resourceQuota, and returns the corresponding resourceQuota object, and an error if there is any.
func (c *FakeResourceQuotas) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ResourceQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(resourcequotasResource, c.ns, name), &v1.ResourceQuota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ResourceQuota), err
}

// List takes label and field selectors, and returns the list of ResourceQuotas that match those selectors.
func (c *FakeResourceQuotas) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ResourceQuotaList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(resourcequotasResource, resourcequotasKind, c.ns, opts), &v1.ResourceQuotaList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ResourceQuotaList{ListMeta: obj.(*v1.ResourceQuotaList).ListMeta}
	for _, item := range obj.(*v1.ResourceQuotaList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested resourceQuotas.
func (c *FakeResourceQuotas) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(resourcequotasResource, c.ns, opts))

}

// Create takes the representation of a resourceQuota and creates it.  Returns the server's representation of the resourceQuota, and an error, if there is any.
func (c *FakeResourceQuotas) Create(ctx context.Context, resourceQuota *v1.ResourceQuota, opts metav1.CreateOptions) (result *v1.ResourceQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(resourcequotasResource, c.ns, resourceQuota), &v1.ResourceQuota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ResourceQuota), err
}

// Update takes the representation of a resourceQuota and updates it. Returns the server's representation of the resourceQuota, and an error, if there is any.
func (c *FakeResourceQuotas) Update(ctx context.Context, resourceQuota *v1.ResourceQuota, opts metav1.UpdateOptions) (result *v1.ResourceQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(resourcequotasResource, c.ns, resourceQuota), &v1.ResourceQuota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ResourceQuota), err
}

// Delete takes name of the resourceQuota and deletes it. Returns an error if one occurs.
func (c *FakeResourceQuotas) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(resourcequotasResource, c.ns, name, opts), &v1.ResourceQuota{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeResourceQuotas) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(resourcequotasResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ResourceQuotaList{})
	return err
}

// Patch applies the patch and returns the patched resourceQuota.
func (c *FakeResourceQuotas) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ResourceQuota, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(resourcequotasResource, c.ns, name, pt, data, subresources...), &v1.ResourceQuota{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ResourceQuota), err
}
