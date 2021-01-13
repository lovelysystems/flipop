// SPDX-License-Identifier: Apache-2.0
//
// Copyright 2021 Digital Ocean, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"time"

	scheme "github.com/digitalocean/flipop/pkg/apis/flipop/generated/clientset/versioned/scheme"
	v1alpha1 "github.com/digitalocean/flipop/pkg/apis/flipop/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NodeDNSRecordSetsGetter has a method to return a NodeDNSRecordSetInterface.
// A group's client should implement this interface.
type NodeDNSRecordSetsGetter interface {
	NodeDNSRecordSets(namespace string) NodeDNSRecordSetInterface
}

// NodeDNSRecordSetInterface has methods to work with NodeDNSRecordSet resources.
type NodeDNSRecordSetInterface interface {
	Create(*v1alpha1.NodeDNSRecordSet) (*v1alpha1.NodeDNSRecordSet, error)
	Update(*v1alpha1.NodeDNSRecordSet) (*v1alpha1.NodeDNSRecordSet, error)
	UpdateStatus(*v1alpha1.NodeDNSRecordSet) (*v1alpha1.NodeDNSRecordSet, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.NodeDNSRecordSet, error)
	List(opts v1.ListOptions) (*v1alpha1.NodeDNSRecordSetList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NodeDNSRecordSet, err error)
	NodeDNSRecordSetExpansion
}

// nodeDNSRecordSets implements NodeDNSRecordSetInterface
type nodeDNSRecordSets struct {
	client rest.Interface
	ns     string
}

// newNodeDNSRecordSets returns a NodeDNSRecordSets
func newNodeDNSRecordSets(c *FlipopV1alpha1Client, namespace string) *nodeDNSRecordSets {
	return &nodeDNSRecordSets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the nodeDNSRecordSet, and returns the corresponding nodeDNSRecordSet object, and an error if there is any.
func (c *nodeDNSRecordSets) Get(name string, options v1.GetOptions) (result *v1alpha1.NodeDNSRecordSet, err error) {
	result = &v1alpha1.NodeDNSRecordSet{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("nodednsrecordsets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NodeDNSRecordSets that match those selectors.
func (c *nodeDNSRecordSets) List(opts v1.ListOptions) (result *v1alpha1.NodeDNSRecordSetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.NodeDNSRecordSetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("nodednsrecordsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested nodeDNSRecordSets.
func (c *nodeDNSRecordSets) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("nodednsrecordsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a nodeDNSRecordSet and creates it.  Returns the server's representation of the nodeDNSRecordSet, and an error, if there is any.
func (c *nodeDNSRecordSets) Create(nodeDNSRecordSet *v1alpha1.NodeDNSRecordSet) (result *v1alpha1.NodeDNSRecordSet, err error) {
	result = &v1alpha1.NodeDNSRecordSet{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("nodednsrecordsets").
		Body(nodeDNSRecordSet).
		Do().
		Into(result)
	return
}

// Update takes the representation of a nodeDNSRecordSet and updates it. Returns the server's representation of the nodeDNSRecordSet, and an error, if there is any.
func (c *nodeDNSRecordSets) Update(nodeDNSRecordSet *v1alpha1.NodeDNSRecordSet) (result *v1alpha1.NodeDNSRecordSet, err error) {
	result = &v1alpha1.NodeDNSRecordSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("nodednsrecordsets").
		Name(nodeDNSRecordSet.Name).
		Body(nodeDNSRecordSet).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *nodeDNSRecordSets) UpdateStatus(nodeDNSRecordSet *v1alpha1.NodeDNSRecordSet) (result *v1alpha1.NodeDNSRecordSet, err error) {
	result = &v1alpha1.NodeDNSRecordSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("nodednsrecordsets").
		Name(nodeDNSRecordSet.Name).
		SubResource("status").
		Body(nodeDNSRecordSet).
		Do().
		Into(result)
	return
}

// Delete takes name of the nodeDNSRecordSet and deletes it. Returns an error if one occurs.
func (c *nodeDNSRecordSets) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("nodednsrecordsets").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *nodeDNSRecordSets) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("nodednsrecordsets").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched nodeDNSRecordSet.
func (c *nodeDNSRecordSets) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NodeDNSRecordSet, err error) {
	result = &v1alpha1.NodeDNSRecordSet{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("nodednsrecordsets").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
