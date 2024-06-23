package messages7

import (
	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type ConReady struct {
	ChunkHeader chunk7.ChunkHeader
}

func (msg *ConReady) MsgId() int {
	return network7.MsgSysConReady
}

func (msg *ConReady) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *ConReady) System() bool {
	return true
}

func (msg *ConReady) Vital() bool {
	return true
}

func (msg *ConReady) Pack() []byte {
	return []byte{}
}

func (msg *ConReady) Unpack(u *packer.Unpacker) error {
	return nil
}

func (msg *ConReady) Header() *chunk7.ChunkHeader {
	return &msg.ChunkHeader
}

func (msg *ConReady) SetHeader(header chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}
