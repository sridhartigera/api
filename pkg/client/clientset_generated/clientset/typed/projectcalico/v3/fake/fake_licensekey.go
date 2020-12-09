// Copyright (c) 2020 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v3 "github.com/tigera/api/pkg/apis/projectcalico/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeLicenseKeys implements LicenseKeyInterface
type FakeLicenseKeys struct {
	Fake *FakeProjectcalicoV3
	ns   string
}

var licensekeysResource = schema.GroupVersionResource{Group: "projectcalico.org", Version: "v3", Resource: "licensekeys"}

var licensekeysKind = schema.GroupVersionKind{Group: "projectcalico.org", Version: "v3", Kind: "LicenseKey"}

// Get takes name of the licenseKey, and returns the corresponding licenseKey object, and an error if there is any.
func (c *FakeLicenseKeys) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.LicenseKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(licensekeysResource, c.ns, name), &v3.LicenseKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.LicenseKey), err
}

// List takes label and field selectors, and returns the list of LicenseKeys that match those selectors.
func (c *FakeLicenseKeys) List(ctx context.Context, opts v1.ListOptions) (result *v3.LicenseKeyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(licensekeysResource, licensekeysKind, c.ns, opts), &v3.LicenseKeyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.LicenseKeyList{ListMeta: obj.(*v3.LicenseKeyList).ListMeta}
	for _, item := range obj.(*v3.LicenseKeyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested licenseKeys.
func (c *FakeLicenseKeys) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(licensekeysResource, c.ns, opts))

}

// Create takes the representation of a licenseKey and creates it.  Returns the server's representation of the licenseKey, and an error, if there is any.
func (c *FakeLicenseKeys) Create(ctx context.Context, licenseKey *v3.LicenseKey, opts v1.CreateOptions) (result *v3.LicenseKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(licensekeysResource, c.ns, licenseKey), &v3.LicenseKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.LicenseKey), err
}

// Update takes the representation of a licenseKey and updates it. Returns the server's representation of the licenseKey, and an error, if there is any.
func (c *FakeLicenseKeys) Update(ctx context.Context, licenseKey *v3.LicenseKey, opts v1.UpdateOptions) (result *v3.LicenseKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(licensekeysResource, c.ns, licenseKey), &v3.LicenseKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.LicenseKey), err
}

// Delete takes name of the licenseKey and deletes it. Returns an error if one occurs.
func (c *FakeLicenseKeys) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(licensekeysResource, c.ns, name), &v3.LicenseKey{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLicenseKeys) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(licensekeysResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v3.LicenseKeyList{})
	return err
}

// Patch applies the patch and returns the patched licenseKey.
func (c *FakeLicenseKeys) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.LicenseKey, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(licensekeysResource, c.ns, name, pt, data, subresources...), &v3.LicenseKey{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.LicenseKey), err
}
