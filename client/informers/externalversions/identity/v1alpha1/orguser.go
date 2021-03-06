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

	identityv1alpha1 "go.bytebuilders.dev/resource-model/apis/identity/v1alpha1"
	versioned "go.bytebuilders.dev/resource-model/client/clientset/versioned"
	internalinterfaces "go.bytebuilders.dev/resource-model/client/informers/externalversions/internalinterfaces"
	v1alpha1 "go.bytebuilders.dev/resource-model/client/listers/identity/v1alpha1"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// OrgUserInformer provides access to a shared informer and lister for
// OrgUsers.
type OrgUserInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.OrgUserLister
}

type orgUserInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewOrgUserInformer constructs a new informer for OrgUser type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewOrgUserInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredOrgUserInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredOrgUserInformer constructs a new informer for OrgUser type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredOrgUserInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.IdentityV1alpha1().OrgUsers().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.IdentityV1alpha1().OrgUsers().Watch(context.TODO(), options)
			},
		},
		&identityv1alpha1.OrgUser{},
		resyncPeriod,
		indexers,
	)
}

func (f *orgUserInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredOrgUserInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *orgUserInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&identityv1alpha1.OrgUser{}, f.defaultInformer)
}

func (f *orgUserInformer) Lister() v1alpha1.OrgUserLister {
	return v1alpha1.NewOrgUserLister(f.Informer().GetIndexer())
}
