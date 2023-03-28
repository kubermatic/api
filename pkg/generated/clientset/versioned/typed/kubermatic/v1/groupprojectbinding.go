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

// GroupProjectBindingsGetter has a method to return a GroupProjectBindingInterface.
// A group's client should implement this interface.
type GroupProjectBindingsGetter interface {
	GroupProjectBindings(namespace string) GroupProjectBindingInterface
}

// GroupProjectBindingInterface has methods to work with GroupProjectBinding resources.
type GroupProjectBindingInterface interface {
	Create(ctx context.Context, groupProjectBinding *v1.GroupProjectBinding, opts metav1.CreateOptions) (*v1.GroupProjectBinding, error)
	Update(ctx context.Context, groupProjectBinding *v1.GroupProjectBinding, opts metav1.UpdateOptions) (*v1.GroupProjectBinding, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.GroupProjectBinding, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.GroupProjectBindingList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.GroupProjectBinding, err error)
	GroupProjectBindingExpansion
}

// groupProjectBindings implements GroupProjectBindingInterface
type groupProjectBindings struct {
	client rest.Interface
	ns     string
}

// newGroupProjectBindings returns a GroupProjectBindings
func newGroupProjectBindings(c *KubermaticV1Client, namespace string) *groupProjectBindings {
	return &groupProjectBindings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the groupProjectBinding, and returns the corresponding groupProjectBinding object, and an error if there is any.
func (c *groupProjectBindings) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.GroupProjectBinding, err error) {
	result = &v1.GroupProjectBinding{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("groupprojectbindings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GroupProjectBindings that match those selectors.
func (c *groupProjectBindings) List(ctx context.Context, opts metav1.ListOptions) (result *v1.GroupProjectBindingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.GroupProjectBindingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("groupprojectbindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested groupProjectBindings.
func (c *groupProjectBindings) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("groupprojectbindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a groupProjectBinding and creates it.  Returns the server's representation of the groupProjectBinding, and an error, if there is any.
func (c *groupProjectBindings) Create(ctx context.Context, groupProjectBinding *v1.GroupProjectBinding, opts metav1.CreateOptions) (result *v1.GroupProjectBinding, err error) {
	result = &v1.GroupProjectBinding{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("groupprojectbindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(groupProjectBinding).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a groupProjectBinding and updates it. Returns the server's representation of the groupProjectBinding, and an error, if there is any.
func (c *groupProjectBindings) Update(ctx context.Context, groupProjectBinding *v1.GroupProjectBinding, opts metav1.UpdateOptions) (result *v1.GroupProjectBinding, err error) {
	result = &v1.GroupProjectBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("groupprojectbindings").
		Name(groupProjectBinding.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(groupProjectBinding).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the groupProjectBinding and deletes it. Returns an error if one occurs.
func (c *groupProjectBindings) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("groupprojectbindings").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *groupProjectBindings) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("groupprojectbindings").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched groupProjectBinding.
func (c *groupProjectBindings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.GroupProjectBinding, err error) {
	result = &v1.GroupProjectBinding{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("groupprojectbindings").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
