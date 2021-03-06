// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types.proto

/*
Package identity is a generated protocol buffer package.

It is generated from these files:
	types.proto

It has these top-level messages:
	Group
	Entity
	Alias
*/
package identity

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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Group represents an identity group.
type Group struct {
	// ID is the unique identifier for this group
	ID string `sentinel:"" protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// Name is the unique name for this group
	Name string `sentinel:"" protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// Policies are the vault policies to be granted to members of this group
	Policies []string `sentinel:"" protobuf:"bytes,3,rep,name=policies" json:"policies,omitempty"`
	// ParentGroupIDs are the identifiers of those groups to which this group is a
	// member of. These will serve as references to the parent group in the
	// hierarchy.
	ParentGroupIDs []string `sentinel:"" protobuf:"bytes,4,rep,name=parent_group_ids,json=parentGroupIds" json:"parent_group_ids,omitempty"`
	// MemberEntityIDs are the identifiers of entities which are members of this
	// group
	MemberEntityIDs []string `sentinel:"" protobuf:"bytes,5,rep,name=member_entity_ids,json=memberEntityIDs" json:"member_entity_ids,omitempty"`
	// Metadata represents the custom data tied with this group
	Metadata map[string]string `sentinel:"" protobuf:"bytes,6,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// CreationTime is the time at which this group was created
	CreationTime *google_protobuf.Timestamp `sentinel:"" protobuf:"bytes,7,opt,name=creation_time,json=creationTime" json:"creation_time,omitempty"`
	// LastUpdateTime is the time at which this group was last modified
	LastUpdateTime *google_protobuf.Timestamp `sentinel:"" protobuf:"bytes,8,opt,name=last_update_time,json=lastUpdateTime" json:"last_update_time,omitempty"`
	// ModifyIndex tracks the number of updates to the group. It is useful to detect
	// updates to the groups.
	ModifyIndex uint64 `sentinel:"" protobuf:"varint,9,opt,name=modify_index,json=modifyIndex" json:"modify_index,omitempty"`
	// BucketKeyHash is the MD5 hash of the storage bucket key into which this
	// group is stored in the underlying storage. This is useful to find all
	// the groups belonging to a particular bucket during invalidation of the
	// storage key.
	BucketKeyHash string `sentinel:"" protobuf:"bytes,10,opt,name=bucket_key_hash,json=bucketKeyHash" json:"bucket_key_hash,omitempty"`
	// Alias is used to mark this group as an internal mapping of a group that
	// is external to the identity store. Alias can only be set if the 'type'
	// is set to 'external'.
	Alias *Alias `sentinel:"" protobuf:"bytes,11,opt,name=alias" json:"alias,omitempty"`
	// Type indicates if this group is an internal group or an external group.
	// Memberships of the internal groups can be managed over the API whereas
	// the memberships on the external group --for which a corresponding alias
	// will be set-- will be managed automatically.
	Type string `sentinel:"" protobuf:"bytes,12,opt,name=type" json:"type,omitempty"`
}

func (m *Group) Reset()                    { *m = Group{} }
func (m *Group) String() string            { return proto.CompactTextString(m) }
func (*Group) ProtoMessage()               {}
func (*Group) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Group) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Group) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Group) GetPolicies() []string {
	if m != nil {
		return m.Policies
	}
	return nil
}

func (m *Group) GetParentGroupIDs() []string {
	if m != nil {
		return m.ParentGroupIDs
	}
	return nil
}

func (m *Group) GetMemberEntityIDs() []string {
	if m != nil {
		return m.MemberEntityIDs
	}
	return nil
}

func (m *Group) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Group) GetCreationTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreationTime
	}
	return nil
}

func (m *Group) GetLastUpdateTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastUpdateTime
	}
	return nil
}

func (m *Group) GetModifyIndex() uint64 {
	if m != nil {
		return m.ModifyIndex
	}
	return 0
}

func (m *Group) GetBucketKeyHash() string {
	if m != nil {
		return m.BucketKeyHash
	}
	return ""
}

func (m *Group) GetAlias() *Alias {
	if m != nil {
		return m.Alias
	}
	return nil
}

func (m *Group) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

// Entity represents an entity that gets persisted and indexed.
// Entity is fundamentally composed of zero or many aliases.
type Entity struct {
	// Aliases are the identities that this entity is made of. This can be
	// empty as well to favor being able to create the entity first and then
	// incrementally adding aliases.
	Aliases []*Alias `sentinel:"" protobuf:"bytes,1,rep,name=aliases" json:"aliases,omitempty"`
	// ID is the unique identifier of the entity which always be a UUID. This
	// should never be allowed to be updated.
	ID string `sentinel:"" protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	// Name is a unique identifier of the entity which is intended to be
	// human-friendly. The default name might not be human friendly since it
	// gets suffixed by a UUID, but it can optionally be updated, unlike the ID
	// field.
	Name string `sentinel:"" protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	// Metadata represents the explicit metadata which is set by the
	// clients.  This is useful to tie any information pertaining to the
	// aliases. This is a non-unique field of entity, meaning multiple
	// entities can have the same metadata set. Entities will be indexed based
	// on this explicit metadata. This enables virtual groupings of entities
	// based on its metadata.
	Metadata map[string]string `sentinel:"" protobuf:"bytes,4,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// CreationTime is the time at which this entity is first created.
	CreationTime *google_protobuf.Timestamp `sentinel:"" protobuf:"bytes,5,opt,name=creation_time,json=creationTime" json:"creation_time,omitempty"`
	// LastUpdateTime is the most recent time at which the properties of this
	// entity got modified. This is helpful in filtering out entities based on
	// its age and to take action on them, if desired.
	LastUpdateTime *google_protobuf.Timestamp `sentinel:"" protobuf:"bytes,6,opt,name=last_update_time,json=lastUpdateTime" json:"last_update_time,omitempty"`
	// MergedEntityIDs are the entities which got merged to this one. Entities
	// will be indexed based on all the entities that got merged into it. This
	// helps to apply the actions on this entity on the tokens that are merged
	// to the merged entities. Merged entities will be deleted entirely and
	// this is the only trackable trail of its earlier presence.
	MergedEntityIDs []string `sentinel:"" protobuf:"bytes,7,rep,name=merged_entity_ids,json=mergedEntityIDs" json:"merged_entity_ids,omitempty"`
	// Policies the entity is entitled to
	Policies []string `sentinel:"" protobuf:"bytes,8,rep,name=policies" json:"policies,omitempty"`
	// BucketKeyHash is the MD5 hash of the storage bucket key into which this
	// entity is stored in the underlying storage. This is useful to find all
	// the entities belonging to a particular bucket during invalidation of the
	// storage key.
	BucketKeyHash string `sentinel:"" protobuf:"bytes,9,opt,name=bucket_key_hash,json=bucketKeyHash" json:"bucket_key_hash,omitempty"`
	// Disabled indicates whether tokens associated with the account should not
	// be able to be used
	Disabled bool `sentinel:"" protobuf:"varint,11,opt,name=disabled" json:"disabled,omitempty"`
}

