package messages7

import (
	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type RconCmdAdd struct {
	ChunkHeader chunk7.ChunkHeader

	Name   string
	Help   string
	Params string
}

func (msg *RconCmdAdd) MsgId() int {
	return network7.MsgSysRconCmdAdd
}

func (msg *RconCmdAdd) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *RconCmdAdd) System() bool {
	return true
}

func (msg *RconCmdAdd) Vital() bool {
	return true
}

func (msg *RconCmdAdd) Pack() []byte {
	p := packer.NewPacker(make([]byte,
		0,
		len(msg.Name)+
			len(msg.Help)+
			len(msg.Params),
	))

	p.AddString(msg.Name)
	p.AddString(msg.Help)
	p.AddString(msg.Params)
	return p.Bytes()
}

func (msg *RconCmdAdd) Unpack(u *packer.Unpacker) (err error) {
	msg.Name, err = u.NextString()
	if err != nil {
		return err
	}
	msg.Help, err = u.NextString()
	if err != nil {
		return err
	}
	msg.Params, err = u.NextString()
	if err != nil {
		return err
	}
	return nil
}

func (msg *RconCmdAdd) Header() *chunk7.ChunkHeader {
	return &msg.ChunkHeader
}

func (msg *RconCmdAdd) SetHeader(header chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}
