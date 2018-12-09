// Code generated by protoc-gen-go. DO NOT EDIT.
// source: actions.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	_ "www.velocidex.com/golang/velociraptor/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// These are used when mapping registry keys to the VFS
type StatEntry_RegistryType int32

const (
	StatEntry_REG_NONE                StatEntry_RegistryType = 0
	StatEntry_REG_SZ                  StatEntry_RegistryType = 1
	StatEntry_REG_EXPAND_SZ           StatEntry_RegistryType = 2
	StatEntry_REG_BINARY              StatEntry_RegistryType = 3
	StatEntry_REG_DWORD               StatEntry_RegistryType = 4
	StatEntry_REG_DWORD_LITTLE_ENDIAN StatEntry_RegistryType = 4
	StatEntry_REG_DWORD_BIG_ENDIAN    StatEntry_RegistryType = 5
	StatEntry_REG_LINK                StatEntry_RegistryType = 6
	StatEntry_REG_MULTI_SZ            StatEntry_RegistryType = 7
	StatEntry_REG_QWORD               StatEntry_RegistryType = 11
)

var StatEntry_RegistryType_name = map[int32]string{
	0: "REG_NONE",
	1: "REG_SZ",
	2: "REG_EXPAND_SZ",
	3: "REG_BINARY",
	4: "REG_DWORD",
	// Duplicate value: 4: "REG_DWORD_LITTLE_ENDIAN",
	5:  "REG_DWORD_BIG_ENDIAN",
	6:  "REG_LINK",
	7:  "REG_MULTI_SZ",
	11: "REG_QWORD",
}

var StatEntry_RegistryType_value = map[string]int32{
	"REG_NONE":                0,
	"REG_SZ":                  1,
	"REG_EXPAND_SZ":           2,
	"REG_BINARY":              3,
	"REG_DWORD":               4,
	"REG_DWORD_LITTLE_ENDIAN": 4,
	"REG_DWORD_BIG_ENDIAN":    5,
	"REG_LINK":                6,
	"REG_MULTI_SZ":            7,
	"REG_QWORD":               11,
}

func (x StatEntry_RegistryType) String() string {
	return proto.EnumName(StatEntry_RegistryType_name, int32(x))
}

func (StatEntry_RegistryType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eeb49063df94c2b8, []int{2, 0}
}

