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

// MLARuleGroupsGetter has a method to return a MLARuleGroupInterface.
// A group's client should implement this interface.
type MLARuleGroupsGetter interface {
	MLARuleGroups(namespace string) MLARuleGroupInterface
}

// MLARuleGroupInterface has methods to work with MLARuleGroup resources.
type MLARuleGroupInterface interface {
	Create(ctx context.Context, mLARuleGroup *v1.MLARuleGroup, opts metav1.CreateOptions) (*v1.MLARuleGroup, error)
	Update(ctx context.Context, mLARuleGroup *v1.MLARuleGroup, opts metav1.UpdateOptions) (*v1.MLARuleGroup, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.MLARuleGroup, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.MLARuleGroupList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.MLARuleGroup, err error)
	MLARuleGroupExpansion
}

// mLARuleGroups implements MLARuleGroupInterface
type mLARuleGroups struct {
	client rest.Interface
	ns     string
}

// newMLARuleGroups returns a MLARuleGroups
func newMLARuleGroups(c *KubermaticV1Client, namespace string) *mLARuleGroups {
	return &mLARuleGroups{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the mLARuleGroup, and returns the corresponding mLARuleGroup object, and an error if there is any.
func (c *mLARuleGroups) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.MLARuleGroup, err error) {
	result = &v1.MLARuleGroup{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mlarulegroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MLARuleGroups that match those selectors.
func (c *mLARuleGroups) List(ctx context.Context, opts metav1.ListOptions) (result *v1.MLARuleGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.MLARuleGroupList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mlarulegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested mLARuleGroups.
func (c *mLARuleGroups) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("mlarulegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a mLARuleGroup and creates it.  Returns the server's representation of the mLARuleGroup, and an error, if there is any.
func (c *mLARuleGroups) Create(ctx context.Context, mLARuleGroup *v1.MLARuleGroup, opts metav1.CreateOptions) (result *v1.MLARuleGroup, err error) {
	result = &v1.MLARuleGroup{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("mlarulegroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(mLARuleGroup).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a mLARuleGroup and updates it. Returns the server's representation of the mLARuleGroup, and an error, if there is any.
func (c *mLARuleGroups) Update(ctx context.Context, mLARuleGroup *v1.MLARuleGroup, opts metav1.UpdateOptions) (result *v1.MLARuleGroup, err error) {
	result = &v1.MLARuleGroup{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mlarulegroups").
		Name(mLARuleGroup.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(mLARuleGroup).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the mLARuleGroup and deletes it. Returns an error if one occurs.
func (c *mLARuleGroups) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mlarulegroups").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *mLARuleGroups) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mlarulegroups").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched mLARuleGroup.
func (c *mLARuleGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.MLARuleGroup, err error) {
	result = &v1.MLARuleGroup{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("mlarulegroups").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
