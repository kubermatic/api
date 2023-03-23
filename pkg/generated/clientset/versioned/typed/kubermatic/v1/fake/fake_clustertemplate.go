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

// FakeClusterTemplates implements ClusterTemplateInterface
type FakeClusterTemplates struct {
	Fake *FakeKubermaticV1
	ns   string
}

var clustertemplatesResource = schema.GroupVersionResource{Group: "kubermatic.k8c.io", Version: "v1", Resource: "clustertemplates"}

var clustertemplatesKind = schema.GroupVersionKind{Group: "kubermatic.k8c.io", Version: "v1", Kind: "ClusterTemplate"}

// Get takes name of the clusterTemplate, and returns the corresponding clusterTemplate object, and an error if there is any.
func (c *FakeClusterTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *kubermaticv1.ClusterTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(clustertemplatesResource, c.ns, name), &kubermaticv1.ClusterTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.ClusterTemplate), err
}

// List takes label and field selectors, and returns the list of ClusterTemplates that match those selectors.
func (c *FakeClusterTemplates) List(ctx context.Context, opts v1.ListOptions) (result *kubermaticv1.ClusterTemplateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(clustertemplatesResource, clustertemplatesKind, c.ns, opts), &kubermaticv1.ClusterTemplateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &kubermaticv1.ClusterTemplateList{ListMeta: obj.(*kubermaticv1.ClusterTemplateList).ListMeta}
	for _, item := range obj.(*kubermaticv1.ClusterTemplateList).Items {
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
func (c *FakeClusterTemplates) Create(ctx context.Context, clusterTemplate *kubermaticv1.ClusterTemplate, opts v1.CreateOptions) (result *kubermaticv1.ClusterTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(clustertemplatesResource, c.ns, clusterTemplate), &kubermaticv1.ClusterTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.ClusterTemplate), err
}

// Update takes the representation of a clusterTemplate and updates it. Returns the server's representation of the clusterTemplate, and an error, if there is any.
func (c *FakeClusterTemplates) Update(ctx context.Context, clusterTemplate *kubermaticv1.ClusterTemplate, opts v1.UpdateOptions) (result *kubermaticv1.ClusterTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(clustertemplatesResource, c.ns, clusterTemplate), &kubermaticv1.ClusterTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.ClusterTemplate), err
}

// Delete takes name of the clusterTemplate and deletes it. Returns an error if one occurs.
func (c *FakeClusterTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(clustertemplatesResource, c.ns, name, opts), &kubermaticv1.ClusterTemplate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(clustertemplatesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &kubermaticv1.ClusterTemplateList{})
	return err
}

// Patch applies the patch and returns the patched clusterTemplate.
func (c *FakeClusterTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *kubermaticv1.ClusterTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(clustertemplatesResource, c.ns, name, pt, data, subresources...), &kubermaticv1.ClusterTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.ClusterTemplate), err
}
