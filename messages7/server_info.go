package messages7

import (
	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
	"github.com/teeworlds-go/go-teeworlds-protocol/varint"
)

type ServerInfo struct {
	ChunkHeader chunk7.ChunkHeader

	Version     string
	Name        string
	Hostname    string
	MapName     string
	GameType    string
	Flags       int
	SkillLevel  int
	PlayerCount int
	PlayerSlots int
	ClientCount int
	MaxClients  int
}

func (msg *ServerInfo) MsgId() int {
	return network7.MsgSysServerInfo
}

func (msg *ServerInfo) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *ServerInfo) System() bool {
	return true
}

func (msg *ServerInfo) Vital() bool {
	return true
}

func (msg *ServerInfo) Pack() []byte {
	p := packer.NewPacker(make([]byte,
		0,
		len(msg.Version)+1+
			len(msg.Name)+1+
			len(msg.Hostname)+1+
			len(msg.MapName)+1+
			len(msg.GameType)+1+
			6*varint.MaxVarintLen32,
	))
	p.AddString(msg.Version)
	p.AddString(msg.Name)
	p.AddString(msg.Hostname)
	p.AddString(msg.MapName)
	p.AddString(msg.GameType)
	p.AddInt(msg.Flags)
	p.AddInt(msg.SkillLevel)
	p.AddInt(msg.PlayerCount)
	p.AddInt(msg.PlayerSlots)
	p.AddInt(msg.ClientCount)
	p.AddInt(msg.MaxClients)

	return p.Bytes()
}

func (msg *ServerInfo) Unpack(u *packer.Unpacker) (err error) {
	msg.Version, err = u.NextString()
	if err != nil {
		return err
	}
	msg.Name, err = u.NextString()
	if err != nil {
		return err
	}
	msg.Hostname, err = u.NextString()
	if err != nil {
		return err
	}
	msg.MapName, err = u.NextString()
	if err != nil {
		return err
	}
	msg.GameType, err = u.NextString()
	if err != nil {
		return err
	}
	msg.Flags, err = u.NextInt()
	if err != nil {
		return err
	}
	msg.SkillLevel, err = u.NextInt()
	if err != nil {
		return err
	}
	msg.PlayerCount, err = u.NextInt()
	if err != nil {
		return err
	}
	msg.PlayerSlots, err = u.NextInt()
	if err != nil {
		return err
	}
	msg.ClientCount, err = u.NextInt()
	if err != nil {
		return err
	}
	msg.MaxClients, err = u.NextInt()
	if err != nil {
		return err
	}

	return nil
}

func (msg *ServerInfo) Header() *chunk7.ChunkHeader {
	return &msg.ChunkHeader
}

func (msg *ServerInfo) SetHeader(header chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}
