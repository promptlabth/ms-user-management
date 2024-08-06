// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: protos/user.proto

package user

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

type GetUserByIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirebaseId string `protobuf:"bytes,1,opt,name=firebase_id,json=firebaseId,proto3" json:"firebase_id,omitempty"`
}

func (x *GetUserByIdReq) Reset() {
	*x = GetUserByIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByIdReq) ProtoMessage() {}

func (x *GetUserByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByIdReq.ProtoReflect.Descriptor instead.
func (*GetUserByIdReq) Descriptor() ([]byte, []int) {
	return file_protos_user_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserByIdReq) GetFirebaseId() string {
	if x != nil {
		return x.FirebaseId
	}
	return ""
}

type GetUserByIdRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             int64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FirebaseId     string      `protobuf:"bytes,2,opt,name=firebase_id,json=firebaseId,proto3" json:"firebase_id,omitempty"`
	Name           string      `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Email          *string     `protobuf:"bytes,4,opt,name=email,proto3,oneof" json:"email,omitempty"`
	ProfilePic     *string     `protobuf:"bytes,5,opt,name=profile_pic,json=profilePic,proto3,oneof" json:"profile_pic,omitempty"`
	Platform       *string     `protobuf:"bytes,6,opt,name=platform,proto3,oneof" json:"platform,omitempty"`
	AccessToken    *string     `protobuf:"bytes,7,opt,name=access_token,json=accessToken,proto3,oneof" json:"access_token,omitempty"`
	StripeId       *string     `protobuf:"bytes,8,opt,name=stripe_id,json=stripeId,proto3,oneof" json:"stripe_id,omitempty"`
	BalanceMessage int64       `protobuf:"varint,9,opt,name=balance_message,json=balanceMessage,proto3" json:"balance_message,omitempty"`
	Plan           *PlanDetail `protobuf:"bytes,10,opt,name=plan,proto3" json:"plan,omitempty"`
}

func (x *GetUserByIdRes) Reset() {
	*x = GetUserByIdRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserByIdRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByIdRes) ProtoMessage() {}

func (x *GetUserByIdRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByIdRes.ProtoReflect.Descriptor instead.
func (*GetUserByIdRes) Descriptor() ([]byte, []int) {
	return file_protos_user_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserByIdRes) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetUserByIdRes) GetFirebaseId() string {
	if x != nil {
		return x.FirebaseId
	}
	return ""
}

func (x *GetUserByIdRes) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetUserByIdRes) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *GetUserByIdRes) GetProfilePic() string {
	if x != nil && x.ProfilePic != nil {
		return *x.ProfilePic
	}
	return ""
}

func (x *GetUserByIdRes) GetPlatform() string {
	if x != nil && x.Platform != nil {
		return *x.Platform
	}
	return ""
}

func (x *GetUserByIdRes) GetAccessToken() string {
	if x != nil && x.AccessToken != nil {
		return *x.AccessToken
	}
	return ""
}

func (x *GetUserByIdRes) GetStripeId() string {
	if x != nil && x.StripeId != nil {
		return *x.StripeId
	}
	return ""
}

func (x *GetUserByIdRes) GetBalanceMessage() int64 {
	if x != nil {
		return x.BalanceMessage
	}
	return 0
}

func (x *GetUserByIdRes) GetPlan() *PlanDetail {
	if x != nil {
		return x.Plan
	}
	return nil
}

type PlanDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	PlanType    string  `protobuf:"bytes,2,opt,name=plan_type,json=planType,proto3" json:"plan_type,omitempty"`
	MaxMessages int64   `protobuf:"varint,3,opt,name=max_messages,json=maxMessages,proto3" json:"max_messages,omitempty"`
	ProductId   *string `protobuf:"bytes,4,opt,name=product_id,json=productId,proto3,oneof" json:"product_id,omitempty"`
}

func (x *PlanDetail) Reset() {
	*x = PlanDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlanDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlanDetail) ProtoMessage() {}

func (x *PlanDetail) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlanDetail.ProtoReflect.Descriptor instead.
func (*PlanDetail) Descriptor() ([]byte, []int) {
	return file_protos_user_proto_rawDescGZIP(), []int{2}
}

func (x *PlanDetail) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PlanDetail) GetPlanType() string {
	if x != nil {
		return x.PlanType
	}
	return ""
}

func (x *PlanDetail) GetMaxMessages() int64 {
	if x != nil {
		return x.MaxMessages
	}
	return 0
}

func (x *PlanDetail) GetProductId() string {
	if x != nil && x.ProductId != nil {
		return *x.ProductId
	}
	return ""
}

type UpsertUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirebaseId  string  `protobuf:"bytes,1,opt,name=firebase_id,json=firebaseId,proto3" json:"firebase_id,omitempty"`
	Name        string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email       *string `protobuf:"bytes,3,opt,name=email,proto3,oneof" json:"email,omitempty"`
	ProfilePic  *string `protobuf:"bytes,4,opt,name=profile_pic,json=profilePic,proto3,oneof" json:"profile_pic,omitempty"`
	Platform    *string `protobuf:"bytes,5,opt,name=platform,proto3,oneof" json:"platform,omitempty"`
	AccessToken *string `protobuf:"bytes,6,opt,name=access_token,json=accessToken,proto3,oneof" json:"access_token,omitempty"`
	StripeId    *string `protobuf:"bytes,8,opt,name=stripe_id,json=stripeId,proto3,oneof" json:"stripe_id,omitempty"`
}

func (x *UpsertUserReq) Reset() {
	*x = UpsertUserReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertUserReq) ProtoMessage() {}

func (x *UpsertUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertUserReq.ProtoReflect.Descriptor instead.
func (*UpsertUserReq) Descriptor() ([]byte, []int) {
	return file_protos_user_proto_rawDescGZIP(), []int{3}
}

func (x *UpsertUserReq) GetFirebaseId() string {
	if x != nil {
		return x.FirebaseId
	}
	return ""
}

func (x *UpsertUserReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpsertUserReq) GetEmail() string {
	if x != nil && x.Email != nil {
		return *x.Email
	}
	return ""
}

func (x *UpsertUserReq) GetProfilePic() string {
	if x != nil && x.ProfilePic != nil {
		return *x.ProfilePic
	}
	return ""
}

func (x *UpsertUserReq) GetPlatform() string {
	if x != nil && x.Platform != nil {
		return *x.Platform
	}
	return ""
}

func (x *UpsertUserReq) GetAccessToken() string {
	if x != nil && x.AccessToken != nil {
		return *x.AccessToken
	}
	return ""
}

func (x *UpsertUserReq) GetStripeId() string {
	if x != nil && x.StripeId != nil {
		return *x.StripeId
	}
	return ""
}

type UpsertUserRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Status  string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *UpsertUserRes) Reset() {
	*x = UpsertUserRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertUserRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertUserRes) ProtoMessage() {}

func (x *UpsertUserRes) ProtoReflect() protoreflect.Message {
	mi := &file_protos_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertUserRes.ProtoReflect.Descriptor instead.
func (*UpsertUserRes) Descriptor() ([]byte, []int) {
	return file_protos_user_proto_rawDescGZIP(), []int{4}
}

func (x *UpsertUserRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UpsertUserRes) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_protos_user_proto protoreflect.FileDescriptor

var file_protos_user_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x31, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x66,
	0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x49, 0x64, 0x22, 0x96, 0x03, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x88, 0x01, 0x01, 0x12,
	0x24, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x69, 0x63, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50,
	0x69, 0x63, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x0b,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x20,
	0x0a, 0x09, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x04, 0x52, 0x08, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x27, 0x0a, 0x0f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x70, 0x6c, 0x61,
	0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x50,
	0x6c, 0x61, 0x6e, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x42,
	0x08, 0x0a, 0x06, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x69, 0x63, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x73, 0x74, 0x72, 0x69,
	0x70, 0x65, 0x5f, 0x69, 0x64, 0x22, 0x8f, 0x01, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x6e, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x6e, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x61, 0x78, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x6d, 0x61, 0x78, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x22, 0xb6, 0x02, 0x0a, 0x0d, 0x55, 0x70, 0x73, 0x65,
	0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x69, 0x72,
	0x65, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x66, 0x69, 0x72, 0x65, 0x62, 0x61, 0x73, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x69, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01,
	0x52, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x50, 0x69, 0x63, 0x88, 0x01, 0x01, 0x12,
	0x1f, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x02, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x88, 0x01, 0x01,
	0x12, 0x26, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x73, 0x74, 0x72, 0x69,
	0x70, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x08, 0x73,
	0x74, 0x72, 0x69, 0x70, 0x65, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x70, 0x69, 0x63, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72,
	0x6d, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x5f, 0x69, 0x64,
	0x22, 0x41, 0x0a, 0x0d, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x32, 0x82, 0x01, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73,
	0x12, 0x36, 0x0a, 0x0a, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x13,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x61, 0x70,
	0x70, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_user_proto_rawDescOnce sync.Once
	file_protos_user_proto_rawDescData = file_protos_user_proto_rawDesc
)

func file_protos_user_proto_rawDescGZIP() []byte {
	file_protos_user_proto_rawDescOnce.Do(func() {
		file_protos_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_user_proto_rawDescData)
	})
	return file_protos_user_proto_rawDescData
}

var file_protos_user_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_user_proto_goTypes = []any{
	(*GetUserByIdReq)(nil), // 0: user.GetUserByIdReq
	(*GetUserByIdRes)(nil), // 1: user.GetUserByIdRes
	(*PlanDetail)(nil),     // 2: user.PlanDetail
	(*UpsertUserReq)(nil),  // 3: user.UpsertUserReq
	(*UpsertUserRes)(nil),  // 4: user.UpsertUserRes
}
var file_protos_user_proto_depIdxs = []int32{
	2, // 0: user.GetUserByIdRes.plan:type_name -> user.PlanDetail
	0, // 1: user.UserService.GetDetailUser:input_type -> user.GetUserByIdReq
	3, // 2: user.UserService.UpsertUser:input_type -> user.UpsertUserReq
	1, // 3: user.UserService.GetDetailUser:output_type -> user.GetUserByIdRes
	4, // 4: user.UserService.UpsertUser:output_type -> user.UpsertUserRes
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_user_proto_init() }
func file_protos_user_proto_init() {
	if File_protos_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_user_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserByIdReq); i {
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
		file_protos_user_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetUserByIdRes); i {
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
		file_protos_user_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*PlanDetail); i {
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
		file_protos_user_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpsertUserReq); i {
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
		file_protos_user_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UpsertUserRes); i {
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
	file_protos_user_proto_msgTypes[1].OneofWrappers = []any{}
	file_protos_user_proto_msgTypes[2].OneofWrappers = []any{}
	file_protos_user_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_user_proto_goTypes,
		DependencyIndexes: file_protos_user_proto_depIdxs,
		MessageInfos:      file_protos_user_proto_msgTypes,
	}.Build()
	File_protos_user_proto = out.File
	file_protos_user_proto_rawDesc = nil
	file_protos_user_proto_goTypes = nil
	file_protos_user_proto_depIdxs = nil
}
