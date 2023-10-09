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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	ResourceKindClusterInfo = "ClusterInfo"
	ResourceClusterInfo     = "clusterinfo"
	ResourceClusterInfos    = "clusterinfos"
)

// +kubebuilder:validation:Enum=AKS;DigitalOcean;EKS;GKE;Linode;Packet;Scaleway;Vultr;Rancher;Generic
type ProviderName string

const (
	ProviderAKS          ProviderName = "AKS"
	ProviderDigitalOcean ProviderName = "DigitalOcean"
	ProviderEKS          ProviderName = "EKS"
	ProviderGKE          ProviderName = "GKE"
	ProviderLinode       ProviderName = "Linode"
	ProviderPacket       ProviderName = "Packet"
	ProviderScaleway     ProviderName = "Scaleway"
	ProviderVultr        ProviderName = "Vultr"
	ProviderRancher      ProviderName = "Rancher"
	ProviderGeneric      ProviderName = "Generic"
	ProviderPrivate      ProviderName = "Private"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=clusterinfos,singular=clusterinfo,shortName=cinfo,scope=Cluster,categories={kubernetes,resource-model,appscode}
// +kubebuilder:subresource:status
type ClusterInfo struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterInfoSpec `json:"spec,omitempty"`
}

type ClusterInfoSpec struct {
	DisplayName string `json:"displayName"`
	Name        string `json:"name"`
	UID         string `json:"uid"`
	OwnerID     int64  `json:"ownerID"`
	//+optional
	ExternalID string `json:"externalID,omitempty"`
	//+optional
	OwnerName string `json:"ownerName,omitempty"`

	//+optional
	Provider ProviderName `json:"provider,omitempty"`
	//+optional
	Endpoint string `json:"endpoint,omitempty"`
	//+optional
	Location string `json:"location,omitempty"`
	//+optional
	Project string `json:"project,omitempty"`
	//+optional
	KubernetesVersion string `json:"kubernetesVersion"`
	//+optional
	NodeCount int32 `json:"nodeCount"`

	//+optional
	CreatedAt int64 `json:"createdAt,omitempty"`

	//+optional
	Age string `json:"age,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ClusterInfoList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterInfo `json:"items,omitempty"`
}
