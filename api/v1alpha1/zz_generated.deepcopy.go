//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022. projectsveltos.io. All rights reserved.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterFeature) DeepCopyInto(out *ClusterFeature) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterFeature.
func (in *ClusterFeature) DeepCopy() *ClusterFeature {
	if in == nil {
		return nil
	}
	out := new(ClusterFeature)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterFeature) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterFeatureList) DeepCopyInto(out *ClusterFeatureList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterFeature, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterFeatureList.
func (in *ClusterFeatureList) DeepCopy() *ClusterFeatureList {
	if in == nil {
		return nil
	}
	out := new(ClusterFeatureList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterFeatureList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterFeatureSpec) DeepCopyInto(out *ClusterFeatureSpec) {
	*out = *in
	if in.WorkloadRoleRefs != nil {
		in, out := &in.WorkloadRoleRefs, &out.WorkloadRoleRefs
		*out = make([]v1.ObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.KyvernoConfiguration != nil {
		in, out := &in.KyvernoConfiguration, &out.KyvernoConfiguration
		*out = new(KyvernoConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.GatekeeperConfiguration != nil {
		in, out := &in.GatekeeperConfiguration, &out.GatekeeperConfiguration
		*out = new(GatekeeperConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.PrometheusConfiguration != nil {
		in, out := &in.PrometheusConfiguration, &out.PrometheusConfiguration
		*out = new(PrometheusConfiguration)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterFeatureSpec.
func (in *ClusterFeatureSpec) DeepCopy() *ClusterFeatureSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterFeatureSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterFeatureStatus) DeepCopyInto(out *ClusterFeatureStatus) {
	*out = *in
	if in.MatchingClusterRefs != nil {
		in, out := &in.MatchingClusterRefs, &out.MatchingClusterRefs
		*out = make([]v1.ObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterFeatureStatus.
func (in *ClusterFeatureStatus) DeepCopy() *ClusterFeatureStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterFeatureStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSummary) DeepCopyInto(out *ClusterSummary) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSummary.
func (in *ClusterSummary) DeepCopy() *ClusterSummary {
	if in == nil {
		return nil
	}
	out := new(ClusterSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterSummary) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSummaryList) DeepCopyInto(out *ClusterSummaryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterSummary, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSummaryList.
func (in *ClusterSummaryList) DeepCopy() *ClusterSummaryList {
	if in == nil {
		return nil
	}
	out := new(ClusterSummaryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterSummaryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSummarySpec) DeepCopyInto(out *ClusterSummarySpec) {
	*out = *in
	in.ClusterFeatureSpec.DeepCopyInto(&out.ClusterFeatureSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSummarySpec.
func (in *ClusterSummarySpec) DeepCopy() *ClusterSummarySpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSummarySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSummaryStatus) DeepCopyInto(out *ClusterSummaryStatus) {
	*out = *in
	if in.FeatureSummaries != nil {
		in, out := &in.FeatureSummaries, &out.FeatureSummaries
		*out = make([]FeatureSummary, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.GatekeeperSortedPolicies != nil {
		in, out := &in.GatekeeperSortedPolicies, &out.GatekeeperSortedPolicies
		*out = make([]v1.ObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSummaryStatus.
func (in *ClusterSummaryStatus) DeepCopy() *ClusterSummaryStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterSummaryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureSummary) DeepCopyInto(out *FeatureSummary) {
	*out = *in
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.FailureReason != nil {
		in, out := &in.FailureReason, &out.FailureReason
		*out = new(string)
		**out = **in
	}
	if in.FailureMessage != nil {
		in, out := &in.FailureMessage, &out.FailureMessage
		*out = new(string)
		**out = **in
	}
	if in.DeployedGroupVersionKind != nil {
		in, out := &in.DeployedGroupVersionKind, &out.DeployedGroupVersionKind
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureSummary.
func (in *FeatureSummary) DeepCopy() *FeatureSummary {
	if in == nil {
		return nil
	}
	out := new(FeatureSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GatekeeperConfiguration) DeepCopyInto(out *GatekeeperConfiguration) {
	*out = *in
	if in.PolicyRefs != nil {
		in, out := &in.PolicyRefs, &out.PolicyRefs
		*out = make([]v1.ObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GatekeeperConfiguration.
func (in *GatekeeperConfiguration) DeepCopy() *GatekeeperConfiguration {
	if in == nil {
		return nil
	}
	out := new(GatekeeperConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KyvernoConfiguration) DeepCopyInto(out *KyvernoConfiguration) {
	*out = *in
	if in.PolicyRefs != nil {
		in, out := &in.PolicyRefs, &out.PolicyRefs
		*out = make([]v1.ObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KyvernoConfiguration.
func (in *KyvernoConfiguration) DeepCopy() *KyvernoConfiguration {
	if in == nil {
		return nil
	}
	out := new(KyvernoConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusConfiguration) DeepCopyInto(out *PrometheusConfiguration) {
	*out = *in
	if in.StorageClassName != nil {
		in, out := &in.StorageClassName, &out.StorageClassName
		*out = new(string)
		**out = **in
	}
	if in.StorageQuantity != nil {
		in, out := &in.StorageQuantity, &out.StorageQuantity
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.PolicyRefs != nil {
		in, out := &in.PolicyRefs, &out.PolicyRefs
		*out = make([]v1.ObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusConfiguration.
func (in *PrometheusConfiguration) DeepCopy() *PrometheusConfiguration {
	if in == nil {
		return nil
	}
	out := new(PrometheusConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadRole) DeepCopyInto(out *WorkloadRole) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadRole.
func (in *WorkloadRole) DeepCopy() *WorkloadRole {
	if in == nil {
		return nil
	}
	out := new(WorkloadRole)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkloadRole) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadRoleList) DeepCopyInto(out *WorkloadRoleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WorkloadRole, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadRoleList.
func (in *WorkloadRoleList) DeepCopy() *WorkloadRoleList {
	if in == nil {
		return nil
	}
	out := new(WorkloadRoleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkloadRoleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadRoleSpec) DeepCopyInto(out *WorkloadRoleSpec) {
	*out = *in
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(string)
		**out = **in
	}
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]rbacv1.PolicyRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.AggregationRule != nil {
		in, out := &in.AggregationRule, &out.AggregationRule
		*out = new(rbacv1.AggregationRule)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadRoleSpec.
func (in *WorkloadRoleSpec) DeepCopy() *WorkloadRoleSpec {
	if in == nil {
		return nil
	}
	out := new(WorkloadRoleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadRoleStatus) DeepCopyInto(out *WorkloadRoleStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadRoleStatus.
func (in *WorkloadRoleStatus) DeepCopy() *WorkloadRoleStatus {
	if in == nil {
		return nil
	}
	out := new(WorkloadRoleStatus)
	in.DeepCopyInto(out)
	return out
}
