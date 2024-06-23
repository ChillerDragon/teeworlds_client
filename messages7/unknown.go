package messages7

import (
	"slices"

	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type Unknown struct {
	ChunkHeader chunk7.ChunkHeader

	// contains entire raw message
	// including message id and chunk header
	// can either be a control message or a game/system message
	Data []byte
	Type network7.MsgType

	// TODO: should this be publicly accessible?
	msgId int
}

func (msg *Unknown) MsgId() int {
	msgId := msg.msgId

	if msg.Type == network7.TypeControl {
		return msg.msgId
	}
	msgId >>= 1
	return msgId
}

func (msg *Unknown) MsgType() network7.MsgType {
	return msg.Type
}

func (msg *Unknown) System() bool {
	if msg.Type == network7.TypeControl {
		return false
	}
	sys := msg.msgId&1 != 0
	return sys
}

func (msg *Unknown) Vital() bool {
	panic("You are not mean't to pack unknown messages. Use msg.Header().Vital instead.")
}

func (msg *Unknown) Pack() []byte {
	return slices.Clone(msg.Data)
}

func (msg *Unknown) Unpack(u *packer.Unpacker) error {
	msg.Data = u.Bytes()

	msgId, err := packer.UnpackInt(msg.Data)
	if err != nil {
		return err
	}
	msg.msgId = msgId
	return nil
}

func (msg *Unknown) Header() *chunk7.ChunkHeader {
	if msg.Type == network7.TypeControl {
		return nil
	}
	return &msg.ChunkHeader
}

func (msg *Unknown) SetHeader(header chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}
