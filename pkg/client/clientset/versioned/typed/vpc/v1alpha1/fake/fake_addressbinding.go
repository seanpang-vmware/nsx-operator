/* Copyright © 2024 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: Apache-2.0 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/vmware-tanzu/nsx-operator/pkg/apis/vpc/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAddressBindings implements AddressBindingInterface
type FakeAddressBindings struct {
	Fake *FakeCrdV1alpha1
	ns   string
}

var addressbindingsResource = v1alpha1.SchemeGroupVersion.WithResource("addressbindings")

var addressbindingsKind = v1alpha1.SchemeGroupVersion.WithKind("AddressBinding")

// Get takes name of the addressBinding, and returns the corresponding addressBinding object, and an error if there is any.
func (c *FakeAddressBindings) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AddressBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(addressbindingsResource, c.ns, name), &v1alpha1.AddressBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AddressBinding), err
}

// List takes label and field selectors, and returns the list of AddressBindings that match those selectors.
func (c *FakeAddressBindings) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AddressBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(addressbindingsResource, addressbindingsKind, c.ns, opts), &v1alpha1.AddressBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AddressBindingList{ListMeta: obj.(*v1alpha1.AddressBindingList).ListMeta}
	for _, item := range obj.(*v1alpha1.AddressBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested addressBindings.
func (c *FakeAddressBindings) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(addressbindingsResource, c.ns, opts))

}

// Create takes the representation of a addressBinding and creates it.  Returns the server's representation of the addressBinding, and an error, if there is any.
func (c *FakeAddressBindings) Create(ctx context.Context, addressBinding *v1alpha1.AddressBinding, opts v1.CreateOptions) (result *v1alpha1.AddressBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(addressbindingsResource, c.ns, addressBinding), &v1alpha1.AddressBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AddressBinding), err
}

// Update takes the representation of a addressBinding and updates it. Returns the server's representation of the addressBinding, and an error, if there is any.
func (c *FakeAddressBindings) Update(ctx context.Context, addressBinding *v1alpha1.AddressBinding, opts v1.UpdateOptions) (result *v1alpha1.AddressBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(addressbindingsResource, c.ns, addressBinding), &v1alpha1.AddressBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AddressBinding), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAddressBindings) UpdateStatus(ctx context.Context, addressBinding *v1alpha1.AddressBinding, opts v1.UpdateOptions) (*v1alpha1.AddressBinding, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(addressbindingsResource, "status", c.ns, addressBinding), &v1alpha1.AddressBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AddressBinding), err
}

// Delete takes name of the addressBinding and deletes it. Returns an error if one occurs.
func (c *FakeAddressBindings) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(addressbindingsResource, c.ns, name, opts), &v1alpha1.AddressBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAddressBindings) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(addressbindingsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AddressBindingList{})
	return err
}

// Patch applies the patch and returns the patched addressBinding.
func (c *FakeAddressBindings) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AddressBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(addressbindingsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AddressBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AddressBinding), err
}
