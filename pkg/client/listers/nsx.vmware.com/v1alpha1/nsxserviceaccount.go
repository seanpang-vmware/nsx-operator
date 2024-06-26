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

// NSXServiceAccountLister helps list NSXServiceAccounts.
// All objects returned here must be treated as read-only.
type NSXServiceAccountLister interface {
	// List lists all NSXServiceAccounts in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.NSXServiceAccount, err error)
	// NSXServiceAccounts returns an object that can list and get NSXServiceAccounts.
	NSXServiceAccounts(namespace string) NSXServiceAccountNamespaceLister
	NSXServiceAccountListerExpansion
}

// nSXServiceAccountLister implements the NSXServiceAccountLister interface.
type nSXServiceAccountLister struct {
	indexer cache.Indexer
}

// NewNSXServiceAccountLister returns a new NSXServiceAccountLister.
func NewNSXServiceAccountLister(indexer cache.Indexer) NSXServiceAccountLister {
	return &nSXServiceAccountLister{indexer: indexer}
}

// List lists all NSXServiceAccounts in the indexer.
func (s *nSXServiceAccountLister) List(selector labels.Selector) (ret []*v1alpha1.NSXServiceAccount, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NSXServiceAccount))
	})
	return ret, err
}

// NSXServiceAccounts returns an object that can list and get NSXServiceAccounts.
func (s *nSXServiceAccountLister) NSXServiceAccounts(namespace string) NSXServiceAccountNamespaceLister {
	return nSXServiceAccountNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NSXServiceAccountNamespaceLister helps list and get NSXServiceAccounts.
// All objects returned here must be treated as read-only.
type NSXServiceAccountNamespaceLister interface {
	// List lists all NSXServiceAccounts in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.NSXServiceAccount, err error)
	// Get retrieves the NSXServiceAccount from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.NSXServiceAccount, error)
	NSXServiceAccountNamespaceListerExpansion
}

// nSXServiceAccountNamespaceLister implements the NSXServiceAccountNamespaceLister
// interface.
type nSXServiceAccountNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NSXServiceAccounts in the indexer for a given namespace.
func (s nSXServiceAccountNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.NSXServiceAccount, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.NSXServiceAccount))
	})
	return ret, err
}

// Get retrieves the NSXServiceAccount from the indexer for a given namespace and name.
func (s nSXServiceAccountNamespaceLister) Get(name string) (*v1alpha1.NSXServiceAccount, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("nsxserviceaccount"), name)
	}
	return obj.(*v1alpha1.NSXServiceAccount), nil
}
