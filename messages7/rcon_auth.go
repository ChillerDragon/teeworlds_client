package messages7

import (
	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type RconAuth struct {
	ChunkHeader chunk7.ChunkHeader

	Password string
}

func (msg *RconAuth) MsgId() int {
	return network7.MsgSysRconAuth
}

func (msg *RconAuth) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *RconAuth) System() bool {
	return true
}

func (msg *RconAuth) Vital() bool {
	return true
}

func (msg *RconAuth) Pack() []byte {
	return packer.PackString(msg.Password)
}

func (msg *RconAuth) Unpack(u *packer.Unpacker) (err error) {
	msg.Password, err = u.NextString()
	return err
}

func (msg *RconAuth) Header() *chunk7.ChunkHeader {
	return &msg.ChunkHeader
}

func (msg *RconAuth) SetHeader(header chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}
