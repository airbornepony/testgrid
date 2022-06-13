/*
Copyright The TestGrid Authors.

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

// Backing state for issues associated with a TestGrid test group.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: issue_state.proto

package issue_state

import (
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

type TargetAndMethods struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetName  string   `protobuf:"bytes,1,opt,name=target_name,json=targetName,proto3" json:"target_name,omitempty"`
	MethodNames []string `protobuf:"bytes,2,rep,name=method_names,json=methodNames,proto3" json:"method_names,omitempty"`
}

func (x *TargetAndMethods) Reset() {
	*x = TargetAndMethods{}
	if protoimpl.UnsafeEnabled {
		mi := &file_issue_state_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TargetAndMethods) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TargetAndMethods) ProtoMessage() {}

func (x *TargetAndMethods) ProtoReflect() protoreflect.Message {
	mi := &file_issue_state_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TargetAndMethods.ProtoReflect.Descriptor instead.
func (*TargetAndMethods) Descriptor() ([]byte, []int) {
	return file_issue_state_proto_rawDescGZIP(), []int{0}
}

func (x *TargetAndMethods) GetTargetName() string {
	if x != nil {
		return x.TargetName
	}
	return ""
}

func (x *TargetAndMethods) GetMethodNames() []string {
	if x != nil {
		return x.MethodNames
	}
	return nil
}

type IssueInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IssueId        string   `protobuf:"bytes,1,opt,name=issue_id,json=issueId,proto3" json:"issue_id,omitempty"`
	Title          string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`                                            // Issue title or description.
	IsAutobug      bool     `protobuf:"varint,3,opt,name=is_autobug,json=isAutobug,proto3" json:"is_autobug,omitempty"`                  // True if auto-created by TestGrid for a failing test.
	IsFlakinessBug bool     `protobuf:"varint,8,opt,name=is_flakiness_bug,json=isFlakinessBug,proto3" json:"is_flakiness_bug,omitempty"` // True if auto-created by TestGrid for a flaky test.
	LastModified   float64  `protobuf:"fixed64,4,opt,name=last_modified,json=lastModified,proto3" json:"last_modified,omitempty"`        // In seconds since epoch.
	RowIds         []string `protobuf:"bytes,5,rep,name=row_ids,json=rowIds,proto3" json:"row_ids,omitempty"`                            // Associated row IDs (mentioned in the issue).
	// Run IDs used to associate this issue with a particular target (in case of
	// repeats, or across runs on different dashboards).
	RunIds []string `protobuf:"bytes,6,rep,name=run_ids,json=runIds,proto3" json:"run_ids,omitempty"`
	// Targets + methods associated with this issue.
	// Only set if test group's `link_bugs_by_test_methods` is True, else all
	// targets + methods will be linked to this issue.
	TargetsAndMethods []*TargetAndMethods `protobuf:"bytes,7,rep,name=targets_and_methods,json=targetsAndMethods,proto3" json:"targets_and_methods,omitempty"`
}

func (x *IssueInfo) Reset() {
	*x = IssueInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_issue_state_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueInfo) ProtoMessage() {}

func (x *IssueInfo) ProtoReflect() protoreflect.Message {
	mi := &file_issue_state_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueInfo.ProtoReflect.Descriptor instead.
func (*IssueInfo) Descriptor() ([]byte, []int) {
	return file_issue_state_proto_rawDescGZIP(), []int{1}
}

func (x *IssueInfo) GetIssueId() string {
	if x != nil {
		return x.IssueId
	}
	return ""
}

func (x *IssueInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *IssueInfo) GetIsAutobug() bool {
	if x != nil {
		return x.IsAutobug
	}
	return false
}

func (x *IssueInfo) GetIsFlakinessBug() bool {
	if x != nil {
		return x.IsFlakinessBug
	}
	return false
}

func (x *IssueInfo) GetLastModified() float64 {
	if x != nil {
		return x.LastModified
	}
	return 0
}

func (x *IssueInfo) GetRowIds() []string {
	if x != nil {
		return x.RowIds
	}
	return nil
}

func (x *IssueInfo) GetRunIds() []string {
	if x != nil {
		return x.RunIds
	}
	return nil
}

func (x *IssueInfo) GetTargetsAndMethods() []*TargetAndMethods {
	if x != nil {
		return x.TargetsAndMethods
	}
	return nil
}

type IssueState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of collected info for bugs.
	IssueInfo []*IssueInfo `protobuf:"bytes,1,rep,name=issue_info,json=issueInfo,proto3" json:"issue_info,omitempty"`
}

func (x *IssueState) Reset() {
	*x = IssueState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_issue_state_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueState) ProtoMessage() {}

func (x *IssueState) ProtoReflect() protoreflect.Message {
	mi := &file_issue_state_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueState.ProtoReflect.Descriptor instead.
func (*IssueState) Descriptor() ([]byte, []int) {
	return file_issue_state_proto_rawDescGZIP(), []int{2}
}

func (x *IssueState) GetIssueInfo() []*IssueInfo {
	if x != nil {
		return x.IssueInfo
	}
	return nil
}

var File_issue_state_proto protoreflect.FileDescriptor

var file_issue_state_proto_rawDesc = []byte{
	0x0a, 0x11, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x10, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x41, 0x6e, 0x64,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x9f, 0x02, 0x0a, 0x09,
	0x49, 0x73, 0x73, 0x75, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x73,
	0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x73, 0x73,
	0x75, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73,
	0x5f, 0x61, 0x75, 0x74, 0x6f, 0x62, 0x75, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x69, 0x73, 0x41, 0x75, 0x74, 0x6f, 0x62, 0x75, 0x67, 0x12, 0x28, 0x0a, 0x10, 0x69, 0x73, 0x5f,
	0x66, 0x6c, 0x61, 0x6b, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x62, 0x75, 0x67, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0e, 0x69, 0x73, 0x46, 0x6c, 0x61, 0x6b, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x42, 0x75, 0x67, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74,
	0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x77, 0x5f,
	0x69, 0x64, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x77, 0x49, 0x64,
	0x73, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x75, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x75, 0x6e, 0x49, 0x64, 0x73, 0x12, 0x41, 0x0a, 0x13, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x73, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74,
	0x41, 0x6e, 0x64, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x52, 0x11, 0x74, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x73, 0x41, 0x6e, 0x64, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x73, 0x22, 0x43, 0x0a,
	0x0a, 0x49, 0x73, 0x73, 0x75, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x29, 0x0a, 0x0a, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x69, 0x73, 0x73,
	0x75, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x4a, 0x04, 0x08, 0x03,
	0x10, 0x04, 0x42, 0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x6c, 0x61, 0x74,
	0x66, 0x6f, 0x72, 0x6d, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x67, 0x72, 0x69, 0x64, 0x2f, 0x70, 0x62,
	0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_issue_state_proto_rawDescOnce sync.Once
	file_issue_state_proto_rawDescData = file_issue_state_proto_rawDesc
)

func file_issue_state_proto_rawDescGZIP() []byte {
	file_issue_state_proto_rawDescOnce.Do(func() {
		file_issue_state_proto_rawDescData = protoimpl.X.CompressGZIP(file_issue_state_proto_rawDescData)
	})
	return file_issue_state_proto_rawDescData
}

var file_issue_state_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_issue_state_proto_goTypes = []interface{}{
	(*TargetAndMethods)(nil), // 0: TargetAndMethods
	(*IssueInfo)(nil),        // 1: IssueInfo
	(*IssueState)(nil),       // 2: IssueState
}
var file_issue_state_proto_depIdxs = []int32{
	0, // 0: IssueInfo.targets_and_methods:type_name -> TargetAndMethods
	1, // 1: IssueState.issue_info:type_name -> IssueInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_issue_state_proto_init() }
func file_issue_state_proto_init() {
	if File_issue_state_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_issue_state_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TargetAndMethods); i {
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
		file_issue_state_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueInfo); i {
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
		file_issue_state_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueState); i {
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
			RawDescriptor: file_issue_state_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_issue_state_proto_goTypes,
		DependencyIndexes: file_issue_state_proto_depIdxs,
		MessageInfos:      file_issue_state_proto_msgTypes,
	}.Build()
	File_issue_state_proto = out.File
	file_issue_state_proto_rawDesc = nil
	file_issue_state_proto_goTypes = nil
	file_issue_state_proto_depIdxs = nil
}
