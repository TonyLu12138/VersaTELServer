//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2020 The KubeSphere Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha2

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceQuota) DeepCopyInto(out *ResourceQuota) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceQuota.
func (in *ResourceQuota) DeepCopy() *ResourceQuota {
	if in == nil {
		return nil
	}
	out := new(ResourceQuota)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceQuota) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceQuotaList) DeepCopyInto(out *ResourceQuotaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceQuota, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceQuotaList.
func (in *ResourceQuotaList) DeepCopy() *ResourceQuotaList {
	if in == nil {
		return nil
	}
	out := new(ResourceQuotaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceQuotaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceQuotaSpec) DeepCopyInto(out *ResourceQuotaSpec) {
	*out = *in
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.Quota.DeepCopyInto(&out.Quota)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceQuotaSpec.
func (in *ResourceQuotaSpec) DeepCopy() *ResourceQuotaSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceQuotaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceQuotaStatus) DeepCopyInto(out *ResourceQuotaStatus) {
	*out = *in
	in.Total.DeepCopyInto(&out.Total)
	if in.Namespaces != nil {
		in, out := &in.Namespaces, &out.Namespaces
		*out = make(ResourceQuotasStatusByNamespace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceQuotaStatus.
func (in *ResourceQuotaStatus) DeepCopy() *ResourceQuotaStatus {
	if in == nil {
		return nil
	}
	out := new(ResourceQuotaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceQuotaStatusByNamespace) DeepCopyInto(out *ResourceQuotaStatusByNamespace) {
	*out = *in
	in.ResourceQuotaStatus.DeepCopyInto(&out.ResourceQuotaStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceQuotaStatusByNamespace.
func (in *ResourceQuotaStatusByNamespace) DeepCopy() *ResourceQuotaStatusByNamespace {
	if in == nil {
		return nil
	}
	out := new(ResourceQuotaStatusByNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ResourceQuotasStatusByNamespace) DeepCopyInto(out *ResourceQuotasStatusByNamespace) {
	{
		in := &in
		*out = make(ResourceQuotasStatusByNamespace, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceQuotasStatusByNamespace.
func (in ResourceQuotasStatusByNamespace) DeepCopy() ResourceQuotasStatusByNamespace {
	if in == nil {
		return nil
	}
	out := new(ResourceQuotasStatusByNamespace)
	in.DeepCopyInto(out)
	return *out
}