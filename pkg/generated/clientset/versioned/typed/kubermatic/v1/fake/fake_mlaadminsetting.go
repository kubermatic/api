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

// FakeMLAAdminSettings implements MLAAdminSettingInterface
type FakeMLAAdminSettings struct {
	Fake *FakeKubermaticV1
	ns   string
}

var mlaadminsettingsResource = schema.GroupVersionResource{Group: "kubermatic.k8c.io", Version: "v1", Resource: "mlaadminsettings"}

var mlaadminsettingsKind = schema.GroupVersionKind{Group: "kubermatic.k8c.io", Version: "v1", Kind: "MLAAdminSetting"}

// Get takes name of the mLAAdminSetting, and returns the corresponding mLAAdminSetting object, and an error if there is any.
func (c *FakeMLAAdminSettings) Get(ctx context.Context, name string, options v1.GetOptions) (result *kubermaticv1.MLAAdminSetting, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(mlaadminsettingsResource, c.ns, name), &kubermaticv1.MLAAdminSetting{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.MLAAdminSetting), err
}

// List takes label and field selectors, and returns the list of MLAAdminSettings that match those selectors.
func (c *FakeMLAAdminSettings) List(ctx context.Context, opts v1.ListOptions) (result *kubermaticv1.MLAAdminSettingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(mlaadminsettingsResource, mlaadminsettingsKind, c.ns, opts), &kubermaticv1.MLAAdminSettingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &kubermaticv1.MLAAdminSettingList{ListMeta: obj.(*kubermaticv1.MLAAdminSettingList).ListMeta}
	for _, item := range obj.(*kubermaticv1.MLAAdminSettingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested mLAAdminSettings.
func (c *FakeMLAAdminSettings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(mlaadminsettingsResource, c.ns, opts))

}

// Create takes the representation of a mLAAdminSetting and creates it.  Returns the server's representation of the mLAAdminSetting, and an error, if there is any.
func (c *FakeMLAAdminSettings) Create(ctx context.Context, mLAAdminSetting *kubermaticv1.MLAAdminSetting, opts v1.CreateOptions) (result *kubermaticv1.MLAAdminSetting, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(mlaadminsettingsResource, c.ns, mLAAdminSetting), &kubermaticv1.MLAAdminSetting{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.MLAAdminSetting), err
}

// Update takes the representation of a mLAAdminSetting and updates it. Returns the server's representation of the mLAAdminSetting, and an error, if there is any.
func (c *FakeMLAAdminSettings) Update(ctx context.Context, mLAAdminSetting *kubermaticv1.MLAAdminSetting, opts v1.UpdateOptions) (result *kubermaticv1.MLAAdminSetting, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(mlaadminsettingsResource, c.ns, mLAAdminSetting), &kubermaticv1.MLAAdminSetting{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.MLAAdminSetting), err
}

// Delete takes name of the mLAAdminSetting and deletes it. Returns an error if one occurs.
func (c *FakeMLAAdminSettings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(mlaadminsettingsResource, c.ns, name, opts), &kubermaticv1.MLAAdminSetting{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMLAAdminSettings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(mlaadminsettingsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &kubermaticv1.MLAAdminSettingList{})
	return err
}

// Patch applies the patch and returns the patched mLAAdminSetting.
func (c *FakeMLAAdminSettings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *kubermaticv1.MLAAdminSetting, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(mlaadminsettingsResource, c.ns, name, pt, data, subresources...), &kubermaticv1.MLAAdminSetting{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kubermaticv1.MLAAdminSetting), err
}
