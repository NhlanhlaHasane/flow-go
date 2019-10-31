// Code generated by capnpc-go. DO NOT EDIT.

package captain

import (
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type BlockProposal struct{ capnp.Struct }

// BlockProposal_TypeID is the unique identifier for the type BlockProposal.
const BlockProposal_TypeID = 0x9c6b5c46ecb4a72e

func NewBlockProposal(s *capnp.Segment) (BlockProposal, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return BlockProposal{st}, err
}

func NewRootBlockProposal(s *capnp.Segment) (BlockProposal, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return BlockProposal{st}, err
}

func ReadRootBlockProposal(msg *capnp.Message) (BlockProposal, error) {
	root, err := msg.RootPtr()
	return BlockProposal{root.Struct()}, err
}

func (s BlockProposal) String() string {
	str, _ := text.Marshal(0x9c6b5c46ecb4a72e, s.Struct)
	return str
}

func (s BlockProposal) Header() (BlockHeader, error) {
	p, err := s.Struct.Ptr(0)
	return BlockHeader{Struct: p.Struct()}, err
}

func (s BlockProposal) HasHeader() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s BlockProposal) SetHeader(v BlockHeader) error {
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewHeader sets the header field to a newly
// allocated BlockHeader struct, preferring placement in s's segment.
func (s BlockProposal) NewHeader() (BlockHeader, error) {
	ss, err := NewBlockHeader(s.Struct.Segment())
	if err != nil {
		return BlockHeader{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

// BlockProposal_List is a list of BlockProposal.
type BlockProposal_List struct{ capnp.List }

// NewBlockProposal creates a new list of BlockProposal.
func NewBlockProposal_List(s *capnp.Segment, sz int32) (BlockProposal_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return BlockProposal_List{l}, err
}

func (s BlockProposal_List) At(i int) BlockProposal { return BlockProposal{s.List.Struct(i)} }

func (s BlockProposal_List) Set(i int, v BlockProposal) error { return s.List.SetStruct(i, v.Struct) }

func (s BlockProposal_List) String() string {
	str, _ := text.MarshalList(0x9c6b5c46ecb4a72e, s.List)
	return str
}

// BlockProposal_Promise is a wrapper for a BlockProposal promised by a client call.
type BlockProposal_Promise struct{ *capnp.Pipeline }

func (p BlockProposal_Promise) Struct() (BlockProposal, error) {
	s, err := p.Pipeline.Struct()
	return BlockProposal{s}, err
}

func (p BlockProposal_Promise) Header() BlockHeader_Promise {
	return BlockHeader_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

type BlockVote struct{ capnp.Struct }

// BlockVote_TypeID is the unique identifier for the type BlockVote.
const BlockVote_TypeID = 0x90ad957dba1800a4

func NewBlockVote(s *capnp.Segment) (BlockVote, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return BlockVote{st}, err
}

func NewRootBlockVote(s *capnp.Segment) (BlockVote, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return BlockVote{st}, err
}

func ReadRootBlockVote(msg *capnp.Message) (BlockVote, error) {
	root, err := msg.RootPtr()
	return BlockVote{root.Struct()}, err
}

func (s BlockVote) String() string {
	str, _ := text.Marshal(0x90ad957dba1800a4, s.Struct)
	return str
}

func (s BlockVote) Hash() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s BlockVote) HasHash() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s BlockVote) SetHash(v []byte) error {
	return s.Struct.SetData(0, v)
}

// BlockVote_List is a list of BlockVote.
type BlockVote_List struct{ capnp.List }

// NewBlockVote creates a new list of BlockVote.
func NewBlockVote_List(s *capnp.Segment, sz int32) (BlockVote_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return BlockVote_List{l}, err
}

func (s BlockVote_List) At(i int) BlockVote { return BlockVote{s.List.Struct(i)} }

func (s BlockVote_List) Set(i int, v BlockVote) error { return s.List.SetStruct(i, v.Struct) }

func (s BlockVote_List) String() string {
	str, _ := text.MarshalList(0x90ad957dba1800a4, s.List)
	return str
}

// BlockVote_Promise is a wrapper for a BlockVote promised by a client call.
type BlockVote_Promise struct{ *capnp.Pipeline }

func (p BlockVote_Promise) Struct() (BlockVote, error) {
	s, err := p.Pipeline.Struct()
	return BlockVote{s}, err
}

type BlockCommit struct{ capnp.Struct }

// BlockCommit_TypeID is the unique identifier for the type BlockCommit.
const BlockCommit_TypeID = 0xa50191d15439269b

func NewBlockCommit(s *capnp.Segment) (BlockCommit, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return BlockCommit{st}, err
}

func NewRootBlockCommit(s *capnp.Segment) (BlockCommit, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1})
	return BlockCommit{st}, err
}

func ReadRootBlockCommit(msg *capnp.Message) (BlockCommit, error) {
	root, err := msg.RootPtr()
	return BlockCommit{root.Struct()}, err
}

func (s BlockCommit) String() string {
	str, _ := text.Marshal(0xa50191d15439269b, s.Struct)
	return str
}

func (s BlockCommit) Hash() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s BlockCommit) HasHash() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s BlockCommit) SetHash(v []byte) error {
	return s.Struct.SetData(0, v)
}

