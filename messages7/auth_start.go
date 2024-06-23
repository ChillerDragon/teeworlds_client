package messages7

import (
	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

// this message is unused in the official 0.7.5 implementation
type AuthStart struct {
	ChunkHeader chunk7.ChunkHeader
}

func (msg *AuthStart) MsgId() int {
	return network7.MsgSysAuthStart
}

func (msg *AuthStart) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *AuthStart) System() bool {
	return true
}

func (msg *AuthStart) Vital() bool {
	return true
}

func (msg *AuthStart) Pack() []byte {
	return []byte{}
}

func (msg *AuthStart) Unpack(u *packer.Unpacker) error {
	return nil
}

func (msg *AuthStart) Header() *chunk7.ChunkHeader {
	return &msg.ChunkHeader
}

func (msg *AuthStart) SetHeader(header chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}
