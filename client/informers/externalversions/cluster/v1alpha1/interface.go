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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "go.bytebuilders.dev/resource-model/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Clusters returns a ClusterInformer.
	Clusters() ClusterInformer
	// ClusterAuthInfoTemplates returns a ClusterAuthInfoTemplateInformer.
	ClusterAuthInfoTemplates() ClusterAuthInfoTemplateInformer
	// ClusterCredentials returns a ClusterCredentialInformer.
	ClusterCredentials() ClusterCredentialInformer
	// ClusterInfos returns a ClusterInfoInformer.
	ClusterInfos() ClusterInfoInformer
	// ClusterUserAuths returns a ClusterUserAuthInformer.
	ClusterUserAuths() ClusterUserAuthInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Clusters returns a ClusterInformer.
func (v *version) Clusters() ClusterInformer {
	return &clusterInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ClusterAuthInfoTemplates returns a ClusterAuthInfoTemplateInformer.
func (v *version) ClusterAuthInfoTemplates() ClusterAuthInfoTemplateInformer {
	return &clusterAuthInfoTemplateInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ClusterCredentials returns a ClusterCredentialInformer.
func (v *version) ClusterCredentials() ClusterCredentialInformer {
	return &clusterCredentialInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ClusterInfos returns a ClusterInfoInformer.
func (v *version) ClusterInfos() ClusterInfoInformer {
	return &clusterInfoInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ClusterUserAuths returns a ClusterUserAuthInformer.
func (v *version) ClusterUserAuths() ClusterUserAuthInformer {
	return &clusterUserAuthInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
