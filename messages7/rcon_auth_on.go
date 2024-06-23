package messages7

import (
	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type RconAuthOn struct {
	ChunkHeader chunk7.ChunkHeader
}

func (msg *RconAuthOn) MsgId() int {
	return network7.MsgSysRconAuthOn
}

func (msg *RconAuthOn) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *RconAuthOn) System() bool {
	return true
}

func (msg *RconAuthOn) Vital() bool {
	return true
}

func (msg *RconAuthOn) Pack() []byte {
	return []byte{}
}

func (msg *RconAuthOn) Unpack(u *packer.Unpacker) error {
	return nil
}

func (msg *RconAuthOn) Header() *chunk7.ChunkHeader {
	return &msg.ChunkHeader
}

func (msg *RconAuthOn) SetHeader(header chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}
