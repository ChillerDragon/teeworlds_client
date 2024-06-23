package messages7

import (
	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type RconAuthOff struct {
	ChunkHeader chunk7.ChunkHeader
}

func (msg *RconAuthOff) MsgId() int {
	return network7.MsgSysRconAuthOff
}

func (msg *RconAuthOff) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *RconAuthOff) System() bool {
	return true
}

func (msg *RconAuthOff) Vital() bool {
	return true
}

func (msg *RconAuthOff) Pack() []byte {
	return []byte{}
}

func (msg *RconAuthOff) Unpack(u *packer.Unpacker) error {
	return nil
}

func (msg *RconAuthOff) Header() *chunk7.ChunkHeader {
	return &msg.ChunkHeader
}

func (msg *RconAuthOff) SetHeader(header chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}
