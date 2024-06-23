package messages7

import (
	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
	"github.com/teeworlds-go/go-teeworlds-protocol/varint"
)

type SvClientInfo struct {
	ChunkHeader chunk7.ChunkHeader

	ClientId              int
	Local                 bool
	Team                  int
	Name                  string
	Clan                  string
	Country               int
	Body                  string
	Marking               string
	Decoration            string
	Hands                 string
	Feet                  string
	Eyes                  string
	CustomColorBody       bool
	CustomColorMarking    bool
	CustomColorDecoration bool
	CustomColorHands      bool
	CustomColorFeet       bool
	CustomColorEyes       bool
	ColorBody             int
	ColorMarking          int
	ColorHands            int
	ColorFeet             int
	ColorEyes             int
	Silent                bool
}

func (info *SvClientInfo) MsgId() int {
	return network7.MsgGameSvClientInfo
}

func (info *SvClientInfo) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (info *SvClientInfo) System() bool {
	return false
}

func (info *SvClientInfo) Vital() bool {
	return true
}

func (msg *SvClientInfo) Pack() []byte {
	p := packer.NewPacker(
		make([]byte,
			0,
			8*varint.MaxVarintLen32+ // int
				8*1+ // bool
				len(msg.Name)+1+
				len(msg.Clan)+1+
				len(msg.Body)+1+
				len(msg.Marking)+1+
				len(msg.Decoration)+1+
				len(msg.Hands)+1+
				len(msg.Feet)+1+
				len(msg.Eyes)+1,
		))

	p.AddInt(msg.ClientId)
	p.AddBool(msg.Local)
	p.AddInt(msg.Team)
	p.AddString(msg.Name)
	p.AddString(msg.Clan)
	p.AddInt(msg.Country)
	p.AddString(msg.Body)
	p.AddString(msg.Marking)
	p.AddString(msg.Decoration)
	p.AddString(msg.Hands)
	p.AddString(msg.Feet)
	p.AddString(msg.Eyes)
	p.AddBool(msg.CustomColorBody)
	p.AddBool(msg.CustomColorMarking)
	p.AddBool(msg.CustomColorDecoration)
	p.AddBool(msg.CustomColorHands)
	p.AddBool(msg.CustomColorFeet)
	p.AddBool(msg.CustomColorEyes)
	p.AddInt(msg.ColorBody)
	p.AddInt(msg.ColorMarking)
	p.AddInt(msg.ColorHands)
	p.AddInt(msg.ColorFeet)
	p.AddInt(msg.ColorEyes)
	p.AddBool(msg.Silent)
	return p.Bytes()
}

func (info *SvClientInfo) Unpack(u *packer.Unpacker) (err error) {

	info.ClientId, err = u.NextInt()
	if err != nil {
		return err
	}
	info.Local, err = u.NextBool()
	if err != nil {
		return err
	}

	info.Team, err = u.NextInt()
	if err != nil {
		return err
	}
	info.Name, err = u.NextString()
	if err != nil {
		return err
	}
	info.Clan, err = u.NextString()
	if err != nil {
		return err
	}
	info.Country, err = u.NextInt()
	if err != nil {
		return err
	}
	info.Body, err = u.NextString()
	if err != nil {
		return err
	}
	info.Marking, err = u.NextString()
	if err != nil {
		return err
	}
	info.Decoration, err = u.NextString()
	if err != nil {
		return err
	}
	info.Hands, err = u.NextString()
	if err != nil {
		return err
	}
	info.Feet, err = u.NextString()
	if err != nil {
		return err
	}
	info.Eyes, err = u.NextString()
	if err != nil {
		return err
	}
	info.CustomColorBody, err = u.NextBool()
	if err != nil {
		return err
	}
	info.CustomColorMarking, err = u.NextBool()
	if err != nil {
		return err
	}
	info.CustomColorDecoration, err = u.NextBool()
	if err != nil {
		return err
	}
	info.CustomColorHands, err = u.NextBool()
	if err != nil {
		return err
	}
	info.CustomColorFeet, err = u.NextBool()
	if err != nil {
		return err
	}
	info.CustomColorEyes, err = u.NextBool()
	if err != nil {
		return err
	}
	info.ColorBody, err = u.NextInt()
	if err != nil {
		return err
	}
	info.ColorMarking, err = u.NextInt()
	if err != nil {
		return err
	}
	info.ColorHands, err = u.NextInt()
	if err != nil {
		return err
	}
	info.ColorFeet, err = u.NextInt()
	if err != nil {
		return err
	}
	info.ColorEyes, err = u.NextInt()
	if err != nil {
		return err
	}
	info.Silent, err = u.NextBool()
	if err != nil {
		return err
	}

	return nil
}

func (msg *SvClientInfo) Header() *chunk7.ChunkHeader {
	return &msg.ChunkHeader
}

func (msg *SvClientInfo) SetHeader(header chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}
