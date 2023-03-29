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

// AdmissionPluginsGetter has a method to return a AdmissionPluginInterface.
// A group's client should implement this interface.
type AdmissionPluginsGetter interface {
	AdmissionPlugins(namespace string) AdmissionPluginInterface
}

// AdmissionPluginInterface has methods to work with AdmissionPlugin resources.
type AdmissionPluginInterface interface {
	Create(ctx context.Context, admissionPlugin *v1.AdmissionPlugin, opts metav1.CreateOptions) (*v1.AdmissionPlugin, error)
	Update(ctx context.Context, admissionPlugin *v1.AdmissionPlugin, opts metav1.UpdateOptions) (*v1.AdmissionPlugin, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.AdmissionPlugin, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.AdmissionPluginList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.AdmissionPlugin, err error)
	AdmissionPluginExpansion
}

// admissionPlugins implements AdmissionPluginInterface
type admissionPlugins struct {
	client rest.Interface
	ns     string
}

// newAdmissionPlugins returns a AdmissionPlugins
func newAdmissionPlugins(c *KubermaticV1Client, namespace string) *admissionPlugins {
	return &admissionPlugins{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the admissionPlugin, and returns the corresponding admissionPlugin object, and an error if there is any.
func (c *admissionPlugins) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.AdmissionPlugin, err error) {
	result = &v1.AdmissionPlugin{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("admissionplugins").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AdmissionPlugins that match those selectors.
func (c *admissionPlugins) List(ctx context.Context, opts metav1.ListOptions) (result *v1.AdmissionPluginList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.AdmissionPluginList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("admissionplugins").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested admissionPlugins.
func (c *admissionPlugins) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("admissionplugins").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a admissionPlugin and creates it.  Returns the server's representation of the admissionPlugin, and an error, if there is any.
func (c *admissionPlugins) Create(ctx context.Context, admissionPlugin *v1.AdmissionPlugin, opts metav1.CreateOptions) (result *v1.AdmissionPlugin, err error) {
	result = &v1.AdmissionPlugin{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("admissionplugins").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(admissionPlugin).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a admissionPlugin and updates it. Returns the server's representation of the admissionPlugin, and an error, if there is any.
func (c *admissionPlugins) Update(ctx context.Context, admissionPlugin *v1.AdmissionPlugin, opts metav1.UpdateOptions) (result *v1.AdmissionPlugin, err error) {
	result = &v1.AdmissionPlugin{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("admissionplugins").
		Name(admissionPlugin.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(admissionPlugin).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the admissionPlugin and deletes it. Returns an error if one occurs.
func (c *admissionPlugins) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("admissionplugins").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *admissionPlugins) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("admissionplugins").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched admissionPlugin.
func (c *admissionPlugins) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.AdmissionPlugin, err error) {
	result = &v1.AdmissionPlugin{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("admissionplugins").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