// Stores information about the GRR client itself
type ClientInformation struct {
	ClientName           string   `protobuf:"bytes,1,opt,name=client_name,json=clientName,proto3" json:"client_name,omitempty"`
	ClientVersion        string   `protobuf:"bytes,7,opt,name=client_version,json=clientVersion,proto3" json:"client_version,omitempty"`
	Revision             uint64   `protobuf:"varint,3,opt,name=revision,proto3" json:"revision,omitempty"`
	BuildTime            string   `protobuf:"bytes,4,opt,name=build_time,json=buildTime,proto3" json:"build_time,omitempty"`
	ClientDescription    string   `protobuf:"bytes,5,opt,name=client_description,json=clientDescription,proto3" json:"client_description,omitempty"`
	Labels               []string `protobuf:"bytes,6,rep,name=labels,proto3" json:"labels,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientInformation) Reset()         { *m = ClientInformation{} }
func (m *ClientInformation) String() string { return proto.CompactTextString(m) }
func (*ClientInformation) ProtoMessage()    {}
func (*ClientInformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeb49063df94c2b8, []int{0}
}

func (m *ClientInformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientInformation.Unmarshal(m, b)
}
func (m *ClientInformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientInformation.Marshal(b, m, deterministic)
}
func (m *ClientInformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientInformation.Merge(m, src)
}
func (m *ClientInformation) XXX_Size() int {
	return xxx_messageInfo_ClientInformation.Size(m)
}
func (m *ClientInformation) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientInformation.DiscardUnknown(m)
}

var xxx_messageInfo_ClientInformation proto.InternalMessageInfo

func (m *ClientInformation) GetClientName() string {
	if m != nil {
		return m.ClientName
	}
	return ""
}

func (m *ClientInformation) GetClientVersion() string {
	if m != nil {
		return m.ClientVersion
	}
	return ""
}

func (m *ClientInformation) GetRevision() uint64 {
	if m != nil {
		return m.Revision
	}
	return 0
}

func (m *ClientInformation) GetBuildTime() string {
	if m != nil {
		return m.BuildTime
	}
	return ""
}

func (m *ClientInformation) GetClientDescription() string {
	if m != nil {
		return m.ClientDescription
	}
	return ""
}

func (m *ClientInformation) GetLabels() []string {
	if m != nil {
		return m.Labels
	}
	return nil
}

// Ask the ListDir action to list all files in path (returns StatEntry)
type ListDirRequest struct {
	Pathspec             *PathSpec `protobuf:"bytes,1,opt,name=pathspec,proto3" json:"pathspec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListDirRequest) Reset()         { *m = ListDirRequest{} }
func (m *ListDirRequest) String() string { return proto.CompactTextString(m) }
func (*ListDirRequest) ProtoMessage()    {}
func (*ListDirRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeb49063df94c2b8, []int{1}
}

func (m *ListDirRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListDirRequest.Unmarshal(m, b)
}
func (m *ListDirRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListDirRequest.Marshal(b, m, deterministic)
}
func (m *ListDirRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDirRequest.Merge(m, src)
}
func (m *ListDirRequest) XXX_Size() int {
	return xxx_messageInfo_ListDirRequest.Size(m)
}
func (m *ListDirRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDirRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListDirRequest proto.InternalMessageInfo

func (m *ListDirRequest) GetPathspec() *PathSpec {
	if m != nil {
		return m.Pathspec
	}
	return nil
}

// A stat() record for a given path
// Next field id: 24.
type StatEntry struct {
	// DEPRECATED
	//  string aff4path = 1;
	StMode       uint64                 `protobuf:"varint,2,opt,name=st_mode,json=stMode,proto3" json:"st_mode,omitempty"`
	StIno        uint32                 `protobuf:"varint,3,opt,name=st_ino,json=stIno,proto3" json:"st_ino,omitempty"`
	StDev        uint32                 `protobuf:"varint,4,opt,name=st_dev,json=stDev,proto3" json:"st_dev,omitempty"`
	StNlink      uint32                 `protobuf:"varint,5,opt,name=st_nlink,json=stNlink,proto3" json:"st_nlink,omitempty"`
	StUid        uint32                 `protobuf:"varint,6,opt,name=st_uid,json=stUid,proto3" json:"st_uid,omitempty"`
	StGid        uint32                 `protobuf:"varint,7,opt,name=st_gid,json=stGid,proto3" json:"st_gid,omitempty"`
	StSize       uint64                 `protobuf:"varint,8,opt,name=st_size,json=stSize,proto3" json:"st_size,omitempty"`
	StAtime      uint64                 `protobuf:"varint,9,opt,name=st_atime,json=stAtime,proto3" json:"st_atime,omitempty"`
	StMtime      uint64                 `protobuf:"varint,10,opt,name=st_mtime,json=stMtime,proto3" json:"st_mtime,omitempty"`
	StCtime      uint64                 `protobuf:"varint,11,opt,name=st_ctime,json=stCtime,proto3" json:"st_ctime,omitempty"`
	StBlocks     uint32                 `protobuf:"varint,12,opt,name=st_blocks,json=stBlocks,proto3" json:"st_blocks,omitempty"`
	StBlksize    uint32                 `protobuf:"varint,13,opt,name=st_blksize,json=stBlksize,proto3" json:"st_blksize,omitempty"`
	StRdev       uint32                 `protobuf:"varint,14,opt,name=st_rdev,json=stRdev,proto3" json:"st_rdev,omitempty"`
	StFlagsOsx   uint32                 `protobuf:"varint,21,opt,name=st_flags_osx,json=stFlagsOsx,proto3" json:"st_flags_osx,omitempty"`
	StFlagsLinux uint32                 `protobuf:"varint,22,opt,name=st_flags_linux,json=stFlagsLinux,proto3" json:"st_flags_linux,omitempty"`
	Symlink      string                 `protobuf:"bytes,15,opt,name=symlink,proto3" json:"symlink,omitempty"`
	RegistryType StatEntry_RegistryType `protobuf:"varint,16,opt,name=registry_type,json=registryType,proto3,enum=proto.StatEntry_RegistryType" json:"registry_type,omitempty"`
	// For very small files their data is stored in this record.
	Resident []byte `protobuf:"bytes,17,opt,name=resident,proto3" json:"resident,omitempty"`
	// The pathspec which the client can use to re-reach this file.
	Pathspec *PathSpec `protobuf:"bytes,18,opt,name=pathspec,proto3" json:"pathspec,omitempty"`
	//   DataBlob registry_data = 19 [(sem_type) = {
	//      description: "If this entry represents a registry value, this field will "
	//     "contain that value encoded according to the correct type.";
	//    }];
	StCrtime             uint64               `protobuf:"varint,20,opt,name=st_crtime,json=stCrtime,proto3" json:"st_crtime,omitempty"`
	ExtAttrs             []*StatEntry_ExtAttr `protobuf:"bytes,23,rep,name=ext_attrs,json=extAttrs,proto3" json:"ext_attrs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *StatEntry) Reset()         { *m = StatEntry{} }
func (m *StatEntry) String() string { return proto.CompactTextString(m) }
func (*StatEntry) ProtoMessage()    {}
func (*StatEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeb49063df94c2b8, []int{2}
}

func (m *StatEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatEntry.Unmarshal(m, b)
}
func (m *StatEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatEntry.Marshal(b, m, deterministic)
}
func (m *StatEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatEntry.Merge(m, src)
}
func (m *StatEntry) XXX_Size() int {
	return xxx_messageInfo_StatEntry.Size(m)
}
func (m *StatEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_StatEntry.DiscardUnknown(m)
}

var xxx_messageInfo_StatEntry proto.InternalMessageInfo

func (m *StatEntry) GetStMode() uint64 {
	if m != nil {
		return m.StMode
	}
	return 0
}

func (m *StatEntry) GetStIno() uint32 {
	if m != nil {
		return m.StIno
	}
	return 0
}

func (m *StatEntry) GetStDev() uint32 {
	if m != nil {
		return m.StDev
	}
	return 0
}

func (m *StatEntry) GetStNlink() uint32 {
	if m != nil {
		return m.StNlink
	}
	return 0
}

func (m *StatEntry) GetStUid() uint32 {
	if m != nil {
		return m.StUid
	}
	return 0
}

func (m *StatEntry) GetStGid() uint32 {
	if m != nil {
		return m.StGid
	}
	return 0
}

func (m *StatEntry) GetStSize() uint64 {
	if m != nil {
		return m.StSize
	}
	return 0
}

func (m *StatEntry) GetStAtime() uint64 {
	if m != nil {
		return m.StAtime
	}
	return 0
}

func (m *StatEntry) GetStMtime() uint64 {
	if m != nil {
		return m.StMtime
	}
	return 0
}

func (m *StatEntry) GetStCtime() uint64 {
	if m != nil {
		return m.StCtime
	}
	return 0
}

func (m *StatEntry) GetStBlocks() uint32 {
	if m != nil {
		return m.StBlocks
	}
	return 0
}

func (m *StatEntry) GetStBlksize() uint32 {
	if m != nil {
		return m.StBlksize
	}
	return 0
}

func (m *StatEntry) GetStRdev() uint32 {
	if m != nil {
		return m.StRdev
	}
	return 0
}

func (m *StatEntry) GetStFlagsOsx() uint32 {
	if m != nil {
		return m.StFlagsOsx
	}
	return 0
}

func (m *StatEntry) GetStFlagsLinux() uint32 {
	if m != nil {
		return m.StFlagsLinux
	}
	return 0
}

func (m *StatEntry) GetSymlink() string {
	if m != nil {
		return m.Symlink
	}
	return ""
}

func (m *StatEntry) GetRegistryType() StatEntry_RegistryType {
	if m != nil {
		return m.RegistryType
	}
	return StatEntry_REG_NONE
}

func (m *StatEntry) GetResident() []byte {
	if m != nil {
		return m.Resident
	}
	return nil
}

func (m *StatEntry) GetPathspec() *PathSpec {
	if m != nil {
		return m.Pathspec
	}
	return nil
}

func (m *StatEntry) GetStCrtime() uint64 {
	if m != nil {
		return m.StCrtime
	}
	return 0
}

func (m *StatEntry) GetExtAttrs() []*StatEntry_ExtAttr {
	if m != nil {
		return m.ExtAttrs
	}
	return nil
}

type StatEntry_ExtAttr struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatEntry_ExtAttr) Reset()         { *m = StatEntry_ExtAttr{} }
func (m *StatEntry_ExtAttr) String() string { return proto.CompactTextString(m) }
func (*StatEntry_ExtAttr) ProtoMessage()    {}
func (*StatEntry_ExtAttr) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeb49063df94c2b8, []int{2, 0}
}

func (m *StatEntry_ExtAttr) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatEntry_ExtAttr.Unmarshal(m, b)
}
func (m *StatEntry_ExtAttr) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatEntry_ExtAttr.Marshal(b, m, deterministic)
}
func (m *StatEntry_ExtAttr) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatEntry_ExtAttr.Merge(m, src)
}
func (m *StatEntry_ExtAttr) XXX_Size() int {
	return xxx_messageInfo_StatEntry_ExtAttr.Size(m)
}
func (m *StatEntry_ExtAttr) XXX_DiscardUnknown() {
	xxx_messageInfo_StatEntry_ExtAttr.DiscardUnknown(m)
}

var xxx_messageInfo_StatEntry_ExtAttr proto.InternalMessageInfo

func (m *StatEntry_ExtAttr) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *StatEntry_ExtAttr) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// A message to encode a filesystem path (maybe for raw access)
// Next field: 15
type PathSpec struct {
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Accessor             string   `protobuf:"bytes,3,opt,name=accessor,proto3" json:"accessor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PathSpec) Reset()         { *m = PathSpec{} }
func (m *PathSpec) String() string { return proto.CompactTextString(m) }
func (*PathSpec) ProtoMessage()    {}
func (*PathSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeb49063df94c2b8, []int{3}
}

func (m *PathSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PathSpec.Unmarshal(m, b)
}
func (m *PathSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PathSpec.Marshal(b, m, deterministic)
}
func (m *PathSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PathSpec.Merge(m, src)
}
func (m *PathSpec) XXX_Size() int {
	return xxx_messageInfo_PathSpec.Size(m)
}
func (m *PathSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_PathSpec.DiscardUnknown(m)
}

var xxx_messageInfo_PathSpec proto.InternalMessageInfo

func (m *PathSpec) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *PathSpec) GetAccessor() string {
	if m != nil {
		return m.Accessor
	}
	return ""
}

// Message to carry uname information.
type Uname struct {
	System               string   `protobuf:"bytes,1,opt,name=system,proto3" json:"system,omitempty"`
	Node                 string   `protobuf:"bytes,2,opt,name=node,proto3" json:"node,omitempty"`
	Release              string   `protobuf:"bytes,3,opt,name=release,proto3" json:"release,omitempty"`
	Version              string   `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	Machine              string   `protobuf:"bytes,5,opt,name=machine,proto3" json:"machine,omitempty"`
	Kernel               string   `protobuf:"bytes,6,opt,name=kernel,proto3" json:"kernel,omitempty"`
	Fqdn                 string   `protobuf:"bytes,7,opt,name=fqdn,proto3" json:"fqdn,omitempty"`
	InstallDate          uint64   `protobuf:"varint,8,opt,name=install_date,json=installDate,proto3" json:"install_date,omitempty"`
	LibcVer              string   `protobuf:"bytes,9,opt,name=libc_ver,json=libcVer,proto3" json:"libc_ver,omitempty"`
	Architecture         string   `protobuf:"bytes,10,opt,name=architecture,proto3" json:"architecture,omitempty"`
	Pep425Tag            string   `protobuf:"bytes,11,opt,name=pep425tag,proto3" json:"pep425tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Uname) Reset()         { *m = Uname{} }
func (m *Uname) String() string { return proto.CompactTextString(m) }
func (*Uname) ProtoMessage()    {}
func (*Uname) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeb49063df94c2b8, []int{4}
}

func (m *Uname) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Uname.Unmarshal(m, b)
}
func (m *Uname) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Uname.Marshal(b, m, deterministic)
}
func (m *Uname) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Uname.Merge(m, src)
}
func (m *Uname) XXX_Size() int {
	return xxx_messageInfo_Uname.Size(m)
}
func (m *Uname) XXX_DiscardUnknown() {
	xxx_messageInfo_Uname.DiscardUnknown(m)
}

