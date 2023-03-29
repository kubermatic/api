// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "k8c.io/api/v3/pkg/apis/ee.apps.kubermatic/v1"
	scheme "k8c.io/api/v3/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ApplicationInstallationsGetter has a method to return a ApplicationInstallationInterface.
// A group's client should implement this interface.
type ApplicationInstallationsGetter interface {
	ApplicationInstallations(namespace string) ApplicationInstallationInterface
}

// ApplicationInstallationInterface has methods to work with ApplicationInstallation resources.
type ApplicationInstallationInterface interface {
	Create(ctx context.Context, applicationInstallation *v1.ApplicationInstallation, opts metav1.CreateOptions) (*v1.ApplicationInstallation, error)
	Update(ctx context.Context, applicationInstallation *v1.ApplicationInstallation, opts metav1.UpdateOptions) (*v1.ApplicationInstallation, error)
	UpdateStatus(ctx context.Context, applicationInstallation *v1.ApplicationInstallation, opts metav1.UpdateOptions) (*v1.ApplicationInstallation, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ApplicationInstallation, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ApplicationInstallationList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ApplicationInstallation, err error)
	ApplicationInstallationExpansion
}

// applicationInstallations implements ApplicationInstallationInterface
type applicationInstallations struct {
	client rest.Interface
	ns     string
}

// newApplicationInstallations returns a ApplicationInstallations
func newApplicationInstallations(c *KubermaticEnterpriseAppsV1Client, namespace string) *applicationInstallations {
	return &applicationInstallations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the applicationInstallation, and returns the corresponding applicationInstallation object, and an error if there is any.
func (c *applicationInstallations) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ApplicationInstallation, err error) {
	result = &v1.ApplicationInstallation{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("applicationinstallations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ApplicationInstallations that match those selectors.
func (c *applicationInstallations) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ApplicationInstallationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.ApplicationInstallationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("applicationinstallations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested applicationInstallations.
func (c *applicationInstallations) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("applicationinstallations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a applicationInstallation and creates it.  Returns the server's representation of the applicationInstallation, and an error, if there is any.
func (c *applicationInstallations) Create(ctx context.Context, applicationInstallation *v1.ApplicationInstallation, opts metav1.CreateOptions) (result *v1.ApplicationInstallation, err error) {
	result = &v1.ApplicationInstallation{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("applicationinstallations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(applicationInstallation).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a applicationInstallation and updates it. Returns the server's representation of the applicationInstallation, and an error, if there is any.
func (c *applicationInstallations) Update(ctx context.Context, applicationInstallation *v1.ApplicationInstallation, opts metav1.UpdateOptions) (result *v1.ApplicationInstallation, err error) {
	result = &v1.ApplicationInstallation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("applicationinstallations").
		Name(applicationInstallation.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(applicationInstallation).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *applicationInstallations) UpdateStatus(ctx context.Context, applicationInstallation *v1.ApplicationInstallation, opts metav1.UpdateOptions) (result *v1.ApplicationInstallation, err error) {
	result = &v1.ApplicationInstallation{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("applicationinstallations").
		Name(applicationInstallation.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(applicationInstallation).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the applicationInstallation and deletes it. Returns an error if one occurs.
func (c *applicationInstallations) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("applicationinstallations").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *applicationInstallations) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("applicationinstallations").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched applicationInstallation.
func (c *applicationInstallations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ApplicationInstallation, err error) {
	result = &v1.ApplicationInstallation{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("applicationinstallations").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}