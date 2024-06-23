package messages7

import (
	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type CtrlKeepAlive struct{}

func (msg *CtrlKeepAlive) MsgId() int {
	return network7.MsgCtrlKeepAlive
}

func (msg *CtrlKeepAlive) MsgType() network7.MsgType {
	return network7.TypeControl
}

func (msg *CtrlKeepAlive) System() bool {
	return false
}

func (msg *CtrlKeepAlive) Vital() bool {
	return false
}

func (msg *CtrlKeepAlive) Pack() []byte {
	return []byte{network7.MsgCtrlKeepAlive}
}

func (msg *CtrlKeepAlive) Unpack(u *packer.Unpacker) error {
	return nil
}

func (msg *CtrlKeepAlive) Header() *chunk7.ChunkHeader {
	return nil
}

func (msg *CtrlKeepAlive) SetHeader(header chunk7.ChunkHeader) {}
