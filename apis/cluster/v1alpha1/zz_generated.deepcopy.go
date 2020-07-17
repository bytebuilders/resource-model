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
func (in *AWSCredential) DeepCopyInto(out *AWSCredential) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSCredential.
func (in *AWSCredential) DeepCopy() *AWSCredential {
	if in == nil {
		return nil
	}
	out := new(AWSCredential)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSProvider) DeepCopyInto(out *AWSProvider) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSProvider.
func (in *AWSProvider) DeepCopy() *AWSProvider {
	if in == nil {
		return nil
	}
	out := new(AWSProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthProviderConfig) DeepCopyInto(out *AuthProviderConfig) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthProviderConfig.
func (in *AuthProviderConfig) DeepCopy() *AuthProviderConfig {
	if in == nil {
		return nil
	}
	out := new(AuthProviderConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureCredential) DeepCopyInto(out *AzureCredential) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureCredential.
func (in *AzureCredential) DeepCopy() *AzureCredential {
	if in == nil {
		return nil
	}
	out := new(AzureCredential)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudCredential) DeepCopyInto(out *CloudCredential) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudCredential.
func (in *CloudCredential) DeepCopy() *CloudCredential {
	if in == nil {
		return nil
	}
	out := new(CloudCredential)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudCredential) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudCredentialList) DeepCopyInto(out *CloudCredentialList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CloudCredential, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudCredentialList.
func (in *CloudCredentialList) DeepCopy() *CloudCredentialList {
	if in == nil {
		return nil
	}
	out := new(CloudCredentialList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CloudCredentialList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudCredentialSpec) DeepCopyInto(out *CloudCredentialSpec) {
	*out = *in
	if in.GCE != nil {
		in, out := &in.GCE, &out.GCE
		*out = new(GCECredential)
		**out = **in
	}
	if in.DigitalOcean != nil {
		in, out := &in.DigitalOcean, &out.DigitalOcean
		*out = new(DigitalOceanCredential)
		**out = **in
	}
	if in.Azure != nil {
		in, out := &in.Azure, &out.Azure
		*out = new(AzureCredential)
		**out = **in
	}
	if in.AWS != nil {
		in, out := &in.AWS, &out.AWS
		*out = new(AWSCredential)
		**out = **in
	}
	if in.Packet != nil {
		in, out := &in.Packet, &out.Packet
		*out = new(PacketCredential)
		**out = **in
	}
	if in.Scaleway != nil {
		in, out := &in.Scaleway, &out.Scaleway
		*out = new(ScalewayCredential)
		**out = **in
	}
	if in.Linode != nil {
		in, out := &in.Linode, &out.Linode
		*out = new(LinodeCredential)
		**out = **in
	}
	if in.Vultr != nil {
		in, out := &in.Vultr, &out.Vultr
		*out = new(VultrCredential)
		**out = **in
	}
	if in.GoogleOAuth != nil {
		in, out := &in.GoogleOAuth, &out.GoogleOAuth
		*out = new(GoogleOAuthCredential)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudCredentialSpec.
func (in *CloudCredentialSpec) DeepCopy() *CloudCredentialSpec {
	if in == nil {
		return nil
	}
	out := new(CloudCredentialSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CloudCredentialStatus) DeepCopyInto(out *CloudCredentialStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CloudCredentialStatus.
func (in *CloudCredentialStatus) DeepCopy() *CloudCredentialStatus {
	if in == nil {
		return nil
	}
	out := new(CloudCredentialStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAuthInfoTemplate) DeepCopyInto(out *ClusterAuthInfoTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAuthInfoTemplate.
func (in *ClusterAuthInfoTemplate) DeepCopy() *ClusterAuthInfoTemplate {
	if in == nil {
		return nil
	}
	out := new(ClusterAuthInfoTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterAuthInfoTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAuthInfoTemplateList) DeepCopyInto(out *ClusterAuthInfoTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterAuthInfoTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAuthInfoTemplateList.
func (in *ClusterAuthInfoTemplateList) DeepCopy() *ClusterAuthInfoTemplateList {
	if in == nil {
		return nil
	}
	out := new(ClusterAuthInfoTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterAuthInfoTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAuthInfoTemplateSpec) DeepCopyInto(out *ClusterAuthInfoTemplateSpec) {
	*out = *in
	if in.CertificateAuthorityData != nil {
		in, out := &in.CertificateAuthorityData, &out.CertificateAuthorityData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.ClientCertificateData != nil {
		in, out := &in.ClientCertificateData, &out.ClientCertificateData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.ClientKeyData != nil {
		in, out := &in.ClientKeyData, &out.ClientKeyData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.ImpersonateGroups != nil {
		in, out := &in.ImpersonateGroups, &out.ImpersonateGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ImpersonateUserExtra != nil {
		in, out := &in.ImpersonateUserExtra, &out.ImpersonateUserExtra
		*out = make(map[string]ExtraValue, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(ExtraValue, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.AuthProvider != nil {
		in, out := &in.AuthProvider, &out.AuthProvider
		*out = new(AuthProviderConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAuthInfoTemplateSpec.
func (in *ClusterAuthInfoTemplateSpec) DeepCopy() *ClusterAuthInfoTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterAuthInfoTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAuthInfoTemplateStatus) DeepCopyInto(out *ClusterAuthInfoTemplateStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAuthInfoTemplateStatus.
func (in *ClusterAuthInfoTemplateStatus) DeepCopy() *ClusterAuthInfoTemplateStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterAuthInfoTemplateStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterInfo) DeepCopyInto(out *ClusterInfo) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterInfo.
func (in *ClusterInfo) DeepCopy() *ClusterInfo {
	if in == nil {
		return nil
	}
	out := new(ClusterInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterInfo) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterInfoList) DeepCopyInto(out *ClusterInfoList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterInfoList.
func (in *ClusterInfoList) DeepCopy() *ClusterInfoList {
	if in == nil {
		return nil
	}
	out := new(ClusterInfoList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterInfoList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterInfoSpec) DeepCopyInto(out *ClusterInfoSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterInfoSpec.
func (in *ClusterInfoSpec) DeepCopy() *ClusterInfoSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterInfoSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterInfoStatus) DeepCopyInto(out *ClusterInfoStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterInfoStatus.
func (in *ClusterInfoStatus) DeepCopy() *ClusterInfoStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterInfoStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterUserAuth) DeepCopyInto(out *ClusterUserAuth) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterUserAuth.
func (in *ClusterUserAuth) DeepCopy() *ClusterUserAuth {
	if in == nil {
		return nil
	}
	out := new(ClusterUserAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterUserAuth) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterUserAuthList) DeepCopyInto(out *ClusterUserAuthList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterUserAuth, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterUserAuthList.
func (in *ClusterUserAuthList) DeepCopy() *ClusterUserAuthList {
	if in == nil {
		return nil
	}
	out := new(ClusterUserAuthList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterUserAuthList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterUserAuthSpec) DeepCopyInto(out *ClusterUserAuthSpec) {
	*out = *in
	if in.ClientCertificateData != nil {
		in, out := &in.ClientCertificateData, &out.ClientCertificateData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.ClientKeyData != nil {
		in, out := &in.ClientKeyData, &out.ClientKeyData
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.ImpersonateGroups != nil {
		in, out := &in.ImpersonateGroups, &out.ImpersonateGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ImpersonateUserExtra != nil {
		in, out := &in.ImpersonateUserExtra, &out.ImpersonateUserExtra
		*out = make(map[string]ExtraValue, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(ExtraValue, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.AuthProvider != nil {
		in, out := &in.AuthProvider, &out.AuthProvider
		*out = new(AuthProviderConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.GoogleOAuth != nil {
		in, out := &in.GoogleOAuth, &out.GoogleOAuth
		*out = new(GoogleOAuthProvider)
		**out = **in
	}
	if in.AWS != nil {
		in, out := &in.AWS, &out.AWS
		*out = new(AWSProvider)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterUserAuthSpec.
func (in *ClusterUserAuthSpec) DeepCopy() *ClusterUserAuthSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterUserAuthSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterUserAuthStatus) DeepCopyInto(out *ClusterUserAuthStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterUserAuthStatus.
func (in *ClusterUserAuthStatus) DeepCopy() *ClusterUserAuthStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterUserAuthStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DigitalOceanCredential) DeepCopyInto(out *DigitalOceanCredential) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DigitalOceanCredential.
func (in *DigitalOceanCredential) DeepCopy() *DigitalOceanCredential {
	if in == nil {
		return nil
	}
	out := new(DigitalOceanCredential)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ExtraValue) DeepCopyInto(out *ExtraValue) {
	{
		in := &in
		*out = make(ExtraValue, len(*in))
		copy(*out, *in)
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExtraValue.
func (in ExtraValue) DeepCopy() ExtraValue {
	if in == nil {
		return nil
	}
	out := new(ExtraValue)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCECredential) DeepCopyInto(out *GCECredential) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCECredential.
func (in *GCECredential) DeepCopy() *GCECredential {
	if in == nil {
		return nil
	}
	out := new(GCECredential)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GoogleOAuthCredential) DeepCopyInto(out *GoogleOAuthCredential) {
	*out = *in
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GoogleOAuthCredential.
func (in *GoogleOAuthCredential) DeepCopy() *GoogleOAuthCredential {
	if in == nil {
		return nil
	}
	out := new(GoogleOAuthCredential)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GoogleOAuthProvider) DeepCopyInto(out *GoogleOAuthProvider) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GoogleOAuthProvider.
func (in *GoogleOAuthProvider) DeepCopy() *GoogleOAuthProvider {
	if in == nil {
		return nil
	}
	out := new(GoogleOAuthProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinodeCredential) DeepCopyInto(out *LinodeCredential) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinodeCredential.
func (in *LinodeCredential) DeepCopy() *LinodeCredential {
	if in == nil {
		return nil
	}
	out := new(LinodeCredential)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PacketCredential) DeepCopyInto(out *PacketCredential) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PacketCredential.
func (in *PacketCredential) DeepCopy() *PacketCredential {
	if in == nil {
		return nil
	}
	out := new(PacketCredential)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScalewayCredential) DeepCopyInto(out *ScalewayCredential) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScalewayCredential.
func (in *ScalewayCredential) DeepCopy() *ScalewayCredential {
	if in == nil {
		return nil
	}
	out := new(ScalewayCredential)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VultrCredential) DeepCopyInto(out *VultrCredential) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VultrCredential.
func (in *VultrCredential) DeepCopy() *VultrCredential {
	if in == nil {
		return nil
	}
	out := new(VultrCredential)
	in.DeepCopyInto(out)
	return out
}
