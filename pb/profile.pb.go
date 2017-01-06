// Code generated by protoc-gen-go.
// source: profile.proto
// DO NOT EDIT!

/*
Package profile is a generated protocol buffer package.

It is generated from these files:
	profile.proto

It has these top-level messages:
	Profile
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Profile struct {
	Handle             string                     `protobuf:"bytes,1,opt,name=handle" json:"handle,omitempty"`
	Name               string                     `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Location           string                     `protobuf:"bytes,3,opt,name=location" json:"location,omitempty"`
	About              string                     `protobuf:"bytes,4,opt,name=about" json:"about,omitempty"`
	ShortDescription   string                     `protobuf:"bytes,5,opt,name=shortDescription" json:"shortDescription,omitempty"`
	Website            string                     `protobuf:"bytes,6,opt,name=website" json:"website,omitempty"`
	Email              string                     `protobuf:"bytes,7,opt,name=email" json:"email,omitempty"`
	PhoneNumber        string                     `protobuf:"bytes,8,opt,name=phoneNumber" json:"phoneNumber,omitempty"`
	Social             []*Profile_SocialAccount   `protobuf:"bytes,9,rep,name=social" json:"social,omitempty"`
	AvgRating          uint32                     `protobuf:"varint,10,opt,name=avgRating" json:"avgRating,omitempty"`
	NumRatings         uint32                     `protobuf:"varint,11,opt,name=numRatings" json:"numRatings,omitempty"`
	Nsfw               bool                       `protobuf:"varint,12,opt,name=nsfw" json:"nsfw,omitempty"`
	Vendor             bool                       `protobuf:"varint,13,opt,name=vendor" json:"vendor,omitempty"`
	Moderator          bool                       `protobuf:"varint,14,opt,name=moderator" json:"moderator,omitempty"`
	ModInfo            *Moderator                 `protobuf:"bytes,15,opt,name=modInfo" json:"modInfo,omitempty"`
	PrimaryColor       string                     `protobuf:"bytes,16,opt,name=primaryColor" json:"primaryColor,omitempty"`
	SecondaryColor     string                     `protobuf:"bytes,17,opt,name=secondaryColor" json:"secondaryColor,omitempty"`
	TextColor          string                     `protobuf:"bytes,18,opt,name=textColor" json:"textColor,omitempty"`
	HighlightColor     string                     `protobuf:"bytes,19,opt,name=highlightColor" json:"highlightColor,omitempty"`
	HighlightTextColor string                     `protobuf:"bytes,20,opt,name=highlightTextColor" json:"highlightTextColor,omitempty"`
	AvatarHashes       *Profile_Image             `protobuf:"bytes,21,opt,name=avatarHashes" json:"avatarHashes,omitempty"`
	HeaderHashes       *Profile_Image             `protobuf:"bytes,22,opt,name=headerHashes" json:"headerHashes,omitempty"`
	FollowerCount      uint32                     `protobuf:"varint,23,opt,name=followerCount" json:"followerCount,omitempty"`
	FollowingCount     uint32                     `protobuf:"varint,24,opt,name=followingCount" json:"followingCount,omitempty"`
	ListingCount       uint32                     `protobuf:"varint,25,opt,name=listingCount" json:"listingCount,omitempty"`
	LastModified       *google_protobuf.Timestamp `protobuf:"bytes,26,opt,name=lastModified" json:"lastModified,omitempty"`
}

func (m *Profile) Reset()                    { *m = Profile{} }
func (m *Profile) String() string            { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()               {}
func (*Profile) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Profile) GetSocial() []*Profile_SocialAccount {
	if m != nil {
		return m.Social
	}
	return nil
}

func (m *Profile) GetModInfo() *Moderator {
	if m != nil {
		return m.ModInfo
	}
	return nil
}

func (m *Profile) GetAvatarHashes() *Profile_Image {
	if m != nil {
		return m.AvatarHashes
	}
	return nil
}

func (m *Profile) GetHeaderHashes() *Profile_Image {
	if m != nil {
		return m.HeaderHashes
	}
	return nil
}

func (m *Profile) GetLastModified() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastModified
	}
	return nil
}

type Profile_SocialAccount struct {
	Type     string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Proof    string `protobuf:"bytes,3,opt,name=proof" json:"proof,omitempty"`
}

func (m *Profile_SocialAccount) Reset()                    { *m = Profile_SocialAccount{} }
func (m *Profile_SocialAccount) String() string            { return proto.CompactTextString(m) }
func (*Profile_SocialAccount) ProtoMessage()               {}
func (*Profile_SocialAccount) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 0} }

type Profile_Image struct {
	Tiny     string `protobuf:"bytes,1,opt,name=tiny" json:"tiny,omitempty"`
	Small    string `protobuf:"bytes,2,opt,name=small" json:"small,omitempty"`
	Medium   string `protobuf:"bytes,3,opt,name=medium" json:"medium,omitempty"`
	Large    string `protobuf:"bytes,4,opt,name=large" json:"large,omitempty"`
	Original string `protobuf:"bytes,5,opt,name=original" json:"original,omitempty"`
}

func (m *Profile_Image) Reset()                    { *m = Profile_Image{} }
func (m *Profile_Image) String() string            { return proto.CompactTextString(m) }
func (*Profile_Image) ProtoMessage()               {}
func (*Profile_Image) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 1} }

func init() {
	proto.RegisterType((*Profile)(nil), "Profile")
	proto.RegisterType((*Profile_SocialAccount)(nil), "Profile.SocialAccount")
	proto.RegisterType((*Profile_Image)(nil), "Profile.Image")
}

var fileDescriptor3 = []byte{
	// 588 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x53, 0x5f, 0x6f, 0xd3, 0x3e,
	0x14, 0xd5, 0x7e, 0xdb, 0xda, 0xce, 0x5d, 0xb7, 0xfd, 0xcc, 0x18, 0x26, 0x42, 0x30, 0x4d, 0x08,
	0x4d, 0x3c, 0x64, 0xd2, 0x78, 0x47, 0x42, 0xe3, 0x81, 0x3d, 0x80, 0x50, 0x18, 0x1f, 0xc0, 0x4d,
	0x9c, 0xc4, 0x92, 0xff, 0x44, 0xb6, 0xd3, 0x52, 0xf1, 0xb1, 0xf9, 0x02, 0xd8, 0xd7, 0x49, 0xda,
	0x8c, 0x3d, 0x54, 0xf5, 0x39, 0xf7, 0xf8, 0xc4, 0xbe, 0xd7, 0x07, 0x2d, 0x1a, 0xa3, 0x4b, 0x2e,
	0x58, 0xea, 0xff, 0x9d, 0x4e, 0xde, 0x54, 0x5a, 0x57, 0x82, 0xdd, 0x00, 0x5a, 0xb6, 0xe5, 0x8d,
	0xe3, 0x92, 0x59, 0x47, 0x65, 0xd3, 0x09, 0x4e, 0xa5, 0x2e, 0x98, 0xa1, 0x4e, 0x9b, 0x48, 0x5c,
	0xfd, 0x99, 0xa1, 0xe9, 0xf7, 0xe8, 0x81, 0x2f, 0xd0, 0xa4, 0xa6, 0xaa, 0x10, 0x8c, 0xec, 0x5d,
	0xee, 0x5d, 0x1f, 0x65, 0x1d, 0xc2, 0x18, 0x1d, 0x28, 0x2a, 0x19, 0xf9, 0x0f, 0x58, 0x58, 0xe3,
	0x04, 0xcd, 0x84, 0xce, 0xa9, 0xe3, 0x5a, 0x91, 0x7d, 0xe0, 0x07, 0x8c, 0xcf, 0xd1, 0x21, 0x5d,
	0xea, 0xd6, 0x91, 0x03, 0x28, 0x44, 0x80, 0xdf, 0xa3, 0x33, 0x5b, 0x6b, 0xe3, 0x3e, 0x33, 0x9b,
	0x1b, 0xde, 0xc0, 0xce, 0x43, 0x10, 0xfc, 0xc3, 0x63, 0x82, 0xa6, 0x6b, 0xb6, 0xb4, 0xdc, 0x31,
	0x32, 0x01, 0x49, 0x0f, 0x83, 0x37, 0x93, 0x94, 0x0b, 0x32, 0x8d, 0xde, 0x00, 0xf0, 0x25, 0x9a,
	0x37, 0xb5, 0x56, 0xec, 0x5b, 0x2b, 0x97, 0xcc, 0x90, 0x19, 0xd4, 0x76, 0x29, 0x9c, 0xa2, 0x89,
	0xd5, 0x39, 0xa7, 0x82, 0x1c, 0x5d, 0xee, 0x5f, 0xcf, 0x6f, 0x2f, 0xd2, 0xee, 0xd6, 0xe9, 0x0f,
	0xa0, 0x3f, 0xe5, 0xb9, 0x6e, 0x95, 0xcb, 0x3a, 0x15, 0x7e, 0x85, 0x8e, 0xe8, 0xaa, 0xca, 0xfc,
	0x85, 0x54, 0x45, 0x90, 0xf7, 0x5b, 0x64, 0x5b, 0x02, 0xbf, 0x46, 0x48, 0xb5, 0x32, 0x02, 0x4b,
	0xe6, 0x50, 0xde, 0x61, 0xa0, 0x63, 0xb6, 0x5c, 0x93, 0x63, 0x5f, 0x99, 0x65, 0xb0, 0x0e, 0xdd,
	0x5d, 0x31, 0x55, 0x68, 0x43, 0x16, 0xc0, 0x76, 0x28, 0x7c, 0x69, 0x18, 0x0a, 0x39, 0x81, 0xd2,
	0x96, 0xc0, 0x6f, 0xd1, 0xd4, 0x83, 0x7b, 0x55, 0x6a, 0x72, 0xea, 0x6b, 0xf3, 0x5b, 0x94, 0x7e,
	0xed, 0x8b, 0x59, 0x5f, 0xc2, 0x57, 0xe8, 0xb8, 0x31, 0x5c, 0x52, 0xb3, 0xb9, 0xd3, 0xc2, 0xdb,
	0x9c, 0x41, 0x03, 0x46, 0x1c, 0x7e, 0x87, 0x4e, 0x2c, 0xcb, 0xb5, 0x2a, 0x06, 0xd5, 0xff, 0xa0,
	0x7a, 0xc4, 0x86, 0xf3, 0x38, 0xf6, 0xcb, 0x45, 0x09, 0x06, 0xc9, 0x96, 0x08, 0x2e, 0x35, 0xaf,
	0x6a, 0xe1, 0x7f, 0x9d, 0xe4, 0x59, 0x74, 0x19, 0xb3, 0xbe, 0xdf, 0x78, 0x60, 0x1e, 0x06, 0xbb,
	0x73, 0xd0, 0x3e, 0x51, 0xc1, 0xb7, 0xe8, 0x98, 0xae, 0xa8, 0xa3, 0xe6, 0x0b, 0xb5, 0x35, 0xb3,
	0xe4, 0x39, 0x5c, 0xf6, 0x64, 0x98, 0xd2, 0xbd, 0xa4, 0x15, 0xcb, 0x46, 0x9a, 0xb0, 0xa7, 0x66,
	0xd4, 0x37, 0xa3, 0xdb, 0x73, 0xf1, 0xf4, 0x9e, 0x5d, 0x8d, 0xef, 0xe7, 0xa2, 0xd4, 0x42, 0xe8,
	0x35, 0x33, 0x77, 0x61, 0xe0, 0xe4, 0x05, 0x0c, 0x6f, 0x4c, 0x86, 0x5b, 0x46, 0xc2, 0x4f, 0x33,
	0xca, 0x08, 0xc8, 0x1e, 0xb1, 0xa1, 0xef, 0x82, 0x5b, 0x37, 0xa8, 0x5e, 0x82, 0x6a, 0xc4, 0xe1,
	0x8f, 0x5e, 0x43, 0xad, 0xf3, 0x53, 0xe3, 0x25, 0x67, 0x05, 0x49, 0xe0, 0x94, 0x49, 0x1a, 0xa3,
	0x9a, 0xf6, 0x51, 0x4d, 0x1f, 0xfa, 0xa8, 0x66, 0x23, 0x7d, 0xf2, 0x13, 0x2d, 0x46, 0x4f, 0x34,
	0x3c, 0x2e, 0xb7, 0x69, 0xfa, 0x90, 0xc2, 0x3a, 0xc4, 0xb1, 0xb5, 0xcc, 0xec, 0xc4, 0x74, 0xc0,
	0x21, 0x32, 0xfe, 0x23, 0xba, 0xec, 0x72, 0x1a, 0x41, 0xf2, 0x1b, 0x1d, 0x42, 0x7f, 0xc0, 0x8e,
	0xab, 0xcd, 0x60, 0xe7, 0xd7, 0x61, 0x8b, 0x95, 0x54, 0x88, 0xce, 0x2b, 0x82, 0xf0, 0x82, 0x25,
	0x2b, 0x78, 0x2b, 0x3b, 0xa7, 0x0e, 0x05, 0xb5, 0xa0, 0xa6, 0x62, 0x7d, 0xde, 0x01, 0x84, 0x23,
	0x69, 0xc3, 0x2b, 0xae, 0x7c, 0xe6, 0x62, 0xce, 0x07, 0xbc, 0x9c, 0xc0, 0xad, 0x3f, 0xfc, 0x0d,
	0x00, 0x00, 0xff, 0xff, 0xf6, 0x28, 0x37, 0x8f, 0xbf, 0x04, 0x00, 0x00,
}
