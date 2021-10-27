// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: node/v1/node.proto

package v1

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

type InjectGovernanceVAARequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Index of the current guardian set.
	CurrentSetIndex uint32 `protobuf:"varint,1,opt,name=current_set_index,json=currentSetIndex,proto3" json:"current_set_index,omitempty"`
	// Sequence number. This is critical for replay protection - make sure the sequence number
	// is unique for every new manually injected governance VAA. Sequences are tracked
	// by emitter, and manually injected VAAs all use a single hardcoded emitter.
	//
	// We use random sequence numbers for the manual emitter.
	Sequence uint64 `protobuf:"varint,2,opt,name=sequence,proto3" json:"sequence,omitempty"`
	// Random nonce for disambiguation. Must be identical across all nodes.
	Nonce uint32 `protobuf:"varint,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// Types that are assignable to Payload:
	//	*InjectGovernanceVAARequest_GuardianSet
	//	*InjectGovernanceVAARequest_ContractUpgrade
	//	*InjectGovernanceVAARequest_BridgeRegisterChain
	//	*InjectGovernanceVAARequest_BridgeContractUpgrade
	Payload isInjectGovernanceVAARequest_Payload `protobuf_oneof:"payload"`
}

func (x *InjectGovernanceVAARequest) Reset() {
	*x = InjectGovernanceVAARequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_v1_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InjectGovernanceVAARequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InjectGovernanceVAARequest) ProtoMessage() {}

func (x *InjectGovernanceVAARequest) ProtoReflect() protoreflect.Message {
	mi := &file_node_v1_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InjectGovernanceVAARequest.ProtoReflect.Descriptor instead.
func (*InjectGovernanceVAARequest) Descriptor() ([]byte, []int) {
	return file_node_v1_node_proto_rawDescGZIP(), []int{0}
}

func (x *InjectGovernanceVAARequest) GetCurrentSetIndex() uint32 {
	if x != nil {
		return x.CurrentSetIndex
	}
	return 0
}

func (x *InjectGovernanceVAARequest) GetSequence() uint64 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

func (x *InjectGovernanceVAARequest) GetNonce() uint32 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (m *InjectGovernanceVAARequest) GetPayload() isInjectGovernanceVAARequest_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *InjectGovernanceVAARequest) GetGuardianSet() *GuardianSetUpdate {
	if x, ok := x.GetPayload().(*InjectGovernanceVAARequest_GuardianSet); ok {
		return x.GuardianSet
	}
	return nil
}

func (x *InjectGovernanceVAARequest) GetContractUpgrade() *ContractUpgrade {
	if x, ok := x.GetPayload().(*InjectGovernanceVAARequest_ContractUpgrade); ok {
		return x.ContractUpgrade
	}
	return nil
}

func (x *InjectGovernanceVAARequest) GetBridgeRegisterChain() *BridgeRegisterChain {
	if x, ok := x.GetPayload().(*InjectGovernanceVAARequest_BridgeRegisterChain); ok {
		return x.BridgeRegisterChain
	}
	return nil
}

func (x *InjectGovernanceVAARequest) GetBridgeContractUpgrade() *BridgeUpgradeContract {
	if x, ok := x.GetPayload().(*InjectGovernanceVAARequest_BridgeContractUpgrade); ok {
		return x.BridgeContractUpgrade
	}
	return nil
}

type isInjectGovernanceVAARequest_Payload interface {
	isInjectGovernanceVAARequest_Payload()
}

type InjectGovernanceVAARequest_GuardianSet struct {
	GuardianSet *GuardianSetUpdate `protobuf:"bytes,10,opt,name=guardian_set,json=guardianSet,proto3,oneof"`
}

type InjectGovernanceVAARequest_ContractUpgrade struct {
	ContractUpgrade *ContractUpgrade `protobuf:"bytes,11,opt,name=contract_upgrade,json=contractUpgrade,proto3,oneof"`
}

type InjectGovernanceVAARequest_BridgeRegisterChain struct {
	BridgeRegisterChain *BridgeRegisterChain `protobuf:"bytes,12,opt,name=bridge_register_chain,json=bridgeRegisterChain,proto3,oneof"`
}

type InjectGovernanceVAARequest_BridgeContractUpgrade struct {
	BridgeContractUpgrade *BridgeUpgradeContract `protobuf:"bytes,13,opt,name=bridge_contract_upgrade,json=bridgeContractUpgrade,proto3,oneof"`
}

func (*InjectGovernanceVAARequest_GuardianSet) isInjectGovernanceVAARequest_Payload() {}

func (*InjectGovernanceVAARequest_ContractUpgrade) isInjectGovernanceVAARequest_Payload() {}

func (*InjectGovernanceVAARequest_BridgeRegisterChain) isInjectGovernanceVAARequest_Payload() {}

func (*InjectGovernanceVAARequest_BridgeContractUpgrade) isInjectGovernanceVAARequest_Payload() {}

type InjectGovernanceVAAResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Canonical digest of the submitted VAA.
	Digest []byte `protobuf:"bytes,1,opt,name=digest,proto3" json:"digest,omitempty"`
}

func (x *InjectGovernanceVAAResponse) Reset() {
	*x = InjectGovernanceVAAResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_v1_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InjectGovernanceVAAResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InjectGovernanceVAAResponse) ProtoMessage() {}

func (x *InjectGovernanceVAAResponse) ProtoReflect() protoreflect.Message {
	mi := &file_node_v1_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InjectGovernanceVAAResponse.ProtoReflect.Descriptor instead.
func (*InjectGovernanceVAAResponse) Descriptor() ([]byte, []int) {
	return file_node_v1_node_proto_rawDescGZIP(), []int{1}
}

func (x *InjectGovernanceVAAResponse) GetDigest() []byte {
	if x != nil {
		return x.Digest
	}
	return nil
}

// GuardianSet represents a new guardian set to be submitted to and signed by the node.
// During the genesis procedure, this data structure will be assembled using off-chain collaborative tooling
// like GitHub using a human-readable encoding, so readability is a concern.
type GuardianSetUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Guardians []*GuardianSetUpdate_Guardian `protobuf:"bytes,3,rep,name=guardians,proto3" json:"guardians,omitempty"`
}

func (x *GuardianSetUpdate) Reset() {
	*x = GuardianSetUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_v1_node_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GuardianSetUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GuardianSetUpdate) ProtoMessage() {}

func (x *GuardianSetUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_node_v1_node_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GuardianSetUpdate.ProtoReflect.Descriptor instead.
func (*GuardianSetUpdate) Descriptor() ([]byte, []int) {
	return file_node_v1_node_proto_rawDescGZIP(), []int{2}
}

func (x *GuardianSetUpdate) GetGuardians() []*GuardianSetUpdate_Guardian {
	if x != nil {
		return x.Guardians
	}
	return nil
}

// GuardianKey specifies the on-disk format for a node's guardian key.
type GuardianKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// data is the binary representation of the secp256k1 private key.
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	// Whether this key is deterministically generated and unsuitable for production mode.
	UnsafeDeterministicKey bool `protobuf:"varint,2,opt,name=unsafe_deterministic_key,json=unsafeDeterministicKey,proto3" json:"unsafe_deterministic_key,omitempty"`
}

func (x *GuardianKey) Reset() {
	*x = GuardianKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_v1_node_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GuardianKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GuardianKey) ProtoMessage() {}

func (x *GuardianKey) ProtoReflect() protoreflect.Message {
	mi := &file_node_v1_node_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GuardianKey.ProtoReflect.Descriptor instead.
func (*GuardianKey) Descriptor() ([]byte, []int) {
	return file_node_v1_node_proto_rawDescGZIP(), []int{3}
}

func (x *GuardianKey) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GuardianKey) GetUnsafeDeterministicKey() bool {
	if x != nil {
		return x.UnsafeDeterministicKey
	}
	return false
}

type BridgeRegisterChain struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Module identifier of the token or NFT bridge (typically "TokenBridge" or "NFTBridge")
	Module string `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
	// ID of the chain to be registered.
	ChainId uint32 `protobuf:"varint,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// Hex-encoded emitter address to be registered (without leading 0x).
	EmitterAddress string `protobuf:"bytes,3,opt,name=emitter_address,json=emitterAddress,proto3" json:"emitter_address,omitempty"`
}

func (x *BridgeRegisterChain) Reset() {
	*x = BridgeRegisterChain{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_v1_node_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BridgeRegisterChain) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BridgeRegisterChain) ProtoMessage() {}

func (x *BridgeRegisterChain) ProtoReflect() protoreflect.Message {
	mi := &file_node_v1_node_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BridgeRegisterChain.ProtoReflect.Descriptor instead.
func (*BridgeRegisterChain) Descriptor() ([]byte, []int) {
	return file_node_v1_node_proto_rawDescGZIP(), []int{4}
}

func (x *BridgeRegisterChain) GetModule() string {
	if x != nil {
		return x.Module
	}
	return ""
}

func (x *BridgeRegisterChain) GetChainId() uint32 {
	if x != nil {
		return x.ChainId
	}
	return 0
}

func (x *BridgeRegisterChain) GetEmitterAddress() string {
	if x != nil {
		return x.EmitterAddress
	}
	return ""
}

// ContractUpgrade represents a Wormhole contract update to be submitted to and signed by the node.
type ContractUpgrade struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID of the chain where the Wormhole contract should be updated (uint8).
	ChainId uint32 `protobuf:"varint,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	// Hex-encoded address (without leading 0x) address of the new program/contract.
	NewContract string `protobuf:"bytes,2,opt,name=new_contract,json=newContract,proto3" json:"new_contract,omitempty"`
}

