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

// ConstraintTemplatesGetter has a method to return a ConstraintTemplateInterface.
// A group's client should implement this interface.
type ConstraintTemplatesGetter interface {
	ConstraintTemplates(namespace string) ConstraintTemplateInterface
}

// ConstraintTemplateInterface has methods to work with ConstraintTemplate resources.
type ConstraintTemplateInterface interface {
	Create(ctx context.Context, constraintTemplate *v1.ConstraintTemplate, opts metav1.CreateOptions) (*v1.ConstraintTemplate, error)
	Update(ctx context.Context, constraintTemplate *v1.ConstraintTemplate, opts metav1.UpdateOptions) (*v1.ConstraintTemplate, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ConstraintTemplate, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ConstraintTemplateList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ConstraintTemplate, err error)
	ConstraintTemplateExpansion
}

// constraintTemplates implements ConstraintTemplateInterface
type constraintTemplates struct {
	client rest.Interface
	ns     string
}

// newConstraintTemplates returns a ConstraintTemplates
func newConstraintTemplates(c *KubermaticV1Client, namespace string) *constraintTemplates {
	return &constraintTemplates{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the constraintTemplate, and returns the corresponding constraintTemplate object, and an error if there is any.
func (c *constraintTemplates) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ConstraintTemplate, err error) {
	result = &v1.ConstraintTemplate{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("constrainttemplates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ConstraintTemplates that match those selectors.
func (c *constraintTemplates) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ConstraintTemplateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ConstraintTemplateList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("constrainttemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested constraintTemplates.
func (c *constraintTemplates) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("constrainttemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a constraintTemplate and creates it.  Returns the server's representation of the constraintTemplate, and an error, if there is any.
func (c *constraintTemplates) Create(ctx context.Context, constraintTemplate *v1.ConstraintTemplate, opts metav1.CreateOptions) (result *v1.ConstraintTemplate, err error) {
	result = &v1.ConstraintTemplate{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("constrainttemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(constraintTemplate).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a constraintTemplate and updates it. Returns the server's representation of the constraintTemplate, and an error, if there is any.
func (c *constraintTemplates) Update(ctx context.Context, constraintTemplate *v1.ConstraintTemplate, opts metav1.UpdateOptions) (result *v1.ConstraintTemplate, err error) {
	result = &v1.ConstraintTemplate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("constrainttemplates").
		Name(constraintTemplate.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(constraintTemplate).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the constraintTemplate and deletes it. Returns an error if one occurs.
func (c *constraintTemplates) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("constrainttemplates").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *constraintTemplates) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("constrainttemplates").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched constraintTemplate.
func (c *constraintTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ConstraintTemplate, err error) {
	result = &v1.ConstraintTemplate{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("constrainttemplates").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}