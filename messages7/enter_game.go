package messages7

import (
	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type EnterGame struct {
	ChunkHeader chunk7.ChunkHeader
}

func (msg *EnterGame) MsgId() int {
	return network7.MsgSysEnterGame
}

func (msg *EnterGame) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *EnterGame) System() bool {
	return true
}

func (msg *EnterGame) Vital() bool {
	return true
}

func (msg *EnterGame) Pack() []byte {
	return []byte{}
}

func (msg *EnterGame) Unpack(u *packer.Unpacker) error {
	return nil

}

func (msg *EnterGame) Header() *chunk7.ChunkHeader {
	return &msg.ChunkHeader
}

func (msg *EnterGame) SetHeader(header chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}