func (x *ContractUpgrade) Reset() {
	*x = ContractUpgrade{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_v1_node_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContractUpgrade) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContractUpgrade) ProtoMessage() {}

func (x *ContractUpgrade) ProtoReflect() protoreflect.Message {
	mi := &file_node_v1_node_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContractUpgrade.ProtoReflect.Descriptor instead.
func (*ContractUpgrade) Descriptor() ([]byte, []int) {
	return file_node_v1_node_proto_rawDescGZIP(), []int{5}
}

func (x *ContractUpgrade) GetChainId() uint32 {
	if x != nil {
		return x.ChainId
	}
	return 0
}

func (x *ContractUpgrade) GetNewContract() string {
	if x != nil {
		return x.NewContract
	}
	return ""
}

type BridgeUpgradeContract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Module identifier of the token or NFT bridge (typically "TokenBridge" or "NFTBridge").
	Module string `protobuf:"bytes,1,opt,name=module,proto3" json:"module,omitempty"`
	// ID of the chain where the bridge contract should be updated (uint16).
	TargetChainId uint32 `protobuf:"varint,2,opt,name=target_chain_id,json=targetChainId,proto3" json:"target_chain_id,omitempty"`
	// Hex-encoded address (without leading 0x) of the new program/contract.
	NewContract string `protobuf:"bytes,3,opt,name=new_contract,json=newContract,proto3" json:"new_contract,omitempty"`
}

