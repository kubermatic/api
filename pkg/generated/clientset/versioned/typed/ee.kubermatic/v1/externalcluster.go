// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "k8c.io/api/v3/pkg/apis/ee.kubermatic/v1"
	scheme "k8c.io/api/v3/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ExternalClustersGetter has a method to return a ExternalClusterInterface.
// A group's client should implement this interface.
type ExternalClustersGetter interface {
	ExternalClusters(namespace string) ExternalClusterInterface
}

// ExternalClusterInterface has methods to work with ExternalCluster resources.
type ExternalClusterInterface interface {
	Create(ctx context.Context, externalCluster *v1.ExternalCluster, opts metav1.CreateOptions) (*v1.ExternalCluster, error)
	Update(ctx context.Context, externalCluster *v1.ExternalCluster, opts metav1.UpdateOptions) (*v1.ExternalCluster, error)
	UpdateStatus(ctx context.Context, externalCluster *v1.ExternalCluster, opts metav1.UpdateOptions) (*v1.ExternalCluster, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ExternalCluster, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ExternalClusterList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ExternalCluster, err error)
	ExternalClusterExpansion
}

// externalClusters implements ExternalClusterInterface
type externalClusters struct {
	client rest.Interface
	ns     string
}

// newExternalClusters returns a ExternalClusters
func newExternalClusters(c *EeKubermaticV1Client, namespace string) *externalClusters {
	return &externalClusters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the externalCluster, and returns the corresponding externalCluster object, and an error if there is any.
func (c *externalClusters) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ExternalCluster, err error) {
	result = &v1.ExternalCluster{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("externalclusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ExternalClusters that match those selectors.
func (c *externalClusters) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ExternalClusterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ExternalClusterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("externalclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested externalClusters.
func (c *externalClusters) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("externalclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a externalCluster and creates it.  Returns the server's representation of the externalCluster, and an error, if there is any.
func (c *externalClusters) Create(ctx context.Context, externalCluster *v1.ExternalCluster, opts metav1.CreateOptions) (result *v1.ExternalCluster, err error) {
	result = &v1.ExternalCluster{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("externalclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(externalCluster).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a externalCluster and updates it. Returns the server's representation of the externalCluster, and an error, if there is any.
func (c *externalClusters) Update(ctx context.Context, externalCluster *v1.ExternalCluster, opts metav1.UpdateOptions) (result *v1.ExternalCluster, err error) {
	result = &v1.ExternalCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("externalclusters").
		Name(externalCluster.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(externalCluster).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *externalClusters) UpdateStatus(ctx context.Context, externalCluster *v1.ExternalCluster, opts metav1.UpdateOptions) (result *v1.ExternalCluster, err error) {
	result = &v1.ExternalCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("externalclusters").
		Name(externalCluster.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(externalCluster).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the externalCluster and deletes it. Returns an error if one occurs.
func (c *externalClusters) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("externalclusters").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *externalClusters) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("externalclusters").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched externalCluster.
func (c *externalClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ExternalCluster, err error) {
	result = &v1.ExternalCluster{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("externalclusters").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
