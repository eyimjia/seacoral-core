/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/upmio/dbscale-kube/pkg/apis/san/v1alpha1"
	scheme "github.com/upmio/dbscale-kube/pkg/client/san/v1alpha1/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// SanSystemsGetter has a method to return a SanSystemInterface.
// A group's client should implement this interface.
type SanSystemsGetter interface {
	SanSystems() SanSystemInterface
}

// SanSystemInterface has methods to work with SanSystem resources.
type SanSystemInterface interface {
	Create(ctx context.Context, sanSystem *v1alpha1.SanSystem, opts v1.CreateOptions) (*v1alpha1.SanSystem, error)
	Update(ctx context.Context, sanSystem *v1alpha1.SanSystem, opts v1.UpdateOptions) (*v1alpha1.SanSystem, error)
	UpdateStatus(ctx context.Context, sanSystem *v1alpha1.SanSystem, opts v1.UpdateOptions) (*v1alpha1.SanSystem, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.SanSystem, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.SanSystemList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SanSystem, err error)
	SanSystemExpansion
}

// sanSystems implements SanSystemInterface
type sanSystems struct {
	client rest.Interface
}

// newSanSystems returns a SanSystems
func newSanSystems(c *SanV1alpha1Client) *sanSystems {
	return &sanSystems{
		client: c.RESTClient(),
	}
}

// Get takes name of the sanSystem, and returns the corresponding sanSystem object, and an error if there is any.
func (c *sanSystems) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SanSystem, err error) {
	result = &v1alpha1.SanSystem{}
	err = c.client.Get().
		Resource("sansystems").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SanSystems that match those selectors.
func (c *sanSystems) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SanSystemList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SanSystemList{}
	err = c.client.Get().
		Resource("sansystems").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested sanSystems.
func (c *sanSystems) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("sansystems").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a sanSystem and creates it.  Returns the server's representation of the sanSystem, and an error, if there is any.
func (c *sanSystems) Create(ctx context.Context, sanSystem *v1alpha1.SanSystem, opts v1.CreateOptions) (result *v1alpha1.SanSystem, err error) {
	result = &v1alpha1.SanSystem{}
	err = c.client.Post().
		Resource("sansystems").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sanSystem).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a sanSystem and updates it. Returns the server's representation of the sanSystem, and an error, if there is any.
func (c *sanSystems) Update(ctx context.Context, sanSystem *v1alpha1.SanSystem, opts v1.UpdateOptions) (result *v1alpha1.SanSystem, err error) {
	result = &v1alpha1.SanSystem{}
	err = c.client.Put().
		Resource("sansystems").
		Name(sanSystem.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sanSystem).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *sanSystems) UpdateStatus(ctx context.Context, sanSystem *v1alpha1.SanSystem, opts v1.UpdateOptions) (result *v1alpha1.SanSystem, err error) {
	result = &v1alpha1.SanSystem{}
	err = c.client.Put().
		Resource("sansystems").
		Name(sanSystem.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(sanSystem).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the sanSystem and deletes it. Returns an error if one occurs.
func (c *sanSystems) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("sansystems").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *sanSystems) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("sansystems").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched sanSystem.
func (c *sanSystems) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SanSystem, err error) {
	result = &v1alpha1.SanSystem{}
	err = c.client.Patch(pt).
		Resource("sansystems").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}