// Copyright (c) 2020 Tigera, Inc. All rights reserved.

// Code generated by informer-gen. DO NOT EDIT.

package projectcalico

import (
	"context"
	time "time"

	apisprojectcalico "github.com/tigera/api/pkg/apis/projectcalico"
	clientset "github.com/tigera/api/pkg/client/clientset_generated/clientset"
	internalinterfaces "github.com/tigera/api/pkg/client/informers_generated/externalversions/internalinterfaces"
	projectcalico "github.com/tigera/api/pkg/client/listers_generated/apis/projectcalico"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ManagedClusterInformer provides access to a shared informer and lister for
// ManagedClusters.
type ManagedClusterInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() projectcalico.ManagedClusterLister
}

type managedClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewManagedClusterInformer constructs a new informer for ManagedCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewManagedClusterInformer(client clientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredManagedClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredManagedClusterInformer constructs a new informer for ManagedCluster type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredManagedClusterInformer(client clientset.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ProjectcalicoProjectcalico().ManagedClusters().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ProjectcalicoProjectcalico().ManagedClusters().Watch(context.TODO(), options)
			},
		},
		&apisprojectcalico.ManagedCluster{},
		resyncPeriod,
		indexers,
	)
}

func (f *managedClusterInformer) defaultInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredManagedClusterInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *managedClusterInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apisprojectcalico.ManagedCluster{}, f.defaultInformer)
}

func (f *managedClusterInformer) Lister() projectcalico.ManagedClusterLister {
	return projectcalico.NewManagedClusterLister(f.Informer().GetIndexer())
}
