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

// SSHKeysGetter has a method to return a SSHKeyInterface.
// A group's client should implement this interface.
type SSHKeysGetter interface {
	SSHKeys(namespace string) SSHKeyInterface
}

// SSHKeyInterface has methods to work with SSHKey resources.
type SSHKeyInterface interface {
	Create(ctx context.Context, sSHKey *v1.SSHKey, opts metav1.CreateOptions) (*v1.SSHKey, error)
	Update(ctx context.Context, sSHKey *v1.SSHKey, opts metav1.UpdateOptions) (*v1.SSHKey, error)
	UpdateStatus(ctx context.Context, sSHKey *v1.SSHKey, opts metav1.UpdateOptions) (*v1.SSHKey, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.SSHKey, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.SSHKeyList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.SSHKey, err error)
	SSHKeyExpansion
}

// sSHKeys implements SSHKeyInterface
type sSHKeys struct {
	client rest.Interface
	ns     string
}

// newSSHKeys returns a SSHKeys
func newSSHKeys(c *KubermaticV1Client, namespace string) *sSHKeys {
	return &sSHKeys{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the sSHKey, and returns the corresponding sSHKey object, and an error if there is any.
func (c *sSHKeys) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.SSHKey, err error) {
	result = &v1.SSHKey{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sshkeys").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SSHKeys that match those selectors.
func (c *sSHKeys) List(ctx context.Context, opts metav1.ListOptions) (result *v1.SSHKeyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.SSHKeyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("sshkeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sSHKeys.
func (c *sSHKeys) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("sshkeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a sSHKey and creates it.  Returns the server's representation of the sSHKey, and an error, if there is any.
func (c *sSHKeys) Create(ctx context.Context, sSHKey *v1.SSHKey, opts metav1.CreateOptions) (result *v1.SSHKey, err error) {
	result = &v1.SSHKey{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("sshkeys").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sSHKey).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a sSHKey and updates it. Returns the server's representation of the sSHKey, and an error, if there is any.
func (c *sSHKeys) Update(ctx context.Context, sSHKey *v1.SSHKey, opts metav1.UpdateOptions) (result *v1.SSHKey, err error) {
	result = &v1.SSHKey{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sshkeys").
		Name(sSHKey.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sSHKey).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *sSHKeys) UpdateStatus(ctx context.Context, sSHKey *v1.SSHKey, opts metav1.UpdateOptions) (result *v1.SSHKey, err error) {
	result = &v1.SSHKey{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("sshkeys").
		Name(sSHKey.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sSHKey).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the sSHKey and deletes it. Returns an error if one occurs.
func (c *sSHKeys) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sshkeys").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sSHKeys) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("sshkeys").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched sSHKey.
func (c *sSHKeys) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.SSHKey, err error) {
	result = &v1.SSHKey{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("sshkeys").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
