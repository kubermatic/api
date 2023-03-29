// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	eeappskubermaticv1 "k8c.io/api/v3/pkg/apis/ee.apps.kubermatic/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeApplicationDefinitions implements ApplicationDefinitionInterface
type FakeApplicationDefinitions struct {
	Fake *FakeKubermaticEnterpriseAppsV1
	ns   string
}

var applicationdefinitionsResource = schema.GroupVersionResource{Group: "ee.apps.kubermatic.k8c.io", Version: "v1", Resource: "applicationdefinitions"}

var applicationdefinitionsKind = schema.GroupVersionKind{Group: "ee.apps.kubermatic.k8c.io", Version: "v1", Kind: "ApplicationDefinition"}

// Get takes name of the applicationDefinition, and returns the corresponding applicationDefinition object, and an error if there is any.
func (c *FakeApplicationDefinitions) Get(ctx context.Context, name string, options v1.GetOptions) (result *eeappskubermaticv1.ApplicationDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(applicationdefinitionsResource, c.ns, name), &eeappskubermaticv1.ApplicationDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*eeappskubermaticv1.ApplicationDefinition), err
}

// List takes label and field selectors, and returns the list of ApplicationDefinitions that match those selectors.
func (c *FakeApplicationDefinitions) List(ctx context.Context, opts v1.ListOptions) (result *eeappskubermaticv1.ApplicationDefinitionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(applicationdefinitionsResource, applicationdefinitionsKind, c.ns, opts), &eeappskubermaticv1.ApplicationDefinitionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &eeappskubermaticv1.ApplicationDefinitionList{ListMeta: obj.(*eeappskubermaticv1.ApplicationDefinitionList).ListMeta}
	for _, item := range obj.(*eeappskubermaticv1.ApplicationDefinitionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested applicationDefinitions.
func (c *FakeApplicationDefinitions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(applicationdefinitionsResource, c.ns, opts))

}

// Create takes the representation of a applicationDefinition and creates it.  Returns the server's representation of the applicationDefinition, and an error, if there is any.
func (c *FakeApplicationDefinitions) Create(ctx context.Context, applicationDefinition *eeappskubermaticv1.ApplicationDefinition, opts v1.CreateOptions) (result *eeappskubermaticv1.ApplicationDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(applicationdefinitionsResource, c.ns, applicationDefinition), &eeappskubermaticv1.ApplicationDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*eeappskubermaticv1.ApplicationDefinition), err
}

// Update takes the representation of a applicationDefinition and updates it. Returns the server's representation of the applicationDefinition, and an error, if there is any.
func (c *FakeApplicationDefinitions) Update(ctx context.Context, applicationDefinition *eeappskubermaticv1.ApplicationDefinition, opts v1.UpdateOptions) (result *eeappskubermaticv1.ApplicationDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(applicationdefinitionsResource, c.ns, applicationDefinition), &eeappskubermaticv1.ApplicationDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*eeappskubermaticv1.ApplicationDefinition), err
}

// Delete takes name of the applicationDefinition and deletes it. Returns an error if one occurs.
func (c *FakeApplicationDefinitions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(applicationdefinitionsResource, c.ns, name, opts), &eeappskubermaticv1.ApplicationDefinition{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeApplicationDefinitions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(applicationdefinitionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &eeappskubermaticv1.ApplicationDefinitionList{})
	return err
}

// Patch applies the patch and returns the patched applicationDefinition.
func (c *FakeApplicationDefinitions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *eeappskubermaticv1.ApplicationDefinition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(applicationdefinitionsResource, c.ns, name, pt, data, subresources...), &eeappskubermaticv1.ApplicationDefinition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*eeappskubermaticv1.ApplicationDefinition), err
}