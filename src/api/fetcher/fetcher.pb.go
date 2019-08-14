// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fetcher.proto

package fetcher

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type Void struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Void) Reset()         { *m = Void{} }
func (m *Void) String() string { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()    {}
func (*Void) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a1f4b766c7849a, []int{0}
}

func (m *Void) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Void.Unmarshal(m, b)
}
func (m *Void) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Void.Marshal(b, m, deterministic)
}
func (m *Void) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Void.Merge(m, src)
}
func (m *Void) XXX_Size() int {
	return xxx_messageInfo_Void.Size(m)
}
func (m *Void) XXX_DiscardUnknown() {
	xxx_messageInfo_Void.DiscardUnknown(m)
}

var xxx_messageInfo_Void proto.InternalMessageInfo

type Msg struct {
	GuildID              string   `protobuf:"bytes,1,opt,name=guildID,proto3" json:"guildID,omitempty"`
	ChannelID            string   `protobuf:"bytes,2,opt,name=channelID,proto3" json:"channelID,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Msg) Reset()         { *m = Msg{} }
func (m *Msg) String() string { return proto.CompactTextString(m) }
func (*Msg) ProtoMessage()    {}
func (*Msg) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a1f4b766c7849a, []int{1}
}

func (m *Msg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Msg.Unmarshal(m, b)
}
func (m *Msg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Msg.Marshal(b, m, deterministic)
}
func (m *Msg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Msg.Merge(m, src)
}
func (m *Msg) XXX_Size() int {
	return xxx_messageInfo_Msg.Size(m)
}
func (m *Msg) XXX_DiscardUnknown() {
	xxx_messageInfo_Msg.DiscardUnknown(m)
}

var xxx_messageInfo_Msg proto.InternalMessageInfo

func (m *Msg) GetGuildID() string {
	if m != nil {
		return m.GuildID
	}
	return ""
}

func (m *Msg) GetChannelID() string {
	if m != nil {
		return m.ChannelID
	}
	return ""
}

func (m *Msg) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type Guilds struct {
	Guilds               []*Guild `protobuf:"bytes,1,rep,name=guilds,proto3" json:"guilds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Guilds) Reset()         { *m = Guilds{} }
func (m *Guilds) String() string { return proto.CompactTextString(m) }
func (*Guilds) ProtoMessage()    {}
func (*Guilds) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a1f4b766c7849a, []int{2}
}

func (m *Guilds) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Guilds.Unmarshal(m, b)
}
func (m *Guilds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Guilds.Marshal(b, m, deterministic)
}
func (m *Guilds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Guilds.Merge(m, src)
}
func (m *Guilds) XXX_Size() int {
	return xxx_messageInfo_Guilds.Size(m)
}
func (m *Guilds) XXX_DiscardUnknown() {
	xxx_messageInfo_Guilds.DiscardUnknown(m)
}

var xxx_messageInfo_Guilds proto.InternalMessageInfo

func (m *Guilds) GetGuilds() []*Guild {
	if m != nil {
		return m.Guilds
	}
	return nil
}

