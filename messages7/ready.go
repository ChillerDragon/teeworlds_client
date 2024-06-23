package messages7

import (
	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type Ready struct {
	ChunkHeader chunk7.ChunkHeader
}

func (msg *Ready) MsgId() int {
	return network7.MsgSysReady
}

func (msg *Ready) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *Ready) System() bool {
	return true
}

func (msg *Ready) Vital() bool {
	return true
}

func (msg *Ready) Pack() []byte {
	return []byte{}
}

func (msg *Ready) Unpack(u *packer.Unpacker) error {
	return nil
}

func (msg *Ready) Header() *chunk7.ChunkHeader {
	return &msg.ChunkHeader
}

func (msg *Ready) SetHeader(header chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}
