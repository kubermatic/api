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

// FakeExternalClusters implements ExternalClusterInterface
type FakeExternalClusters struct {
	Fake *FakeKubermaticV1
	ns   string
}

var externalclustersResource = v1.SchemeGroupVersion.WithResource("externalclusters")

var externalclustersKind = v1.SchemeGroupVersion.WithKind("ExternalCluster")

// Get takes name of the externalCluster, and returns the corresponding externalCluster object, and an error if there is any.
func (c *FakeExternalClusters) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ExternalCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(externalclustersResource, c.ns, name), &v1.ExternalCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ExternalCluster), err
}

// List takes label and field selectors, and returns the list of ExternalClusters that match those selectors.
func (c *FakeExternalClusters) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ExternalClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(externalclustersResource, externalclustersKind, c.ns, opts), &v1.ExternalClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ExternalClusterList{ListMeta: obj.(*v1.ExternalClusterList).ListMeta}
	for _, item := range obj.(*v1.ExternalClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested externalClusters.
func (c *FakeExternalClusters) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(externalclustersResource, c.ns, opts))

}

// Create takes the representation of a externalCluster and creates it.  Returns the server's representation of the externalCluster, and an error, if there is any.
func (c *FakeExternalClusters) Create(ctx context.Context, externalCluster *v1.ExternalCluster, opts metav1.CreateOptions) (result *v1.ExternalCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(externalclustersResource, c.ns, externalCluster), &v1.ExternalCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ExternalCluster), err
}

// Update takes the representation of a externalCluster and updates it. Returns the server's representation of the externalCluster, and an error, if there is any.
func (c *FakeExternalClusters) Update(ctx context.Context, externalCluster *v1.ExternalCluster, opts metav1.UpdateOptions) (result *v1.ExternalCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(externalclustersResource, c.ns, externalCluster), &v1.ExternalCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ExternalCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeExternalClusters) UpdateStatus(ctx context.Context, externalCluster *v1.ExternalCluster, opts metav1.UpdateOptions) (*v1.ExternalCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(externalclustersResource, "status", c.ns, externalCluster), &v1.ExternalCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ExternalCluster), err
}

// Delete takes name of the externalCluster and deletes it. Returns an error if one occurs.
func (c *FakeExternalClusters) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(externalclustersResource, c.ns, name, opts), &v1.ExternalCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeExternalClusters) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(externalclustersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ExternalClusterList{})
	return err
}

// Patch applies the patch and returns the patched externalCluster.
func (c *FakeExternalClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ExternalCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(externalclustersResource, c.ns, name, pt, data, subresources...), &v1.ExternalCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.ExternalCluster), err
}
