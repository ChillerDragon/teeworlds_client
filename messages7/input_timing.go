package messages7

import (
	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
	"github.com/teeworlds-go/go-teeworlds-protocol/varint"
)

type InputTiming struct {
	ChunkHeader chunk7.ChunkHeader

	IntendedPredTick int
	TimeLeft         int
}

func (msg *InputTiming) MsgId() int {
	return network7.MsgSysInputTiming
}

func (msg *InputTiming) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *InputTiming) System() bool {
	return true
}

func (msg *InputTiming) Vital() bool {
	return false
}

func (msg *InputTiming) Pack() []byte {
	p := packer.NewPacker(make([]byte, 0, 2*varint.MaxVarintLen32))
	p.AddInt(msg.IntendedPredTick)
	p.AddInt(msg.TimeLeft)
	return p.Bytes()
}

func (msg *InputTiming) Unpack(u *packer.Unpacker) (err error) {
	msg.IntendedPredTick, err = u.NextInt()
	if err != nil {
		return err
	}
	msg.TimeLeft, err = u.NextInt()
	if err != nil {
		return err
	}
	return nil
}

func (msg *InputTiming) Header() *chunk7.ChunkHeader {
	return &msg.ChunkHeader
}

func (msg *InputTiming) SetHeader(header chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}
