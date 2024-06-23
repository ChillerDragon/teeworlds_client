package messages7

import (
	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type PingReply struct {
	ChunkHeader chunk7.ChunkHeader
}

func (msg *PingReply) MsgId() int {
	return network7.MsgSysPingReply
}

func (msg *PingReply) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *PingReply) System() bool {
	return true
}

func (msg *PingReply) Vital() bool {
	return true
}

func (msg *PingReply) Pack() []byte {
	return []byte{}
}

func (msg *PingReply) Unpack(u *packer.Unpacker) error {
	return nil
}

func (msg *PingReply) Header() *chunk7.ChunkHeader {
	return &msg.ChunkHeader
}

func (msg *PingReply) SetHeader(header chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}
