package messages7

import (
	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

// this message is unused in the official 0.7.5 implementation
type Error struct {
	ChunkHeader chunk7.ChunkHeader
}

func (msg *Error) MsgId() int {
	return network7.MsgSysError
}

func (msg *Error) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *Error) System() bool {
	return true
}

func (msg *Error) Vital() bool {
	return true
}

func (msg *Error) Pack() []byte {
	return []byte{}
}

func (msg *Error) Unpack(u *packer.Unpacker) error {
	return nil
}

func (msg *Error) Header() *chunk7.ChunkHeader {
	return &msg.ChunkHeader
}

func (msg *Error) SetHeader(header chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}
