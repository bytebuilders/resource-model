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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "go.bytebuilders.dev/resource-model/apis/identity/v1alpha1"
	scheme "go.bytebuilders.dev/resource-model/client/clientset/versioned/scheme"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// OrgUsersGetter has a method to return a OrgUserInterface.
// A group's client should implement this interface.
type OrgUsersGetter interface {
	OrgUsers() OrgUserInterface
}

// OrgUserInterface has methods to work with OrgUser resources.
type OrgUserInterface interface {
	Create(ctx context.Context, orgUser *v1alpha1.OrgUser, opts v1.CreateOptions) (*v1alpha1.OrgUser, error)
	Update(ctx context.Context, orgUser *v1alpha1.OrgUser, opts v1.UpdateOptions) (*v1alpha1.OrgUser, error)
	UpdateStatus(ctx context.Context, orgUser *v1alpha1.OrgUser, opts v1.UpdateOptions) (*v1alpha1.OrgUser, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.OrgUser, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.OrgUserList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.OrgUser, err error)
	OrgUserExpansion
}

// orgUsers implements OrgUserInterface
type orgUsers struct {
	client rest.Interface
}

// newOrgUsers returns a OrgUsers
func newOrgUsers(c *IdentityV1alpha1Client) *orgUsers {
	return &orgUsers{
		client: c.RESTClient(),
	}
}

// Get takes name of the orgUser, and returns the corresponding orgUser object, and an error if there is any.
func (c *orgUsers) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.OrgUser, err error) {
	result = &v1alpha1.OrgUser{}
	err = c.client.Get().
		Resource("orgusers").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of OrgUsers that match those selectors.
func (c *orgUsers) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.OrgUserList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.OrgUserList{}
	err = c.client.Get().
		Resource("orgusers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested orgUsers.
func (c *orgUsers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("orgusers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a orgUser and creates it.  Returns the server's representation of the orgUser, and an error, if there is any.
func (c *orgUsers) Create(ctx context.Context, orgUser *v1alpha1.OrgUser, opts v1.CreateOptions) (result *v1alpha1.OrgUser, err error) {
	result = &v1alpha1.OrgUser{}
	err = c.client.Post().
		Resource("orgusers").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(orgUser).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a orgUser and updates it. Returns the server's representation of the orgUser, and an error, if there is any.
func (c *orgUsers) Update(ctx context.Context, orgUser *v1alpha1.OrgUser, opts v1.UpdateOptions) (result *v1alpha1.OrgUser, err error) {
	result = &v1alpha1.OrgUser{}
	err = c.client.Put().
		Resource("orgusers").
		Name(orgUser.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(orgUser).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *orgUsers) UpdateStatus(ctx context.Context, orgUser *v1alpha1.OrgUser, opts v1.UpdateOptions) (result *v1alpha1.OrgUser, err error) {
	result = &v1alpha1.OrgUser{}
	err = c.client.Put().
		Resource("orgusers").
		Name(orgUser.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(orgUser).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the orgUser and deletes it. Returns an error if one occurs.
func (c *orgUsers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("orgusers").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *orgUsers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("orgusers").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched orgUser.
func (c *orgUsers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.OrgUser, err error) {
	result = &v1alpha1.OrgUser{}
	err = c.client.Patch(pt).
		Resource("orgusers").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}