// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.17.3
// source: stroeer/page/stage/v1/stage.proto

package stage

import (
	v1 "github.com/stroeer/tapir/go/stroeer/core/v1"
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

// A stream stage with companions and the main content area.
// Embedded items can be editorial articles, advertisement and/or stages (only one level deep).
type Stage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Configuration  *Stage_Configuration `protobuf:"bytes,1,opt,name=configuration,proto3" json:"configuration,omitempty"`
	StreamItems    []*Stage_Item        `protobuf:"bytes,2,rep,name=stream_items,json=streamItems,proto3" json:"stream_items,omitempty"`
	CompanionItems []*Stage_Item        `protobuf:"bytes,3,rep,name=companion_items,json=companionItems,proto3" json:"companion_items,omitempty"`
}

func (x *Stage) Reset() {
	*x = Stage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stroeer_page_stage_v1_stage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stage) ProtoMessage() {}

func (x *Stage) ProtoReflect() protoreflect.Message {
	mi := &file_stroeer_page_stage_v1_stage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stage.ProtoReflect.Descriptor instead.
func (*Stage) Descriptor() ([]byte, []int) {
	return file_stroeer_page_stage_v1_stage_proto_rawDescGZIP(), []int{0}
}

func (x *Stage) GetConfiguration() *Stage_Configuration {
	if x != nil {
		return x.Configuration
	}
	return nil
}

func (x *Stage) GetStreamItems() []*Stage_Item {
	if x != nil {
		return x.StreamItems
	}
	return nil
}

func (x *Stage) GetCompanionItems() []*Stage_Item {
	if x != nil {
		return x.CompanionItems
	}
	return nil
}

// Stage and companion layouts must be coordinated with UI department, especially the `stream` layout.
// The stage title is in the `references`, also when not linked.
type Stage_Configuration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// * `stage_layout`: represents a visual layout, which may affect teasers, colors, and other styles.
	// If entry `stage_layout` is missing, use layout `stream`.
	// the stage_layout field can be used to configure full page width commercials (e.g. nativendo, t-online empfiehlt)
	Fields map[string]string `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// All kind of references as part of the stage. Filter reference by
	// `Reference.type`:
	// * `ref_stage_title`
	// * `ref_sub_navigation`
	References []*v1.Reference `protobuf:"bytes,2,rep,name=references,proto3" json:"references,omitempty"`
}

func (x *Stage_Configuration) Reset() {
	*x = Stage_Configuration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stroeer_page_stage_v1_stage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stage_Configuration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stage_Configuration) ProtoMessage() {}

func (x *Stage_Configuration) ProtoReflect() protoreflect.Message {
	mi := &file_stroeer_page_stage_v1_stage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stage_Configuration.ProtoReflect.Descriptor instead.
func (*Stage_Configuration) Descriptor() ([]byte, []int) {
	return file_stroeer_page_stage_v1_stage_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Stage_Configuration) GetFields() map[string]string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *Stage_Configuration) GetReferences() []*v1.Reference {
	if x != nil {
		return x.References
	}
	return nil
}

// A generic model to represent what can be part of a stage.
type Stage_Item struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional flags/hints for customization independent of the `item_type`
	Fields map[string]string `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Types that are assignable to ItemType:
	//	*Stage_Item_Stage
	//	*Stage_Item_ArticleTeaser_
	//	*Stage_Item_Commercial_
	ItemType isStage_Item_ItemType `protobuf_oneof:"item_type"`
}

func (x *Stage_Item) Reset() {
	*x = Stage_Item{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stroeer_page_stage_v1_stage_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stage_Item) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stage_Item) ProtoMessage() {}

