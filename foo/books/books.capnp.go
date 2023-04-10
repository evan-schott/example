// Code generated by capnpc-go. DO NOT EDIT.

package books

import (
	capnp "capnproto.org/go/capnp/v3"
	text "capnproto.org/go/capnp/v3/encoding/text"
	schemas "capnproto.org/go/capnp/v3/schemas"
)

type Book capnp.Struct

// Book_TypeID is the unique identifier for the type Book.
const Book_TypeID = 0x8100cc88d7d4d47c

func NewBook(s *capnp.Segment) (Book, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Book(st), err
}

func NewRootBook(s *capnp.Segment) (Book, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1})
	return Book(st), err
}

func ReadRootBook(msg *capnp.Message) (Book, error) {
	root, err := msg.Root()
	return Book(root.Struct()), err
}

func (s Book) String() string {
	str, _ := text.Marshal(0x8100cc88d7d4d47c, capnp.Struct(s))
	return str
}

func (s Book) EncodeAsPtr(seg *capnp.Segment) capnp.Ptr {
	return capnp.Struct(s).EncodeAsPtr(seg)
}

func (Book) DecodeFromPtr(p capnp.Ptr) Book {
	return Book(capnp.Struct{}.DecodeFromPtr(p))
}

func (s Book) ToPtr() capnp.Ptr {
	return capnp.Struct(s).ToPtr()
}
func (s Book) IsValid() bool {
	return capnp.Struct(s).IsValid()
}

func (s Book) Message() *capnp.Message {
	return capnp.Struct(s).Message()
}

func (s Book) Segment() *capnp.Segment {
	return capnp.Struct(s).Segment()
}
func (s Book) Title() (string, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.Text(), err
}

func (s Book) HasTitle() bool {
	return capnp.Struct(s).HasPtr(0)
}

func (s Book) TitleBytes() ([]byte, error) {
	p, err := capnp.Struct(s).Ptr(0)
	return p.TextBytes(), err
}

func (s Book) SetTitle(v string) error {
	return capnp.Struct(s).SetText(0, v)
}

func (s Book) PageCount() int32 {
	return int32(capnp.Struct(s).Uint32(0))
}

func (s Book) SetPageCount(v int32) {
	capnp.Struct(s).SetUint32(0, uint32(v))
}

// Book_List is a list of Book.
type Book_List = capnp.StructList[Book]

// NewBook creates a new list of Book.
func NewBook_List(s *capnp.Segment, sz int32) (Book_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 1}, sz)
	return capnp.StructList[Book](l), err
}

// Book_Future is a wrapper for a Book promised by a client call.
type Book_Future struct{ *capnp.Future }

func (f Book_Future) Struct() (Book, error) {
	p, err := f.Future.Ptr()
	return Book(p.Struct()), err
}

const schema_85d3acc39d94e0f8 = "x\xda\x12Ht`1\xe4\xdd\xcf\xc8\xc0\x14(\xc2\xca" +
	"\xb6\xbf\xe6\xca\x95\xeb\x1dg\x1a\x03\x05\x18\x19\xff\xffx" +
	"0e\xee\xe15\x97[\x19X\x19\xd9\x19\x18\x04\x8f\xae" +
	"\x12<\x0b\xa2O\x963\xe8\xfeO\xcb\xcf\xd7O\xca\xcf" +
	"\xcff,\xd6KN,\xc8+\xb0\xe2w\xca\xcf\xcf\x0e" +
	"`d\x0c\xe4`fa``ad`\x10\xd44b" +
	"`\x08Taf\x0c4`bdd\x14a\x04\x89\xe9" +
	"\x0610\x04\xea03\x06Z01\xca\x97d\x96\xe4" +
	"\xa42\xf2001\xf200\xfe/HLOu\xce" +
	"/\xcdc`,ada`bda`\x04\x04\x00" +
	"\x00\xff\xff(3&b"

func init() {
	schemas.Register(schema_85d3acc39d94e0f8,
		0x8100cc88d7d4d47c)
}
