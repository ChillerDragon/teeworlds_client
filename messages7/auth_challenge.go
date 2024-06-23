package messages7

import (
	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

// this message is unused in the official 0.7.5 implementation
type AuthChallenge struct {
	ChunkHeader chunk7.ChunkHeader
}

func (msg *AuthChallenge) MsgId() int {
	return network7.MsgSysAuthChallenge
}

func (msg *AuthChallenge) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *AuthChallenge) System() bool {
	return true
}

func (msg *AuthChallenge) Vital() bool {
	return true
}

func (msg *AuthChallenge) Pack() []byte {
	return []byte{}
}

func (msg *AuthChallenge) Unpack(u *packer.Unpacker) error {
	return nil
}

func (msg *AuthChallenge) Header() *chunk7.ChunkHeader {
	return &msg.ChunkHeader
}

func (msg *AuthChallenge) SetHeader(header chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}