// BlockCommit_List is a list of BlockCommit.
type BlockCommit_List struct{ capnp.List }

// NewBlockCommit creates a new list of BlockCommit.
func NewBlockCommit_List(s *capnp.Segment, sz int32) (BlockCommit_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 1}, sz)
	return BlockCommit_List{l}, err
}

func (s BlockCommit_List) At(i int) BlockCommit { return BlockCommit{s.List.Struct(i)} }

func (s BlockCommit_List) Set(i int, v BlockCommit) error { return s.List.SetStruct(i, v.Struct) }

func (s BlockCommit_List) String() string {
	str, _ := text.MarshalList(0xa50191d15439269b, s.List)
	return str
}

// BlockCommit_Promise is a wrapper for a BlockCommit promised by a client call.
type BlockCommit_Promise struct{ *capnp.Pipeline }

func (p BlockCommit_Promise) Struct() (BlockCommit, error) {
	s, err := p.Pipeline.Struct()
	return BlockCommit{s}, err
}

const schema_ca03600dec38188d = "x\xda2\x08ct`2d\x9d\xcf\xc1\xc0\x10\x98\xc2" +
	"\xca\xf6w\x89\xc4\xae\xda\xa9k'\x08\x0a0\xfe\xef\x95" +
	"\xb0x\xc3\x9b\xc0|\x8a\x81\x95\x91\x9d\x81A\xd8\x94\xe5" +
	"\x92\xb0#\x0b\x88e\xcbb\xcf\xc0\xf8_o\xf9\x967" +
	"n1\xd9s\x18\xb0\xa8\x8de\xf9$\x9c\x09V\x9b\x0a" +
	"V;[\xcd2\xe4\xe2D\xc6\xa5\xd8\xd4\xb6\xb2<\x12" +
	"\x9e\x08V\xdb\xcbb\xcf\xf0\xf9\x7fr~NJqI" +
	"i\x1aS\x9a^rbA^\x81\x95SN~rv" +
	"X~\x09cj\x00#c \x0b3\x0b\x03\x03\x0b#" +
	"\x03\x83 \xaf\x16\x03C \x073c\xa0\x08\x13#\x7f" +
	"Fbq\x06#/\x03\x13#/\x03#\x0e#\x02\x8a" +
	"\xf2\xe5\x0b\xf2\x8b\x13s\xd0\x8c\xb1B\x18c\x9f\x91\x9a" +
	"\x98\x92Z\xc4(\xf0?\xb3\xa0\xe3\xf9\x09]\xbd\xad\x0c" +
	"\x0c\x8c\x8c\x028\x8dt\xce\xcfe\xcf\xcd,!\xd2]" +
	"\x80\x00\x00\x00\xff\xff\x81\xf2Y\xd8"

func init() {
	schemas.Register(schema_ca03600dec38188d,
		0x90ad957dba1800a4,
		0x9c6b5c46ecb4a72e,
		0xa50191d15439269b)
}
