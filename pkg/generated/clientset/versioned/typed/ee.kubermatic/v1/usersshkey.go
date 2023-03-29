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

// UserSSHKeysGetter has a method to return a UserSSHKeyInterface.
// A group's client should implement this interface.
type UserSSHKeysGetter interface {
	UserSSHKeys(namespace string) UserSSHKeyInterface
}

// UserSSHKeyInterface has methods to work with UserSSHKey resources.
type UserSSHKeyInterface interface {
	Create(ctx context.Context, userSSHKey *v1.UserSSHKey, opts metav1.CreateOptions) (*v1.UserSSHKey, error)
	Update(ctx context.Context, userSSHKey *v1.UserSSHKey, opts metav1.UpdateOptions) (*v1.UserSSHKey, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.UserSSHKey, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.UserSSHKeyList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.UserSSHKey, err error)
	UserSSHKeyExpansion
}

// userSSHKeys implements UserSSHKeyInterface
type userSSHKeys struct {
	client rest.Interface
	ns     string
}

// newUserSSHKeys returns a UserSSHKeys
func newUserSSHKeys(c *KubermaticEnterpriseV1Client, namespace string) *userSSHKeys {
	return &userSSHKeys{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the userSSHKey, and returns the corresponding userSSHKey object, and an error if there is any.
func (c *userSSHKeys) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.UserSSHKey, err error) {
	result = &v1.UserSSHKey{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("usersshkeys").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of UserSSHKeys that match those selectors.
func (c *userSSHKeys) List(ctx context.Context, opts metav1.ListOptions) (result *v1.UserSSHKeyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.UserSSHKeyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("usersshkeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested userSSHKeys.
func (c *userSSHKeys) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("usersshkeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a userSSHKey and creates it.  Returns the server's representation of the userSSHKey, and an error, if there is any.
func (c *userSSHKeys) Create(ctx context.Context, userSSHKey *v1.UserSSHKey, opts metav1.CreateOptions) (result *v1.UserSSHKey, err error) {
	result = &v1.UserSSHKey{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("usersshkeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(userSSHKey).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a userSSHKey and updates it. Returns the server's representation of the userSSHKey, and an error, if there is any.
func (c *userSSHKeys) Update(ctx context.Context, userSSHKey *v1.UserSSHKey, opts metav1.UpdateOptions) (result *v1.UserSSHKey, err error) {
	result = &v1.UserSSHKey{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("usersshkeys").
		Name(userSSHKey.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(userSSHKey).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the userSSHKey and deletes it. Returns an error if one occurs.
func (c *userSSHKeys) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("usersshkeys").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *userSSHKeys) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("usersshkeys").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched userSSHKey.
func (c *userSSHKeys) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.UserSSHKey, err error) {
	result = &v1.UserSSHKey{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("usersshkeys").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}