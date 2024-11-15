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

// FakeSubnetConnectionBindingMaps implements SubnetConnectionBindingMapInterface
type FakeSubnetConnectionBindingMaps struct {
	Fake *FakeCrdV1alpha1
	ns   string
}

var subnetconnectionbindingmapsResource = v1alpha1.SchemeGroupVersion.WithResource("subnetconnectionbindingmaps")

var subnetconnectionbindingmapsKind = v1alpha1.SchemeGroupVersion.WithKind("SubnetConnectionBindingMap")

// Get takes name of the subnetConnectionBindingMap, and returns the corresponding subnetConnectionBindingMap object, and an error if there is any.
func (c *FakeSubnetConnectionBindingMaps) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SubnetConnectionBindingMap, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(subnetconnectionbindingmapsResource, c.ns, name), &v1alpha1.SubnetConnectionBindingMap{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SubnetConnectionBindingMap), err
}

// List takes label and field selectors, and returns the list of SubnetConnectionBindingMaps that match those selectors.
func (c *FakeSubnetConnectionBindingMaps) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SubnetConnectionBindingMapList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(subnetconnectionbindingmapsResource, subnetconnectionbindingmapsKind, c.ns, opts), &v1alpha1.SubnetConnectionBindingMapList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SubnetConnectionBindingMapList{ListMeta: obj.(*v1alpha1.SubnetConnectionBindingMapList).ListMeta}
	for _, item := range obj.(*v1alpha1.SubnetConnectionBindingMapList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested subnetConnectionBindingMaps.
func (c *FakeSubnetConnectionBindingMaps) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(subnetconnectionbindingmapsResource, c.ns, opts))

}

// Create takes the representation of a subnetConnectionBindingMap and creates it.  Returns the server's representation of the subnetConnectionBindingMap, and an error, if there is any.
func (c *FakeSubnetConnectionBindingMaps) Create(ctx context.Context, subnetConnectionBindingMap *v1alpha1.SubnetConnectionBindingMap, opts v1.CreateOptions) (result *v1alpha1.SubnetConnectionBindingMap, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(subnetconnectionbindingmapsResource, c.ns, subnetConnectionBindingMap), &v1alpha1.SubnetConnectionBindingMap{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SubnetConnectionBindingMap), err
}

// Update takes the representation of a subnetConnectionBindingMap and updates it. Returns the server's representation of the subnetConnectionBindingMap, and an error, if there is any.
func (c *FakeSubnetConnectionBindingMaps) Update(ctx context.Context, subnetConnectionBindingMap *v1alpha1.SubnetConnectionBindingMap, opts v1.UpdateOptions) (result *v1alpha1.SubnetConnectionBindingMap, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(subnetconnectionbindingmapsResource, c.ns, subnetConnectionBindingMap), &v1alpha1.SubnetConnectionBindingMap{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SubnetConnectionBindingMap), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSubnetConnectionBindingMaps) UpdateStatus(ctx context.Context, subnetConnectionBindingMap *v1alpha1.SubnetConnectionBindingMap, opts v1.UpdateOptions) (*v1alpha1.SubnetConnectionBindingMap, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(subnetconnectionbindingmapsResource, "status", c.ns, subnetConnectionBindingMap), &v1alpha1.SubnetConnectionBindingMap{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SubnetConnectionBindingMap), err
}

// Delete takes name of the subnetConnectionBindingMap and deletes it. Returns an error if one occurs.
func (c *FakeSubnetConnectionBindingMaps) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(subnetconnectionbindingmapsResource, c.ns, name, opts), &v1alpha1.SubnetConnectionBindingMap{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSubnetConnectionBindingMaps) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(subnetconnectionbindingmapsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SubnetConnectionBindingMapList{})
	return err
}

// Patch applies the patch and returns the patched subnetConnectionBindingMap.
func (c *FakeSubnetConnectionBindingMaps) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SubnetConnectionBindingMap, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(subnetconnectionbindingmapsResource, c.ns, name, pt, data, subresources...), &v1alpha1.SubnetConnectionBindingMap{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SubnetConnectionBindingMap), err
}