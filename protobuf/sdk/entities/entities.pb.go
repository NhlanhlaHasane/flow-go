// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sdk/entities/entities.proto

package entities

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type TransactionStatus int32

const (
	TransactionStatus_STATUS_UNKNOWN   TransactionStatus = 0
	TransactionStatus_STATUS_PENDING   TransactionStatus = 1
	TransactionStatus_STATUS_FINALIZED TransactionStatus = 2
	TransactionStatus_STATUS_REVERTED  TransactionStatus = 3
	TransactionStatus_STATUS_SEALED    TransactionStatus = 4
)

var TransactionStatus_name = map[int32]string{
	0: "STATUS_UNKNOWN",
	1: "STATUS_PENDING",
	2: "STATUS_FINALIZED",
	3: "STATUS_REVERTED",
	4: "STATUS_SEALED",
}

var TransactionStatus_value = map[string]int32{
	"STATUS_UNKNOWN":   0,
	"STATUS_PENDING":   1,
	"STATUS_FINALIZED": 2,
	"STATUS_REVERTED":  3,
	"STATUS_SEALED":    4,
}

func (x TransactionStatus) String() string {
	return proto.EnumName(TransactionStatus_name, int32(x))
}

func (TransactionStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_9ccd9f92419ce3f6, []int{0}
}

type BlockHeader struct {
	ID                   []byte   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	PreviousBlockHash    []byte   `protobuf:"bytes,2,opt,name=previousBlockHash,proto3" json:"previousBlockHash,omitempty"`
	Height               uint64   `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockHeader) Reset()         { *m = BlockHeader{} }
func (m *BlockHeader) String() string { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()    {}
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ccd9f92419ce3f6, []int{0}
}

func (m *BlockHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockHeader.Unmarshal(m, b)
}
func (m *BlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockHeader.Marshal(b, m, deterministic)
}
func (m *BlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeader.Merge(m, src)
}
func (m *BlockHeader) XXX_Size() int {
	return xxx_messageInfo_BlockHeader.Size(m)
}
func (m *BlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeader proto.InternalMessageInfo

func (m *BlockHeader) GetID() []byte {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *BlockHeader) GetPreviousBlockHash() []byte {
	if m != nil {
		return m.PreviousBlockHash
	}
	return nil
}

func (m *BlockHeader) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type Block struct {
	ChainID                string                  `protobuf:"bytes,1,opt,name=chainID,proto3" json:"chainID,omitempty"`
	Height                 uint64                  `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	PreviousBlockHash      []byte                  `protobuf:"bytes,3,opt,name=previousBlockHash,proto3" json:"previousBlockHash,omitempty"`
	Timestamp              *timestamp.Timestamp    `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	SignedCollectionHashes []*SignedCollectionHash `protobuf:"bytes,5,rep,name=signedCollectionHashes,proto3" json:"signedCollectionHashes,omitempty"`
	BlockSeals             []*BlockSeal            `protobuf:"bytes,6,rep,name=blockSeals,proto3" json:"blockSeals,omitempty"`
	Signatures             [][]byte                `protobuf:"bytes,7,rep,name=signatures,proto3" json:"signatures,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}                `json:"-"`
	XXX_unrecognized       []byte                  `json:"-"`
	XXX_sizecache          int32                   `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ccd9f92419ce3f6, []int{1}
}

func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetChainID() string {
	if m != nil {
		return m.ChainID
	}
	return ""
}

func (m *Block) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Block) GetPreviousBlockHash() []byte {
	if m != nil {
		return m.PreviousBlockHash
	}
	return nil
}

func (m *Block) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *Block) GetSignedCollectionHashes() []*SignedCollectionHash {
	if m != nil {
		return m.SignedCollectionHashes
	}
	return nil
}

func (m *Block) GetBlockSeals() []*BlockSeal {
	if m != nil {
		return m.BlockSeals
	}
	return nil
}

func (m *Block) GetSignatures() [][]byte {
	if m != nil {
		return m.Signatures
	}
	return nil
}

// When computation results for past blocks become available, a corresponding execution receipt
// (issued by execution nodes) and result approval (issued by access nodes) is issued.
//
// We want to persist the execution receipt (assuming the results are consistent) and
// the result approval in subsequent blocks. This step is called _sealing a block_.
//
// This is a sub-message for a block seal. This will allow us later to (optionally) seal
// multiple past blocks in one new block without revising the structure.
//
type BlockSeal struct {
	BlockHash                  []byte   `protobuf:"bytes,1,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	ExecutionReceiptHash       []byte   `protobuf:"bytes,2,opt,name=executionReceiptHash,proto3" json:"executionReceiptHash,omitempty"`
	ExecutionReceiptSignatures [][]byte `protobuf:"bytes,3,rep,name=executionReceiptSignatures,proto3" json:"executionReceiptSignatures,omitempty"`
	ResultApprovalSignatures   [][]byte `protobuf:"bytes,4,rep,name=resultApprovalSignatures,proto3" json:"resultApprovalSignatures,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *BlockSeal) Reset()         { *m = BlockSeal{} }
func (m *BlockSeal) String() string { return proto.CompactTextString(m) }
func (*BlockSeal) ProtoMessage()    {}
func (*BlockSeal) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ccd9f92419ce3f6, []int{2}
}

func (m *BlockSeal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockSeal.Unmarshal(m, b)
}
func (m *BlockSeal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockSeal.Marshal(b, m, deterministic)
}
func (m *BlockSeal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockSeal.Merge(m, src)
}
func (m *BlockSeal) XXX_Size() int {
	return xxx_messageInfo_BlockSeal.Size(m)
}
func (m *BlockSeal) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockSeal.DiscardUnknown(m)
}

var xxx_messageInfo_BlockSeal proto.InternalMessageInfo

func (m *BlockSeal) GetBlockHash() []byte {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

func (m *BlockSeal) GetExecutionReceiptHash() []byte {
	if m != nil {
		return m.ExecutionReceiptHash
	}
	return nil
}

func (m *BlockSeal) GetExecutionReceiptSignatures() [][]byte {
	if m != nil {
		return m.ExecutionReceiptSignatures
	}
	return nil
}

func (m *BlockSeal) GetResultApprovalSignatures() [][]byte {
	if m != nil {
		return m.ResultApprovalSignatures
	}
	return nil
}

type IntermediateRegisters struct {
	TransactionHash []byte `protobuf:"bytes,1,opt,name=transactionHash,proto3" json:"transactionHash,omitempty"`
	// the register changes at the end of computing the transaction
	Registers            []*Register `protobuf:"bytes,2,rep,name=registers,proto3" json:"registers,omitempty"`
	ComputeUsed          uint64      `protobuf:"varint,3,opt,name=computeUsed,proto3" json:"computeUsed,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *IntermediateRegisters) Reset()         { *m = IntermediateRegisters{} }
func (m *IntermediateRegisters) String() string { return proto.CompactTextString(m) }
func (*IntermediateRegisters) ProtoMessage()    {}
func (*IntermediateRegisters) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ccd9f92419ce3f6, []int{3}
}