func (x *BridgeUpgradeContract) Reset() {
	*x = BridgeUpgradeContract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_v1_node_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BridgeUpgradeContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BridgeUpgradeContract) ProtoMessage() {}

func (x *BridgeUpgradeContract) ProtoReflect() protoreflect.Message {
	mi := &file_node_v1_node_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BridgeUpgradeContract.ProtoReflect.Descriptor instead.
func (*BridgeUpgradeContract) Descriptor() ([]byte, []int) {
	return file_node_v1_node_proto_rawDescGZIP(), []int{6}
}

func (x *BridgeUpgradeContract) GetModule() string {
	if x != nil {
		return x.Module
	}
	return ""
}

func (x *BridgeUpgradeContract) GetTargetChainId() uint32 {
	if x != nil {
		return x.TargetChainId
	}
	return 0
}

func (x *BridgeUpgradeContract) GetNewContract() string {
	if x != nil {
		return x.NewContract
	}
	return ""
}

type FindMissingMessagesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Emitter chain ID to iterate.
	EmitterChain uint32 `protobuf:"varint,1,opt,name=emitter_chain,json=emitterChain,proto3" json:"emitter_chain,omitempty"`
	// Hex-encoded (without leading 0x) emitter address to iterate.
	EmitterAddress string `protobuf:"bytes,2,opt,name=emitter_address,json=emitterAddress,proto3" json:"emitter_address,omitempty"`
}

func (x *FindMissingMessagesRequest) Reset() {
	*x = FindMissingMessagesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_v1_node_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindMissingMessagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindMissingMessagesRequest) ProtoMessage() {}

