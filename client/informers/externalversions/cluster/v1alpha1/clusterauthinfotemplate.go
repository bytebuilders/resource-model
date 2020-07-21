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
	"context"
	time "time"

	clusterv1alpha1 "go.bytebuilders.dev/resource-model/apis/cluster/v1alpha1"
	versioned "go.bytebuilders.dev/resource-model/client/clientset/versioned"
	internalinterfaces "go.bytebuilders.dev/resource-model/client/informers/externalversions/internalinterfaces"
	v1alpha1 "go.bytebuilders.dev/resource-model/client/listers/cluster/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ClusterAuthInfoTemplateInformer provides access to a shared informer and lister for
// ClusterAuthInfoTemplates.
type ClusterAuthInfoTemplateInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ClusterAuthInfoTemplateLister
}

type clusterAuthInfoTemplateInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewClusterAuthInfoTemplateInformer constructs a new informer for ClusterAuthInfoTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterAuthInfoTemplateInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredClusterAuthInfoTemplateInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredClusterAuthInfoTemplateInformer constructs a new informer for ClusterAuthInfoTemplate type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredClusterAuthInfoTemplateInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ClusterV1alpha1().ClusterAuthInfoTemplates().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ClusterV1alpha1().ClusterAuthInfoTemplates().Watch(context.TODO(), options)
			},
		},
		&clusterv1alpha1.ClusterAuthInfoTemplate{},
		resyncPeriod,
		indexers,
	)
}

func (f *clusterAuthInfoTemplateInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredClusterAuthInfoTemplateInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *clusterAuthInfoTemplateInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&clusterv1alpha1.ClusterAuthInfoTemplate{}, f.defaultInformer)
}

func (f *clusterAuthInfoTemplateInformer) Lister() v1alpha1.ClusterAuthInfoTemplateLister {
	return v1alpha1.NewClusterAuthInfoTemplateLister(f.Informer().GetIndexer())
}
