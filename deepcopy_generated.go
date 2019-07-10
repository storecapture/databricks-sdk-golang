// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package databricks

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DBClientOption) DeepCopyInto(out *DBClientOption) {
	*out = *in
	if in.DefaultHeaders != nil {
		in, out := &in.DefaultHeaders, &out.DefaultHeaders
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DBClientOption.
func (in *DBClientOption) DeepCopy() *DBClientOption {
	if in == nil {
		return nil
	}
	out := new(DBClientOption)
	in.DeepCopyInto(out)
	return out
}