func (x *FindMissingMessagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_node_v1_node_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindMissingMessagesRequest.ProtoReflect.Descriptor instead.
func (*FindMissingMessagesRequest) Descriptor() ([]byte, []int) {
	return file_node_v1_node_proto_rawDescGZIP(), []int{7}
}

func (x *FindMissingMessagesRequest) GetEmitterChain() uint32 {
	if x != nil {
		return x.EmitterChain
	}
	return 0
}

func (x *FindMissingMessagesRequest) GetEmitterAddress() string {
	if x != nil {
		return x.EmitterAddress
	}
	return ""
}

type FindMissingMessagesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of missing sequence numbers.
	MissingMessages []string `protobuf:"bytes,1,rep,name=missing_messages,json=missingMessages,proto3" json:"missing_messages,omitempty"`
	// Range processed
	FirstSequence uint64 `protobuf:"varint,2,opt,name=first_sequence,json=firstSequence,proto3" json:"first_sequence,omitempty"`
	LastSequence  uint64 `protobuf:"varint,3,opt,name=last_sequence,json=lastSequence,proto3" json:"last_sequence,omitempty"`
}

func (x *FindMissingMessagesResponse) Reset() {
	*x = FindMissingMessagesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_v1_node_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindMissingMessagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindMissingMessagesResponse) ProtoMessage() {}

func (x *FindMissingMessagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_node_v1_node_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindMissingMessagesResponse.ProtoReflect.Descriptor instead.
func (*FindMissingMessagesResponse) Descriptor() ([]byte, []int) {
	return file_node_v1_node_proto_rawDescGZIP(), []int{8}
}

func (x *FindMissingMessagesResponse) GetMissingMessages() []string {
	if x != nil {
		return x.MissingMessages
	}
	return nil
}

func (x *FindMissingMessagesResponse) GetFirstSequence() uint64 {
	if x != nil {
		return x.FirstSequence
	}
	return 0
}

func (x *FindMissingMessagesResponse) GetLastSequence() uint64 {
	if x != nil {
		return x.LastSequence
	}
	return 0
}

// List of guardian set members.
type GuardianSetUpdate_Guardian struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Guardian key pubkey. Stored as hex string with 0x prefix for human readability -
	// this is the canonical Ethereum representation.
	Pubkey string `protobuf:"bytes,1,opt,name=pubkey,proto3" json:"pubkey,omitempty"`
	// Optional descriptive name. Not stored on any chain, purely informational.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GuardianSetUpdate_Guardian) Reset() {
	*x = GuardianSetUpdate_Guardian{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_v1_node_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GuardianSetUpdate_Guardian) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GuardianSetUpdate_Guardian) ProtoMessage() {}

