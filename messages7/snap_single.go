package messages7

import (
	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
	"github.com/teeworlds-go/go-teeworlds-protocol/varint"
)

type SnapSingle struct {
	ChunkHeader chunk7.ChunkHeader

	GameTick  int
	DeltaTick int
	Crc       int
	PartSize  int
	Data      []byte
}

func (msg *SnapSingle) MsgId() int {
	return network7.MsgSysSnapSingle
}

func (msg *SnapSingle) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *SnapSingle) System() bool {
	return true
}

func (msg *SnapSingle) Vital() bool {
	return false
}

func (msg *SnapSingle) Pack() []byte {
	p := packer.NewPacker(make([]byte,
		0,
		4*varint.MaxVarintLen32+
			len(msg.Data)),
	)
	p.AddInt(msg.GameTick)
	p.AddInt(msg.DeltaTick)
	p.AddInt(msg.Crc)
	p.AddInt(msg.PartSize)
	p.AddBytes(msg.Data)
	return p.Bytes()
}

func (msg *SnapSingle) Unpack(u *packer.Unpacker) (err error) {
	msg.GameTick, err = u.NextInt()
	if err != nil {
		return err
	}
	msg.DeltaTick, err = u.NextInt()
	if err != nil {
		return err
	}
	msg.Crc, err = u.NextInt()
	if err != nil {
		return err
	}
	msg.PartSize, err = u.NextInt()
	if err != nil {
		return err
	}
	msg.Data = u.Bytes()

	return nil
}

func (msg *SnapSingle) Header() *chunk7.ChunkHeader {
	return &msg.ChunkHeader
}

func (msg *SnapSingle) SetHeader(header chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}