func (m *IntermediateRegisters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntermediateRegisters.Unmarshal(m, b)
}
func (m *IntermediateRegisters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntermediateRegisters.Marshal(b, m, deterministic)
}
func (m *IntermediateRegisters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntermediateRegisters.Merge(m, src)
}
func (m *IntermediateRegisters) XXX_Size() int {
	return xxx_messageInfo_IntermediateRegisters.Size(m)
}
func (m *IntermediateRegisters) XXX_DiscardUnknown() {
	xxx_messageInfo_IntermediateRegisters.DiscardUnknown(m)
}

var xxx_messageInfo_IntermediateRegisters proto.InternalMessageInfo

func (m *IntermediateRegisters) GetTransactionHash() []byte {
	if m != nil {
		return m.TransactionHash
	}
	return nil
}

func (m *IntermediateRegisters) GetRegisters() []*Register {
	if m != nil {
		return m.Registers
	}
	return nil
}

func (m *IntermediateRegisters) GetComputeUsed() uint64 {
	if m != nil {
		return m.ComputeUsed
	}
	return 0
}

// Generated by each execution node individually; submitted to access nodes.
type ExecutionReceipt struct {
	PreviousReceiptHash       []byte                   `protobuf:"bytes,1,opt,name=previousReceiptHash,proto3" json:"previousReceiptHash,omitempty"`
	BlockHash                 []byte                   `protobuf:"bytes,2,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	InitialRegisters          []*Register              `protobuf:"bytes,3,rep,name=initialRegisters,proto3" json:"initialRegisters,omitempty"`
	IntermediateRegistersList []*IntermediateRegisters `protobuf:"bytes,4,rep,name=intermediateRegistersList,proto3" json:"intermediateRegistersList,omitempty"`
	Signatures                [][]byte                 `protobuf:"bytes,5,rep,name=signatures,proto3" json:"signatures,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}                 `json:"-"`
	XXX_unrecognized          []byte                   `json:"-"`
	XXX_sizecache             int32                    `json:"-"`
}

func (m *ExecutionReceipt) Reset()         { *m = ExecutionReceipt{} }
func (m *ExecutionReceipt) String() string { return proto.CompactTextString(m) }
func (*ExecutionReceipt) ProtoMessage()    {}
func (*ExecutionReceipt) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ccd9f92419ce3f6, []int{4}
}

