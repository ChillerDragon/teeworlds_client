package messages7

import (
	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
	"github.com/teeworlds-go/go-teeworlds-protocol/varint"
)

type SvChat struct {
	ChunkHeader chunk7.ChunkHeader

	Mode     network7.ChatMode
	ClientId int
	TargetId int
	Message  string
}

func (msg *SvChat) MsgId() int {
	return network7.MsgGameSvChat
}

func (msg *SvChat) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *SvChat) System() bool {
	return false
}

func (msg *SvChat) Vital() bool {
	return true
}

func (msg *SvChat) Pack() []byte {
	p := packer.NewPacker(
		make([]byte,
			0,
			3*varint.MaxVarintLen32+
				len(msg.Message)+1,
		))
	p.AddInt(int(msg.Mode))
	p.AddInt(msg.ClientId)
	p.AddInt(msg.TargetId)
	return p.Bytes()
}

func (msg *SvChat) Unpack(u *packer.Unpacker) (err error) {

	chatMode, err := u.NextInt()
	if err != nil {
		return err
	}
	msg.Mode = network7.ChatMode(chatMode)

	msg.ClientId, err = u.NextInt()
	if err != nil {
		return err
	}
	msg.TargetId, err = u.NextInt()
	if err != nil {
		return err
	}
	msg.Message, err = u.NextString()
	if err != nil {
		return err
	}
	return nil
}

func (msg *SvChat) Header() *chunk7.ChunkHeader {
	return &msg.ChunkHeader
}

func (msg *SvChat) SetHeader(header chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}
