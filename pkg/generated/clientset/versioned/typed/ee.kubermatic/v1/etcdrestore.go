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

// EtcdRestoresGetter has a method to return a EtcdRestoreInterface.
// A group's client should implement this interface.
type EtcdRestoresGetter interface {
	EtcdRestores(namespace string) EtcdRestoreInterface
}

// EtcdRestoreInterface has methods to work with EtcdRestore resources.
type EtcdRestoreInterface interface {
	Create(ctx context.Context, etcdRestore *v1.EtcdRestore, opts metav1.CreateOptions) (*v1.EtcdRestore, error)
	Update(ctx context.Context, etcdRestore *v1.EtcdRestore, opts metav1.UpdateOptions) (*v1.EtcdRestore, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.EtcdRestore, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.EtcdRestoreList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.EtcdRestore, err error)
	EtcdRestoreExpansion
}

// etcdRestores implements EtcdRestoreInterface
type etcdRestores struct {
	client rest.Interface
	ns     string
}

// newEtcdRestores returns a EtcdRestores
func newEtcdRestores(c *KubermaticEnterpriseV1Client, namespace string) *etcdRestores {
	return &etcdRestores{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the etcdRestore, and returns the corresponding etcdRestore object, and an error if there is any.
func (c *etcdRestores) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.EtcdRestore, err error) {
	result = &v1.EtcdRestore{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("etcdrestores").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of EtcdRestores that match those selectors.
func (c *etcdRestores) List(ctx context.Context, opts metav1.ListOptions) (result *v1.EtcdRestoreList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.EtcdRestoreList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("etcdrestores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested etcdRestores.
func (c *etcdRestores) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("etcdrestores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a etcdRestore and creates it.  Returns the server's representation of the etcdRestore, and an error, if there is any.
func (c *etcdRestores) Create(ctx context.Context, etcdRestore *v1.EtcdRestore, opts metav1.CreateOptions) (result *v1.EtcdRestore, err error) {
	result = &v1.EtcdRestore{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("etcdrestores").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(etcdRestore).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a etcdRestore and updates it. Returns the server's representation of the etcdRestore, and an error, if there is any.
func (c *etcdRestores) Update(ctx context.Context, etcdRestore *v1.EtcdRestore, opts metav1.UpdateOptions) (result *v1.EtcdRestore, err error) {
	result = &v1.EtcdRestore{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("etcdrestores").
		Name(etcdRestore.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(etcdRestore).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the etcdRestore and deletes it. Returns an error if one occurs.
func (c *etcdRestores) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("etcdrestores").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *etcdRestores) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("etcdrestores").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched etcdRestore.
func (c *etcdRestores) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.EtcdRestore, err error) {
	result = &v1.EtcdRestore{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("etcdrestores").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
