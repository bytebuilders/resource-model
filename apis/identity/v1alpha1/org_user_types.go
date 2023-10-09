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
	ResourceKindOrgUser = "OrgUser"
	ResourceOrgUser     = "orguser"
	ResourceOrgUsers    = "orgusers"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// +kubebuilder:object:root=true
// +kubebuilder:resource:path=orgusers,singular=orguser,scope=Cluster,categories={kubernetes,resource-model,appscode}
// +kubebuilder:subresource:status
type OrgUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OrgUserSpec   `json:"spec,omitempty"`
	Status            OrgUserStatus `json:"status,omitempty"`
}

type OrgUserSpec struct {
	UserID int64 `json:"userID"`
	// +optional
	UserName string `json:"userName,omitempty"`
	// +optional
	FullName string `json:"fullName,omitempty"`
	// +optional
	Email string `json:"email,omitempty"`
	// +optional
	AvatarURL string `json:"avatarURL,omitempty"`
	OrgID     int64  `json:"orgID"`
	// +optional
	OrgName  string `json:"orgName,omitempty"`
	IsPublic bool   `json:"isPublic"`
	// +optional
	TeamName string `json:"teamName"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type OrgUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OrgUser `json:"items,omitempty"`
}

type OrgUserStatus struct {
	// ObservedGeneration is the most recent generation observed for this resource. It corresponds to the
	// resource's generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
}
