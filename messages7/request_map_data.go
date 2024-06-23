package messages7

import (
	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type RequestMapData struct {
	ChunkHeader chunk7.ChunkHeader
}

func (msg *RequestMapData) MsgId() int {
	return network7.MsgSysRequestMapData
}

func (msg *RequestMapData) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *RequestMapData) System() bool {
	return true
}

func (msg *RequestMapData) Vital() bool {
	return true
}

func (msg *RequestMapData) Pack() []byte {
	return []byte{}
}

func (msg *RequestMapData) Unpack(u *packer.Unpacker) error {
	return nil
}

func (msg *RequestMapData) Header() *chunk7.ChunkHeader {
	return &msg.ChunkHeader
}

func (msg *RequestMapData) SetHeader(header chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}
