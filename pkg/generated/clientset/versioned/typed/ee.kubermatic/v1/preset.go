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

// PresetsGetter has a method to return a PresetInterface.
// A group's client should implement this interface.
type PresetsGetter interface {
	Presets(namespace string) PresetInterface
}

// PresetInterface has methods to work with Preset resources.
type PresetInterface interface {
	Create(ctx context.Context, preset *v1.Preset, opts metav1.CreateOptions) (*v1.Preset, error)
	Update(ctx context.Context, preset *v1.Preset, opts metav1.UpdateOptions) (*v1.Preset, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Preset, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.PresetList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Preset, err error)
	PresetExpansion
}

// presets implements PresetInterface
type presets struct {
	client rest.Interface
	ns     string
}

// newPresets returns a Presets
func newPresets(c *KubermaticEnterpriseV1Client, namespace string) *presets {
	return &presets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the preset, and returns the corresponding preset object, and an error if there is any.
func (c *presets) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Preset, err error) {
	result = &v1.Preset{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("presets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Presets that match those selectors.
func (c *presets) List(ctx context.Context, opts metav1.ListOptions) (result *v1.PresetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.PresetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("presets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested presets.
func (c *presets) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("presets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a preset and creates it.  Returns the server's representation of the preset, and an error, if there is any.
func (c *presets) Create(ctx context.Context, preset *v1.Preset, opts metav1.CreateOptions) (result *v1.Preset, err error) {
	result = &v1.Preset{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("presets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(preset).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a preset and updates it. Returns the server's representation of the preset, and an error, if there is any.
func (c *presets) Update(ctx context.Context, preset *v1.Preset, opts metav1.UpdateOptions) (result *v1.Preset, err error) {
	result = &v1.Preset{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("presets").
		Name(preset.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(preset).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the preset and deletes it. Returns an error if one occurs.
func (c *presets) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("presets").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *presets) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("presets").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched preset.
func (c *presets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Preset, err error) {
	result = &v1.Preset{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("presets").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
