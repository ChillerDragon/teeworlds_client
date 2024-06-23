package messages7

import (
	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
	"github.com/teeworlds-go/go-teeworlds-protocol/varint"
)

type SvTeam struct {
	ChunkHeader chunk7.ChunkHeader

	ClientId     int
	Silent       bool
	CooldownTick int
}

func (msg *SvTeam) MsgId() int {
	return network7.MsgGameSvTeam
}

func (msg *SvTeam) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *SvTeam) System() bool {
	return false
}

func (msg *SvTeam) Vital() bool {
	return true
}

func (msg *SvTeam) Pack() []byte {
	p := packer.NewPacker(make([]byte, 0, 2*varint.MaxVarintLen32+1))
	p.AddInt(msg.ClientId)
	p.AddBool(msg.Silent)
	p.AddInt(msg.CooldownTick)
	return p.Bytes()
}

func (msg *SvTeam) Unpack(u *packer.Unpacker) (err error) {
	msg.ClientId, err = u.NextInt()
	if err != nil {
		return err
	}
	msg.Silent, err = u.NextBool()
	if err != nil {
		return err
	}
	msg.CooldownTick, err = u.NextInt()
	if err != nil {
		return err
	}
	return nil
}

func (msg *SvTeam) Header() *chunk7.ChunkHeader {
	return &msg.ChunkHeader
}

func (msg *SvTeam) SetHeader(header chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}
