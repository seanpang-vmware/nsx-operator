/* Copyright © 2024 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: Apache-2.0 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/vmware-tanzu/nsx-operator/pkg/apis/nsx.vmware.com/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIPAddressAllocations implements IPAddressAllocationInterface
type FakeIPAddressAllocations struct {
	Fake *FakeNsxV1alpha1
	ns   string
}

var ipaddressallocationsResource = v1alpha1.SchemeGroupVersion.WithResource("ipaddressallocations")

var ipaddressallocationsKind = v1alpha1.SchemeGroupVersion.WithKind("IPAddressAllocation")

// Get takes name of the iPAddressAllocation, and returns the corresponding iPAddressAllocation object, and an error if there is any.
func (c *FakeIPAddressAllocations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.IPAddressAllocation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ipaddressallocationsResource, c.ns, name), &v1alpha1.IPAddressAllocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IPAddressAllocation), err
}

// List takes label and field selectors, and returns the list of IPAddressAllocations that match those selectors.
func (c *FakeIPAddressAllocations) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.IPAddressAllocationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ipaddressallocationsResource, ipaddressallocationsKind, c.ns, opts), &v1alpha1.IPAddressAllocationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IPAddressAllocationList{ListMeta: obj.(*v1alpha1.IPAddressAllocationList).ListMeta}
	for _, item := range obj.(*v1alpha1.IPAddressAllocationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iPAddressAllocations.
func (c *FakeIPAddressAllocations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ipaddressallocationsResource, c.ns, opts))

}

// Create takes the representation of a iPAddressAllocation and creates it.  Returns the server's representation of the iPAddressAllocation, and an error, if there is any.
func (c *FakeIPAddressAllocations) Create(ctx context.Context, iPAddressAllocation *v1alpha1.IPAddressAllocation, opts v1.CreateOptions) (result *v1alpha1.IPAddressAllocation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ipaddressallocationsResource, c.ns, iPAddressAllocation), &v1alpha1.IPAddressAllocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IPAddressAllocation), err
}

// Update takes the representation of a iPAddressAllocation and updates it. Returns the server's representation of the iPAddressAllocation, and an error, if there is any.
func (c *FakeIPAddressAllocations) Update(ctx context.Context, iPAddressAllocation *v1alpha1.IPAddressAllocation, opts v1.UpdateOptions) (result *v1alpha1.IPAddressAllocation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ipaddressallocationsResource, c.ns, iPAddressAllocation), &v1alpha1.IPAddressAllocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IPAddressAllocation), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIPAddressAllocations) UpdateStatus(ctx context.Context, iPAddressAllocation *v1alpha1.IPAddressAllocation, opts v1.UpdateOptions) (*v1alpha1.IPAddressAllocation, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ipaddressallocationsResource, "status", c.ns, iPAddressAllocation), &v1alpha1.IPAddressAllocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IPAddressAllocation), err
}

// Delete takes name of the iPAddressAllocation and deletes it. Returns an error if one occurs.
func (c *FakeIPAddressAllocations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(ipaddressallocationsResource, c.ns, name, opts), &v1alpha1.IPAddressAllocation{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIPAddressAllocations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ipaddressallocationsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.IPAddressAllocationList{})
	return err
}

// Patch applies the patch and returns the patched iPAddressAllocation.
func (c *FakeIPAddressAllocations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.IPAddressAllocation, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ipaddressallocationsResource, c.ns, name, pt, data, subresources...), &v1alpha1.IPAddressAllocation{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.IPAddressAllocation), err
}
