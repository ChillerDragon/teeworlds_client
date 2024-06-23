package messages7

import (
	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type CtrlClose struct {
	Reason string
}

func (msg *CtrlClose) MsgId() int {
	return network7.MsgCtrlClose
}

func (msg *CtrlClose) MsgType() network7.MsgType {
	return network7.TypeControl
}

func (msg *CtrlClose) System() bool {
	return false
}

func (msg *CtrlClose) Vital() bool {
	return false
}

func (msg *CtrlClose) Pack() []byte {
	p := packer.NewPacker(make([]byte, 0, 1+len(msg.Reason)+1))
	p.AddByte(network7.MsgCtrlClose)
	p.AddString(msg.Reason)
	return p.Bytes()
}

func (msg *CtrlClose) Unpack(u *packer.Unpacker) (err error) {
	msg.Reason, err = u.NextString()
	return err
}

func (msg *CtrlClose) Header() *chunk7.ChunkHeader {
	return nil
}

func (msg *CtrlClose) SetHeader(header chunk7.ChunkHeader) {}
