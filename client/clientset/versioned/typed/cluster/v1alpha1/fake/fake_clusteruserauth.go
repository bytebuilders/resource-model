/*
Copyright 2020 AppsCode Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "go.bytebuilders.dev/resource-model/apis/cluster/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterUserAuths implements ClusterUserAuthInterface
type FakeClusterUserAuths struct {
	Fake *FakeClusterV1alpha1
}

var clusteruserauthsResource = schema.GroupVersionResource{Group: "cluster.bytebuilders.dev", Version: "v1alpha1", Resource: "clusteruserauths"}

var clusteruserauthsKind = schema.GroupVersionKind{Group: "cluster.bytebuilders.dev", Version: "v1alpha1", Kind: "ClusterUserAuth"}

// Get takes name of the clusterUserAuth, and returns the corresponding clusterUserAuth object, and an error if there is any.
func (c *FakeClusterUserAuths) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ClusterUserAuth, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusteruserauthsResource, name), &v1alpha1.ClusterUserAuth{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterUserAuth), err
}

// List takes label and field selectors, and returns the list of ClusterUserAuths that match those selectors.
func (c *FakeClusterUserAuths) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ClusterUserAuthList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusteruserauthsResource, clusteruserauthsKind, opts), &v1alpha1.ClusterUserAuthList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterUserAuthList{ListMeta: obj.(*v1alpha1.ClusterUserAuthList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterUserAuthList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterUserAuths.
func (c *FakeClusterUserAuths) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusteruserauthsResource, opts))
}

// Create takes the representation of a clusterUserAuth and creates it.  Returns the server's representation of the clusterUserAuth, and an error, if there is any.
func (c *FakeClusterUserAuths) Create(ctx context.Context, clusterUserAuth *v1alpha1.ClusterUserAuth, opts v1.CreateOptions) (result *v1alpha1.ClusterUserAuth, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusteruserauthsResource, clusterUserAuth), &v1alpha1.ClusterUserAuth{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterUserAuth), err
}

// Update takes the representation of a clusterUserAuth and updates it. Returns the server's representation of the clusterUserAuth, and an error, if there is any.
func (c *FakeClusterUserAuths) Update(ctx context.Context, clusterUserAuth *v1alpha1.ClusterUserAuth, opts v1.UpdateOptions) (result *v1alpha1.ClusterUserAuth, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusteruserauthsResource, clusterUserAuth), &v1alpha1.ClusterUserAuth{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterUserAuth), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterUserAuths) UpdateStatus(ctx context.Context, clusterUserAuth *v1alpha1.ClusterUserAuth, opts v1.UpdateOptions) (*v1alpha1.ClusterUserAuth, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clusteruserauthsResource, "status", clusterUserAuth), &v1alpha1.ClusterUserAuth{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterUserAuth), err
}

// Delete takes name of the clusterUserAuth and deletes it. Returns an error if one occurs.
func (c *FakeClusterUserAuths) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(clusteruserauthsResource, name), &v1alpha1.ClusterUserAuth{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterUserAuths) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusteruserauthsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterUserAuthList{})
	return err
}

// Patch applies the patch and returns the patched clusterUserAuth.
func (c *FakeClusterUserAuths) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ClusterUserAuth, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusteruserauthsResource, name, pt, data, subresources...), &v1alpha1.ClusterUserAuth{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterUserAuth), err
}
