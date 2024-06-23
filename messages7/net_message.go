package messages7

import (
	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type NetMessage interface {
	MsgId() int
	MsgType() network7.MsgType
	System() bool
	Vital() bool
	Pack() []byte
	Unpack(u *packer.Unpacker) error

	Header() *chunk7.ChunkHeader
	SetHeader(header chunk7.ChunkHeader)
}
