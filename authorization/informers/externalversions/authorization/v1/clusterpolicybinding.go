// This file was automatically generated by informer-gen

package v1

import (
	authorization_v1 "github.com/openshift/api/authorization/v1"
	versioned "github.com/openshift/client-go/authorization/clientset/versioned"
	internalinterfaces "github.com/openshift/client-go/authorization/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift/client-go/authorization/listers/authorization/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// ClusterPolicyBindingInformer provides access to a shared informer and lister for
// ClusterPolicyBindings.
type ClusterPolicyBindingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.ClusterPolicyBindingLister
}

type clusterPolicyBindingInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

// NewClusterPolicyBindingInformer constructs a new informer for ClusterPolicyBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewClusterPolicyBindingInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				return client.AuthorizationV1().ClusterPolicyBindings().List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				return client.AuthorizationV1().ClusterPolicyBindings().Watch(options)
			},
		},
		&authorization_v1.ClusterPolicyBinding{},
		resyncPeriod,
		indexers,
	)
}

func defaultClusterPolicyBindingInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewClusterPolicyBindingInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *clusterPolicyBindingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&authorization_v1.ClusterPolicyBinding{}, defaultClusterPolicyBindingInformer)
}

func (f *clusterPolicyBindingInformer) Lister() v1.ClusterPolicyBindingLister {
	return v1.NewClusterPolicyBindingLister(f.Informer().GetIndexer())
}
