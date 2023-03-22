/*
Copyright 2023 The Kubermatic Kubernetes Platform contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "k8c.io/api/v2/pkg/apis/kubermatic/v1"
	scheme "k8c.io/api/v2/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// UserProjectBindingsGetter has a method to return a UserProjectBindingInterface.
// A group's client should implement this interface.
type UserProjectBindingsGetter interface {
	UserProjectBindings(namespace string) UserProjectBindingInterface
}

// UserProjectBindingInterface has methods to work with UserProjectBinding resources.
type UserProjectBindingInterface interface {
	Create(ctx context.Context, userProjectBinding *v1.UserProjectBinding, opts metav1.CreateOptions) (*v1.UserProjectBinding, error)
	Update(ctx context.Context, userProjectBinding *v1.UserProjectBinding, opts metav1.UpdateOptions) (*v1.UserProjectBinding, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.UserProjectBinding, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.UserProjectBindingList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.UserProjectBinding, err error)
	UserProjectBindingExpansion
}

// userProjectBindings implements UserProjectBindingInterface
type userProjectBindings struct {
	client rest.Interface
	ns     string
}

// newUserProjectBindings returns a UserProjectBindings
func newUserProjectBindings(c *KubermaticV1Client, namespace string) *userProjectBindings {
	return &userProjectBindings{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the userProjectBinding, and returns the corresponding userProjectBinding object, and an error if there is any.
func (c *userProjectBindings) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.UserProjectBinding, err error) {
	result = &v1.UserProjectBinding{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("userprojectbindings").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of UserProjectBindings that match those selectors.
func (c *userProjectBindings) List(ctx context.Context, opts metav1.ListOptions) (result *v1.UserProjectBindingList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.UserProjectBindingList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("userprojectbindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested userProjectBindings.
func (c *userProjectBindings) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("userprojectbindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a userProjectBinding and creates it.  Returns the server's representation of the userProjectBinding, and an error, if there is any.
func (c *userProjectBindings) Create(ctx context.Context, userProjectBinding *v1.UserProjectBinding, opts metav1.CreateOptions) (result *v1.UserProjectBinding, err error) {
	result = &v1.UserProjectBinding{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("userprojectbindings").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(userProjectBinding).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a userProjectBinding and updates it. Returns the server's representation of the userProjectBinding, and an error, if there is any.
func (c *userProjectBindings) Update(ctx context.Context, userProjectBinding *v1.UserProjectBinding, opts metav1.UpdateOptions) (result *v1.UserProjectBinding, err error) {
	result = &v1.UserProjectBinding{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("userprojectbindings").
		Name(userProjectBinding.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(userProjectBinding).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the userProjectBinding and deletes it. Returns an error if one occurs.
func (c *userProjectBindings) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("userprojectbindings").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *userProjectBindings) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("userprojectbindings").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched userProjectBinding.
func (c *userProjectBindings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.UserProjectBinding, err error) {
	result = &v1.UserProjectBinding{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("userprojectbindings").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