func (x *Stage_Item) ProtoReflect() protoreflect.Message {
	mi := &file_stroeer_page_stage_v1_stage_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stage_Item.ProtoReflect.Descriptor instead.
func (*Stage_Item) Descriptor() ([]byte, []int) {
	return file_stroeer_page_stage_v1_stage_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Stage_Item) GetFields() map[string]string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (m *Stage_Item) GetItemType() isStage_Item_ItemType {
	if m != nil {
		return m.ItemType
	}
	return nil
}

func (x *Stage_Item) GetStage() *Stage {
	if x, ok := x.GetItemType().(*Stage_Item_Stage); ok {
		return x.Stage
	}
	return nil
}

func (x *Stage_Item) GetArticleTeaser() *Stage_Item_ArticleTeaser {
	if x, ok := x.GetItemType().(*Stage_Item_ArticleTeaser_); ok {
		return x.ArticleTeaser
	}
	return nil
}

func (x *Stage_Item) GetCommercial() *Stage_Item_Commercial {
	if x, ok := x.GetItemType().(*Stage_Item_Commercial_); ok {
		return x.Commercial
	}
	return nil
}

type isStage_Item_ItemType interface {
	isStage_Item_ItemType()
}

type Stage_Item_Stage struct {
	// We don't support recursion and limit the depth at this point to 1 in the services
	Stage *Stage `protobuf:"bytes,2,opt,name=stage,proto3,oneof"`
}

type Stage_Item_ArticleTeaser_ struct {
	ArticleTeaser *Stage_Item_ArticleTeaser `protobuf:"bytes,3,opt,name=article_teaser,json=articleTeaser,proto3,oneof"`
}

type Stage_Item_Commercial_ struct {
	Commercial *Stage_Item_Commercial `protobuf:"bytes,4,opt,name=commercial,proto3,oneof"`
}

func (*Stage_Item_Stage) isStage_Item_ItemType() {}

func (*Stage_Item_ArticleTeaser_) isStage_Item_ItemType() {}

func (*Stage_Item_Commercial_) isStage_Item_ItemType() {}

// Bundles the data needed to render a core article as a teaser as part of a stage.
// Coordinate teaser layouts, especially the `stream` layout with UI department.
type Stage_Item_ArticleTeaser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// `fields` may contain additional rendering information.
	// * `teaser_layout`: marker to use a particular layout in the rendering teaser template.
	// * `teaser_variant`: marker to use also a particular variant of the provided `teaser_layout`.
	// If `teaser_layout` entries is missing use layout `stream`.
	Fields map[string]string `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// `Article.body` is set to `null` to reduce overhead.
	Article *v1.Article `protobuf:"bytes,2,opt,name=article,proto3" json:"article,omitempty"`
}

func (x *Stage_Item_ArticleTeaser) Reset() {
	*x = Stage_Item_ArticleTeaser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stroeer_page_stage_v1_stage_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stage_Item_ArticleTeaser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stage_Item_ArticleTeaser) ProtoMessage() {}

func (x *Stage_Item_ArticleTeaser) ProtoReflect() protoreflect.Message {
	mi := &file_stroeer_page_stage_v1_stage_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stage_Item_ArticleTeaser.ProtoReflect.Descriptor instead.
func (*Stage_Item_ArticleTeaser) Descriptor() ([]byte, []int) {
	return file_stroeer_page_stage_v1_stage_proto_rawDescGZIP(), []int{0, 1, 1}
}

func (x *Stage_Item_ArticleTeaser) GetFields() map[string]string {
	if x != nil {
		return x.Fields
	}
	return nil
}

func (x *Stage_Item_ArticleTeaser) GetArticle() *v1.Article {
	if x != nil {
		return x.Article
	}
	return nil
}

// Bundles the data to render a commercial as part of a stage.
type Stage_Item_Commercial struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// e.g. advertisement booking ids. Data may differ between commercial types.
	Fields map[string]string `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Stage_Item_Commercial) Reset() {
	*x = Stage_Item_Commercial{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stroeer_page_stage_v1_stage_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stage_Item_Commercial) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stage_Item_Commercial) ProtoMessage() {}

