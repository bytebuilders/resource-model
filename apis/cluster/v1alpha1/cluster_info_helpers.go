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
	"strconv"

	"go.bytebuilders.dev/resource-model/apis/cluster"
	"go.bytebuilders.dev/resource-model/crds"

	"k8s.io/apimachinery/pkg/fields"
	"kmodules.xyz/client-go/apiextensions"
)

type ClusterOptions struct {
	ResourceName         string
	Provider             string
	UserID               int64
	CID                  string
	OwnerID              int64
	ImportType           string
	ConnectorProductName string
}

func (_ ClusterInfo) CustomResourceDefinition() *apiextensions.CustomResourceDefinition {
	return crds.MustCustomResourceDefinition(SchemeGroupVersion.WithResource(ResourceClusterInfos))
}

func (clusterInfo *ClusterInfo) SetLabels(opts ClusterOptions) {
	labelMap := map[string]string{
		cluster.LabelResourceName:      opts.ResourceName,
		cluster.LabelClusterUID:        opts.CID,
		cluster.LabelClusterOwnerID:    strconv.FormatInt(opts.OwnerID, 10),
		cluster.LabelClusterProvider:   opts.Provider,
		cluster.LabelClusterImportType: opts.ImportType,
	}

	if len(opts.ConnectorProductName) > 0 {
		labelMap[cluster.LabelClusterImportType] = cluster.ClusterImportTypePrivate
		labelMap[cluster.LabelClusterConnectorProductName] = opts.ConnectorProductName
	}

	clusterInfo.ObjectMeta.SetLabels(labelMap)
}

func (_ ClusterInfo) FormatLabels(opts ClusterOptions) string {
	labelMap := make(map[string]string)
	if opts.ResourceName != "" {
		labelMap[cluster.LabelResourceName] = opts.ResourceName
	}
	if opts.CID != "" {
		labelMap[cluster.LabelClusterUID] = opts.CID
	}
	if opts.OwnerID != 0 {
		labelMap[cluster.LabelClusterOwnerID] = strconv.FormatInt(opts.OwnerID, 10)
	}
	if opts.Provider != "" {
		labelMap[cluster.LabelClusterProvider] = opts.Provider
	}
	if opts.ImportType != "" {
		labelMap[cluster.LabelClusterImportType] = opts.ImportType
	}
	if opts.ConnectorProductName != "" {
		labelMap[cluster.LabelClusterConnectorProductName] = opts.ConnectorProductName
	}

	return fields.SelectorFromSet(labelMap).String()
}