func (m *ExecutionReceipt) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecutionReceipt.Unmarshal(m, b)
}
func (m *ExecutionReceipt) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecutionReceipt.Marshal(b, m, deterministic)
}
func (m *ExecutionReceipt) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutionReceipt.Merge(m, src)
}
func (m *ExecutionReceipt) XXX_Size() int {
	return xxx_messageInfo_ExecutionReceipt.Size(m)
}
func (m *ExecutionReceipt) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutionReceipt.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutionReceipt proto.InternalMessageInfo

func (m *ExecutionReceipt) GetPreviousReceiptHash() []byte {
	if m != nil {
		return m.PreviousReceiptHash
	}
	return nil
}

func (m *ExecutionReceipt) GetBlockHash() []byte {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

func (m *ExecutionReceipt) GetInitialRegisters() []*Register {
	if m != nil {
		return m.InitialRegisters
	}
	return nil
}

func (m *ExecutionReceipt) GetIntermediateRegistersList() []*IntermediateRegisters {
	if m != nil {
		return m.IntermediateRegistersList
	}
	return nil
}

func (m *ExecutionReceipt) GetSignatures() [][]byte {
	if m != nil {
		return m.Signatures
	}
	return nil
}

// Issued by the access nodes once they agree with the values
// for the computation that they were assigned to check.
type ResultApproval struct {
	BlockHeight             uint64   `protobuf:"varint,1,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	ExecutionReceiptHash    []byte   `protobuf:"bytes,2,opt,name=executionReceiptHash,proto3" json:"executionReceiptHash,omitempty"`
	ResultApprovalSignature []byte   `protobuf:"bytes,3,opt,name=resultApprovalSignature,proto3" json:"resultApprovalSignature,omitempty"`
	Proof                   uint64   `protobuf:"varint,4,opt,name=proof,proto3" json:"proof,omitempty"`
	Signature               []byte   `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *ResultApproval) Reset()         { *m = ResultApproval{} }
func (m *ResultApproval) String() string { return proto.CompactTextString(m) }
func (*ResultApproval) ProtoMessage()    {}
func (*ResultApproval) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ccd9f92419ce3f6, []int{5}
}

func (m *ResultApproval) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResultApproval.Unmarshal(m, b)
}
func (m *ResultApproval) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResultApproval.Marshal(b, m, deterministic)
}
func (m *ResultApproval) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResultApproval.Merge(m, src)
}
func (m *ResultApproval) XXX_Size() int {
	return xxx_messageInfo_ResultApproval.Size(m)
}
func (m *ResultApproval) XXX_DiscardUnknown() {
	xxx_messageInfo_ResultApproval.DiscardUnknown(m)
}

var xxx_messageInfo_ResultApproval proto.InternalMessageInfo

func (m *ResultApproval) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *ResultApproval) GetExecutionReceiptHash() []byte {
	if m != nil {
		return m.ExecutionReceiptHash
	}
	return nil
}

func (m *ResultApproval) GetResultApprovalSignature() []byte {
	if m != nil {
		return m.ResultApprovalSignature
	}
	return nil
}

func (m *ResultApproval) GetProof() uint64 {
	if m != nil {
		return m.Proof
	}
	return 0
}

func (m *ResultApproval) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Register struct {
	Id                   []byte   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Register) Reset()         { *m = Register{} }
func (m *Register) String() string { return proto.CompactTextString(m) }
func (*Register) ProtoMessage()    {}
func (*Register) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ccd9f92419ce3f6, []int{6}
}

func (m *Register) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Register.Unmarshal(m, b)
}
func (m *Register) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Register.Marshal(b, m, deterministic)
}
func (m *Register) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Register.Merge(m, src)
}
func (m *Register) XXX_Size() int {
	return xxx_messageInfo_Register.Size(m)
}
func (m *Register) XXX_DiscardUnknown() {
	xxx_messageInfo_Register.DiscardUnknown(m)
}

var xxx_messageInfo_Register proto.InternalMessageInfo

func (m *Register) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Register) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type AccountSignature struct {
	Account              []byte   `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Signature            []byte   `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountSignature) Reset()         { *m = AccountSignature{} }
func (m *AccountSignature) String() string { return proto.CompactTextString(m) }
func (*AccountSignature) ProtoMessage()    {}
func (*AccountSignature) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ccd9f92419ce3f6, []int{7}
}

func (m *AccountSignature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountSignature.Unmarshal(m, b)
}
func (m *AccountSignature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountSignature.Marshal(b, m, deterministic)
}
func (m *AccountSignature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountSignature.Merge(m, src)
}
func (m *AccountSignature) XXX_Size() int {
	return xxx_messageInfo_AccountSignature.Size(m)
}
func (m *AccountSignature) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountSignature.DiscardUnknown(m)
}

var xxx_messageInfo_AccountSignature proto.InternalMessageInfo