type Guild struct {
	Channels             []*Channel `protobuf:"bytes,1,rep,name=channels,proto3" json:"channels,omitempty"`
	Members              []*Member  `protobuf:"bytes,2,rep,name=members,proto3" json:"members,omitempty"`
	Roles                []*Role    `protobuf:"bytes,3,rep,name=roles,proto3" json:"roles,omitempty"`
	CreatedAt            string     `protobuf:"bytes,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	Description          string     `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	Icon                 string     `protobuf:"bytes,6,opt,name=icon,proto3" json:"icon,omitempty"`
	Id                   string     `protobuf:"bytes,7,opt,name=id,proto3" json:"id,omitempty"`
	MemberCount          int32      `protobuf:"varint,8,opt,name=memberCount,proto3" json:"memberCount,omitempty"`
	Name                 string     `protobuf:"bytes,9,opt,name=name,proto3" json:"name,omitempty"`
	Region               string     `protobuf:"bytes,10,opt,name=region,proto3" json:"region,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Guild) Reset()         { *m = Guild{} }
func (m *Guild) String() string { return proto.CompactTextString(m) }
func (*Guild) ProtoMessage()    {}
func (*Guild) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a1f4b766c7849a, []int{3}
}

func (m *Guild) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Guild.Unmarshal(m, b)
}
func (m *Guild) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Guild.Marshal(b, m, deterministic)
}
func (m *Guild) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Guild.Merge(m, src)
}
func (m *Guild) XXX_Size() int {
	return xxx_messageInfo_Guild.Size(m)
}
func (m *Guild) XXX_DiscardUnknown() {
	xxx_messageInfo_Guild.DiscardUnknown(m)
}

var xxx_messageInfo_Guild proto.InternalMessageInfo

func (m *Guild) GetChannels() []*Channel {
	if m != nil {
		return m.Channels
	}
	return nil
}

func (m *Guild) GetMembers() []*Member {
	if m != nil {
		return m.Members
	}
	return nil
}

func (m *Guild) GetRoles() []*Role {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *Guild) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Guild) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Guild) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *Guild) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Guild) GetMemberCount() int32 {
	if m != nil {
		return m.MemberCount
	}
	return 0
}

func (m *Guild) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Guild) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

type GuildRequest struct {
	Id                   []string `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
	Gql                  string   `protobuf:"bytes,2,opt,name=gql,proto3" json:"gql,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GuildRequest) Reset()         { *m = GuildRequest{} }
func (m *GuildRequest) String() string { return proto.CompactTextString(m) }
func (*GuildRequest) ProtoMessage()    {}
func (*GuildRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a1f4b766c7849a, []int{4}
}

func (m *GuildRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GuildRequest.Unmarshal(m, b)
}
func (m *GuildRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GuildRequest.Marshal(b, m, deterministic)
}
func (m *GuildRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GuildRequest.Merge(m, src)
}
func (m *GuildRequest) XXX_Size() int {
	return xxx_messageInfo_GuildRequest.Size(m)
}
func (m *GuildRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GuildRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GuildRequest proto.InternalMessageInfo

func (m *GuildRequest) GetId() []string {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *GuildRequest) GetGql() string {
	if m != nil {
		return m.Gql
	}
	return ""
}

type Member struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XRoles               []string `protobuf:"bytes,2,rep,name=_roles,json=Roles,proto3" json:"_roles,omitempty"`
	DisplayName          string   `protobuf:"bytes,3,opt,name=displayName,proto3" json:"displayName,omitempty"`
	JoinedAt             string   `protobuf:"bytes,4,opt,name=joinedAt,proto3" json:"joinedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Member) Reset()         { *m = Member{} }
func (m *Member) String() string { return proto.CompactTextString(m) }
func (*Member) ProtoMessage()    {}
func (*Member) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a1f4b766c7849a, []int{5}
}

func (m *Member) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Member.Unmarshal(m, b)
}
func (m *Member) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Member.Marshal(b, m, deterministic)
}
func (m *Member) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Member.Merge(m, src)
}
func (m *Member) XXX_Size() int {
	return xxx_messageInfo_Member.Size(m)
}
func (m *Member) XXX_DiscardUnknown() {
	xxx_messageInfo_Member.DiscardUnknown(m)
}

var xxx_messageInfo_Member proto.InternalMessageInfo

func (m *Member) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *Member) GetXRoles() []string {
	if m != nil {
		return m.XRoles
	}
	return nil
}

func (m *Member) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Member) GetJoinedAt() string {
	if m != nil {
		return m.JoinedAt
	}
	return ""
}

