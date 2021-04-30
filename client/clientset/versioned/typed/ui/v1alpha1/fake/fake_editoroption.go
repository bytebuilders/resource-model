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

	v1alpha1 "go.bytebuilders.dev/resource-model/apis/ui/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEditorOptions implements EditorOptionInterface
type FakeEditorOptions struct {
	Fake *FakeUiV1alpha1
}

var editoroptionsResource = schema.GroupVersionResource{Group: "ui.bytebuilders.dev", Version: "v1alpha1", Resource: "editoroptions"}

var editoroptionsKind = schema.GroupVersionKind{Group: "ui.bytebuilders.dev", Version: "v1alpha1", Kind: "EditorOption"}

// Get takes name of the editorOption, and returns the corresponding editorOption object, and an error if there is any.
func (c *FakeEditorOptions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.EditorOption, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(editoroptionsResource, name), &v1alpha1.EditorOption{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EditorOption), err
}

// List takes label and field selectors, and returns the list of EditorOptions that match those selectors.
func (c *FakeEditorOptions) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.EditorOptionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(editoroptionsResource, editoroptionsKind, opts), &v1alpha1.EditorOptionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EditorOptionList{ListMeta: obj.(*v1alpha1.EditorOptionList).ListMeta}
	for _, item := range obj.(*v1alpha1.EditorOptionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested editorOptions.
func (c *FakeEditorOptions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(editoroptionsResource, opts))
}

// Create takes the representation of a editorOption and creates it.  Returns the server's representation of the editorOption, and an error, if there is any.
func (c *FakeEditorOptions) Create(ctx context.Context, editorOption *v1alpha1.EditorOption, opts v1.CreateOptions) (result *v1alpha1.EditorOption, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(editoroptionsResource, editorOption), &v1alpha1.EditorOption{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EditorOption), err
}

// Update takes the representation of a editorOption and updates it. Returns the server's representation of the editorOption, and an error, if there is any.
func (c *FakeEditorOptions) Update(ctx context.Context, editorOption *v1alpha1.EditorOption, opts v1.UpdateOptions) (result *v1alpha1.EditorOption, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(editoroptionsResource, editorOption), &v1alpha1.EditorOption{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EditorOption), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEditorOptions) UpdateStatus(ctx context.Context, editorOption *v1alpha1.EditorOption, opts v1.UpdateOptions) (*v1alpha1.EditorOption, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(editoroptionsResource, "status", editorOption), &v1alpha1.EditorOption{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EditorOption), err
}

// Delete takes name of the editorOption and deletes it. Returns an error if one occurs.
func (c *FakeEditorOptions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(editoroptionsResource, name), &v1alpha1.EditorOption{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEditorOptions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(editoroptionsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.EditorOptionList{})
	return err
}

// Patch applies the patch and returns the patched editorOption.
func (c *FakeEditorOptions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.EditorOption, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(editoroptionsResource, name, pt, data, subresources...), &v1alpha1.EditorOption{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.EditorOption), err
}
