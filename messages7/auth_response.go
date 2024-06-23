package messages7

import (
	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

// this message is unused in the official 0.7.5 implementation
type AuthResponse struct {
	ChunkHeader chunk7.ChunkHeader
}

func (msg *AuthResponse) MsgId() int {
	return network7.MsgSysAuthResponse
}

func (msg *AuthResponse) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *AuthResponse) System() bool {
	return true
}

func (msg *AuthResponse) Vital() bool {
	return true
}

func (msg *AuthResponse) Pack() []byte {
	return []byte{}
}

func (msg *AuthResponse) Unpack(u *packer.Unpacker) error {
	return nil
}

func (msg *AuthResponse) Header() *chunk7.ChunkHeader {
	return &msg.ChunkHeader
}

func (msg *AuthResponse) SetHeader(header chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}
