// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentConfig) DeepCopyInto(out *EnvironmentConfig) {
	*out = *in
	if in.EnvFrom != nil {
		in, out := &in.EnvFrom, &out.EnvFrom
		*out = make([]v1.EnvFromSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentConfig.
func (in *EnvironmentConfig) DeepCopy() *EnvironmentConfig {
	if in == nil {
		return nil
	}
	out := new(EnvironmentConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlinkApplication) DeepCopyInto(out *FlinkApplication) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlinkApplication.
func (in *FlinkApplication) DeepCopy() *FlinkApplication {
	if in == nil {
		return nil
	}
	out := new(FlinkApplication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlinkApplication) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlinkApplicationList) DeepCopyInto(out *FlinkApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FlinkApplication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlinkApplicationList.
func (in *FlinkApplicationList) DeepCopy() *FlinkApplicationList {
	if in == nil {
		return nil
	}
	out := new(FlinkApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FlinkApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlinkApplicationSpec) DeepCopyInto(out *FlinkApplicationSpec) {
	*out = *in
	in.TaskManagerConfig.DeepCopyInto(&out.TaskManagerConfig)
	in.JobManagerConfig.DeepCopyInto(&out.JobManagerConfig)
	out.FlinkJob = in.FlinkJob
	if in.RpcPort != nil {
		in, out := &in.RpcPort, &out.RpcPort
		*out = new(int32)
		**out = **in
	}
	if in.BlobPort != nil {
		in, out := &in.BlobPort, &out.BlobPort
		*out = new(int32)
		**out = **in
	}
	if in.QueryPort != nil {
		in, out := &in.QueryPort, &out.QueryPort
		*out = new(int32)
		**out = **in
	}
	if in.UiPort != nil {
		in, out := &in.UiPort, &out.UiPort
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlinkApplicationSpec.
func (in *FlinkApplicationSpec) DeepCopy() *FlinkApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(FlinkApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlinkApplicationStatus) DeepCopyInto(out *FlinkApplicationStatus) {
	*out = *in
	if in.StartedAt != nil {
		in, out := &in.StartedAt, &out.StartedAt
		*out = (*in).DeepCopy()
	}
	if in.StoppedAt != nil {
		in, out := &in.StoppedAt, &out.StoppedAt
		*out = (*in).DeepCopy()
	}
	if in.LastUpdatedAt != nil {
		in, out := &in.LastUpdatedAt, &out.LastUpdatedAt
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlinkApplicationStatus.
func (in *FlinkApplicationStatus) DeepCopy() *FlinkApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(FlinkApplicationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlinkJobInfo) DeepCopyInto(out *FlinkJobInfo) {
	*out = *in
	out.SavepointInfo = in.SavepointInfo
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlinkJobInfo.
func (in *FlinkJobInfo) DeepCopy() *FlinkJobInfo {
	if in == nil {
		return nil
	}
	out := new(FlinkJobInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JobManagerConfig) DeepCopyInto(out *JobManagerConfig) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	in.Environment.DeepCopyInto(&out.Environment)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JobManagerConfig.
func (in *JobManagerConfig) DeepCopy() *JobManagerConfig {
	if in == nil {
		return nil
	}
	out := new(JobManagerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SavepointInfo) DeepCopyInto(out *SavepointInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SavepointInfo.
func (in *SavepointInfo) DeepCopy() *SavepointInfo {
	if in == nil {
		return nil
	}
	out := new(SavepointInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TaskManagerConfig) DeepCopyInto(out *TaskManagerConfig) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	in.Environment.DeepCopyInto(&out.Environment)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TaskManagerConfig.
func (in *TaskManagerConfig) DeepCopy() *TaskManagerConfig {
	if in == nil {
		return nil
	}
	out := new(TaskManagerConfig)
	in.DeepCopyInto(out)
	return out
}