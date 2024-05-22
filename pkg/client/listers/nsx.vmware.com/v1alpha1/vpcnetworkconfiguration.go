/* Copyright © 2024 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: Apache-2.0 */

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/vmware-tanzu/nsx-operator/pkg/apis/nsx.vmware.com/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VPCNetworkConfigurationLister helps list VPCNetworkConfigurations.
// All objects returned here must be treated as read-only.
type VPCNetworkConfigurationLister interface {
	// List lists all VPCNetworkConfigurations in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.VPCNetworkConfiguration, err error)
	// Get retrieves the VPCNetworkConfiguration from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.VPCNetworkConfiguration, error)
	VPCNetworkConfigurationListerExpansion
}

// vPCNetworkConfigurationLister implements the VPCNetworkConfigurationLister interface.
type vPCNetworkConfigurationLister struct {
	indexer cache.Indexer
}

// NewVPCNetworkConfigurationLister returns a new VPCNetworkConfigurationLister.
func NewVPCNetworkConfigurationLister(indexer cache.Indexer) VPCNetworkConfigurationLister {
	return &vPCNetworkConfigurationLister{indexer: indexer}
}

// List lists all VPCNetworkConfigurations in the indexer.
func (s *vPCNetworkConfigurationLister) List(selector labels.Selector) (ret []*v1alpha1.VPCNetworkConfiguration, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.VPCNetworkConfiguration))
	})
	return ret, err
}

// Get retrieves the VPCNetworkConfiguration from the index for a given name.
func (s *vPCNetworkConfigurationLister) Get(name string) (*v1alpha1.VPCNetworkConfiguration, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("vpcnetworkconfiguration"), name)
	}
	return obj.(*v1alpha1.VPCNetworkConfiguration), nil
}
