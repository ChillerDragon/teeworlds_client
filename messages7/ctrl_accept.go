package messages7

import (
	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type CtrlAccept struct {
	Token [4]byte
}

func (msg *CtrlAccept) MsgId() int {
	return network7.MsgCtrlAccept
}

func (msg *CtrlAccept) MsgType() network7.MsgType {
	return network7.TypeControl
}

func (msg *CtrlAccept) System() bool {
	return false
}

func (msg *CtrlAccept) Vital() bool {
	return false
}

func (msg *CtrlAccept) Pack() []byte {
	p := packer.NewPacker(make([]byte, 0, 1+len(msg.Token)+len(TokenPadding)))
	p.AddByte(network7.MsgCtrlAccept)
	p.AddBytes(msg.Token[:])
	p.AddBytes(TokenPadding)
	return p.Bytes()
}

func (msg *CtrlAccept) Unpack(u *packer.Unpacker) error {
	return nil
}

func (msg *CtrlAccept) Header() *chunk7.ChunkHeader {
	return nil
}

func (msg *CtrlAccept) SetHeader(header chunk7.ChunkHeader) {}
