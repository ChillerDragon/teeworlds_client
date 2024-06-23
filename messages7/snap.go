package messages7

import (
	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
	"github.com/teeworlds-go/go-teeworlds-protocol/varint"
)

type Snap struct {
	ChunkHeader chunk7.ChunkHeader

	GameTick  int
	DeltaTick int
	NumParts  int
	Part      int
	Crc       int
	PartSize  int
	Data      []byte
}

func (msg *Snap) MsgId() int {
	return network7.MsgSysSnap
}

func (msg *Snap) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *Snap) System() bool {
	return true
}

func (msg *Snap) Vital() bool {
	return false
}

func (msg *Snap) Pack() []byte {
	p := packer.NewPacker(make([]byte,
		0,
		6*varint.MaxVarintLen32+
			len(msg.Data),
	))
	p.AddInt(msg.GameTick)
	p.AddInt(msg.DeltaTick)
	p.AddInt(msg.NumParts)
	p.AddInt(msg.Part)
	p.AddInt(msg.Crc)
	p.AddInt(msg.PartSize)
	p.AddBytes(msg.Data)

	return p.Bytes()
}

func (msg *Snap) Unpack(u *packer.Unpacker) (err error) {
	msg.GameTick, err = u.NextInt()
	if err != nil {
		return err
	}
	msg.DeltaTick, err = u.NextInt()
	if err != nil {
		return err
	}
	msg.NumParts, err = u.NextInt()
	if err != nil {
		return err
	}
	msg.Part, err = u.NextInt()
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

func (msg *Snap) Header() *chunk7.ChunkHeader {
	return &msg.ChunkHeader
}

func (msg *Snap) SetHeader(header chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}
