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

// FakeEtcdBackupConfigs implements EtcdBackupConfigInterface
type FakeEtcdBackupConfigs struct {
	Fake *FakeKubermaticEnterpriseV1
	ns   string
}

var etcdbackupconfigsResource = schema.GroupVersionResource{Group: "ee.kubermatic.k8c.io", Version: "v1", Resource: "etcdbackupconfigs"}

var etcdbackupconfigsKind = schema.GroupVersionKind{Group: "ee.kubermatic.k8c.io", Version: "v1", Kind: "EtcdBackupConfig"}

// Get takes name of the etcdBackupConfig, and returns the corresponding etcdBackupConfig object, and an error if there is any.
func (c *FakeEtcdBackupConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *eekubermaticv1.EtcdBackupConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(etcdbackupconfigsResource, c.ns, name), &eekubermaticv1.EtcdBackupConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*eekubermaticv1.EtcdBackupConfig), err
}

// List takes label and field selectors, and returns the list of EtcdBackupConfigs that match those selectors.
func (c *FakeEtcdBackupConfigs) List(ctx context.Context, opts v1.ListOptions) (result *eekubermaticv1.EtcdBackupConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(etcdbackupconfigsResource, etcdbackupconfigsKind, c.ns, opts), &eekubermaticv1.EtcdBackupConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &eekubermaticv1.EtcdBackupConfigList{ListMeta: obj.(*eekubermaticv1.EtcdBackupConfigList).ListMeta}
	for _, item := range obj.(*eekubermaticv1.EtcdBackupConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested etcdBackupConfigs.
func (c *FakeEtcdBackupConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(etcdbackupconfigsResource, c.ns, opts))

}

// Create takes the representation of a etcdBackupConfig and creates it.  Returns the server's representation of the etcdBackupConfig, and an error, if there is any.
func (c *FakeEtcdBackupConfigs) Create(ctx context.Context, etcdBackupConfig *eekubermaticv1.EtcdBackupConfig, opts v1.CreateOptions) (result *eekubermaticv1.EtcdBackupConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(etcdbackupconfigsResource, c.ns, etcdBackupConfig), &eekubermaticv1.EtcdBackupConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*eekubermaticv1.EtcdBackupConfig), err
}

// Update takes the representation of a etcdBackupConfig and updates it. Returns the server's representation of the etcdBackupConfig, and an error, if there is any.
func (c *FakeEtcdBackupConfigs) Update(ctx context.Context, etcdBackupConfig *eekubermaticv1.EtcdBackupConfig, opts v1.UpdateOptions) (result *eekubermaticv1.EtcdBackupConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(etcdbackupconfigsResource, c.ns, etcdBackupConfig), &eekubermaticv1.EtcdBackupConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*eekubermaticv1.EtcdBackupConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEtcdBackupConfigs) UpdateStatus(ctx context.Context, etcdBackupConfig *eekubermaticv1.EtcdBackupConfig, opts v1.UpdateOptions) (*eekubermaticv1.EtcdBackupConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(etcdbackupconfigsResource, "status", c.ns, etcdBackupConfig), &eekubermaticv1.EtcdBackupConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*eekubermaticv1.EtcdBackupConfig), err
}

// Delete takes name of the etcdBackupConfig and deletes it. Returns an error if one occurs.
func (c *FakeEtcdBackupConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(etcdbackupconfigsResource, c.ns, name, opts), &eekubermaticv1.EtcdBackupConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEtcdBackupConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(etcdbackupconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &eekubermaticv1.EtcdBackupConfigList{})
	return err
}

// Patch applies the patch and returns the patched etcdBackupConfig.
func (c *FakeEtcdBackupConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *eekubermaticv1.EtcdBackupConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(etcdbackupconfigsResource, c.ns, name, pt, data, subresources...), &eekubermaticv1.EtcdBackupConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*eekubermaticv1.EtcdBackupConfig), err
}