var xxx_messageInfo_Uname proto.InternalMessageInfo

func (m *Uname) GetSystem() string {
	if m != nil {
		return m.System
	}
	return ""
}

func (m *Uname) GetNode() string {
	if m != nil {
		return m.Node
	}
	return ""
}

func (m *Uname) GetRelease() string {
	if m != nil {
		return m.Release
	}
	return ""
}

func (m *Uname) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Uname) GetMachine() string {
	if m != nil {
		return m.Machine
	}
	return ""
}

func (m *Uname) GetKernel() string {
	if m != nil {
		return m.Kernel
	}
	return ""
}

func (m *Uname) GetFqdn() string {
	if m != nil {
		return m.Fqdn
	}
	return ""
}

func (m *Uname) GetInstallDate() uint64 {
	if m != nil {
		return m.InstallDate
	}
	return 0
}

func (m *Uname) GetLibcVer() string {
	if m != nil {
		return m.LibcVer
	}
	return ""
}

func (m *Uname) GetArchitecture() string {
	if m != nil {
		return m.Architecture
	}
	return ""
}

func (m *Uname) GetPep425Tag() string {
	if m != nil {
		return m.Pep425Tag
	}
	return ""
}

func init() {
	proto.RegisterEnum("proto.StatEntry_RegistryType", StatEntry_RegistryType_name, StatEntry_RegistryType_value)
	proto.RegisterType((*ClientInformation)(nil), "proto.ClientInformation")
	proto.RegisterType((*ListDirRequest)(nil), "proto.ListDirRequest")
	proto.RegisterType((*StatEntry)(nil), "proto.StatEntry")
	proto.RegisterType((*StatEntry_ExtAttr)(nil), "proto.StatEntry.ExtAttr")
	proto.RegisterType((*PathSpec)(nil), "proto.PathSpec")
	proto.RegisterType((*Uname)(nil), "proto.Uname")
}

