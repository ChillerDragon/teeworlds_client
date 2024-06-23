package messages7

import (
	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type MaplistEntryAdd struct {
	ChunkHeader chunk7.ChunkHeader

	MapName string
}

func (msg *MaplistEntryAdd) MsgId() int {
	return network7.MsgSysMaplistEntryAdd
}

func (msg *MaplistEntryAdd) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *MaplistEntryAdd) System() bool {
	return true
}

func (msg *MaplistEntryAdd) Vital() bool {
	return true
}

func (msg *MaplistEntryAdd) Pack() []byte {
	return packer.PackString(msg.MapName)
}

func (msg *MaplistEntryAdd) Unpack(u *packer.Unpacker) (err error) {
	msg.MapName, err = u.NextString()
	return err
}

func (msg *MaplistEntryAdd) Header() *chunk7.ChunkHeader {
	return &msg.ChunkHeader
}

func (msg *MaplistEntryAdd) SetHeader(header chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}
