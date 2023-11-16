//go:build !ignore_autogenerated

// Copyright 2019-2023 The Liqo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterIdentity) DeepCopyInto(out *ClusterIdentity) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterIdentity.
func (in *ClusterIdentity) DeepCopy() *ClusterIdentity {
	if in == nil {
		return nil
	}
	out := new(ClusterIdentity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ForeignCluster) DeepCopyInto(out *ForeignCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ForeignCluster.
func (in *ForeignCluster) DeepCopy() *ForeignCluster {
	if in == nil {
		return nil
	}
	out := new(ForeignCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ForeignCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ForeignClusterList) DeepCopyInto(out *ForeignClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ForeignCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ForeignClusterList.
func (in *ForeignClusterList) DeepCopy() *ForeignClusterList {
	if in == nil {
		return nil
	}
	out := new(ForeignClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ForeignClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ForeignClusterSpec) DeepCopyInto(out *ForeignClusterSpec) {
	*out = *in
	out.ClusterIdentity = in.ClusterIdentity
	if in.InsecureSkipTLSVerify != nil {
		in, out := &in.InsecureSkipTLSVerify, &out.InsecureSkipTLSVerify
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ForeignClusterSpec.
func (in *ForeignClusterSpec) DeepCopy() *ForeignClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ForeignClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ForeignClusterStatus) DeepCopyInto(out *ForeignClusterStatus) {
	*out = *in
	out.TenantNamespace = in.TenantNamespace
	if in.PeeringConditions != nil {
		in, out := &in.PeeringConditions, &out.PeeringConditions
		*out = make([]PeeringCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ForeignClusterStatus.
func (in *ForeignClusterStatus) DeepCopy() *ForeignClusterStatus {
	if in == nil {
		return nil
	}
	out := new(ForeignClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PeeringCondition) DeepCopyInto(out *PeeringCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PeeringCondition.
func (in *PeeringCondition) DeepCopy() *PeeringCondition {
	if in == nil {
		return nil
	}
	out := new(PeeringCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRequest) DeepCopyInto(out *ResourceRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRequest.
func (in *ResourceRequest) DeepCopy() *ResourceRequest {
	if in == nil {
		return nil
	}
	out := new(ResourceRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRequestList) DeepCopyInto(out *ResourceRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRequestList.
func (in *ResourceRequestList) DeepCopy() *ResourceRequestList {
	if in == nil {
		return nil
	}
	out := new(ResourceRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRequestSpec) DeepCopyInto(out *ResourceRequestSpec) {
	*out = *in
	out.ClusterIdentity = in.ClusterIdentity
	if in.WithdrawalTimestamp != nil {
		in, out := &in.WithdrawalTimestamp, &out.WithdrawalTimestamp
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRequestSpec.
func (in *ResourceRequestSpec) DeepCopy() *ResourceRequestSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRequestStatus) DeepCopyInto(out *ResourceRequestStatus) {
	*out = *in
	if in.OfferWithdrawalTimestamp != nil {
		in, out := &in.OfferWithdrawalTimestamp, &out.OfferWithdrawalTimestamp
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRequestStatus.
func (in *ResourceRequestStatus) DeepCopy() *ResourceRequestStatus {
	if in == nil {
		return nil
	}
	out := new(ResourceRequestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TenantNamespaceType) DeepCopyInto(out *TenantNamespaceType) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TenantNamespaceType.
func (in *TenantNamespaceType) DeepCopy() *TenantNamespaceType {
	if in == nil {
		return nil
	}
	out := new(TenantNamespaceType)
	in.DeepCopyInto(out)
	return out
}