func init() { proto.RegisterFile("actions.proto", fileDescriptor_eeb49063df94c2b8) }

var fileDescriptor_eeb49063df94c2b8 = []byte{
	// 1720 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x56, 0x4f, 0x73, 0x1b, 0x49,
	0x15, 0x47, 0x89, 0xad, 0x3f, 0xed, 0x3f, 0xb1, 0x9b, 0x64, 0x2d, 0xbc, 0x1b, 0xe8, 0x72, 0x2d,
	0x15, 0x99, 0x38, 0x63, 0x5b, 0x76, 0xbc, 0x21, 0x05, 0x07, 0xcb, 0x52, 0x8c, 0x88, 0x2d, 0x9b,
	0x91, 0xb3, 0x59, 0xf6, 0xa2, 0xb4, 0x66, 0x9e, 0xa4, 0xae, 0xcc, 0x74, 0xcb, 0xdd, 0x2d, 0x4b,
	0xda, 0xda, 0x2a, 0x6a, 0x6f, 0x7c, 0x01, 0xaa, 0xb8, 0xf1, 0x65, 0x72, 0xe5, 0x13, 0x50, 0x5c,
	0xe0, 0x0a, 0xdf, 0x80, 0x03, 0xd5, 0x6f, 0x66, 0x64, 0xb9, 0x30, 0xa7, 0x99, 0x7e, 0x7f, 0x7e,
	0xef, 0x6f, 0xbf, 0xd7, 0x64, 0x85, 0x07, 0x56, 0x28, 0x69, 0xbc, 0xa1, 0x56, 0x56, 0xd1, 0x45,
	0xfc, 0x6c, 0xbe, 0x1e, 0x8f, 0xc7, 0xde, 0x0d, 0x44, 0x2a, 0x10, 0x21, 0x4c, 0xbc, 0x40, 0xc5,
	0xbb, 0x7d, 0x15, 0x71, 0xd9, 0xdf, 0x4d, 0x88, 0x9a, 0x0f, 0xad, 0xd2, 0xbb, 0x28, 0xbc, 0x6b,
	0x20, 0xe6, 0xd2, 0x8a, 0x20, 0x81, 0xd8, 0xfa, 0x7b, 0x8e, 0xac, 0x9f, 0x44, 0x02, 0xa4, 0x6d,
	0xca, 0x9e, 0xd2, 0x31, 0x77, 0xf8, 0xf4, 0x67, 0x64, 0x29, 0x40, 0x62, 0x47, 0xf2, 0x18, 0xca,
	0x39, 0x96, 0xab, 0x94, 0x7c, 0x92, 0x90, 0x5a, 0x3c, 0x06, 0xfa, 0x73, 0xb2, 0x9a, 0x0a, 0xdc,
	0x80, 0x36, 0x42, 0xc9, 0x72, 0x01, 0x65, 0x56, 0x12, 0xea, 0xd7, 0x09, 0x91, 0x6e, 0x92, 0xa2,
	0x86, 0x1b, 0x81, 0x02, 0x0f, 0x59, 0xae, 0xb2, 0xe0, 0xcf, 0xce, 0xf4, 0x29, 0x21, 0xdd, 0x91,
	0x88, 0xc2, 0x8e, 0x15, 0x31, 0x94, 0x17, 0x50, 0xbd, 0x84, 0x94, 0x2b, 0x11, 0x03, 0x7d, 0x41,
	0x68, 0x6a, 0x21, 0x04, 0x13, 0x68, 0x31, 0x74, 0x8e, 0x95, 0x17, 0x51, 0x6c, 0x3d, 0xe1, 0xd4,
	0x6f, 0x19, 0xf4, 0x33, 0x92, 0x8f, 0x78, 0x17, 0x22, 0x53, 0xce, 0xb3, 0x87, 0x95, 0x92, 0x9f,
	0x9e, 0xb6, 0x7e, 0x4d, 0x56, 0xcf, 0x84, 0xb1, 0x75, 0xa1, 0x7d, 0xb8, 0x1e, 0x81, 0xb1, 0xf4,
	0x39, 0x29, 0x0e, 0xb9, 0x1d, 0x98, 0x21, 0x04, 0x18, 0xd8, 0x52, 0xf5, 0x51, 0x92, 0x0b, 0xef,
	0x92, 0xdb, 0x41, 0x7b, 0x08, 0x81, 0x3f, 0x13, 0xd8, 0xfa, 0xf7, 0x12, 0x29, 0xb5, 0x2d, 0xb7,
	0x0d, 0x69, 0xf5, 0x94, 0x9e, 0x92, 0x82, 0xb1, 0x9d, 0x58, 0x85, 0x50, 0x7e, 0xe0, 0xa2, 0xa9,
	0x79, 0xff, 0xf8, 0xcf, 0x3f, 0x3f, 0xe5, 0x2a, 0xa4, 0xe8, 0x64, 0xce, 0x55, 0x08, 0xf4, 0x8b,
	0x63, 0x36, 0x92, 0x62, 0xc2, 0x7a, 0x22, 0x02, 0x36, 0x04, 0x1d, 0x0b, 0xe3, 0xa2, 0x65, 0x4e,
	0xc9, 0xf3, 0xf3, 0x26, 0x91, 0x7a, 0x42, 0xf2, 0xc6, 0x76, 0x84, 0x54, 0x98, 0x95, 0x15, 0x7f,
	0xd1, 0xd8, 0xa6, 0x54, 0x29, 0x39, 0x84, 0x1b, 0x4c, 0x07, 0x92, 0xeb, 0x70, 0x43, 0x7f, 0x42,
	0x8a, 0xc6, 0x76, 0x64, 0x24, 0xe4, 0x47, 0x4c, 0xc0, 0x8a, 0x5f, 0x30, 0xb6, 0xe5, 0x8e, 0xa9,
	0xc6, 0x48, 0x84, 0xe5, 0x7c, 0xa6, 0xf1, 0x4e, 0x84, 0x29, 0xb9, 0x2f, 0x42, 0x2c, 0x0b, 0x92,
	0x4f, 0x45, 0x48, 0x37, 0xd0, 0x7f, 0x23, 0xbe, 0x83, 0x72, 0x11, 0xab, 0x91, 0x37, 0xb6, 0x2d,
	0xbe, 0x03, 0xfa, 0x1b, 0xb4, 0xc0, 0xb1, 0x12, 0x25, 0x8c, 0xec, 0x05, 0x46, 0xf6, 0x8c, 0x50,
	0xbf, 0xfe, 0xa6, 0xce, 0x2d, 0x38, 0x56, 0x1b, 0x02, 0x25, 0x43, 0x43, 0xd7, 0xcf, 0xb8, 0xb1,
	0x8c, 0x07, 0x01, 0x18, 0xc3, 0x1c, 0xc3, 0x73, 0x0e, 0x1d, 0xbb, 0x3f, 0xfa, 0x5b, 0x44, 0x8a,
	0x11, 0x89, 0x20, 0xd2, 0x2e, 0x22, 0x6d, 0xdf, 0x8b, 0xf4, 0x63, 0x44, 0x8a, 0x55, 0x28, 0x7a,
	0x02, 0xc2, 0x5b, 0xac, 0x73, 0xc4, 0x6a, 0x21, 0x56, 0x80, 0x58, 0x4b, 0x88, 0x75, 0x80, 0x58,
	0x2f, 0xee, 0xc5, 0xda, 0x40, 0x2c, 0x21, 0x55, 0x08, 0x2c, 0x18, 0x70, 0xd9, 0x87, 0x5b, 0xbc,
	0x13, 0xc4, 0xfb, 0x9c, 0x94, 0x8c, 0xed, 0x74, 0x23, 0x15, 0x7c, 0x34, 0xe5, 0x65, 0x4c, 0x4c,
	0xd1, 0xd8, 0x1a, 0x9e, 0x5d, 0x3b, 0x22, 0xf3, 0x23, 0xa6, 0x67, 0x05, 0xb9, 0x25, 0xc7, 0x45,
	0x42, 0x9a, 0x3a, 0xed, 0x6a, 0xb3, 0x8a, 0xbc, 0xbc, 0xb1, 0x7e, 0x08, 0x37, 0x54, 0x92, 0x65,
	0x63, 0x3b, 0xbd, 0x88, 0xf7, 0x4d, 0x47, 0x99, 0x49, 0xf9, 0x89, 0xe3, 0xd6, 0xce, 0xd0, 0xd1,
	0x37, 0xe4, 0x11, 0x36, 0xcf, 0xc4, 0xbe, 0x71, 0xfc, 0x0b, 0x33, 0xa1, 0x07, 0x30, 0xb1, 0x20,
	0x43, 0x08, 0x93, 0x0e, 0xe1, 0xd6, 0x6a, 0xd1, 0x1d, 0x59, 0x30, 0xac, 0xa7, 0x34, 0x3b, 0xe7,
	0x01, 0xab, 0x18, 0xb0, 0xac, 0x3b, 0x65, 0x1f, 0x82, 0x01, 0xa2, 0x7e, 0xd8, 0xf6, 0x89, 0xb9,
	0x05, 0xb8, 0x26, 0xab, 0x33, 0x7b, 0x91, 0x90, 0xa3, 0x49, 0xf9, 0x33, 0xb4, 0xf8, 0x16, 0x2d,
	0x36, 0xc8, 0xfa, 0xbc, 0xc5, 0x33, 0x27, 0x40, 0xf7, 0xfe, 0xaf, 0x4d, 0x6e, 0x98, 0x86, 0xa1,
	0xd2, 0x16, 0x42, 0x34, 0x19, 0x19, 0xc7, 0xfc, 0xe0, 0x2f, 0x9b, 0x79, 0xfd, 0x32, 0x29, 0x98,
	0x69, 0x8c, 0xed, 0xf7, 0x08, 0xef, 0x5f, 0x76, 0xa4, 0x7f, 0xcd, 0x91, 0x15, 0x0d, 0x7d, 0x61,
	0xac, 0x9e, 0x76, 0xec, 0x74, 0x08, 0xe5, 0x35, 0x96, 0xab, 0xac, 0x56, 0x9f, 0xa6, 0x37, 0x6a,
	0x76, 0x75, 0x3c, 0x3f, 0x95, 0xba, 0x9a, 0x0e, 0xa1, 0xf6, 0xc7, 0x1c, 0x3a, 0xfb, 0x43, 0x8e,
	0xfe, 0xa1, 0xd9, 0x63, 0x76, 0x20, 0x0c, 0x13, 0x86, 0x71, 0x66, 0x2c, 0xb7, 0x4c, 0xf5, 0x18,
	0x67, 0x19, 0x28, 0xbb, 0xe1, 0xd1, 0x08, 0x76, 0x12, 0x99, 0x9e, 0x80, 0x28, 0x64, 0x81, 0x92,
	0x96, 0x0b, 0x69, 0x98, 0x1d, 0x00, 0x73, 0x36, 0x9d, 0x06, 0xf2, 0x51, 0xd6, 0x63, 0x57, 0x03,
	0x40, 0x21, 0x90, 0x96, 0x8d, 0x45, 0x14, 0x31, 0x1e, 0x19, 0xc5, 0xba, 0xc0, 0x40, 0x06, 0xca,
	0x65, 0x41, 0x48, 0xa6, 0xc1, 0x88, 0x10, 0xac, 0xe7, 0x2f, 0xeb, 0x39, 0xc7, 0x92, 0x79, 0xe5,
	0x38, 0xd2, 0x96, 0xd7, 0x59, 0xae, 0xb2, 0xec, 0xcf, 0xce, 0xb4, 0x3f, 0x37, 0x37, 0xe8, 0xbd,
	0x73, 0xa3, 0xf6, 0x1a, 0xc3, 0x3a, 0xa4, 0x55, 0xe7, 0x82, 0x13, 0x66, 0x4e, 0x5a, 0xf4, 0x44,
	0x80, 0x63, 0x95, 0x59, 0x95, 0x05, 0x12, 0x01, 0x73, 0x67, 0xe7, 0x29, 0x8e, 0x32, 0xef, 0x76,
	0xe6, 0xd0, 0x53, 0x6c, 0xd3, 0x40, 0x63, 0xdf, 0x3f, 0xc6, 0xbe, 0xff, 0x05, 0x02, 0x7f, 0x79,
	0x6f, 0xdf, 0xaf, 0x9e, 0x68, 0x48, 0xe1, 0xb1, 0xdd, 0x8b, 0xc6, 0x9e, 0xa0, 0x2e, 0x7d, 0x49,
	0x4a, 0x30, 0x71, 0xd7, 0xda, 0x6a, 0x53, 0xde, 0x60, 0x0f, 0x2b, 0x4b, 0xd5, 0xf2, 0xff, 0x14,
	0xa6, 0x31, 0xb1, 0xc7, 0xd6, 0x6a, 0xbf, 0x08, 0xc9, 0x8f, 0xd9, 0x3c, 0x20, 0x85, 0x94, 0x48,
	0x29, 0x59, 0x98, 0x5b, 0x00, 0xf8, 0x4f, 0x1f, 0x93, 0x45, 0x4c, 0x34, 0x8e, 0xc0, 0x92, 0x9f,
	0x1c, 0xb6, 0x3e, 0xe5, 0xc8, 0xf2, 0x7c, 0x8d, 0xe9, 0x32, 0x29, 0xfa, 0x8d, 0xd3, 0x4e, 0xeb,
	0xa2, 0xd5, 0x58, 0xfb, 0x11, 0x25, 0x24, 0xef, 0x4e, 0xed, 0x6f, 0xd7, 0x72, 0x74, 0x9d, 0xac,
	0xb8, 0xff, 0xc6, 0x37, 0x97, 0xc7, 0xad, 0xba, 0x23, 0x3d, 0xa0, 0xab, 0x84, 0x38, 0x52, 0xad,
	0xd9, 0x3a, 0xf6, 0x7f, 0xbf, 0xf6, 0x90, 0xae, 0x90, 0x92, 0x3b, 0xd7, 0xdf, 0x5f, 0xf8, 0xf5,
	0xb5, 0x05, 0xfa, 0x39, 0xd9, 0x98, 0x1d, 0x3b, 0x67, 0xcd, 0xab, 0xab, 0xb3, 0x46, 0xa7, 0xd1,
	0xaa, 0x37, 0x8f, 0x5b, 0x6b, 0x0b, 0xb4, 0x4c, 0x1e, 0xdf, 0x32, 0x6b, 0xcd, 0xd3, 0x8c, 0xb3,
	0x98, 0xb9, 0x70, 0xd6, 0x6c, 0xbd, 0x5d, 0xcb, 0xd3, 0x35, 0xb2, 0xec, 0x4e, 0xe7, 0xef, 0xce,
	0xae, 0x9a, 0xce, 0x6a, 0x21, 0xb3, 0xf2, 0x3b, 0xb4, 0xb2, 0xb4, 0xf9, 0x60, 0x2d, 0xb7, 0xf5,
	0xb7, 0x1c, 0x29, 0x66, 0xe5, 0xa4, 0x3f, 0xe4, 0xc8, 0x82, 0xab, 0x4a, 0x12, 0x69, 0x2d, 0xc6,
	0x22, 0xf4, 0x29, 0xcc, 0xaa, 0x3b, 0xe4, 0xc6, 0xb8, 0xb1, 0xa5, 0xb0, 0x92, 0x6f, 0x44, 0x04,
	0x66, 0x6a, 0x2c, 0xc4, 0xe9, 0x6c, 0x54, 0xda, 0x35, 0x63, 0xd6, 0x98, 0xae, 0xc5, 0x85, 0xb4,
	0xa0, 0x87, 0x1a, 0xd2, 0xfb, 0xe6, 0xb4, 0x32, 0x51, 0xd7, 0x95, 0xc2, 0x1a, 0xa6, 0xc6, 0x92,
	0x8d, 0xf9, 0xd4, 0xf3, 0xd1, 0x34, 0x7d, 0x4b, 0x8a, 0x99, 0x04, 0xee, 0x8a, 0x52, 0x36, 0x4f,
	0xe9, 0xb3, 0xab, 0x79, 0xed, 0x51, 0xea, 0x88, 0x06, 0xab, 0x05, 0xdc, 0x00, 0x62, 0xbb, 0x3e,
	0xf3, 0xfc, 0x19, 0xc0, 0xd6, 0xbf, 0x0a, 0x64, 0xf1, 0x1d, 0x16, 0xf1, 0x2d, 0xc9, 0x27, 0xae,
	0x26, 0xa5, 0xcd, 0x06, 0x2b, 0x7d, 0xee, 0x40, 0xd3, 0x20, 0x86, 0x11, 0xb7, 0xee, 0x49, 0xc0,
	0x2a, 0xef, 0x85, 0x0c, 0xd5, 0xd8, 0x7c, 0x5f, 0xe7, 0x7a, 0x2c, 0xe4, 0xf7, 0x38, 0x15, 0xb6,
	0xdd, 0x36, 0x43, 0x41, 0xfa, 0x8a, 0x2c, 0xc8, 0x6c, 0x27, 0x96, 0x6a, 0x5f, 0x22, 0xd4, 0x4f,
	0xe9, 0x17, 0x0e, 0x6a, 0xa0, 0x8c, 0x75, 0x06, 0x67, 0x77, 0x34, 0x51, 0xf1, 0x7c, 0xd4, 0xa0,
	0x97, 0xa4, 0xa0, 0x21, 0x02, 0x6e, 0x20, 0x0d, 0xee, 0x08, 0x95, 0xf7, 0xa8, 0xe7, 0x94, 0x2f,
	0xda, 0x2c, 0xe5, 0x32, 0xbc, 0x82, 0x6e, 0x49, 0x68, 0x06, 0x5e, 0xdf, 0x63, 0x5f, 0xed, 0xb0,
	0x8b, 0xf6, 0x37, 0x3b, 0x2c, 0x84, 0xae, 0xe0, 0xd2, 0xf3, 0x33, 0x18, 0x7a, 0x45, 0x0a, 0xd9,
	0x8b, 0x04, 0x9f, 0x14, 0x77, 0xef, 0xe4, 0x45, 0x9b, 0xa5, 0x5c, 0xd6, 0xac, 0x27, 0x48, 0x47,
	0xde, 0xbe, 0xf7, 0xd5, 0xd1, 0xde, 0x7e, 0xfb, 0x72, 0x7f, 0x87, 0xed, 0xef, 0x79, 0xbf, 0xf4,
	0xaa, 0x3b, 0x6c, 0xff, 0xd0, 0xdb, 0x3b, 0xf4, 0xfc, 0x0c, 0x8a, 0x9e, 0x93, 0x42, 0xcc, 0x83,
	0x81, 0x90, 0x90, 0xbc, 0x40, 0xee, 0xcd, 0x17, 0xd7, 0xc1, 0x40, 0x58, 0x08, 0xec, 0x48, 0x43,
	0x82, 0x7d, 0x7c, 0x5e, 0x3f, 0x3a, 0xdc, 0x61, 0x93, 0x57, 0x47, 0x9d, 0x23, 0x07, 0x97, 0x62,
	0xd0, 0x6f, 0x49, 0xfe, 0x23, 0x68, 0x09, 0x11, 0x6e, 0xed, 0x52, 0xad, 0x86, 0x68, 0xbf, 0xa2,
	0xaf, 0x1d, 0x5a, 0xc2, 0x99, 0xf9, 0x69, 0xac, 0x16, 0xb2, 0x7f, 0xd7, 0xd7, 0x1d, 0xb6, 0x7f,
	0xe0, 0xed, 0x7b, 0x7b, 0x3b, 0xec, 0xc0, 0xdb, 0x7f, 0xf9, 0x42, 0x07, 0x55, 0xcf, 0x4f, 0x11,
	0x69, 0x83, 0x2c, 0xf4, 0xae, 0xc3, 0xf4, 0x3d, 0x56, 0xdb, 0x47, 0xe4, 0xe7, 0x74, 0xfb, 0xd6,
	0xcf, 0x67, 0x86, 0xf5, 0x46, 0x51, 0x34, 0x65, 0xd7, 0x23, 0x1e, 0x25, 0x7b, 0x37, 0x54, 0x31,
	0x17, 0x92, 0xb9, 0x42, 0x79, 0x3e, 0xaa, 0x53, 0x9f, 0x2c, 0x0b, 0x69, 0x2c, 0x8f, 0xa2, 0x4e,
	0xc8, 0x6d, 0xfa, 0x5e, 0x98, 0xed, 0xf2, 0xa5, 0xb9, 0x39, 0x44, 0x37, 0xdf, 0x0f, 0x40, 0x66,
	0x49, 0x18, 0x73, 0xd7, 0xe5, 0xa8, 0x08, 0xa1, 0xe7, 0x2f, 0xa5, 0xff, 0x4e, 0x98, 0xbe, 0x22,
	0xc5, 0x48, 0x74, 0x03, 0xf7, 0x64, 0xc4, 0x57, 0x46, 0xa9, 0xf6, 0x14, 0xf1, 0x36, 0xe8, 0x13,
	0xe7, 0xde, 0x09, 0x8b, 0x44, 0x57, 0x73, 0x37, 0xf4, 0x93, 0xd8, 0xfd, 0x82, 0x13, 0xff, 0x1a,
	0x34, 0x75, 0xd3, 0x65, 0x3e, 0xbd, 0xf8, 0xb4, 0x28, 0xd5, 0xfe, 0x92, 0xec, 0x91, 0x3f, 0xe7,
	0xe8, 0x9f, 0x72, 0x78, 0x1b, 0xe6, 0x2b, 0x90, 0x75, 0x5c, 0x57, 0x48, 0xae, 0xa7, 0x1e, 0xab,
	0xb4, 0x94, 0x85, 0x84, 0x14, 0x70, 0xe9, 0xf6, 0x41, 0x28, 0x7a, 0x3d, 0xd0, 0x6e, 0x4d, 0xf4,
	0xb4, 0x8a, 0xf1, 0xc6, 0xa4, 0x15, 0xba, 0x8b, 0x24, 0xd2, 0x49, 0xed, 0x1a, 0x11, 0xb7, 0xd3,
	0x41, 0x95, 0x75, 0x85, 0x4d, 0x91, 0x99, 0x1e, 0x49, 0xe9, 0x4a, 0xa4, 0x24, 0xe3, 0xec, 0xe8,
	0x10, 0x59, 0x49, 0x36, 0xb6, 0xfd, 0x3b, 0x5e, 0xd3, 0x98, 0x94, 0x86, 0x30, 0x3c, 0xac, 0xbe,
	0xb4, 0xbc, 0x8f, 0x2f, 0x9a, 0x52, 0xed, 0x02, 0x23, 0x68, 0xd2, 0x53, 0xe7, 0xff, 0x48, 0x8a,
	0xeb, 0x11, 0x30, 0x23, 0xfa, 0x92, 0xdf, 0x89, 0x61, 0x38, 0xb5, 0x03, 0x35, 0xcb, 0x71, 0x85,
	0x1b, 0x16, 0x42, 0x4f, 0xc8, 0x64, 0x8a, 0x5c, 0x36, 0x2e, 0x0f, 0xab, 0x2f, 0x99, 0xe5, 0x7d,
	0xb3, 0xed, 0xf9, 0xb7, 0x16, 0xba, 0x79, 0x9c, 0xf5, 0x07, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff,
	0x59, 0x68, 0x8e, 0x72, 0x36, 0x0c, 0x00, 0x00,
}