func (x *Stage_Item_Commercial) ProtoReflect() protoreflect.Message {
	mi := &file_stroeer_page_stage_v1_stage_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stage_Item_Commercial.ProtoReflect.Descriptor instead.
func (*Stage_Item_Commercial) Descriptor() ([]byte, []int) {
	return file_stroeer_page_stage_v1_stage_proto_rawDescGZIP(), []int{0, 1, 2}
}

func (x *Stage_Item_Commercial) GetFields() map[string]string {
	if x != nil {
		return x.Fields
	}
	return nil
}

var File_stroeer_page_stage_v1_stage_proto protoreflect.FileDescriptor

var file_stroeer_page_stage_v1_stage_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x73,
	0x74, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x15, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x67,
	0x65, 0x2e, 0x73, 0x74, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1d, 0x73, 0x74, 0x72, 0x6f,
	0x65, 0x65, 0x72, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x72, 0x74, 0x69,
	0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x73, 0x74, 0x72, 0x6f, 0x65,
	0x65, 0x72, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xae, 0x09, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x67,
	0x65, 0x12, 0x50, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x65,
	0x65, 0x72, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x44, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5f, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x74, 0x72, 0x6f,
	0x65, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x67, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0b, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x4a, 0x0a, 0x0f, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x67,
	0x65, 0x2e, 0x73, 0x74, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x67, 0x65,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x69, 0x6f, 0x6e,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0xd6, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4e, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65,
	0x72, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x74, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x3a, 0x0a, 0x0a, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x73, 0x74,
	0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x0a, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0xe7,
	0x05, 0x0a, 0x04, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x45, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65,
	0x72, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x74, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x34,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x74, 0x61,
	0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x67, 0x65, 0x12, 0x58, 0x0a, 0x0e, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x5f,
	0x74, 0x65, 0x61, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x73,
	0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x67,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x2e,
	0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x65, 0x61, 0x73, 0x65, 0x72, 0x48, 0x00, 0x52,
	0x0d, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x65, 0x61, 0x73, 0x65, 0x72, 0x12, 0x4e,
	0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x69, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x67,
	0x65, 0x2e, 0x73, 0x74, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x67, 0x65,
	0x2e, 0x49, 0x74, 0x65, 0x6d, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x69, 0x61, 0x6c,
	0x48, 0x00, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x69, 0x61, 0x6c, 0x1a, 0x39,
	0x0a, 0x0b, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0xd3, 0x01, 0x0a, 0x0d, 0x41, 0x72,
	0x74, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x65, 0x61, 0x73, 0x65, 0x72, 0x12, 0x53, 0x0a, 0x06, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3b, 0x2e, 0x73, 0x74,
	0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x67, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x74, 0x65, 0x6d, 0x2e, 0x41,
	0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x54, 0x65, 0x61, 0x73, 0x65, 0x72, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x12, 0x32, 0x0a, 0x07, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x52, 0x07, 0x61, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x1a, 0x39, 0x0a, 0x0b, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a,
	0x99, 0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x69, 0x61, 0x6c, 0x12, 0x50,
	0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38,
	0x2e, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x74,
	0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x74, 0x65,
	0x6d, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x69, 0x61, 0x6c, 0x2e, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73,
	0x1a, 0x39, 0x0a, 0x0b, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x69,
	0x74, 0x65, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x55, 0x0a, 0x18, 0x64, 0x65, 0x2e, 0x73,
	0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x67,
	0x65, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2f, 0x74, 0x61, 0x70, 0x69, 0x72,
	0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x74, 0x72, 0x6f, 0x65, 0x65, 0x72, 0x2f, 0x70, 0x61, 0x67, 0x65,
	0x2f, 0x73, 0x74, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x74, 0x61, 0x67, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stroeer_page_stage_v1_stage_proto_rawDescOnce sync.Once
	file_stroeer_page_stage_v1_stage_proto_rawDescData = file_stroeer_page_stage_v1_stage_proto_rawDesc
)

func file_stroeer_page_stage_v1_stage_proto_rawDescGZIP() []byte {
	file_stroeer_page_stage_v1_stage_proto_rawDescOnce.Do(func() {
		file_stroeer_page_stage_v1_stage_proto_rawDescData = protoimpl.X.CompressGZIP(file_stroeer_page_stage_v1_stage_proto_rawDescData)
	})
	return file_stroeer_page_stage_v1_stage_proto_rawDescData
}

