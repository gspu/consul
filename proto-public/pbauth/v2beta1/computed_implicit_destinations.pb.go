// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: pbauth/v2beta1/computed_implicit_destinations.proto

package authv2beta1

import (
	pbresource "github.com/hashicorp/consul/proto-public/pbresource"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ImplicitDestinations tracks destination services for a given workload identity.
type ComputedImplicitDestinations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// destinations is the list of destinations.
	Destinations []*ImplicitDestination `protobuf:"bytes,1,rep,name=destinations,proto3" json:"destinations,omitempty"`
	// BoundReferences is a slice of mixed type references of resources that were
	// involved in the formulation of this resource.
	BoundReferences []*pbresource.Reference `protobuf:"bytes,2,rep,name=bound_references,json=boundReferences,proto3" json:"bound_references,omitempty"`
}

func (x *ComputedImplicitDestinations) Reset() {
	*x = ComputedImplicitDestinations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbauth_v2beta1_computed_implicit_destinations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ComputedImplicitDestinations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ComputedImplicitDestinations) ProtoMessage() {}

func (x *ComputedImplicitDestinations) ProtoReflect() protoreflect.Message {
	mi := &file_pbauth_v2beta1_computed_implicit_destinations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ComputedImplicitDestinations.ProtoReflect.Descriptor instead.
func (*ComputedImplicitDestinations) Descriptor() ([]byte, []int) {
	return file_pbauth_v2beta1_computed_implicit_destinations_proto_rawDescGZIP(), []int{0}
}

func (x *ComputedImplicitDestinations) GetDestinations() []*ImplicitDestination {
	if x != nil {
		return x.Destinations
	}
	return nil
}

func (x *ComputedImplicitDestinations) GetBoundReferences() []*pbresource.Reference {
	if x != nil {
		return x.BoundReferences
	}
	return nil
}

// Destination contains a reference to a catalog service, any workload identity
// references this service points to and a list of port names that are allowed by TrafficPermissions.
type ImplicitDestination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServiceRef   *pbresource.Reference   `protobuf:"bytes,1,opt,name=service_ref,json=serviceRef,proto3" json:"service_ref,omitempty"`
	IdentityRefs []*pbresource.Reference `protobuf:"bytes,2,rep,name=identity_refs,json=identityRefs,proto3" json:"identity_refs,omitempty"`
	Ports        []string                `protobuf:"bytes,3,rep,name=ports,proto3" json:"ports,omitempty"`
}

func (x *ImplicitDestination) Reset() {
	*x = ImplicitDestination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pbauth_v2beta1_computed_implicit_destinations_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImplicitDestination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImplicitDestination) ProtoMessage() {}