type User struct {
	Avatar               string   `protobuf:"bytes,1,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Tag                  string   `protobuf:"bytes,3,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a1f4b766c7849a, []int{6}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

type Channel struct {
	CreatedAt            string   `protobuf:"bytes,1,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Type                 string   `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Position             int32    `protobuf:"varint,5,opt,name=position,proto3" json:"position,omitempty"`
	Topic                string   `protobuf:"bytes,6,opt,name=topic,proto3" json:"topic,omitempty"`
	Nsfw                 bool     `protobuf:"varint,7,opt,name=nsfw,proto3" json:"nsfw,omitempty"`
	Bitrate              int32    `protobuf:"varint,8,opt,name=bitrate,proto3" json:"bitrate,omitempty"`
	UserLimit            int32    `protobuf:"varint,9,opt,name=userLimit,proto3" json:"userLimit,omitempty"`
	ParentID             string   `protobuf:"bytes,10,opt,name=parentID,proto3" json:"parentID,omitempty"`
	RateLimitPerUser     int32    `protobuf:"varint,11,opt,name=rateLimitPerUser,proto3" json:"rateLimitPerUser,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Channel) Reset()         { *m = Channel{} }
func (m *Channel) String() string { return proto.CompactTextString(m) }
func (*Channel) ProtoMessage()    {}
func (*Channel) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a1f4b766c7849a, []int{7}
}

func (m *Channel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Channel.Unmarshal(m, b)
}
func (m *Channel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Channel.Marshal(b, m, deterministic)
}
func (m *Channel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Channel.Merge(m, src)
}
func (m *Channel) XXX_Size() int {
	return xxx_messageInfo_Channel.Size(m)
}
func (m *Channel) XXX_DiscardUnknown() {
	xxx_messageInfo_Channel.DiscardUnknown(m)
}

var xxx_messageInfo_Channel proto.InternalMessageInfo

func (m *Channel) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *Channel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Channel) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Channel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Channel) GetPosition() int32 {
	if m != nil {
		return m.Position
	}
	return 0
}

func (m *Channel) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *Channel) GetNsfw() bool {
	if m != nil {
		return m.Nsfw
	}
	return false
}

func (m *Channel) GetBitrate() int32 {
	if m != nil {
		return m.Bitrate
	}
	return 0
}

func (m *Channel) GetUserLimit() int32 {
	if m != nil {
		return m.UserLimit
	}
	return 0
}

func (m *Channel) GetParentID() string {
	if m != nil {
		return m.ParentID
	}
	return ""
}

func (m *Channel) GetRateLimitPerUser() int32 {
	if m != nil {
		return m.RateLimitPerUser
	}
	return 0
}

type Role struct {
	Color                int32    `protobuf:"varint,1,opt,name=color,proto3" json:"color,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Position             int32    `protobuf:"varint,3,opt,name=position,proto3" json:"position,omitempty"`
	Id                   string   `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	Permissions          int32    `protobuf:"varint,5,opt,name=permissions,proto3" json:"permissions,omitempty"`
	Hoist                bool     `protobuf:"varint,6,opt,name=hoist,proto3" json:"hoist,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Role) Reset()         { *m = Role{} }
func (m *Role) String() string { return proto.CompactTextString(m) }
func (*Role) ProtoMessage()    {}
func (*Role) Descriptor() ([]byte, []int) {
	return fileDescriptor_b9a1f4b766c7849a, []int{8}
}

func (m *Role) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Role.Unmarshal(m, b)
}
func (m *Role) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Role.Marshal(b, m, deterministic)
}
func (m *Role) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Role.Merge(m, src)
}
func (m *Role) XXX_Size() int {
	return xxx_messageInfo_Role.Size(m)
}
func (m *Role) XXX_DiscardUnknown() {
	xxx_messageInfo_Role.DiscardUnknown(m)
}

var xxx_messageInfo_Role proto.InternalMessageInfo

func (m *Role) GetColor() int32 {
	if m != nil {
		return m.Color
	}
	return 0
}

func (m *Role) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Role) GetPosition() int32 {
	if m != nil {
		return m.Position
	}
	return 0
}

func (m *Role) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Role) GetPermissions() int32 {
	if m != nil {
		return m.Permissions
	}
	return 0
}

func (m *Role) GetHoist() bool {
	if m != nil {
		return m.Hoist
	}
	return false
}

func init() {
	proto.RegisterType((*Void)(nil), "fetcher.Void")
	proto.RegisterType((*Msg)(nil), "fetcher.Msg")
	proto.RegisterType((*Guilds)(nil), "fetcher.Guilds")
	proto.RegisterType((*Guild)(nil), "fetcher.Guild")
	proto.RegisterType((*GuildRequest)(nil), "fetcher.GuildRequest")
	proto.RegisterType((*Member)(nil), "fetcher.Member")
	proto.RegisterType((*User)(nil), "fetcher.User")
	proto.RegisterType((*Channel)(nil), "fetcher.Channel")
	proto.RegisterType((*Role)(nil), "fetcher.Role")
}

