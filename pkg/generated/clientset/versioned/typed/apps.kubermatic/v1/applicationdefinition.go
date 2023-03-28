// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "k8c.io/api/v3/pkg/apis/apps.kubermatic/v1"
	scheme "k8c.io/api/v3/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ApplicationDefinitionsGetter has a method to return a ApplicationDefinitionInterface.
// A group's client should implement this interface.
type ApplicationDefinitionsGetter interface {
	ApplicationDefinitions(namespace string) ApplicationDefinitionInterface
}

// ApplicationDefinitionInterface has methods to work with ApplicationDefinition resources.
type ApplicationDefinitionInterface interface {
	Create(ctx context.Context, applicationDefinition *v1.ApplicationDefinition, opts metav1.CreateOptions) (*v1.ApplicationDefinition, error)
	Update(ctx context.Context, applicationDefinition *v1.ApplicationDefinition, opts metav1.UpdateOptions) (*v1.ApplicationDefinition, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ApplicationDefinition, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ApplicationDefinitionList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ApplicationDefinition, err error)
	ApplicationDefinitionExpansion
}

// applicationDefinitions implements ApplicationDefinitionInterface
type applicationDefinitions struct {
	client rest.Interface
	ns     string
}

// newApplicationDefinitions returns a ApplicationDefinitions
func newApplicationDefinitions(c *AppsKubermaticV1Client, namespace string) *applicationDefinitions {
	return &applicationDefinitions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the applicationDefinition, and returns the corresponding applicationDefinition object, and an error if there is any.
func (c *applicationDefinitions) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ApplicationDefinition, err error) {
	result = &v1.ApplicationDefinition{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("applicationdefinitions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ApplicationDefinitions that match those selectors.
func (c *applicationDefinitions) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ApplicationDefinitionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ApplicationDefinitionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("applicationdefinitions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested applicationDefinitions.
func (c *applicationDefinitions) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("applicationdefinitions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a applicationDefinition and creates it.  Returns the server's representation of the applicationDefinition, and an error, if there is any.
func (c *applicationDefinitions) Create(ctx context.Context, applicationDefinition *v1.ApplicationDefinition, opts metav1.CreateOptions) (result *v1.ApplicationDefinition, err error) {
	result = &v1.ApplicationDefinition{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("applicationdefinitions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(applicationDefinition).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a applicationDefinition and updates it. Returns the server's representation of the applicationDefinition, and an error, if there is any.
func (c *applicationDefinitions) Update(ctx context.Context, applicationDefinition *v1.ApplicationDefinition, opts metav1.UpdateOptions) (result *v1.ApplicationDefinition, err error) {
	result = &v1.ApplicationDefinition{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("applicationdefinitions").
		Name(applicationDefinition.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(applicationDefinition).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the applicationDefinition and deletes it. Returns an error if one occurs.
func (c *applicationDefinitions) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("applicationdefinitions").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *applicationDefinitions) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("applicationdefinitions").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched applicationDefinition.
func (c *applicationDefinitions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ApplicationDefinition, err error) {
	result = &v1.ApplicationDefinition{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("applicationdefinitions").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