func (x *GuardianSetUpdate_Guardian) ProtoReflect() protoreflect.Message {
	mi := &file_node_v1_node_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GuardianSetUpdate_Guardian.ProtoReflect.Descriptor instead.
func (*GuardianSetUpdate_Guardian) Descriptor() ([]byte, []int) {
	return file_node_v1_node_proto_rawDescGZIP(), []int{2, 0}
}

func (x *GuardianSetUpdate_Guardian) GetPubkey() string {
	if x != nil {
		return x.Pubkey
	}
	return ""
}

func (x *GuardianSetUpdate_Guardian) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_node_v1_node_proto protoreflect.FileDescriptor

var file_node_v1_node_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x22, 0xbb, 0x03,
	0x0a, 0x1a, 0x49, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x6f, 0x76, 0x65, 0x72, 0x6e, 0x61, 0x6e,
	0x63, 0x65, 0x56, 0x41, 0x41, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x11,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x53, 0x65, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75,
	0x65, 0x6e, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x67, 0x75,
	0x61, 0x72, 0x64, 0x69, 0x61, 0x6e, 0x5f, 0x73, 0x65, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x75, 0x61, 0x72, 0x64,
	0x69, 0x61, 0x6e, 0x53, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x00, 0x52, 0x0b,
	0x67, 0x75, 0x61, 0x72, 0x64, 0x69, 0x61, 0x6e, 0x53, 0x65, 0x74, 0x12, 0x45, 0x0a, 0x10, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x48,
	0x00, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x55, 0x70, 0x67, 0x72, 0x61,
	0x64, 0x65, 0x12, 0x52, 0x0a, 0x15, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x72, 0x69, 0x64,
	0x67, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x48,
	0x00, 0x52, 0x13, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x58, 0x0a, 0x17, 0x62, 0x72, 0x69, 0x64, 0x67, 0x65,
	0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5f, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64,
	0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x48, 0x00, 0x52, 0x15, 0x62, 0x72, 0x69, 0x64, 0x67,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65,
	0x42, 0x09, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x35, 0x0a, 0x1b, 0x49,
	0x6e, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x6f, 0x76, 0x65, 0x72, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x56,
	0x41, 0x41, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x69,
	0x67, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x64, 0x69, 0x67, 0x65,
	0x73, 0x74, 0x22, 0x8e, 0x01, 0x0a, 0x11, 0x47, 0x75, 0x61, 0x72, 0x64, 0x69, 0x61, 0x6e, 0x53,
	0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x41, 0x0a, 0x09, 0x67, 0x75, 0x61, 0x72,
	0x64, 0x69, 0x61, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6e, 0x6f,
	0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x75, 0x61, 0x72, 0x64, 0x69, 0x61, 0x6e, 0x53, 0x65,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x47, 0x75, 0x61, 0x72, 0x64, 0x69, 0x61, 0x6e,
	0x52, 0x09, 0x67, 0x75, 0x61, 0x72, 0x64, 0x69, 0x61, 0x6e, 0x73, 0x1a, 0x36, 0x0a, 0x08, 0x47,
	0x75, 0x61, 0x72, 0x64, 0x69, 0x61, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x75, 0x62, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x5b, 0x0a, 0x0b, 0x47, 0x75, 0x61, 0x72, 0x64, 0x69, 0x61, 0x6e, 0x4b,
	0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x38, 0x0a, 0x18, 0x75, 0x6e, 0x73, 0x61, 0x66, 0x65,
	0x5f, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x69, 0x63, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x75, 0x6e, 0x73, 0x61, 0x66, 0x65,
	0x44, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x69, 0x73, 0x74, 0x69, 0x63, 0x4b, 0x65, 0x79,
	0x22, 0x71, 0x0a, 0x13, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x6d,
	0x69, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0x4f, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x55,
	0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x22, 0x7a, 0x0a, 0x15, 0x42, 0x72, 0x69, 0x64, 0x67, 0x65, 0x55, 0x70,
	0x67, 0x72, 0x61, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d,
	0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x5f,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d,
	0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x6e, 0x65, 0x77, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x22, 0x6a, 0x0a, 0x1a, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23,
	0x0a, 0x0d, 0x65, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x65, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x43, 0x68,
	0x61, 0x69, 0x6e, 0x12, 0x27, 0x0a, 0x0f, 0x65, 0x6d, 0x69, 0x74, 0x74, 0x65, 0x72, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x6d,
	0x69, 0x74, 0x74, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x94, 0x01, 0x0a,
	0x1b, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x10,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x5f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0d, 0x66, 0x69, 0x72, 0x73, 0x74, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x53, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x65, 0x32, 0xdb, 0x01, 0x0a, 0x15, 0x4e, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x69, 0x76,
	0x69, 0x6c, 0x65, 0x67, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x60, 0x0a,
	0x13, 0x49, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x6f, 0x76, 0x65, 0x72, 0x6e, 0x61, 0x6e, 0x63,
	0x65, 0x56, 0x41, 0x41, 0x12, 0x23, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x6e, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x6f, 0x76, 0x65, 0x72, 0x6e, 0x61, 0x6e, 0x63, 0x65, 0x56,
	0x41, 0x41, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6e, 0x6f, 0x64, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x6a, 0x65, 0x63, 0x74, 0x47, 0x6f, 0x76, 0x65, 0x72, 0x6e,
	0x61, 0x6e, 0x63, 0x65, 0x56, 0x41, 0x41, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x60, 0x0a, 0x13, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x23, 0x2e, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x6e, 0x6f,
	0x64, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e,
	0x67, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x49, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x65, 0x76, 0x67, 0x65, 0x6e, 0x69, 0x79, 0x2d, 0x73, 0x63, 0x68, 0x65, 0x72, 0x62, 0x69, 0x6e,
	0x61, 0x2f, 0x73, 0x61, 0x6e, 0x64, 0x62, 0x6f, 0x78, 0x2f, 0x77, 0x6f, 0x72, 0x6d, 0x68, 0x6f,
	0x6c, 0x65, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_node_v1_node_proto_rawDescOnce sync.Once
	file_node_v1_node_proto_rawDescData = file_node_v1_node_proto_rawDesc
)