func (x *ImplicitDestination) ProtoReflect() protoreflect.Message {
	mi := &file_pbauth_v2beta1_computed_implicit_destinations_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImplicitDestination.ProtoReflect.Descriptor instead.
func (*ImplicitDestination) Descriptor() ([]byte, []int) {
	return file_pbauth_v2beta1_computed_implicit_destinations_proto_rawDescGZIP(), []int{1}
}

func (x *ImplicitDestination) GetServiceRef() *pbresource.Reference {
	if x != nil {
		return x.ServiceRef
	}
	return nil
}

func (x *ImplicitDestination) GetIdentityRefs() []*pbresource.Reference {
	if x != nil {
		return x.IdentityRefs
	}
	return nil
}

func (x *ImplicitDestination) GetPorts() []string {
	if x != nil {
		return x.Ports
	}
	return nil
}

var File_pbauth_v2beta1_computed_implicit_destinations_proto protoreflect.FileDescriptor

var file_pbauth_v2beta1_computed_implicit_destinations_proto_rawDesc = []byte{
	0x0a, 0x33, 0x70, 0x62, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x5f, 0x69, 0x6d, 0x70, 0x6c, 0x69, 0x63,
	0x69, 0x74, 0x5f, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70,
	0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x32, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x1a, 0x1c, 0x70, 0x62, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x70, 0x62, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x01,
	0x0a, 0x1c, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x49, 0x6d, 0x70, 0x6c, 0x69, 0x63,
	0x69, 0x74, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x56,
	0x0a, 0x0c, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70,
	0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x32, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x49, 0x6d, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x44, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4f, 0x0a, 0x10, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x5f,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x24, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6c, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x0f, 0x62, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x3a, 0x06, 0xa2, 0x93, 0x04, 0x02, 0x08, 0x03, 0x22,
	0xc5, 0x01, 0x0a, 0x13, 0x49, 0x6d, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x44, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x68,
	0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x66, 0x12, 0x49,
	0x0a, 0x0d, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x72, 0x65, 0x66, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72,
	0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x0c, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x65, 0x66, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x72,
	0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x3a,
	0x06, 0xa2, 0x93, 0x04, 0x02, 0x08, 0x03, 0x42, 0xa2, 0x02, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e,
	0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x42, 0x21, 0x43,
	0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x64, 0x49, 0x6d, 0x70, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x44,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x43, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68,
	0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x70, 0x62, 0x61,
	0x75, 0x74, 0x68, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x61, 0x75, 0x74, 0x68,
	0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x48, 0x43, 0x41, 0xaa, 0x02, 0x1d,
	0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c,
	0x2e, 0x41, 0x75, 0x74, 0x68, 0x2e, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x1d,
	0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c,
	0x5c, 0x41, 0x75, 0x74, 0x68, 0x5c, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0xe2, 0x02, 0x29,
	0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c,
	0x5c, 0x41, 0x75, 0x74, 0x68, 0x5c, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x20, 0x48, 0x61, 0x73, 0x68,
	0x69, 0x63, 0x6f, 0x72, 0x70, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x3a, 0x3a, 0x41,
	0x75, 0x74, 0x68, 0x3a, 0x3a, 0x56, 0x32, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pbauth_v2beta1_computed_implicit_destinations_proto_rawDescOnce sync.Once
	file_pbauth_v2beta1_computed_implicit_destinations_proto_rawDescData = file_pbauth_v2beta1_computed_implicit_destinations_proto_rawDesc
)

func file_pbauth_v2beta1_computed_implicit_destinations_proto_rawDescGZIP() []byte {
	file_pbauth_v2beta1_computed_implicit_destinations_proto_rawDescOnce.Do(func() {
		file_pbauth_v2beta1_computed_implicit_destinations_proto_rawDescData = protoimpl.X.CompressGZIP(file_pbauth_v2beta1_computed_implicit_destinations_proto_rawDescData)
	})
	return file_pbauth_v2beta1_computed_implicit_destinations_proto_rawDescData
}

var file_pbauth_v2beta1_computed_implicit_destinations_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pbauth_v2beta1_computed_implicit_destinations_proto_goTypes = []interface{}{
	(*ComputedImplicitDestinations)(nil), // 0: hashicorp.consul.auth.v2beta1.ComputedImplicitDestinations
	(*ImplicitDestination)(nil),          // 1: hashicorp.consul.auth.v2beta1.ImplicitDestination
	(*pbresource.Reference)(nil),         // 2: hashicorp.consul.resource.Reference
}
var file_pbauth_v2beta1_computed_implicit_destinations_proto_depIdxs = []int32{
	1, // 0: hashicorp.consul.auth.v2beta1.ComputedImplicitDestinations.destinations:type_name -> hashicorp.consul.auth.v2beta1.ImplicitDestination
	2, // 1: hashicorp.consul.auth.v2beta1.ComputedImplicitDestinations.bound_references:type_name -> hashicorp.consul.resource.Reference
	2, // 2: hashicorp.consul.auth.v2beta1.ImplicitDestination.service_ref:type_name -> hashicorp.consul.resource.Reference
	2, // 3: hashicorp.consul.auth.v2beta1.ImplicitDestination.identity_refs:type_name -> hashicorp.consul.resource.Reference
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pbauth_v2beta1_computed_implicit_destinations_proto_init() }
func file_pbauth_v2beta1_computed_implicit_destinations_proto_init() {
	if File_pbauth_v2beta1_computed_implicit_destinations_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pbauth_v2beta1_computed_implicit_destinations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ComputedImplicitDestinations); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pbauth_v2beta1_computed_implicit_destinations_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImplicitDestination); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pbauth_v2beta1_computed_implicit_destinations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pbauth_v2beta1_computed_implicit_destinations_proto_goTypes,
		DependencyIndexes: file_pbauth_v2beta1_computed_implicit_destinations_proto_depIdxs,
		MessageInfos:      file_pbauth_v2beta1_computed_implicit_destinations_proto_msgTypes,
	}.Build()
	File_pbauth_v2beta1_computed_implicit_destinations_proto = out.File
	file_pbauth_v2beta1_computed_implicit_destinations_proto_rawDesc = nil
	file_pbauth_v2beta1_computed_implicit_destinations_proto_goTypes = nil
	file_pbauth_v2beta1_computed_implicit_destinations_proto_depIdxs = nil
}