func (m *AccountSignature) GetAccount() []byte {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *AccountSignature) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Transaction struct {
	Script               []byte              `protobuf:"bytes,1,opt,name=script,proto3" json:"script,omitempty"`
	ReferenceBlockHash   []byte              `protobuf:"bytes,2,opt,name=referenceBlockHash,proto3" json:"referenceBlockHash,omitempty"`
	Nonce                uint64              `protobuf:"varint,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	ComputeLimit         uint64              `protobuf:"varint,4,opt,name=computeLimit,proto3" json:"computeLimit,omitempty"`
	PayerAccount         []byte              `protobuf:"bytes,5,opt,name=PayerAccount,proto3" json:"PayerAccount,omitempty"`
	ScriptAccounts       [][]byte            `protobuf:"bytes,6,rep,name=ScriptAccounts,proto3" json:"ScriptAccounts,omitempty"`
	Signatures           []*AccountSignature `protobuf:"bytes,7,rep,name=signatures,proto3" json:"signatures,omitempty"`
	Status               TransactionStatus   `protobuf:"varint,8,opt,name=status,proto3,enum=flow.sdk.entities.TransactionStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Transaction) Reset()         { *m = Transaction{} }
func (m *Transaction) String() string { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()    {}
func (*Transaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ccd9f92419ce3f6, []int{8}
}

func (m *Transaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transaction.Unmarshal(m, b)
}
func (m *Transaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transaction.Marshal(b, m, deterministic)
}
func (m *Transaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transaction.Merge(m, src)
}
func (m *Transaction) XXX_Size() int {
	return xxx_messageInfo_Transaction.Size(m)
}
func (m *Transaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Transaction.DiscardUnknown(m)
}

var xxx_messageInfo_Transaction proto.InternalMessageInfo

func (m *Transaction) GetScript() []byte {
	if m != nil {
		return m.Script
	}
	return nil
}

func (m *Transaction) GetReferenceBlockHash() []byte {
	if m != nil {
		return m.ReferenceBlockHash
	}
	return nil
}

func (m *Transaction) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *Transaction) GetComputeLimit() uint64 {
	if m != nil {
		return m.ComputeLimit
	}
	return 0
}

func (m *Transaction) GetPayerAccount() []byte {
	if m != nil {
		return m.PayerAccount
	}
	return nil
}

func (m *Transaction) GetScriptAccounts() [][]byte {
	if m != nil {
		return m.ScriptAccounts
	}
	return nil
}

func (m *Transaction) GetSignatures() []*AccountSignature {
	if m != nil {
		return m.Signatures
	}
	return nil
}

func (m *Transaction) GetStatus() TransactionStatus {
	if m != nil {
		return m.Status
	}
	return TransactionStatus_STATUS_UNKNOWN
}

type Collection struct {
	Transactions         []*Transaction `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
	FoundationBlockHash  []byte         `protobuf:"bytes,2,opt,name=foundationBlockHash,proto3" json:"foundationBlockHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Collection) Reset()         { *m = Collection{} }
func (m *Collection) String() string { return proto.CompactTextString(m) }
func (*Collection) ProtoMessage()    {}
func (*Collection) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ccd9f92419ce3f6, []int{9}
}

func (m *Collection) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Collection.Unmarshal(m, b)
}
func (m *Collection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Collection.Marshal(b, m, deterministic)
}
func (m *Collection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Collection.Merge(m, src)
}
func (m *Collection) XXX_Size() int {
	return xxx_messageInfo_Collection.Size(m)
}
func (m *Collection) XXX_DiscardUnknown() {
	xxx_messageInfo_Collection.DiscardUnknown(m)
}

var xxx_messageInfo_Collection proto.InternalMessageInfo

func (m *Collection) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func (m *Collection) GetFoundationBlockHash() []byte {
	if m != nil {
		return m.FoundationBlockHash
	}
	return nil
}

type SignedCollectionHash struct {
	CollectionHash       []byte   `protobuf:"bytes,1,opt,name=collectionHash,proto3" json:"collectionHash,omitempty"`
	Signatures           [][]byte `protobuf:"bytes,2,rep,name=signatures,proto3" json:"signatures,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignedCollectionHash) Reset()         { *m = SignedCollectionHash{} }
func (m *SignedCollectionHash) String() string { return proto.CompactTextString(m) }
func (*SignedCollectionHash) ProtoMessage()    {}
func (*SignedCollectionHash) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ccd9f92419ce3f6, []int{10}
}

func (m *SignedCollectionHash) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignedCollectionHash.Unmarshal(m, b)
}
func (m *SignedCollectionHash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignedCollectionHash.Marshal(b, m, deterministic)
}
func (m *SignedCollectionHash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignedCollectionHash.Merge(m, src)
}
func (m *SignedCollectionHash) XXX_Size() int {
	return xxx_messageInfo_SignedCollectionHash.Size(m)
}
func (m *SignedCollectionHash) XXX_DiscardUnknown() {
	xxx_messageInfo_SignedCollectionHash.DiscardUnknown(m)
}

var xxx_messageInfo_SignedCollectionHash proto.InternalMessageInfo

func (m *SignedCollectionHash) GetCollectionHash() []byte {
	if m != nil {
		return m.CollectionHash
	}
	return nil
}

func (m *SignedCollectionHash) GetSignatures() [][]byte {
	if m != nil {
		return m.Signatures
	}
	return nil
}

type Account struct {
	Address              []byte              `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Balance              uint64              `protobuf:"varint,2,opt,name=balance,proto3" json:"balance,omitempty"`
	Code                 []byte              `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Keys                 []*AccountPublicKey `protobuf:"bytes,4,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ccd9f92419ce3f6, []int{11}
}

