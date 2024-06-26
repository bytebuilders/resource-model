/*
Copyright AppsCode Inc. and Contributors

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
	cloudv1alpha1 "go.bytebuilders.dev/resource-model/apis/cloud/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AceSetupConfig is the Schema for the kubestashconfigs API

// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type AceSetupConfig struct {
	metav1.TypeMeta `json:",inline"`

	DeploymentType         string     `json:"deploymentType,omitempty"`
	Nats                   NatsConfig `json:"nats"`
	ImporterServiceAccount string     `json:"importerServiceAccount,omitempty"`
	AceSetupInlineConfig   `json:",inline"`
}

// NatsConfig holds the NATS-related fields.
type NatsConfig struct {
	Exports              bool `json:"exports"`
	ReloadNatsAccounts   bool `json:"reloadNatsAccounts"`
	CreateNatsStream     bool `json:"createNatsStream,omitempty"`
	RefactorNatsAccounts bool `json:"refactorNatsAccounts,omitempty"`
	Migrate              bool `json:"migrate,omitempty"`
}

type AceSetupInlineConfig struct {
	// +optional
	Admin AcePlatformAdmin `json:"admin"`
	// +optional
	SelfManagement SelfManagement `json:"selfManagement"`
	// +optional
	InitialOrgName string `json:"initialOrgName"`
	// +optional
	MarketplaceInfo MarketplaceConfig `json:"marketplaceConfig"`
}

type AcePlatformAdmin struct {
	// +optional
	Username string `json:"username"`
	// +optional
	Password string `json:"password"`
	// +optional
	Email string `json:"email"`
	// +optional
	DisplayName string `json:"displayName"`
}

type SelfManagement struct {
	// +optional
	Import bool `json:"import"`
	// +optional
	EnableFeatures []string `json:"enableFeatures"`
	// +optional
	DisableFeatures []string `json:"disableFeatures"`
}

type MarketplaceConfig struct {
	AWS                  cloudv1alpha1.AWSCredential         `json:"aws,omitempty"`
	Azure                cloudv1alpha1.AzureCredential       `json:"azure,omitempty"`
	GoogleCloud          cloudv1alpha1.GoogleCloudCredential `json:"googleCloud,omitempty"`
	ClusterCreateOptions CAPIClusterCreateOptions            `json:"clusterCreateOptions,omitempty"`
}

type CAPIClusterCreateOptions struct {
	ClusterName       string `json:"clusterName,omitempty"`
	Region            string `json:"region,omitempty"`
	ProjectID         string `json:"projectID,omitempty"`
	CIDRRange         string `json:"CIDRRange,omitempty"`
	KubernetesVersion string `json:"kubernetesVersion,omitempty"`
	MachineType       string `json:"machineType,omitempty"`
}
