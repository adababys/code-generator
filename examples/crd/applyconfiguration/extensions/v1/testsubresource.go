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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	metav1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// TestSubresourceApplyConfiguration represents a declarative configuration of the TestSubresource type for use
// with apply.
type TestSubresourceApplyConfiguration struct {
	metav1.TypeMetaApplyConfiguration `json:",inline"`
	Name                              *string `json:"name,omitempty"`
}

// TestSubresourceApplyConfiguration constructs a declarative configuration of the TestSubresource type for use with
// apply.
func TestSubresource() *TestSubresourceApplyConfiguration {
	b := &TestSubresourceApplyConfiguration{}
	b.WithKind("TestSubresource")
	b.WithAPIVersion("extensions.test.crd.code-generator.k8s.io/v1")
	return b
}

// WithKind sets the Kind field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Kind field is set to the value of the last call.
func (b *TestSubresourceApplyConfiguration) WithKind(value string) *TestSubresourceApplyConfiguration {
	b.TypeMetaApplyConfiguration.Kind = &value
	return b
}

// WithAPIVersion sets the APIVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the APIVersion field is set to the value of the last call.
func (b *TestSubresourceApplyConfiguration) WithAPIVersion(value string) *TestSubresourceApplyConfiguration {
	b.TypeMetaApplyConfiguration.APIVersion = &value
	return b
}

// WithName sets the Name field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Name field is set to the value of the last call.
func (b *TestSubresourceApplyConfiguration) WithName(value string) *TestSubresourceApplyConfiguration {
	b.Name = &value
	return b
}