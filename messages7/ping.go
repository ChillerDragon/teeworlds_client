package messages7

import (
	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type Ping struct {
	ChunkHeader chunk7.ChunkHeader
}

func (msg *Ping) MsgId() int {
	return network7.MsgSysPing
}

func (msg *Ping) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *Ping) System() bool {
	return true
}

func (msg *Ping) Vital() bool {
	return true
}

func (msg *Ping) Pack() []byte {
	return []byte{}
}

func (msg *Ping) Unpack(u *packer.Unpacker) error {
	return nil
}

func (msg *Ping) Header() *chunk7.ChunkHeader {
	return &msg.ChunkHeader
}

func (msg *Ping) SetHeader(header chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}
