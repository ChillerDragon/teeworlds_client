package messages7

import (
	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type MaplistEntryRem struct {
	ChunkHeader chunk7.ChunkHeader

	MapName string
}

func (msg *MaplistEntryRem) MsgId() int {
	return network7.MsgSysMaplistEntryRem
}

func (msg *MaplistEntryRem) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *MaplistEntryRem) System() bool {
	return true
}

func (msg *MaplistEntryRem) Vital() bool {
	return true
}

func (msg *MaplistEntryRem) Pack() []byte {
	return packer.PackString(msg.MapName)
}

func (msg *MaplistEntryRem) Unpack(u *packer.Unpacker) (err error) {
	msg.MapName, err = u.NextString()
	return err
}

func (msg *MaplistEntryRem) Header() *chunk7.ChunkHeader {
	return &msg.ChunkHeader
}

func (msg *MaplistEntryRem) SetHeader(header chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}
