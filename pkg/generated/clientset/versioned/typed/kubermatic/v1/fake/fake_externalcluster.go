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

// FakeExternalClusters implements ExternalClusterInterface
type FakeExternalClusters struct {
	Fake *FakeKubermaticV1
	ns   string
}

var externalclustersResource = schema.GroupVersionResource{Group: "kubermatic.k8c.io", Version: "v1", Resource: "externalclusters"}

var externalclustersKind = schema.GroupVersionKind{Group: "kubermatic.k8c.io", Version: "v1", Kind: "ExternalCluster"}

// Get takes name of the externalCluster, and returns the corresponding externalCluster object, and an error if there is any.
func (c *FakeExternalClusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *kubermaticv1.ExternalCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(externalclustersResource, c.ns, name), &kubermaticv1.ExternalCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.ExternalCluster), err
}

// List takes label and field selectors, and returns the list of ExternalClusters that match those selectors.
func (c *FakeExternalClusters) List(ctx context.Context, opts v1.ListOptions) (result *kubermaticv1.ExternalClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(externalclustersResource, externalclustersKind, c.ns, opts), &kubermaticv1.ExternalClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &kubermaticv1.ExternalClusterList{ListMeta: obj.(*kubermaticv1.ExternalClusterList).ListMeta}
	for _, item := range obj.(*kubermaticv1.ExternalClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested externalClusters.
func (c *FakeExternalClusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(externalclustersResource, c.ns, opts))

}

// Create takes the representation of a externalCluster and creates it.  Returns the server's representation of the externalCluster, and an error, if there is any.
func (c *FakeExternalClusters) Create(ctx context.Context, externalCluster *kubermaticv1.ExternalCluster, opts v1.CreateOptions) (result *kubermaticv1.ExternalCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(externalclustersResource, c.ns, externalCluster), &kubermaticv1.ExternalCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.ExternalCluster), err
}

// Update takes the representation of a externalCluster and updates it. Returns the server's representation of the externalCluster, and an error, if there is any.
func (c *FakeExternalClusters) Update(ctx context.Context, externalCluster *kubermaticv1.ExternalCluster, opts v1.UpdateOptions) (result *kubermaticv1.ExternalCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(externalclustersResource, c.ns, externalCluster), &kubermaticv1.ExternalCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.ExternalCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeExternalClusters) UpdateStatus(ctx context.Context, externalCluster *kubermaticv1.ExternalCluster, opts v1.UpdateOptions) (*kubermaticv1.ExternalCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(externalclustersResource, "status", c.ns, externalCluster), &kubermaticv1.ExternalCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.ExternalCluster), err
}

// Delete takes name of the externalCluster and deletes it. Returns an error if one occurs.
func (c *FakeExternalClusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(externalclustersResource, c.ns, name, opts), &kubermaticv1.ExternalCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeExternalClusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(externalclustersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &kubermaticv1.ExternalClusterList{})
	return err
}

// Patch applies the patch and returns the patched externalCluster.
func (c *FakeExternalClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *kubermaticv1.ExternalCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(externalclustersResource, c.ns, name, pt, data, subresources...), &kubermaticv1.ExternalCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.ExternalCluster), err
}
