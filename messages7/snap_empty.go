package messages7

import (
	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
	"github.com/teeworlds-go/go-teeworlds-protocol/varint"
)

type SnapEmpty struct {
	ChunkHeader chunk7.ChunkHeader

	GameTick  int
	DeltaTick int
}

func (msg *SnapEmpty) MsgId() int {
	return network7.MsgSysSnapEmpty
}

func (msg *SnapEmpty) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *SnapEmpty) System() bool {
	return true
}

func (msg *SnapEmpty) Vital() bool {
	return false
}

func (msg *SnapEmpty) Pack() []byte {
	p := packer.NewPacker(make([]byte, 0, 2*varint.MaxVarintLen32))
	p.AddInt(msg.GameTick)
	p.AddInt(msg.DeltaTick)
	return p.Bytes()
}

func (msg *SnapEmpty) Unpack(u *packer.Unpacker) (err error) {
	msg.GameTick, err = u.NextInt()
	if err != nil {
		return err
	}
	msg.DeltaTick, err = u.NextInt()
	if err != nil {
		return err
	}
	return nil
}

func (msg *SnapEmpty) Header() *chunk7.ChunkHeader {
	return &msg.ChunkHeader
}

func (msg *SnapEmpty) SetHeader(header chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}
