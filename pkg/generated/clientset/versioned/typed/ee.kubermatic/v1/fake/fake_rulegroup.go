// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	eekubermaticv1 "k8c.io/api/v3/pkg/apis/ee.kubermatic/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRuleGroups implements RuleGroupInterface
type FakeRuleGroups struct {
	Fake *FakeEeKubermaticV1
	ns   string
}

var rulegroupsResource = schema.GroupVersionResource{Group: "ee.kubermatic.k8c.io", Version: "v1", Resource: "rulegroups"}

var rulegroupsKind = schema.GroupVersionKind{Group: "ee.kubermatic.k8c.io", Version: "v1", Kind: "RuleGroup"}

// Get takes name of the ruleGroup, and returns the corresponding ruleGroup object, and an error if there is any.
func (c *FakeRuleGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *eekubermaticv1.RuleGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(rulegroupsResource, c.ns, name), &eekubermaticv1.RuleGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*eekubermaticv1.RuleGroup), err
}

// List takes label and field selectors, and returns the list of RuleGroups that match those selectors.
func (c *FakeRuleGroups) List(ctx context.Context, opts v1.ListOptions) (result *eekubermaticv1.RuleGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(rulegroupsResource, rulegroupsKind, c.ns, opts), &eekubermaticv1.RuleGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &eekubermaticv1.RuleGroupList{ListMeta: obj.(*eekubermaticv1.RuleGroupList).ListMeta}
	for _, item := range obj.(*eekubermaticv1.RuleGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ruleGroups.
func (c *FakeRuleGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(rulegroupsResource, c.ns, opts))

}

// Create takes the representation of a ruleGroup and creates it.  Returns the server's representation of the ruleGroup, and an error, if there is any.
func (c *FakeRuleGroups) Create(ctx context.Context, ruleGroup *eekubermaticv1.RuleGroup, opts v1.CreateOptions) (result *eekubermaticv1.RuleGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(rulegroupsResource, c.ns, ruleGroup), &eekubermaticv1.RuleGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*eekubermaticv1.RuleGroup), err
}

// Update takes the representation of a ruleGroup and updates it. Returns the server's representation of the ruleGroup, and an error, if there is any.
func (c *FakeRuleGroups) Update(ctx context.Context, ruleGroup *eekubermaticv1.RuleGroup, opts v1.UpdateOptions) (result *eekubermaticv1.RuleGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(rulegroupsResource, c.ns, ruleGroup), &eekubermaticv1.RuleGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*eekubermaticv1.RuleGroup), err
}

// Delete takes name of the ruleGroup and deletes it. Returns an error if one occurs.
func (c *FakeRuleGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(rulegroupsResource, c.ns, name, opts), &eekubermaticv1.RuleGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRuleGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(rulegroupsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &eekubermaticv1.RuleGroupList{})
	return err
}

// Patch applies the patch and returns the patched ruleGroup.
func (c *FakeRuleGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *eekubermaticv1.RuleGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(rulegroupsResource, c.ns, name, pt, data, subresources...), &eekubermaticv1.RuleGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*eekubermaticv1.RuleGroup), err
}
