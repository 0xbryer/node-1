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

// Code generated by informer-gen. DO NOT EDIT.

package v2beta1

import (
	internalinterfaces "github.com/ovrclk/akash/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Inventories returns a InventoryInformer.
	Inventories() InventoryInformer
	// InventoryRequests returns a InventoryRequestInformer.
	InventoryRequests() InventoryRequestInformer
	// Manifests returns a ManifestInformer.
	Manifests() ManifestInformer
	// ProviderHosts returns a ProviderHostInformer.
	ProviderHosts() ProviderHostInformer
	// ProviderLeasedIPs returns a ProviderLeasedIPInformer.
	ProviderLeasedIPs() ProviderLeasedIPInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Inventories returns a InventoryInformer.
func (v *version) Inventories() InventoryInformer {
	return &inventoryInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// InventoryRequests returns a InventoryRequestInformer.
func (v *version) InventoryRequests() InventoryRequestInformer {
	return &inventoryRequestInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// Manifests returns a ManifestInformer.
func (v *version) Manifests() ManifestInformer {
	return &manifestInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ProviderHosts returns a ProviderHostInformer.
func (v *version) ProviderHosts() ProviderHostInformer {
	return &providerHostInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// ProviderLeasedIPs returns a ProviderLeasedIPInformer.
func (v *version) ProviderLeasedIPs() ProviderLeasedIPInformer {
	return &providerLeasedIPInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