func (m *Entity) Reset()                    { *m = Entity{} }
func (m *Entity) String() string            { return proto.CompactTextString(m) }
func (*Entity) ProtoMessage()               {}
func (*Entity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Entity) GetAliases() []*Alias {
	if m != nil {
		return m.Aliases
	}
	return nil
}

func (m *Entity) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Entity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Entity) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Entity) GetCreationTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreationTime
	}
	return nil
}

func (m *Entity) GetLastUpdateTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastUpdateTime
	}
	return nil
}

func (m *Entity) GetMergedEntityIDs() []string {
	if m != nil {
		return m.MergedEntityIDs
	}
	return nil
}

func (m *Entity) GetPolicies() []string {
	if m != nil {
		return m.Policies
	}
	return nil
}

func (m *Entity) GetBucketKeyHash() string {
	if m != nil {
		return m.BucketKeyHash
	}
	return ""
}

func (m *Entity) GetDisabled() bool {
	if m != nil {
		return m.Disabled
	}
	return false
}

// Alias represents the alias that gets stored inside of the
// entity object in storage and also represents in an in-memory index of an
// alias object.
type Alias struct {
	// ID is the unique identifier that represents this alias
	ID string `sentinel:"" protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	// CanonicalID is the entity identifier to which this alias belongs to
	CanonicalID string `sentinel:"" protobuf:"bytes,2,opt,name=canonical_id,json=canonicalId" json:"canonical_id,omitempty"`
	// MountType is the backend mount's type to which this alias belongs to.
	// This enables categorically querying aliases of specific backend types.
	MountType string `sentinel:"" protobuf:"bytes,3,opt,name=mount_type,json=mountType" json:"mount_type,omitempty"`
	// MountAccessor is the backend mount's accessor to which this alias
	// belongs to.
	MountAccessor string `sentinel:"" protobuf:"bytes,4,opt,name=mount_accessor,json=mountAccessor" json:"mount_accessor,omitempty"`
	// MountPath is the backend mount's path to which the Maccessor belongs to. This
	// field is not used for any operational purposes. This is only returned when
	// alias is read, only as a nicety.
	MountPath string `sentinel:"" protobuf:"bytes,5,opt,name=mount_path,json=mountPath" json:"mount_path,omitempty"`
	// Metadata is the explicit metadata that clients set against an entity
	// which enables virtual grouping of aliases. Aliases will be indexed
	// against their metadata.
	Metadata map[string]string `sentinel:"" protobuf:"bytes,6,rep,name=metadata" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Name is the identifier of this alias in its authentication source.
	// This does not uniquely identify an alias in Vault. This in conjunction
	// with MountAccessor form to be the factors that represent an alias in a
	// unique way. Aliases will be indexed based on this combined uniqueness
	// factor.
	Name string `sentinel:"" protobuf:"bytes,7,opt,name=name" json:"name,omitempty"`
	// CreationTime is the time at which this alias was first created
	CreationTime *google_protobuf.Timestamp `sentinel:"" protobuf:"bytes,8,opt,name=creation_time,json=creationTime" json:"creation_time,omitempty"`
	// LastUpdateTime is the most recent time at which the properties of this
	// alias got modified. This is helpful in filtering out aliases based
	// on its age and to take action on them, if desired.
	LastUpdateTime *google_protobuf.Timestamp `sentinel:"" protobuf:"bytes,9,opt,name=last_update_time,json=lastUpdateTime" json:"last_update_time,omitempty"`
	// MergedFromCanonicalIDs is the FIFO history of merging activity
	MergedFromCanonicalIDs []string `sentinel:"" protobuf:"bytes,10,rep,name=merged_from_canonical_ids,json=mergedFromCanonicalIds" json:"merged_from_canonical_ids,omitempty"`
}

