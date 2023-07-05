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

// MLAClusterConfigurationsGetter has a method to return a MLAClusterConfigurationInterface.
// A group's client should implement this interface.
type MLAClusterConfigurationsGetter interface {
	MLAClusterConfigurations(namespace string) MLAClusterConfigurationInterface
}

// MLAClusterConfigurationInterface has methods to work with MLAClusterConfiguration resources.
type MLAClusterConfigurationInterface interface {
	Create(ctx context.Context, mLAClusterConfiguration *v1.MLAClusterConfiguration, opts metav1.CreateOptions) (*v1.MLAClusterConfiguration, error)
	Update(ctx context.Context, mLAClusterConfiguration *v1.MLAClusterConfiguration, opts metav1.UpdateOptions) (*v1.MLAClusterConfiguration, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.MLAClusterConfiguration, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.MLAClusterConfigurationList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.MLAClusterConfiguration, err error)
	MLAClusterConfigurationExpansion
}

// mLAClusterConfigurations implements MLAClusterConfigurationInterface
type mLAClusterConfigurations struct {
	client rest.Interface
	ns     string
}

// newMLAClusterConfigurations returns a MLAClusterConfigurations
func newMLAClusterConfigurations(c *KubermaticEnterpriseV1Client, namespace string) *mLAClusterConfigurations {
	return &mLAClusterConfigurations{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the mLAClusterConfiguration, and returns the corresponding mLAClusterConfiguration object, and an error if there is any.
func (c *mLAClusterConfigurations) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.MLAClusterConfiguration, err error) {
	result = &v1.MLAClusterConfiguration{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mlaclusterconfigurations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MLAClusterConfigurations that match those selectors.
func (c *mLAClusterConfigurations) List(ctx context.Context, opts metav1.ListOptions) (result *v1.MLAClusterConfigurationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.MLAClusterConfigurationList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mlaclusterconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested mLAClusterConfigurations.
func (c *mLAClusterConfigurations) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("mlaclusterconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a mLAClusterConfiguration and creates it.  Returns the server's representation of the mLAClusterConfiguration, and an error, if there is any.
func (c *mLAClusterConfigurations) Create(ctx context.Context, mLAClusterConfiguration *v1.MLAClusterConfiguration, opts metav1.CreateOptions) (result *v1.MLAClusterConfiguration, err error) {
	result = &v1.MLAClusterConfiguration{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("mlaclusterconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(mLAClusterConfiguration).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a mLAClusterConfiguration and updates it. Returns the server's representation of the mLAClusterConfiguration, and an error, if there is any.
func (c *mLAClusterConfigurations) Update(ctx context.Context, mLAClusterConfiguration *v1.MLAClusterConfiguration, opts metav1.UpdateOptions) (result *v1.MLAClusterConfiguration, err error) {
	result = &v1.MLAClusterConfiguration{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mlaclusterconfigurations").
		Name(mLAClusterConfiguration.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(mLAClusterConfiguration).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the mLAClusterConfiguration and deletes it. Returns an error if one occurs.
func (c *mLAClusterConfigurations) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mlaclusterconfigurations").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *mLAClusterConfigurations) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mlaclusterconfigurations").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched mLAClusterConfiguration.
func (c *mLAClusterConfigurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.MLAClusterConfiguration, err error) {
	result = &v1.MLAClusterConfiguration{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("mlaclusterconfigurations").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
