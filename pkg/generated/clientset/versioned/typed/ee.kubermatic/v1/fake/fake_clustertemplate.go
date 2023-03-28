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

// FakeClusterTemplates implements ClusterTemplateInterface
type FakeClusterTemplates struct {
	Fake *FakeEeKubermaticV1
	ns   string
}

var clustertemplatesResource = schema.GroupVersionResource{Group: "ee.kubermatic.k8c.io", Version: "v1", Resource: "clustertemplates"}

var clustertemplatesKind = schema.GroupVersionKind{Group: "ee.kubermatic.k8c.io", Version: "v1", Kind: "ClusterTemplate"}

// Get takes name of the clusterTemplate, and returns the corresponding clusterTemplate object, and an error if there is any.
func (c *FakeClusterTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *eekubermaticv1.ClusterTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clustertemplatesResource, c.ns, name), &eekubermaticv1.ClusterTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*eekubermaticv1.ClusterTemplate), err
}

// List takes label and field selectors, and returns the list of ClusterTemplates that match those selectors.
func (c *FakeClusterTemplates) List(ctx context.Context, opts v1.ListOptions) (result *eekubermaticv1.ClusterTemplateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clustertemplatesResource, clustertemplatesKind, c.ns, opts), &eekubermaticv1.ClusterTemplateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &eekubermaticv1.ClusterTemplateList{ListMeta: obj.(*eekubermaticv1.ClusterTemplateList).ListMeta}
	for _, item := range obj.(*eekubermaticv1.ClusterTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterTemplates.
func (c *FakeClusterTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(clustertemplatesResource, c.ns, opts))

}

// Create takes the representation of a clusterTemplate and creates it.  Returns the server's representation of the clusterTemplate, and an error, if there is any.
func (c *FakeClusterTemplates) Create(ctx context.Context, clusterTemplate *eekubermaticv1.ClusterTemplate, opts v1.CreateOptions) (result *eekubermaticv1.ClusterTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clustertemplatesResource, c.ns, clusterTemplate), &eekubermaticv1.ClusterTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*eekubermaticv1.ClusterTemplate), err
}

// Update takes the representation of a clusterTemplate and updates it. Returns the server's representation of the clusterTemplate, and an error, if there is any.
func (c *FakeClusterTemplates) Update(ctx context.Context, clusterTemplate *eekubermaticv1.ClusterTemplate, opts v1.UpdateOptions) (result *eekubermaticv1.ClusterTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clustertemplatesResource, c.ns, clusterTemplate), &eekubermaticv1.ClusterTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*eekubermaticv1.ClusterTemplate), err
}

// Delete takes name of the clusterTemplate and deletes it. Returns an error if one occurs.
func (c *FakeClusterTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(clustertemplatesResource, c.ns, name, opts), &eekubermaticv1.ClusterTemplate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clustertemplatesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &eekubermaticv1.ClusterTemplateList{})
	return err
}

// Patch applies the patch and returns the patched clusterTemplate.
func (c *FakeClusterTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *eekubermaticv1.ClusterTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clustertemplatesResource, c.ns, name, pt, data, subresources...), &eekubermaticv1.ClusterTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*eekubermaticv1.ClusterTemplate), err
}