func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Account) GetBalance() uint64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *Account) GetCode() []byte {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *Account) GetKeys() []*AccountPublicKey {
	if m != nil {
		return m.Keys
	}
	return nil
}

type AccountPublicKey struct {
	PublicKey            []byte   `protobuf:"bytes,1,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	SignAlgo             uint32   `protobuf:"varint,2,opt,name=signAlgo,proto3" json:"signAlgo,omitempty"`
	HashAlgo             uint32   `protobuf:"varint,3,opt,name=hashAlgo,proto3" json:"hashAlgo,omitempty"`
	Weight               uint32   `protobuf:"varint,4,opt,name=weight,proto3" json:"weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountPublicKey) Reset()         { *m = AccountPublicKey{} }
func (m *AccountPublicKey) String() string { return proto.CompactTextString(m) }
func (*AccountPublicKey) ProtoMessage()    {}
func (*AccountPublicKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ccd9f92419ce3f6, []int{12}
}

func (m *AccountPublicKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountPublicKey.Unmarshal(m, b)
}
func (m *AccountPublicKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountPublicKey.Marshal(b, m, deterministic)
}
func (m *AccountPublicKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountPublicKey.Merge(m, src)
}
func (m *AccountPublicKey) XXX_Size() int {
	return xxx_messageInfo_AccountPublicKey.Size(m)
}
func (m *AccountPublicKey) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountPublicKey.DiscardUnknown(m)
}

var xxx_messageInfo_AccountPublicKey proto.InternalMessageInfo

