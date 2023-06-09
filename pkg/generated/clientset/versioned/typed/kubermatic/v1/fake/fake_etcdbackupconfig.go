// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8c.io/api/v3/pkg/apis/kubermatic/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEtcdBackupConfigs implements EtcdBackupConfigInterface
type FakeEtcdBackupConfigs struct {
	Fake *FakeKubermaticV1
	ns   string
}

var etcdbackupconfigsResource = v1.SchemeGroupVersion.WithResource("etcdbackupconfigs")

var etcdbackupconfigsKind = v1.SchemeGroupVersion.WithKind("EtcdBackupConfig")

// Get takes name of the etcdBackupConfig, and returns the corresponding etcdBackupConfig object, and an error if there is any.
func (c *FakeEtcdBackupConfigs) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.EtcdBackupConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(etcdbackupconfigsResource, c.ns, name), &v1.EtcdBackupConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.EtcdBackupConfig), err
}

// List takes label and field selectors, and returns the list of EtcdBackupConfigs that match those selectors.
func (c *FakeEtcdBackupConfigs) List(ctx context.Context, opts metav1.ListOptions) (result *v1.EtcdBackupConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(etcdbackupconfigsResource, etcdbackupconfigsKind, c.ns, opts), &v1.EtcdBackupConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.EtcdBackupConfigList{ListMeta: obj.(*v1.EtcdBackupConfigList).ListMeta}
	for _, item := range obj.(*v1.EtcdBackupConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested etcdBackupConfigs.
func (c *FakeEtcdBackupConfigs) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(etcdbackupconfigsResource, c.ns, opts))

}

// Create takes the representation of a etcdBackupConfig and creates it.  Returns the server's representation of the etcdBackupConfig, and an error, if there is any.
func (c *FakeEtcdBackupConfigs) Create(ctx context.Context, etcdBackupConfig *v1.EtcdBackupConfig, opts metav1.CreateOptions) (result *v1.EtcdBackupConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(etcdbackupconfigsResource, c.ns, etcdBackupConfig), &v1.EtcdBackupConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.EtcdBackupConfig), err
}

// Update takes the representation of a etcdBackupConfig and updates it. Returns the server's representation of the etcdBackupConfig, and an error, if there is any.
func (c *FakeEtcdBackupConfigs) Update(ctx context.Context, etcdBackupConfig *v1.EtcdBackupConfig, opts metav1.UpdateOptions) (result *v1.EtcdBackupConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(etcdbackupconfigsResource, c.ns, etcdBackupConfig), &v1.EtcdBackupConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.EtcdBackupConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEtcdBackupConfigs) UpdateStatus(ctx context.Context, etcdBackupConfig *v1.EtcdBackupConfig, opts metav1.UpdateOptions) (*v1.EtcdBackupConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(etcdbackupconfigsResource, "status", c.ns, etcdBackupConfig), &v1.EtcdBackupConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.EtcdBackupConfig), err
}

// Delete takes name of the etcdBackupConfig and deletes it. Returns an error if one occurs.
func (c *FakeEtcdBackupConfigs) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(etcdbackupconfigsResource, c.ns, name, opts), &v1.EtcdBackupConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEtcdBackupConfigs) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(etcdbackupconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.EtcdBackupConfigList{})
	return err
}

// Patch applies the patch and returns the patched etcdBackupConfig.
func (c *FakeEtcdBackupConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.EtcdBackupConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(etcdbackupconfigsResource, c.ns, name, pt, data, subresources...), &v1.EtcdBackupConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.EtcdBackupConfig), err
}