var file_stroeer_page_stage_v1_stage_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_stroeer_page_stage_v1_stage_proto_goTypes = []interface{}{
	(*Stage)(nil),                    // 0: stroeer.page.stage.v1.Stage
	(*Stage_Configuration)(nil),      // 1: stroeer.page.stage.v1.Stage.Configuration
	(*Stage_Item)(nil),               // 2: stroeer.page.stage.v1.Stage.Item
	nil,                              // 3: stroeer.page.stage.v1.Stage.Configuration.FieldsEntry
	nil,                              // 4: stroeer.page.stage.v1.Stage.Item.FieldsEntry
	(*Stage_Item_ArticleTeaser)(nil), // 5: stroeer.page.stage.v1.Stage.Item.ArticleTeaser
	(*Stage_Item_Commercial)(nil),    // 6: stroeer.page.stage.v1.Stage.Item.Commercial
	nil,                              // 7: stroeer.page.stage.v1.Stage.Item.ArticleTeaser.FieldsEntry
	nil,                              // 8: stroeer.page.stage.v1.Stage.Item.Commercial.FieldsEntry
	(*v1.Reference)(nil),             // 9: stroeer.core.v1.Reference
	(*v1.Article)(nil),               // 10: stroeer.core.v1.Article
}
var file_stroeer_page_stage_v1_stage_proto_depIdxs = []int32{
	1,  // 0: stroeer.page.stage.v1.Stage.configuration:type_name -> stroeer.page.stage.v1.Stage.Configuration
	2,  // 1: stroeer.page.stage.v1.Stage.stream_items:type_name -> stroeer.page.stage.v1.Stage.Item
	2,  // 2: stroeer.page.stage.v1.Stage.companion_items:type_name -> stroeer.page.stage.v1.Stage.Item
	3,  // 3: stroeer.page.stage.v1.Stage.Configuration.fields:type_name -> stroeer.page.stage.v1.Stage.Configuration.FieldsEntry
	9,  // 4: stroeer.page.stage.v1.Stage.Configuration.references:type_name -> stroeer.core.v1.Reference
	4,  // 5: stroeer.page.stage.v1.Stage.Item.fields:type_name -> stroeer.page.stage.v1.Stage.Item.FieldsEntry
	0,  // 6: stroeer.page.stage.v1.Stage.Item.stage:type_name -> stroeer.page.stage.v1.Stage
	5,  // 7: stroeer.page.stage.v1.Stage.Item.article_teaser:type_name -> stroeer.page.stage.v1.Stage.Item.ArticleTeaser
	6,  // 8: stroeer.page.stage.v1.Stage.Item.commercial:type_name -> stroeer.page.stage.v1.Stage.Item.Commercial
	7,  // 9: stroeer.page.stage.v1.Stage.Item.ArticleTeaser.fields:type_name -> stroeer.page.stage.v1.Stage.Item.ArticleTeaser.FieldsEntry
	10, // 10: stroeer.page.stage.v1.Stage.Item.ArticleTeaser.article:type_name -> stroeer.core.v1.Article
	8,  // 11: stroeer.page.stage.v1.Stage.Item.Commercial.fields:type_name -> stroeer.page.stage.v1.Stage.Item.Commercial.FieldsEntry
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_stroeer_page_stage_v1_stage_proto_init() }
func file_stroeer_page_stage_v1_stage_proto_init() {
	if File_stroeer_page_stage_v1_stage_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stroeer_page_stage_v1_stage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stage); i {
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
		file_stroeer_page_stage_v1_stage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stage_Configuration); i {
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
		file_stroeer_page_stage_v1_stage_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stage_Item); i {
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
		file_stroeer_page_stage_v1_stage_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stage_Item_ArticleTeaser); i {
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
		file_stroeer_page_stage_v1_stage_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stage_Item_Commercial); i {
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
	file_stroeer_page_stage_v1_stage_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Stage_Item_Stage)(nil),
		(*Stage_Item_ArticleTeaser_)(nil),
		(*Stage_Item_Commercial_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_stroeer_page_stage_v1_stage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_stroeer_page_stage_v1_stage_proto_goTypes,
		DependencyIndexes: file_stroeer_page_stage_v1_stage_proto_depIdxs,
		MessageInfos:      file_stroeer_page_stage_v1_stage_proto_msgTypes,
	}.Build()
	File_stroeer_page_stage_v1_stage_proto = out.File
	file_stroeer_page_stage_v1_stage_proto_rawDesc = nil
	file_stroeer_page_stage_v1_stage_proto_goTypes = nil
	file_stroeer_page_stage_v1_stage_proto_depIdxs = nil
}