func init() { proto.RegisterFile("fetcher.proto", fileDescriptor_b9a1f4b766c7849a) }

var fileDescriptor_b9a1f4b766c7849a = []byte{
	// 637 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0x6d, 0x6b, 0x14, 0x31,
	0x10, 0xee, 0xbe, 0xde, 0x75, 0xae, 0x2f, 0x47, 0xb0, 0x12, 0x0e, 0x3f, 0x9c, 0xab, 0x48, 0x15,
	0x29, 0xa5, 0x22, 0x7e, 0x55, 0x5a, 0x94, 0x82, 0x15, 0x09, 0xa8, 0x1f, 0x65, 0xbb, 0x97, 0x6e,
	0x23, 0x7b, 0x9b, 0x6d, 0x92, 0x53, 0xee, 0x9b, 0x3f, 0xc2, 0x9f, 0xa0, 0xbf, 0xc1, 0xbf, 0x27,
	0x99, 0x64, 0x5f, 0x7a, 0xd2, 0x6f, 0x79, 0x9e, 0x99, 0xcc, 0x64, 0x9e, 0x99, 0x0c, 0xec, 0x5e,
	0x71, 0x53, 0x5c, 0x73, 0x75, 0xd4, 0x28, 0x69, 0x24, 0x19, 0x79, 0x98, 0xa5, 0x10, 0x7f, 0x96,
	0x62, 0x91, 0x7d, 0x81, 0xe8, 0x42, 0x97, 0x84, 0xc2, 0xa8, 0x5c, 0x89, 0x6a, 0x71, 0x7e, 0x46,
	0x83, 0x79, 0x70, 0xb8, 0xcd, 0x5a, 0x48, 0x1e, 0xc0, 0x76, 0x71, 0x9d, 0xd7, 0x35, 0xaf, 0xce,
	0xcf, 0x68, 0x88, 0xb6, 0x9e, 0xb0, 0xf7, 0x0a, 0x59, 0x1b, 0x5e, 0x1b, 0x1a, 0xb9, 0x7b, 0x1e,
	0x66, 0xc7, 0x90, 0xbe, 0xb3, 0x21, 0x34, 0x79, 0x02, 0x29, 0x06, 0xd3, 0x34, 0x98, 0x47, 0x87,
	0x93, 0x93, 0xbd, 0xa3, 0xf6, 0x4d, 0xe8, 0xc0, 0xbc, 0x35, 0xfb, 0x1b, 0x42, 0x82, 0x0c, 0x79,
	0x0e, 0x63, 0x9f, 0xa2, 0xbd, 0x33, 0xed, 0xee, 0x9c, 0x3a, 0x03, 0xeb, 0x3c, 0xc8, 0x53, 0x18,
	0x2d, 0xf9, 0xf2, 0x92, 0x2b, 0x4d, 0x43, 0x74, 0xde, 0xef, 0x9c, 0x2f, 0x90, 0x67, 0xad, 0x9d,
	0x3c, 0x82, 0x44, 0xc9, 0x8a, 0x6b, 0x1a, 0xa1, 0xe3, 0x6e, 0xe7, 0xc8, 0x64, 0xc5, 0x99, 0xb3,
	0x61, 0xc5, 0x8a, 0xe7, 0x86, 0x2f, 0xde, 0x18, 0x1a, 0xfb, 0x8a, 0x5b, 0x82, 0xcc, 0x61, 0xb2,
	0xe0, 0xba, 0x50, 0xa2, 0x31, 0x42, 0xd6, 0x34, 0x41, 0xfb, 0x90, 0x22, 0x04, 0x62, 0x51, 0xc8,
	0x9a, 0xa6, 0x68, 0xc2, 0x33, 0xd9, 0x83, 0x50, 0x2c, 0xe8, 0x08, 0x99, 0x50, 0x2c, 0x6c, 0x14,
	0xf7, 0xa6, 0x53, 0xb9, 0xaa, 0x0d, 0x1d, 0xcf, 0x83, 0xc3, 0x84, 0x0d, 0x29, 0x1b, 0xa5, 0xce,
	0x97, 0x9c, 0x6e, 0xbb, 0x28, 0xf6, 0x4c, 0xee, 0x43, 0xaa, 0x78, 0x69, 0xd3, 0x02, 0xb2, 0x1e,
	0x65, 0xc7, 0xb0, 0xe3, 0xa4, 0xe4, 0x37, 0x2b, 0xae, 0x8d, 0xcf, 0x66, 0x95, 0x73, 0xd9, 0xa6,
	0x10, 0x95, 0x37, 0x95, 0xef, 0x9e, 0x3d, 0x66, 0x3f, 0x03, 0x48, 0x9d, 0x38, 0xe4, 0x21, 0xc4,
	0x2b, 0xcd, 0x15, 0xf6, 0x7d, 0x28, 0xc9, 0x27, 0xcd, 0x15, 0x43, 0x13, 0x39, 0x80, 0xf4, 0xab,
	0xd3, 0x2d, 0xc4, 0x98, 0x09, 0x43, 0xa1, 0xac, 0x14, 0x42, 0x37, 0x55, 0xbe, 0xfe, 0x60, 0x5f,
	0x1a, 0x79, 0x29, 0x7a, 0x8a, 0xcc, 0x60, 0xfc, 0x4d, 0x8a, 0x7a, 0xa0, 0x64, 0x87, 0xb3, 0xd7,
	0x10, 0xdb, 0x14, 0xb6, 0xa8, 0xfc, 0x7b, 0x6e, 0x72, 0xe5, 0x27, 0xcf, 0x23, 0x5f, 0x44, 0xd8,
	0x49, 0x36, 0x85, 0xc8, 0xe4, 0xa5, 0xcf, 0x62, 0x8f, 0xd9, 0x9f, 0x10, 0x46, 0x7e, 0x1c, 0x6e,
	0x37, 0x2d, 0xd8, 0x6c, 0xda, 0x66, 0x2c, 0x02, 0xb1, 0x59, 0x37, 0xed, 0x93, 0xf1, 0xdc, 0x09,
	0x1e, 0x0f, 0x04, 0x9f, 0xc1, 0xb8, 0x91, 0x5a, 0x74, 0x9d, 0x4e, 0x58, 0x87, 0xc9, 0x3d, 0x48,
	0x8c, 0x6c, 0x44, 0xe1, 0xfb, 0xec, 0x00, 0x46, 0xd1, 0x57, 0x3f, 0xb0, 0xd5, 0x63, 0x86, 0x67,
	0xfb, 0x49, 0x2e, 0x85, 0x51, 0xb9, 0xe1, 0xbe, 0xd1, 0x2d, 0xb4, 0xaf, 0xb6, 0x02, 0xbf, 0x17,
	0x4b, 0x61, 0xb0, 0xd3, 0x09, 0xeb, 0x09, 0xcc, 0x9e, 0x2b, 0x5e, 0x9b, 0xf3, 0x33, 0xdf, 0xf0,
	0x0e, 0x93, 0x67, 0x30, 0xb5, 0x11, 0xd0, 0xf1, 0x23, 0x57, 0x56, 0x49, 0x3a, 0xc1, 0x00, 0xff,
	0xf1, 0xd9, 0xaf, 0x00, 0x62, 0xdb, 0x31, 0xfb, 0xe4, 0x42, 0x56, 0xd2, 0x29, 0x9d, 0x30, 0x07,
	0xba, 0xc2, 0xc3, 0x3b, 0x0a, 0x8f, 0x36, 0x0a, 0x77, 0x62, 0xc6, 0xc3, 0x59, 0x6e, 0xb8, 0x5a,
	0x0a, 0xad, 0x85, 0xac, 0xb5, 0xd7, 0x69, 0x48, 0xd9, 0xbc, 0xd7, 0x52, 0x68, 0x83, 0x52, 0x8d,
	0x99, 0x03, 0x27, 0xbf, 0x03, 0x3f, 0xb6, 0x6f, 0xdd, 0xc4, 0x91, 0x57, 0x30, 0xc1, 0xe1, 0xf3,
	0x7b, 0xe3, 0x60, 0x63, 0x4f, 0xb8, 0xe1, 0x9e, 0xed, 0xdf, 0xa6, 0x75, 0xb6, 0x45, 0x5e, 0x02,
	0xf4, 0x17, 0xef, 0xba, 0xb7, 0xb1, 0x76, 0xb2, 0x2d, 0xf2, 0x18, 0x22, 0x9d, 0xaf, 0xc9, 0x4e,
	0xbf, 0x2e, 0x74, 0x39, 0xeb, 0x3f, 0x00, 0xee, 0xc7, 0xad, 0xcb, 0x14, 0x37, 0xe7, 0x8b, 0x7f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x22, 0x4b, 0xc3, 0x4a, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GuildFetcherClient is the client API for GuildFetcher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GuildFetcherClient interface {
	FetchGuilds(ctx context.Context, in *GuildRequest, opts ...grpc.CallOption) (*Guilds, error)
	FetchGuild(ctx context.Context, in *GuildRequest, opts ...grpc.CallOption) (*Guild, error)
	Say(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Void, error)
}

type guildFetcherClient struct {
	cc *grpc.ClientConn
}

func NewGuildFetcherClient(cc *grpc.ClientConn) GuildFetcherClient {
	return &guildFetcherClient{cc}
}

func (c *guildFetcherClient) FetchGuilds(ctx context.Context, in *GuildRequest, opts ...grpc.CallOption) (*Guilds, error) {
	out := new(Guilds)
	err := c.cc.Invoke(ctx, "/fetcher.GuildFetcher/fetchGuilds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guildFetcherClient) FetchGuild(ctx context.Context, in *GuildRequest, opts ...grpc.CallOption) (*Guild, error) {
	out := new(Guild)
	err := c.cc.Invoke(ctx, "/fetcher.GuildFetcher/fetchGuild", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *guildFetcherClient) Say(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/fetcher.GuildFetcher/say", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GuildFetcherServer is the server API for GuildFetcher service.
type GuildFetcherServer interface {
	FetchGuilds(context.Context, *GuildRequest) (*Guilds, error)
	FetchGuild(context.Context, *GuildRequest) (*Guild, error)
	Say(context.Context, *Msg) (*Void, error)
}

// UnimplementedGuildFetcherServer can be embedded to have forward compatible implementations.
type UnimplementedGuildFetcherServer struct {
}

func (*UnimplementedGuildFetcherServer) FetchGuilds(ctx context.Context, req *GuildRequest) (*Guilds, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchGuilds not implemented")
}
func (*UnimplementedGuildFetcherServer) FetchGuild(ctx context.Context, req *GuildRequest) (*Guild, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchGuild not implemented")
}
func (*UnimplementedGuildFetcherServer) Say(ctx context.Context, req *Msg) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Say not implemented")
}

func RegisterGuildFetcherServer(s *grpc.Server, srv GuildFetcherServer) {
	s.RegisterService(&_GuildFetcher_serviceDesc, srv)
}

func _GuildFetcher_FetchGuilds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuildFetcherServer).FetchGuilds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fetcher.GuildFetcher/FetchGuilds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuildFetcherServer).FetchGuilds(ctx, req.(*GuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuildFetcher_FetchGuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GuildRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuildFetcherServer).FetchGuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fetcher.GuildFetcher/FetchGuild",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuildFetcherServer).FetchGuild(ctx, req.(*GuildRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GuildFetcher_Say_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Msg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GuildFetcherServer).Say(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fetcher.GuildFetcher/Say",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GuildFetcherServer).Say(ctx, req.(*Msg))
	}
	return interceptor(ctx, in, info, handler)
}

var _GuildFetcher_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fetcher.GuildFetcher",
	HandlerType: (*GuildFetcherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "fetchGuilds",
			Handler:    _GuildFetcher_FetchGuilds_Handler,
		},
		{
			MethodName: "fetchGuild",
			Handler:    _GuildFetcher_FetchGuild_Handler,
		},
		{
			MethodName: "say",
			Handler:    _GuildFetcher_Say_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fetcher.proto",
}