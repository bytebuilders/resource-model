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

package cluster

const (
	LabelResourceName      = "byte.builders/resource-name"
	LabelClusterUID        = "byte.builders/cluster-uid"
	LabelClusterUserID     = "byte.builders/cluster-user-id"
	LabelClusterOwnerID    = "byte.builders/cluster-owner-id"
	LabelClusterProvider   = "byte.builders/cluster-provider"
	LabelClusterImportType = "byte.builders/cluster-import-type"
	LabelClusterExternalID = "byte.builders/cluster-external-id"

	LabelTricksterReference = "byte.builders/cluster"

	LabelClusterManagerACE                         = "clustermanager.appscode.com/ace"
	LabelClusterManagerOCMHub                      = "clustermanager.appscode.com/ocm-hub"
	LabelClusterManagerOCMSpoke                    = "clustermanager.appscode.com/ocm-spoke"
	LabelClusterManagerOCMMulticlusterControlplane = "clustermanager.appscode.com/ocm-mc"
	LabelClusterManagerRancher                     = "clustermanager.appscode.com/rancher"
	LabelClusterManagerOpenShift                   = "clustermanager.appscode.com/openshift"
	LabelClusterManagerVCluster                    = "clustermanager.appscode.com/vcluster"
)

const (
	ClusterImportTypePublic  = "public"
	ClusterImportTypePrivate = "private"
)
