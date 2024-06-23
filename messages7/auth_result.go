package messages7

import (
	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

// this message is unused in the official 0.7.5 implementation
type AuthResult struct {
	ChunkHeader chunk7.ChunkHeader
}

func (msg *AuthResult) MsgId() int {
	return network7.MsgSysAuthResult
}

func (msg *AuthResult) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *AuthResult) System() bool {
	return true
}

func (msg *AuthResult) Vital() bool {
	return true
}

func (msg *AuthResult) Pack() []byte {
	return []byte{}
}

func (msg *AuthResult) Unpack(u *packer.Unpacker) error {
	return nil
}

func (msg *AuthResult) Header() *chunk7.ChunkHeader {
	return &msg.ChunkHeader
}

func (msg *AuthResult) SetHeader(header chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}