func file_node_v1_node_proto_rawDescGZIP() []byte {
	file_node_v1_node_proto_rawDescOnce.Do(func() {
		file_node_v1_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_node_v1_node_proto_rawDescData)
	})
	return file_node_v1_node_proto_rawDescData
}

var file_node_v1_node_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_node_v1_node_proto_goTypes = []interface{}{
	(*InjectGovernanceVAARequest)(nil),  // 0: node.v1.InjectGovernanceVAARequest
	(*InjectGovernanceVAAResponse)(nil), // 1: node.v1.InjectGovernanceVAAResponse
	(*GuardianSetUpdate)(nil),           // 2: node.v1.GuardianSetUpdate
	(*GuardianKey)(nil),                 // 3: node.v1.GuardianKey
	(*BridgeRegisterChain)(nil),         // 4: node.v1.BridgeRegisterChain
	(*ContractUpgrade)(nil),             // 5: node.v1.ContractUpgrade
	(*BridgeUpgradeContract)(nil),       // 6: node.v1.BridgeUpgradeContract
	(*FindMissingMessagesRequest)(nil),  // 7: node.v1.FindMissingMessagesRequest
	(*FindMissingMessagesResponse)(nil), // 8: node.v1.FindMissingMessagesResponse
	(*GuardianSetUpdate_Guardian)(nil),  // 9: node.v1.GuardianSetUpdate.Guardian
}
var file_node_v1_node_proto_depIdxs = []int32{
	2, // 0: node.v1.InjectGovernanceVAARequest.guardian_set:type_name -> node.v1.GuardianSetUpdate
	5, // 1: node.v1.InjectGovernanceVAARequest.contract_upgrade:type_name -> node.v1.ContractUpgrade
	4, // 2: node.v1.InjectGovernanceVAARequest.bridge_register_chain:type_name -> node.v1.BridgeRegisterChain
	6, // 3: node.v1.InjectGovernanceVAARequest.bridge_contract_upgrade:type_name -> node.v1.BridgeUpgradeContract
	9, // 4: node.v1.GuardianSetUpdate.guardians:type_name -> node.v1.GuardianSetUpdate.Guardian
	0, // 5: node.v1.NodePrivilegedService.InjectGovernanceVAA:input_type -> node.v1.InjectGovernanceVAARequest
	7, // 6: node.v1.NodePrivilegedService.FindMissingMessages:input_type -> node.v1.FindMissingMessagesRequest
	1, // 7: node.v1.NodePrivilegedService.InjectGovernanceVAA:output_type -> node.v1.InjectGovernanceVAAResponse
	8, // 8: node.v1.NodePrivilegedService.FindMissingMessages:output_type -> node.v1.FindMissingMessagesResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_node_v1_node_proto_init() }
func file_node_v1_node_proto_init() {
	if File_node_v1_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_node_v1_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InjectGovernanceVAARequest); i {
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
		file_node_v1_node_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InjectGovernanceVAAResponse); i {
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
		file_node_v1_node_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GuardianSetUpdate); i {
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
		file_node_v1_node_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GuardianKey); i {
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
		file_node_v1_node_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BridgeRegisterChain); i {
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
		file_node_v1_node_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContractUpgrade); i {
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
		file_node_v1_node_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BridgeUpgradeContract); i {
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
		file_node_v1_node_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindMissingMessagesRequest); i {
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
		file_node_v1_node_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindMissingMessagesResponse); i {
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
		file_node_v1_node_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GuardianSetUpdate_Guardian); i {
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
	file_node_v1_node_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*InjectGovernanceVAARequest_GuardianSet)(nil),
		(*InjectGovernanceVAARequest_ContractUpgrade)(nil),
		(*InjectGovernanceVAARequest_BridgeRegisterChain)(nil),
		(*InjectGovernanceVAARequest_BridgeContractUpgrade)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_node_v1_node_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_node_v1_node_proto_goTypes,
		DependencyIndexes: file_node_v1_node_proto_depIdxs,
		MessageInfos:      file_node_v1_node_proto_msgTypes,
	}.Build()
	File_node_v1_node_proto = out.File
	file_node_v1_node_proto_rawDesc = nil
	file_node_v1_node_proto_goTypes = nil
	file_node_v1_node_proto_depIdxs = nil
}