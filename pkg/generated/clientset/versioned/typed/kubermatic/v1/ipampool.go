// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "k8c.io/api/v3/pkg/apis/kubermatic/v1"
	scheme "k8c.io/api/v3/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// IPAMPoolsGetter has a method to return a IPAMPoolInterface.
// A group's client should implement this interface.
type IPAMPoolsGetter interface {
	IPAMPools(namespace string) IPAMPoolInterface
}

// IPAMPoolInterface has methods to work with IPAMPool resources.
type IPAMPoolInterface interface {
	Create(ctx context.Context, iPAMPool *v1.IPAMPool, opts metav1.CreateOptions) (*v1.IPAMPool, error)
	Update(ctx context.Context, iPAMPool *v1.IPAMPool, opts metav1.UpdateOptions) (*v1.IPAMPool, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.IPAMPool, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.IPAMPoolList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.IPAMPool, err error)
	IPAMPoolExpansion
}

// iPAMPools implements IPAMPoolInterface
type iPAMPools struct {
	client rest.Interface
	ns     string
}

// newIPAMPools returns a IPAMPools
func newIPAMPools(c *KubermaticV1Client, namespace string) *iPAMPools {
	return &iPAMPools{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the iPAMPool, and returns the corresponding iPAMPool object, and an error if there is any.
func (c *iPAMPools) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.IPAMPool, err error) {
	result = &v1.IPAMPool{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ipampools").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of IPAMPools that match those selectors.
func (c *iPAMPools) List(ctx context.Context, opts metav1.ListOptions) (result *v1.IPAMPoolList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.IPAMPoolList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("ipampools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested iPAMPools.
func (c *iPAMPools) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("ipampools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a iPAMPool and creates it.  Returns the server's representation of the iPAMPool, and an error, if there is any.
func (c *iPAMPools) Create(ctx context.Context, iPAMPool *v1.IPAMPool, opts metav1.CreateOptions) (result *v1.IPAMPool, err error) {
	result = &v1.IPAMPool{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("ipampools").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(iPAMPool).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a iPAMPool and updates it. Returns the server's representation of the iPAMPool, and an error, if there is any.
func (c *iPAMPools) Update(ctx context.Context, iPAMPool *v1.IPAMPool, opts metav1.UpdateOptions) (result *v1.IPAMPool, err error) {
	result = &v1.IPAMPool{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("ipampools").
		Name(iPAMPool.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(iPAMPool).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the iPAMPool and deletes it. Returns an error if one occurs.
func (c *iPAMPools) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ipampools").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *iPAMPools) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("ipampools").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched iPAMPool.
func (c *iPAMPools) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.IPAMPool, err error) {
	result = &v1.IPAMPool{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("ipampools").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
