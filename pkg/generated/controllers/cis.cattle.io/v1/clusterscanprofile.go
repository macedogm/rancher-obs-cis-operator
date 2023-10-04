/*
Copyright 2023 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	v1 "github.com/rancher/cis-operator/pkg/apis/cis.cattle.io/v1"
	"github.com/rancher/lasso/pkg/client"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/generic"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

type ClusterScanProfileHandler func(string, *v1.ClusterScanProfile) (*v1.ClusterScanProfile, error)

type ClusterScanProfileController interface {
	generic.ControllerMeta
	ClusterScanProfileClient

	OnChange(ctx context.Context, name string, sync ClusterScanProfileHandler)
	OnRemove(ctx context.Context, name string, sync ClusterScanProfileHandler)
	Enqueue(name string)
	EnqueueAfter(name string, duration time.Duration)

	Cache() ClusterScanProfileCache
}

type ClusterScanProfileClient interface {
	Create(*v1.ClusterScanProfile) (*v1.ClusterScanProfile, error)
	Update(*v1.ClusterScanProfile) (*v1.ClusterScanProfile, error)

	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*v1.ClusterScanProfile, error)
	List(opts metav1.ListOptions) (*v1.ClusterScanProfileList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ClusterScanProfile, err error)
}

type ClusterScanProfileCache interface {
	Get(name string) (*v1.ClusterScanProfile, error)
	List(selector labels.Selector) ([]*v1.ClusterScanProfile, error)

	AddIndexer(indexName string, indexer ClusterScanProfileIndexer)
	GetByIndex(indexName, key string) ([]*v1.ClusterScanProfile, error)
}

type ClusterScanProfileIndexer func(obj *v1.ClusterScanProfile) ([]string, error)

type clusterScanProfileController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewClusterScanProfileController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) ClusterScanProfileController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &clusterScanProfileController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromClusterScanProfileHandlerToHandler(sync ClusterScanProfileHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1.ClusterScanProfile
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1.ClusterScanProfile))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *clusterScanProfileController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1.ClusterScanProfile))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateClusterScanProfileDeepCopyOnChange(client ClusterScanProfileClient, obj *v1.ClusterScanProfile, handler func(obj *v1.ClusterScanProfile) (*v1.ClusterScanProfile, error)) (*v1.ClusterScanProfile, error) {
	if obj == nil {
		return obj, nil
	}

	copyObj := obj.DeepCopy()
	newObj, err := handler(copyObj)
	if newObj != nil {
		copyObj = newObj
	}
	if obj.ResourceVersion == copyObj.ResourceVersion && !equality.Semantic.DeepEqual(obj, copyObj) {
		return client.Update(copyObj)
	}

	return copyObj, err
}

func (c *clusterScanProfileController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *clusterScanProfileController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *clusterScanProfileController) OnChange(ctx context.Context, name string, sync ClusterScanProfileHandler) {
	c.AddGenericHandler(ctx, name, FromClusterScanProfileHandlerToHandler(sync))
}

func (c *clusterScanProfileController) OnRemove(ctx context.Context, name string, sync ClusterScanProfileHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromClusterScanProfileHandlerToHandler(sync)))
}

func (c *clusterScanProfileController) Enqueue(name string) {
	c.controller.Enqueue("", name)
}

func (c *clusterScanProfileController) EnqueueAfter(name string, duration time.Duration) {
	c.controller.EnqueueAfter("", name, duration)
}

func (c *clusterScanProfileController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *clusterScanProfileController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *clusterScanProfileController) Cache() ClusterScanProfileCache {
	return &clusterScanProfileCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *clusterScanProfileController) Create(obj *v1.ClusterScanProfile) (*v1.ClusterScanProfile, error) {
	result := &v1.ClusterScanProfile{}
	return result, c.client.Create(context.TODO(), "", obj, result, metav1.CreateOptions{})
}

func (c *clusterScanProfileController) Update(obj *v1.ClusterScanProfile) (*v1.ClusterScanProfile, error) {
	result := &v1.ClusterScanProfile{}
	return result, c.client.Update(context.TODO(), "", obj, result, metav1.UpdateOptions{})
}

func (c *clusterScanProfileController) Delete(name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), "", name, *options)
}

func (c *clusterScanProfileController) Get(name string, options metav1.GetOptions) (*v1.ClusterScanProfile, error) {
	result := &v1.ClusterScanProfile{}
	return result, c.client.Get(context.TODO(), "", name, result, options)
}

func (c *clusterScanProfileController) List(opts metav1.ListOptions) (*v1.ClusterScanProfileList, error) {
	result := &v1.ClusterScanProfileList{}
	return result, c.client.List(context.TODO(), "", result, opts)
}

func (c *clusterScanProfileController) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), "", opts)
}

func (c *clusterScanProfileController) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (*v1.ClusterScanProfile, error) {
	result := &v1.ClusterScanProfile{}
	return result, c.client.Patch(context.TODO(), "", name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type clusterScanProfileCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *clusterScanProfileCache) Get(name string) (*v1.ClusterScanProfile, error) {
	obj, exists, err := c.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v1.ClusterScanProfile), nil
}

func (c *clusterScanProfileCache) List(selector labels.Selector) (ret []*v1.ClusterScanProfile, err error) {

	err = cache.ListAll(c.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ClusterScanProfile))
	})

	return ret, err
}

func (c *clusterScanProfileCache) AddIndexer(indexName string, indexer ClusterScanProfileIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1.ClusterScanProfile))
		},
	}))
}

func (c *clusterScanProfileCache) GetByIndex(indexName, key string) (result []*v1.ClusterScanProfile, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v1.ClusterScanProfile, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v1.ClusterScanProfile))
	}
	return result, nil
}