func (m *AccountPublicKey) GetPublicKey() []byte {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

func (m *AccountPublicKey) GetSignAlgo() uint32 {
	if m != nil {
		return m.SignAlgo
	}
	return 0
}

func (m *AccountPublicKey) GetHashAlgo() uint32 {
	if m != nil {
		return m.HashAlgo
	}
	return 0
}

func (m *AccountPublicKey) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

type Event struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	TransactionHash      []byte   `protobuf:"bytes,2,opt,name=transactionHash,proto3" json:"transactionHash,omitempty"`
	Index                uint32   `protobuf:"varint,3,opt,name=index,proto3" json:"index,omitempty"`
	Payload              []byte   `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_9ccd9f92419ce3f6, []int{13}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Event) GetTransactionHash() []byte {
	if m != nil {
		return m.TransactionHash
	}
	return nil
}

func (m *Event) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Event) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterEnum("flow.sdk.entities.TransactionStatus", TransactionStatus_name, TransactionStatus_value)
	proto.RegisterType((*BlockHeader)(nil), "flow.sdk.entities.BlockHeader")
	proto.RegisterType((*Block)(nil), "flow.sdk.entities.Block")
	proto.RegisterType((*BlockSeal)(nil), "flow.sdk.entities.BlockSeal")
	proto.RegisterType((*IntermediateRegisters)(nil), "flow.sdk.entities.IntermediateRegisters")
	proto.RegisterType((*ExecutionReceipt)(nil), "flow.sdk.entities.ExecutionReceipt")
	proto.RegisterType((*ResultApproval)(nil), "flow.sdk.entities.ResultApproval")
	proto.RegisterType((*Register)(nil), "flow.sdk.entities.Register")
	proto.RegisterType((*AccountSignature)(nil), "flow.sdk.entities.AccountSignature")
	proto.RegisterType((*Transaction)(nil), "flow.sdk.entities.Transaction")
	proto.RegisterType((*Collection)(nil), "flow.sdk.entities.Collection")
	proto.RegisterType((*SignedCollectionHash)(nil), "flow.sdk.entities.SignedCollectionHash")
	proto.RegisterType((*Account)(nil), "flow.sdk.entities.Account")
	proto.RegisterType((*AccountPublicKey)(nil), "flow.sdk.entities.AccountPublicKey")
	proto.RegisterType((*Event)(nil), "flow.sdk.entities.Event")
}

func init() { proto.RegisterFile("sdk/entities/entities.proto", fileDescriptor_9ccd9f92419ce3f6) }

var fileDescriptor_9ccd9f92419ce3f6 = []byte{
	// 1028 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xdd, 0x8e, 0xdb, 0xc4,
	0x17, 0xff, 0xdb, 0x49, 0xf6, 0xe3, 0x24, 0xdd, 0x66, 0xa7, 0xf9, 0x17, 0xb3, 0xad, 0x4a, 0x64,
	0x10, 0x44, 0x08, 0x65, 0xab, 0x70, 0x41, 0x41, 0x15, 0x52, 0xb6, 0x09, 0x25, 0x74, 0x15, 0x56,
	0x93, 0x2c, 0x48, 0xbd, 0xa0, 0x9a, 0xd8, 0x93, 0x64, 0xb4, 0x8e, 0x6d, 0x79, 0xc6, 0xdb, 0xe6,
	0x0e, 0x71, 0xc9, 0x43, 0xf0, 0x12, 0xbc, 0x0a, 0x37, 0xdc, 0x20, 0x1e, 0x05, 0xcd, 0x78, 0xec,
	0x38, 0x8e, 0xb3, 0x48, 0xdc, 0xf9, 0xfc, 0xce, 0xc7, 0x9c, 0xf3, 0x3b, 0x67, 0xce, 0x18, 0x1e,
	0x71, 0xf7, 0xe6, 0x9c, 0xfa, 0x82, 0x09, 0x46, 0x79, 0xf6, 0xd1, 0x0d, 0xa3, 0x40, 0x04, 0xe8,
	0x74, 0xee, 0x05, 0x6f, 0xbb, 0xdc, 0xbd, 0xe9, 0xa6, 0x8a, 0xb3, 0x0f, 0x16, 0x41, 0xb0, 0xf0,
	0xe8, 0xb9, 0x32, 0x98, 0xc5, 0xf3, 0x73, 0xc1, 0x56, 0x94, 0x0b, 0xb2, 0x0a, 0x13, 0x1f, 0xdb,
	0x81, 0xfa, 0x85, 0x17, 0x38, 0x37, 0xdf, 0x52, 0xe2, 0xd2, 0x08, 0x9d, 0x80, 0x39, 0x1a, 0x58,
	0x46, 0xdb, 0xe8, 0x34, 0xb0, 0x39, 0x1a, 0xa0, 0xcf, 0xe0, 0x34, 0x8c, 0xe8, 0x2d, 0x0b, 0x62,
	0x9e, 0x98, 0x11, 0xbe, 0xb4, 0x4c, 0xa5, 0xde, 0x55, 0xa0, 0x87, 0x70, 0xb0, 0xa4, 0x6c, 0xb1,
	0x14, 0x56, 0xa5, 0x6d, 0x74, 0xaa, 0x58, 0x4b, 0xf6, 0xdf, 0x26, 0xd4, 0x94, 0x15, 0xb2, 0xe0,
	0xd0, 0x59, 0x12, 0xe6, 0xeb, 0x43, 0x8e, 0x71, 0x2a, 0xe6, 0x7c, 0xcd, 0xbc, 0x6f, 0x79, 0x06,
	0x95, 0x7d, 0x19, 0x3c, 0x83, 0xe3, 0xac, 0x42, 0xab, 0xda, 0x36, 0x3a, 0xf5, 0xde, 0x59, 0x37,
	0xe1, 0xa0, 0x9b, 0x72, 0xd0, 0x9d, 0xa6, 0x16, 0x78, 0x63, 0x8c, 0xde, 0xc0, 0x43, 0xce, 0x16,
	0x3e, 0x75, 0x5f, 0x04, 0x9e, 0x47, 0x1d, 0xc1, 0x02, 0x5f, 0x46, 0xa4, 0xdc, 0xaa, 0xb5, 0x2b,
	0x9d, 0x7a, 0xef, 0x93, 0xee, 0x0e, 0xbb, 0xdd, 0x49, 0x89, 0x03, 0xde, 0x13, 0x06, 0x3d, 0x07,
	0x98, 0xc9, 0x3c, 0x27, 0x94, 0x78, 0xdc, 0x3a, 0x50, 0x41, 0x1f, 0x97, 0x04, 0xbd, 0x48, 0x8d,
	0x70, 0xce, 0x1e, 0x3d, 0x01, 0x90, 0x71, 0x89, 0x88, 0x23, 0xca, 0xad, 0xc3, 0x76, 0xa5, 0xd3,
	0xc0, 0x39, 0xc4, 0xfe, 0xd3, 0x80, 0xe3, 0xcc, 0x13, 0x3d, 0x86, 0xe3, 0x59, 0x46, 0x56, 0xd2,
	0xcd, 0x0d, 0x80, 0x7a, 0xd0, 0xa2, 0xef, 0xa8, 0x13, 0xcb, 0xe4, 0x30, 0x75, 0x28, 0x0b, 0x45,
	0xae, 0xaf, 0xa5, 0x3a, 0xf4, 0x35, 0x9c, 0x15, 0xf1, 0xc9, 0x26, 0x9f, 0x8a, 0xca, 0xe7, 0x0e,
	0x0b, 0xf4, 0x15, 0x58, 0x11, 0xe5, 0xb1, 0x27, 0xfa, 0x61, 0x18, 0x05, 0xb7, 0xc4, 0xcb, 0x79,
	0x57, 0x95, 0xf7, 0x5e, 0xbd, 0xfd, 0x9b, 0x01, 0xff, 0x1f, 0xf9, 0x82, 0x46, 0x2b, 0xea, 0x32,
	0x22, 0x28, 0xa6, 0x0b, 0xc6, 0x05, 0x8d, 0x38, 0xea, 0xc0, 0x7d, 0x11, 0x11, 0x9f, 0x93, 0x8c,
	0x68, 0x5d, 0x6d, 0x11, 0x46, 0x5f, 0xc2, 0x71, 0x94, 0xba, 0x59, 0xa6, 0x22, 0xff, 0x51, 0x09,
	0xf9, 0x69, 0x68, 0xbc, 0xb1, 0x46, 0x6d, 0xa8, 0x3b, 0xc1, 0x2a, 0x8c, 0x05, 0xbd, 0xe6, 0xd4,
	0xd5, 0xa3, 0x9d, 0x87, 0xec, 0xdf, 0x4d, 0x68, 0x0e, 0x0b, 0xb5, 0xa3, 0xa7, 0xf0, 0x20, 0x9d,
	0xcf, 0x3c, 0xc9, 0x49, 0x7e, 0x65, 0xaa, 0xed, 0xae, 0x99, 0xc5, 0xae, 0xbd, 0x84, 0x26, 0xf3,
	0x99, 0x60, 0xc4, 0xcb, 0xea, 0x57, 0xbc, 0xff, 0x4b, 0x21, 0x3b, 0x4e, 0x68, 0x0e, 0xef, 0xb3,
	0x32, 0x36, 0x2f, 0x19, 0x17, 0xaa, 0x17, 0xf5, 0x5e, 0xa7, 0x24, 0x62, 0x69, 0x07, 0xf0, 0xfe,
	0x50, 0x85, 0x91, 0xad, 0xed, 0x8c, 0xec, 0x1f, 0x06, 0x9c, 0xe0, 0xad, 0x9e, 0x4b, 0xaa, 0x93,
	0x82, 0x93, 0x4d, 0x60, 0x24, 0x54, 0xe7, 0xa0, 0xff, 0x34, 0xbb, 0xcf, 0xe0, 0xbd, 0x3d, 0xb3,
	0xa5, 0x17, 0xc9, 0x3e, 0x35, 0x6a, 0x41, 0x2d, 0x8c, 0x82, 0x60, 0xae, 0x56, 0x49, 0x15, 0x27,
	0x82, 0xec, 0x53, 0x56, 0x86, 0x55, 0x4b, 0xfa, 0x94, 0x01, 0xf6, 0x53, 0x38, 0x4a, 0x79, 0x90,
	0xeb, 0x94, 0xb9, 0xe9, 0x3a, 0x65, 0xae, 0x8c, 0x77, 0x4b, 0xbc, 0x98, 0xea, 0x74, 0x13, 0xc1,
	0xfe, 0x0e, 0x9a, 0x7d, 0xc7, 0x09, 0x62, 0x7f, 0x73, 0x61, 0xe4, 0xa2, 0x24, 0x09, 0xa6, 0xdd,
	0x53, 0x71, 0xfb, 0x74, 0xb3, 0x78, 0xfa, 0x5f, 0x26, 0xd4, 0xa7, 0x9b, 0xd9, 0x97, 0x6b, 0x95,
	0x3b, 0x11, 0x0b, 0xd3, 0x30, 0x5a, 0x42, 0x5d, 0x40, 0x11, 0x9d, 0xd3, 0x88, 0xfa, 0x0e, 0x2d,
	0x6e, 0xf6, 0x12, 0x8d, 0xcc, 0xdc, 0x0f, 0x7c, 0x87, 0xea, 0xf1, 0x4f, 0x04, 0x64, 0x43, 0x43,
	0xdf, 0x83, 0x4b, 0xb6, 0x62, 0x42, 0xd3, 0xb4, 0x85, 0x49, 0x9b, 0x2b, 0xb2, 0xa6, 0x91, 0x2e,
	0x51, 0x13, 0xb6, 0x85, 0xa1, 0x8f, 0xe1, 0x64, 0xa2, 0xf2, 0xd2, 0x40, 0xb2, 0x1f, 0x1b, 0xb8,
	0x80, 0xa2, 0x17, 0x3b, 0x5b, 0xb0, 0xde, 0xfb, 0xb0, 0x64, 0x56, 0x8b, 0x74, 0xe6, 0xe7, 0x0e,
	0x3d, 0x87, 0x03, 0x2e, 0x88, 0x88, 0xb9, 0x75, 0xd4, 0x36, 0x3a, 0x27, 0xbd, 0x8f, 0x4a, 0x02,
	0xe4, 0x28, 0x9c, 0x28, 0x5b, 0xac, 0x7d, 0xec, 0x5f, 0x0c, 0x80, 0xcd, 0x6e, 0x47, 0x17, 0xd0,
	0xc8, 0xad, 0x1a, 0x6e, 0x19, 0x2a, 0xa7, 0x27, 0x77, 0x87, 0xc4, 0x5b, 0x3e, 0x72, 0x53, 0xcc,
	0x83, 0xd8, 0x77, 0x89, 0x14, 0x8b, 0xcd, 0x28, 0x53, 0xd9, 0x3f, 0x41, 0xab, 0xec, 0xed, 0x91,
	0x3c, 0x3a, 0x5b, 0x88, 0xee, 0x7a, 0x01, 0x2d, 0x5c, 0x4d, 0x73, 0xe7, 0x6a, 0xfe, 0x6a, 0xc0,
	0x61, 0xda, 0x1b, 0x39, 0x89, 0xae, 0x1b, 0x51, 0xce, 0xb3, 0x49, 0x4c, 0x44, 0xa9, 0x99, 0x11,
	0x8f, 0xc8, 0xa9, 0x48, 0xde, 0xec, 0x54, 0x44, 0x08, 0xaa, 0x4e, 0xe0, 0xa6, 0xd7, 0x4b, 0x7d,
	0xa3, 0x2f, 0xa0, 0x7a, 0x43, 0xd7, 0x5c, 0x6f, 0x98, 0x3b, 0xba, 0x76, 0x15, 0xcf, 0x3c, 0xe6,
	0xbc, 0xa2, 0x6b, 0xac, 0x1c, 0xec, 0x9f, 0x8d, 0xec, 0x7e, 0x64, 0x2a, 0x79, 0x0b, 0xc2, 0x54,
	0x48, 0x5f, 0xb8, 0x0c, 0x40, 0x67, 0x70, 0x24, 0xab, 0xe9, 0x7b, 0x8b, 0x40, 0xa5, 0x76, 0x0f,
	0x67, 0xb2, 0xd4, 0x2d, 0x09, 0x5f, 0x2a, 0x5d, 0x25, 0xd1, 0xa5, 0xb2, 0xbc, 0x2d, 0x6f, 0x93,
	0xd5, 0x53, 0x55, 0x1a, 0x2d, 0xd9, 0x31, 0xd4, 0x86, 0xb7, 0xd4, 0x17, 0xb2, 0x30, 0xb1, 0x0e,
	0xa9, 0xfe, 0x79, 0x51, 0xdf, 0x65, 0x8f, 0x90, 0x59, 0xfe, 0x08, 0xb5, 0xa0, 0xc6, 0x7c, 0x97,
	0xbe, 0xd3, 0xe7, 0x26, 0x82, 0xa4, 0x31, 0x24, 0x6b, 0x2f, 0x20, 0xae, 0x3a, 0xb5, 0x81, 0x53,
	0xf1, 0xd3, 0x35, 0x9c, 0xee, 0x0c, 0x22, 0x42, 0x70, 0x32, 0x99, 0xf6, 0xa7, 0xd7, 0x93, 0x37,
	0xd7, 0xe3, 0x57, 0xe3, 0xef, 0x7f, 0x1c, 0x37, 0xff, 0x97, 0xc3, 0xae, 0x86, 0xe3, 0xc1, 0x68,
	0xfc, 0xb2, 0x69, 0xa0, 0x16, 0x34, 0x35, 0xf6, 0xcd, 0x68, 0xdc, 0xbf, 0x1c, 0xbd, 0x1e, 0x0e,
	0x9a, 0x26, 0x7a, 0x00, 0xf7, 0x35, 0x8a, 0x87, 0x3f, 0x0c, 0xf1, 0x74, 0x38, 0x68, 0x56, 0xd0,
	0x29, 0xdc, 0xd3, 0xe0, 0x64, 0xd8, 0xbf, 0x1c, 0x0e, 0x9a, 0xd5, 0x0b, 0x78, 0x7d, 0x94, 0xf6,
	0x65, 0x76, 0xa0, 0xfe, 0x9c, 0x3e, 0xff, 0x27, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x40, 0x0a, 0x28,
	0x7d, 0x0a, 0x00, 0x00,
}
