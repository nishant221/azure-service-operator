//go:build !ignore_autogenerated

/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

// Code generated by controller-gen. DO NOT EDIT.

package core

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DestinationExpression) DeepCopyInto(out *DestinationExpression) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DestinationExpression.
func (in *DestinationExpression) DeepCopy() *DestinationExpression {
	if in == nil {
		return nil
	}
	out := new(DestinationExpression)
	in.DeepCopyInto(out)
	return out
}