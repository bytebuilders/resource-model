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

package credential

import (
	"go.bytebuilders.dev/resource-model/apis/cloud"
	"go.bytebuilders.dev/resource-model/apis/cloud/v1alpha1"

	"github.com/spf13/pflag"
)

type AWS struct {
	*v1alpha1.AWSCredential
	Region string
}

func (c *AWS) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&c.Region, cloud.AWS+"."+AWSRegion, c.Region, "provide this flag when provider is aws")
	fs.StringVar(&c.AccessKeyID, cloud.AWS+"."+AWSAccessKeyID, c.AccessKeyID, "provide this flag when provider is aws")
	fs.StringVar(&c.SecretAccessKey, cloud.AWS+"."+AWSSecretAccessKey, c.SecretAccessKey, "provide this flag when provider is aws")
}

func (c AWS) RequiredFlags() []string {
	return []string{
		cloud.AWS + "." + AWSRegion,
		cloud.AWS + "." + AWSAccessKeyID,
		cloud.AWS + "." + AWSSecretAccessKey,
	}
}
