// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	kubermaticv1 "k8c.io/api/v3/pkg/apis/kubermatic/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKubermaticConfigurations implements KubermaticConfigurationInterface
type FakeKubermaticConfigurations struct {
	Fake *FakeKubermaticV1
	ns   string
}

var kubermaticconfigurationsResource = schema.GroupVersionResource{Group: "kubermatic.k8c.io", Version: "v1", Resource: "kubermaticconfigurations"}

var kubermaticconfigurationsKind = schema.GroupVersionKind{Group: "kubermatic.k8c.io", Version: "v1", Kind: "KubermaticConfiguration"}

// Get takes name of the kubermaticConfiguration, and returns the corresponding kubermaticConfiguration object, and an error if there is any.
func (c *FakeKubermaticConfigurations) Get(ctx context.Context, name string, options v1.GetOptions) (result *kubermaticv1.KubermaticConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kubermaticconfigurationsResource, c.ns, name), &kubermaticv1.KubermaticConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.KubermaticConfiguration), err
}

// List takes label and field selectors, and returns the list of KubermaticConfigurations that match those selectors.
func (c *FakeKubermaticConfigurations) List(ctx context.Context, opts v1.ListOptions) (result *kubermaticv1.KubermaticConfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kubermaticconfigurationsResource, kubermaticconfigurationsKind, c.ns, opts), &kubermaticv1.KubermaticConfigurationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &kubermaticv1.KubermaticConfigurationList{ListMeta: obj.(*kubermaticv1.KubermaticConfigurationList).ListMeta}
	for _, item := range obj.(*kubermaticv1.KubermaticConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kubermaticConfigurations.
func (c *FakeKubermaticConfigurations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kubermaticconfigurationsResource, c.ns, opts))

}

// Create takes the representation of a kubermaticConfiguration and creates it.  Returns the server's representation of the kubermaticConfiguration, and an error, if there is any.
func (c *FakeKubermaticConfigurations) Create(ctx context.Context, kubermaticConfiguration *kubermaticv1.KubermaticConfiguration, opts v1.CreateOptions) (result *kubermaticv1.KubermaticConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kubermaticconfigurationsResource, c.ns, kubermaticConfiguration), &kubermaticv1.KubermaticConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.KubermaticConfiguration), err
}

// Update takes the representation of a kubermaticConfiguration and updates it. Returns the server's representation of the kubermaticConfiguration, and an error, if there is any.
func (c *FakeKubermaticConfigurations) Update(ctx context.Context, kubermaticConfiguration *kubermaticv1.KubermaticConfiguration, opts v1.UpdateOptions) (result *kubermaticv1.KubermaticConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kubermaticconfigurationsResource, c.ns, kubermaticConfiguration), &kubermaticv1.KubermaticConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.KubermaticConfiguration), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKubermaticConfigurations) UpdateStatus(ctx context.Context, kubermaticConfiguration *kubermaticv1.KubermaticConfiguration, opts v1.UpdateOptions) (*kubermaticv1.KubermaticConfiguration, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(kubermaticconfigurationsResource, "status", c.ns, kubermaticConfiguration), &kubermaticv1.KubermaticConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.KubermaticConfiguration), err
}

// Delete takes name of the kubermaticConfiguration and deletes it. Returns an error if one occurs.
func (c *FakeKubermaticConfigurations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(kubermaticconfigurationsResource, c.ns, name, opts), &kubermaticv1.KubermaticConfiguration{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKubermaticConfigurations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kubermaticconfigurationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &kubermaticv1.KubermaticConfigurationList{})
	return err
}

// Patch applies the patch and returns the patched kubermaticConfiguration.
func (c *FakeKubermaticConfigurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *kubermaticv1.KubermaticConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kubermaticconfigurationsResource, c.ns, name, pt, data, subresources...), &kubermaticv1.KubermaticConfiguration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.KubermaticConfiguration), err
}