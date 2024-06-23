package messages7

import (
	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type ReadyToEnter struct {
	ChunkHeader chunk7.ChunkHeader
}

func (msg *ReadyToEnter) MsgId() int {
	return network7.MsgGameReadyToEnter
}

func (msg *ReadyToEnter) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *ReadyToEnter) System() bool {
	return false
}

func (msg *ReadyToEnter) Vital() bool {
	return true
}

func (msg *ReadyToEnter) Pack() []byte {
	return []byte{}
}

func (msg *ReadyToEnter) Unpack(u *packer.Unpacker) error {
	return nil
}

func (msg *ReadyToEnter) Header() *chunk7.ChunkHeader {
	return &msg.ChunkHeader
}

func (msg *ReadyToEnter) SetHeader(header chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}
