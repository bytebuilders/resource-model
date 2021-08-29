//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChartRepoRef) DeepCopyInto(out *ChartRepoRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartRepoRef.
func (in *ChartRepoRef) DeepCopy() *ChartRepoRef {
	if in == nil {
		return nil
	}
	out := new(ChartRepoRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OptionsEditor) DeepCopyInto(out *OptionsEditor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OptionsEditor.
func (in *OptionsEditor) DeepCopy() *OptionsEditor {
	if in == nil {
		return nil
	}
	out := new(OptionsEditor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OptionsEditor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OptionsEditorList) DeepCopyInto(out *OptionsEditorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OptionsEditor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OptionsEditorList.
func (in *OptionsEditorList) DeepCopy() *OptionsEditorList {
	if in == nil {
		return nil
	}
	out := new(OptionsEditorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OptionsEditorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OptionsEditorSpec) DeepCopyInto(out *OptionsEditorSpec) {
	*out = *in
	out.Resource = in.Resource
	out.ChartRef = in.ChartRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OptionsEditorSpec.
func (in *OptionsEditorSpec) DeepCopy() *OptionsEditorSpec {
	if in == nil {
		return nil
	}
	out := new(OptionsEditorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OptionsEditorStatus) DeepCopyInto(out *OptionsEditorStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OptionsEditorStatus.
func (in *OptionsEditorStatus) DeepCopy() *OptionsEditorStatus {
	if in == nil {
		return nil
	}
	out := new(OptionsEditorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceID) DeepCopyInto(out *ResourceID) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceID.
func (in *ResourceID) DeepCopy() *ResourceID {
	if in == nil {
		return nil
	}
	out := new(ResourceID)
	in.DeepCopyInto(out)
	return out
}
