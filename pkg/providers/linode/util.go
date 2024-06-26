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

package linode

import (
	"fmt"

	"go.bytebuilders.dev/resource-model/apis/cloud"
	"go.bytebuilders.dev/resource-model/apis/cloud/v1alpha1"
	"go.bytebuilders.dev/resource-model/pkg/util"

	"github.com/linode/linodego"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ParseRegion(in *linodego.Region) *v1alpha1.Region {
	return &v1alpha1.Region{
		Location: in.Country,
		Region:   in.ID,
		Zones: []string{
			in.ID,
		},
	}
}

func ParseInstance(in *linodego.LinodeType) (*v1alpha1.MachineType, error) {
	return &v1alpha1.MachineType{
		ObjectMeta: metav1.ObjectMeta{
			Name: util.Sanitize(cloud.Linode + "-" + in.ID),
			Labels: map[string]string{
				cloud.KeyCloudProvider: cloud.Linode,
			},
		},
		Spec: v1alpha1.MachineTypeSpec{
			SKU:         in.ID,
			Description: in.Label,
			CPU:         resource.NewQuantity(int64(in.VCPUs), resource.DecimalExponent),
			RAM:         util.QuantityP(resource.MustParse(fmt.Sprintf("%dMi", in.Memory))),
			Disk:        util.QuantityP(resource.MustParse(fmt.Sprintf("%dMi", in.Disk))),
		},
	}, nil
}
