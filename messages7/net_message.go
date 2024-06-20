package messages7

import (
	"github.com/teeworlds-go/teeworlds/network7"
	"github.com/teeworlds-go/teeworlds/packer"
)

type NetMessage interface {
	MsgId() int
	MsgType() network7.MsgType
	System() bool
	Vital() bool
	Pack() []byte
	Unpack(u *packer.Unpacker)
}