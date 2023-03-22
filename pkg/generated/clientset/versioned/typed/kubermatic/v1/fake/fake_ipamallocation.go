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

package fake

import (
	"context"

	kubermaticv1 "k8c.io/api/v2/pkg/apis/kubermatic/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIPAMAllocations implements IPAMAllocationInterface
type FakeIPAMAllocations struct {
	Fake *FakeKubermaticV1
	ns   string
}

var ipamallocationsResource = schema.GroupVersionResource{Group: "kubermatic.k8c.io", Version: "v1", Resource: "ipamallocations"}

var ipamallocationsKind = schema.GroupVersionKind{Group: "kubermatic.k8c.io", Version: "v1", Kind: "IPAMAllocation"}

// Get takes name of the iPAMAllocation, and returns the corresponding iPAMAllocation object, and an error if there is any.
func (c *FakeIPAMAllocations) Get(ctx context.Context, name string, options v1.GetOptions) (result *kubermaticv1.IPAMAllocation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ipamallocationsResource, c.ns, name), &kubermaticv1.IPAMAllocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.IPAMAllocation), err
}

// List takes label and field selectors, and returns the list of IPAMAllocations that match those selectors.
func (c *FakeIPAMAllocations) List(ctx context.Context, opts v1.ListOptions) (result *kubermaticv1.IPAMAllocationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ipamallocationsResource, ipamallocationsKind, c.ns, opts), &kubermaticv1.IPAMAllocationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &kubermaticv1.IPAMAllocationList{ListMeta: obj.(*kubermaticv1.IPAMAllocationList).ListMeta}
	for _, item := range obj.(*kubermaticv1.IPAMAllocationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iPAMAllocations.
func (c *FakeIPAMAllocations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ipamallocationsResource, c.ns, opts))

}

// Create takes the representation of a iPAMAllocation and creates it.  Returns the server's representation of the iPAMAllocation, and an error, if there is any.
func (c *FakeIPAMAllocations) Create(ctx context.Context, iPAMAllocation *kubermaticv1.IPAMAllocation, opts v1.CreateOptions) (result *kubermaticv1.IPAMAllocation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ipamallocationsResource, c.ns, iPAMAllocation), &kubermaticv1.IPAMAllocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.IPAMAllocation), err
}

// Update takes the representation of a iPAMAllocation and updates it. Returns the server's representation of the iPAMAllocation, and an error, if there is any.
func (c *FakeIPAMAllocations) Update(ctx context.Context, iPAMAllocation *kubermaticv1.IPAMAllocation, opts v1.UpdateOptions) (result *kubermaticv1.IPAMAllocation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ipamallocationsResource, c.ns, iPAMAllocation), &kubermaticv1.IPAMAllocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.IPAMAllocation), err
}

// Delete takes name of the iPAMAllocation and deletes it. Returns an error if one occurs.
func (c *FakeIPAMAllocations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(ipamallocationsResource, c.ns, name, opts), &kubermaticv1.IPAMAllocation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIPAMAllocations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ipamallocationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &kubermaticv1.IPAMAllocationList{})
	return err
}

// Patch applies the patch and returns the patched iPAMAllocation.
func (c *FakeIPAMAllocations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *kubermaticv1.IPAMAllocation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ipamallocationsResource, c.ns, name, pt, data, subresources...), &kubermaticv1.IPAMAllocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.IPAMAllocation), err
}
