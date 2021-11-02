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

// NodeSelectorTermApplyConfiguration represents an declarative configuration of the NodeSelectorTerm type for use
// with apply.
type NodeSelectorTermApplyConfiguration struct {
	MatchExpressions []NodeSelectorRequirementApplyConfiguration `json:"matchExpressions,omitempty"`
	MatchFields      []NodeSelectorRequirementApplyConfiguration `json:"matchFields,omitempty"`
}

// NodeSelectorTermApplyConfiguration constructs an declarative configuration of the NodeSelectorTerm type for use with
// apply.
func NodeSelectorTerm() *NodeSelectorTermApplyConfiguration {
	return &NodeSelectorTermApplyConfiguration{}
}

// WithMatchExpressions adds the given value to the MatchExpressions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the MatchExpressions field.
func (b *NodeSelectorTermApplyConfiguration) WithMatchExpressions(values ...*NodeSelectorRequirementApplyConfiguration) *NodeSelectorTermApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithMatchExpressions")
		}
		b.MatchExpressions = append(b.MatchExpressions, *values[i])
	}
	return b
}

// WithMatchFields adds the given value to the MatchFields field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the MatchFields field.
func (b *NodeSelectorTermApplyConfiguration) WithMatchFields(values ...*NodeSelectorRequirementApplyConfiguration) *NodeSelectorTermApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithMatchFields")
		}
		b.MatchFields = append(b.MatchFields, *values[i])
	}
	return b
}