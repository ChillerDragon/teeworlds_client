package messages7

import (
	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

// this message is unused in the official 0.7.5 implementation
type SnapSmall struct {
	ChunkHeader chunk7.ChunkHeader
}

func (msg *SnapSmall) MsgId() int {
	return network7.MsgSysSnapSmall
}

func (msg *SnapSmall) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *SnapSmall) System() bool {
	return true
}

func (msg *SnapSmall) Vital() bool {
	return false
}

func (msg *SnapSmall) Pack() []byte {
	return []byte{}
}

func (msg *SnapSmall) Unpack(u *packer.Unpacker) error {
	return nil
}

func (msg *SnapSmall) Header() *chunk7.ChunkHeader {
	return &msg.ChunkHeader
}

func (msg *SnapSmall) SetHeader(header chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}
