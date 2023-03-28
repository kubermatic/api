// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	appskubermaticv1 "k8c.io/api/v3/pkg/apis/apps.kubermatic/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeApplicationInstallations implements ApplicationInstallationInterface
type FakeApplicationInstallations struct {
	Fake *FakeAppsV1
	ns   string
}

var applicationinstallationsResource = schema.GroupVersionResource{Group: "apps.kubermatic.k8c.io", Version: "v1", Resource: "applicationinstallations"}

var applicationinstallationsKind = schema.GroupVersionKind{Group: "apps.kubermatic.k8c.io", Version: "v1", Kind: "ApplicationInstallation"}

// Get takes name of the applicationInstallation, and returns the corresponding applicationInstallation object, and an error if there is any.
func (c *FakeApplicationInstallations) Get(ctx context.Context, name string, options v1.GetOptions) (result *appskubermaticv1.ApplicationInstallation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(applicationinstallationsResource, c.ns, name), &appskubermaticv1.ApplicationInstallation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*appskubermaticv1.ApplicationInstallation), err
}

// List takes label and field selectors, and returns the list of ApplicationInstallations that match those selectors.
func (c *FakeApplicationInstallations) List(ctx context.Context, opts v1.ListOptions) (result *appskubermaticv1.ApplicationInstallationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(applicationinstallationsResource, applicationinstallationsKind, c.ns, opts), &appskubermaticv1.ApplicationInstallationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &appskubermaticv1.ApplicationInstallationList{ListMeta: obj.(*appskubermaticv1.ApplicationInstallationList).ListMeta}
	for _, item := range obj.(*appskubermaticv1.ApplicationInstallationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested applicationInstallations.
func (c *FakeApplicationInstallations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(applicationinstallationsResource, c.ns, opts))

}

// Create takes the representation of a applicationInstallation and creates it.  Returns the server's representation of the applicationInstallation, and an error, if there is any.
func (c *FakeApplicationInstallations) Create(ctx context.Context, applicationInstallation *appskubermaticv1.ApplicationInstallation, opts v1.CreateOptions) (result *appskubermaticv1.ApplicationInstallation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(applicationinstallationsResource, c.ns, applicationInstallation), &appskubermaticv1.ApplicationInstallation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*appskubermaticv1.ApplicationInstallation), err
}

// Update takes the representation of a applicationInstallation and updates it. Returns the server's representation of the applicationInstallation, and an error, if there is any.
func (c *FakeApplicationInstallations) Update(ctx context.Context, applicationInstallation *appskubermaticv1.ApplicationInstallation, opts v1.UpdateOptions) (result *appskubermaticv1.ApplicationInstallation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(applicationinstallationsResource, c.ns, applicationInstallation), &appskubermaticv1.ApplicationInstallation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*appskubermaticv1.ApplicationInstallation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeApplicationInstallations) UpdateStatus(ctx context.Context, applicationInstallation *appskubermaticv1.ApplicationInstallation, opts v1.UpdateOptions) (*appskubermaticv1.ApplicationInstallation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(applicationinstallationsResource, "status", c.ns, applicationInstallation), &appskubermaticv1.ApplicationInstallation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*appskubermaticv1.ApplicationInstallation), err
}

// Delete takes name of the applicationInstallation and deletes it. Returns an error if one occurs.
func (c *FakeApplicationInstallations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(applicationinstallationsResource, c.ns, name, opts), &appskubermaticv1.ApplicationInstallation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApplicationInstallations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(applicationinstallationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &appskubermaticv1.ApplicationInstallationList{})
	return err
}

// Patch applies the patch and returns the patched applicationInstallation.
func (c *FakeApplicationInstallations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *appskubermaticv1.ApplicationInstallation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(applicationinstallationsResource, c.ns, name, pt, data, subresources...), &appskubermaticv1.ApplicationInstallation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*appskubermaticv1.ApplicationInstallation), err
}