func (m *Alias) Reset()                    { *m = Alias{} }
func (m *Alias) String() string            { return proto.CompactTextString(m) }
func (*Alias) ProtoMessage()               {}
func (*Alias) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Alias) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Alias) GetCanonicalID() string {
	if m != nil {
		return m.CanonicalID
	}
	return ""
}

func (m *Alias) GetMountType() string {
	if m != nil {
		return m.MountType
	}
	return ""
}

func (m *Alias) GetMountAccessor() string {
	if m != nil {
		return m.MountAccessor
	}
	return ""
}

func (m *Alias) GetMountPath() string {
	if m != nil {
		return m.MountPath
	}
	return ""
}

func (m *Alias) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Alias) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Alias) GetCreationTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.CreationTime
	}
	return nil
}

func (m *Alias) GetLastUpdateTime() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastUpdateTime
	}
	return nil
}

func (m *Alias) GetMergedFromCanonicalIDs() []string {
	if m != nil {
		return m.MergedFromCanonicalIDs
	}
	return nil
}

func init() {
	proto.RegisterType((*Group)(nil), "identity.Group")
	proto.RegisterType((*Entity)(nil), "identity.Entity")
	proto.RegisterType((*Alias)(nil), "identity.Alias")
}

func init() { proto.RegisterFile("types.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 617 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0xc7, 0xd5, 0xa6, 0x1f, 0xe9, 0x69, 0xd7, 0x0d, 0x0b, 0x21, 0x53, 0x69, 0xd0, 0x4d, 0x1a,
	0x2a, 0x5c, 0x64, 0xd2, 0xb8, 0x61, 0xe3, 0x02, 0x4d, 0x30, 0x60, 0x42, 0x48, 0x28, 0x1a, 0xd7,
	0x91, 0x1b, 0x7b, 0xad, 0xb5, 0x24, 0x8e, 0x62, 0x17, 0x91, 0xd7, 0xe1, 0xd5, 0xb8, 0xe6, 0x1d,
	0x90, 0x8f, 0x9b, 0x36, 0xd0, 0xf1, 0x31, 0x6d, 0x77, 0xf6, 0xff, 0x1c, 0x1f, 0x1f, 0x9f, 0xff,
	0x2f, 0x81, 0xbe, 0x29, 0x73, 0xa1, 0x83, 0xbc, 0x50, 0x46, 0x11, 0x5f, 0x72, 0x91, 0x19, 0x69,
	0xca, 0xd1, 0xe3, 0x99, 0x52, 0xb3, 0x44, 0x1c, 0xa2, 0x3e, 0x5d, 0x5c, 0x1e, 0x1a, 0x99, 0x0a,
	0x6d, 0x58, 0x9a, 0xbb, 0xd4, 0xfd, 0x6f, 0x2d, 0x68, 0xbf, 0x2b, 0xd4, 0x22, 0x27, 0x43, 0x68,
	0x4a, 0x4e, 0x1b, 0xe3, 0xc6, 0xa4, 0x17, 0x36, 0x25, 0x27, 0x04, 0x5a, 0x19, 0x4b, 0x05, 0x6d,
	0xa2, 0x82, 0x6b, 0x32, 0x02, 0x3f, 0x57, 0x89, 0x8c, 0xa5, 0xd0, 0xd4, 0x1b, 0x7b, 0x93, 0x5e,
	0xb8, 0xda, 0x93, 0x09, 0xec, 0xe4, 0xac, 0x10, 0x99, 0x89, 0x66, 0xb6, 0x5e, 0x24, 0xb9, 0xa6,
	0x2d, 0xcc, 0x19, 0x3a, 0x1d, 0xaf, 0x39, 0xe7, 0x9a, 0x3c, 0x83, 0x7b, 0xa9, 0x48, 0xa7, 0xa2,
	0x88, 0x5c, 0x97, 0x98, 0xda, 0xc6, 0xd4, 0x6d, 0x17, 0x38, 0x43, 0xdd, 0xe6, 0x1e, 0x83, 0x9f,
	0x0a, 0xc3, 0x38, 0x33, 0x8c, 0x76, 0xc6, 0xde, 0xa4, 0x7f, 0xb4, 0x1b, 0x54, 0xaf, 0x0b, 0xb0,
	0x62, 0xf0, 0x71, 0x19, 0x3f, 0xcb, 0x4c, 0x51, 0x86, 0xab, 0x74, 0xf2, 0x0a, 0xb6, 0xe2, 0x42,
	0x30, 0x23, 0x55, 0x16, 0xd9, 0x67, 0xd3, 0xee, 0xb8, 0x31, 0xe9, 0x1f, 0x8d, 0x02, 0x37, 0x93,
	0xa0, 0x9a, 0x49, 0x70, 0x51, 0xcd, 0x24, 0x1c, 0x54, 0x07, 0xac, 0x44, 0xde, 0xc0, 0x4e, 0xc2,
	0xb4, 0x89, 0x16, 0x39, 0x67, 0x46, 0xb8, 0x1a, 0xfe, 0x3f, 0x6b, 0x0c, 0xed, 0x99, 0xcf, 0x78,
	0x04, 0xab, 0xec, 0xc1, 0x20, 0x55, 0x5c, 0x5e, 0x96, 0x91, 0xcc, 0xb8, 0xf8, 0x4a, 0x7b, 0xe3,
	0xc6, 0xa4, 0x15, 0xf6, 0x9d, 0x76, 0x6e, 0x25, 0xf2, 0x04, 0xb6, 0xa7, 0x8b, 0xf8, 0x4a, 0x98,
	0xe8, 0x4a, 0x94, 0xd1, 0x9c, 0xe9, 0x39, 0x05, 0x9c, 0xfa, 0x96, 0x93, 0x3f, 0x88, 0xf2, 0x3d,
	0xd3, 0x73, 0x72, 0x00, 0x6d, 0x96, 0x48, 0xa6, 0x69, 0x1f, 0xbb, 0xd8, 0x5e, 0x4f, 0xe2, 0xd4,
	0xca, 0xa1, 0x8b, 0x5a, 0xe7, 0x2c, 0x0d, 0x74, 0xe0, 0x9c, 0xb3, 0xeb, 0xd1, 0x4b, 0xd8, 0xfa,
	0x65, 0x4e, 0x64, 0x07, 0xbc, 0x2b, 0x51, 0x2e, 0xfd, 0xb6, 0x4b, 0x72, 0x1f, 0xda, 0x5f, 0x58,
	0xb2, 0xa8, 0x1c, 0x77, 0x9b, 0x93, 0xe6, 0x8b, 0xc6, 0xfe, 0x77, 0x0f, 0x3a, 0xce, 0x12, 0xf2,
	0x14, 0xba, 0x78, 0x89, 0xd0, 0xb4, 0x81, 0x76, 0x6c, 0x34, 0x51, 0xc5, 0x97, 0x40, 0x35, 0x37,
	0x80, 0xf2, 0x6a, 0x40, 0x9d, 0xd4, 0xec, 0x6d, 0x61, 0xbd, 0x47, 0xeb, 0x7a, 0xee, 0xca, 0xff,
	0xf7, 0xb7, 0x7d, 0x07, 0xfe, 0x76, 0x6e, 0xec, 0x2f, 0xd2, 0x5c, 0xcc, 0x04, 0xaf, 0xd3, 0xdc,
	0xad, 0x68, 0xb6, 0x81, 0x35, 0xcd, 0xf5, 0xef, 0xc7, 0xff, 0xed, 0xfb, 0xb9, 0x06, 0x82, 0xde,
	0x75, 0x10, 0x8c, 0xc0, 0xe7, 0x52, 0xb3, 0x69, 0x22, 0x38, 0x72, 0xe0, 0x87, 0xab, 0xfd, 0xed,
	0x5c, 0xfe, 0xe1, 0x41, 0x1b, 0x2d, 0xdc, 0xf8, 0x15, 0xec, 0xc1, 0x20, 0x66, 0x99, 0xca, 0x64,
	0xcc, 0x92, 0x68, 0xe5, 0x69, 0x7f, 0xa5, 0x9d, 0x73, 0xb2, 0x0b, 0x90, 0xaa, 0x45, 0x66, 0x22,
	0x24, 0xcf, 0x59, 0xdc, 0x43, 0xe5, 0xa2, 0xcc, 0x05, 0x39, 0x80, 0xa1, 0x0b, 0xb3, 0x38, 0x16,
	0x5a, 0xab, 0x82, 0xb6, 0xdc, 0xdb, 0x50, 0x3d, 0x5d, 0x8a, 0xeb, 0x2a, 0x39, 0x33, 0x73, 0xf4,
	0xb3, 0xaa, 0xf2, 0x89, 0x99, 0xf9, 0xdf, 0x7f, 0x06, 0xd8, 0xfa, 0x1f, 0x61, 0xa9, 0xe0, 0xeb,
	0xd6, 0xe0, 0xdb, 0x00, 0xc8, 0xbf, 0x03, 0x80, 0x7a, 0x37, 0x06, 0xe8, 0x18, 0x1e, 0x2e, 0x01,
	0xba, 0x2c, 0x54, 0x1a, 0xd5, 0x27, 0xad, 0x29, 0x20, 0x25, 0x0f, 0x5c, 0xc2, 0xdb, 0x42, 0xa5,
	0xaf, 0xd7, 0x43, 0xd7, 0xb7, 0xf2, 0x7b, 0xda, 0xc1, 0xde, 0x9e, 0xff, 0x0c, 0x00, 0x00, 0xff,
	0xff, 0x60, 0x0b, 0xc9, 0x74, 0x3b, 0x06, 0x00, 0x00,
}
