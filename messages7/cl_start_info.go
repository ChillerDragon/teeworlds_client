package messages7

import (
	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
	"github.com/teeworlds-go/go-teeworlds-protocol/varint"
)

type ClStartInfo struct {
	ChunkHeader chunk7.ChunkHeader

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
	ColorDecoration       int
	ColorHands            int
	ColorFeet             int
	ColorEyes             int
}

func (info ClStartInfo) MsgId() int {
	return network7.MsgGameClStartInfo
}

func (info ClStartInfo) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (info ClStartInfo) System() bool {
	return false
}

func (info ClStartInfo) Vital() bool {
	return true
}

func (info ClStartInfo) Pack() []byte {
	p := packer.NewPacker(make([]byte,
		0,
		7*varint.MaxVarintLen32+ // int
			6+ // bool
			len(info.Name)+1+
			len(info.Clan)+1+
			len(info.Body)+1+
			len(info.Marking)+1+
			len(info.Decoration)+1+
			len(info.Hands)+1+
			len(info.Feet)+1+
			len(info.Eyes)+1,
	))

	p.AddString(info.Name)
	p.AddString(info.Clan)
	p.AddInt(info.Country)
	p.AddString(info.Body)
	p.AddString(info.Marking)
	p.AddString(info.Decoration)
	p.AddString(info.Hands)
	p.AddString(info.Feet)
	p.AddString(info.Eyes)
	p.AddBool(info.CustomColorBody)
	p.AddBool(info.CustomColorMarking)
	p.AddBool(info.CustomColorDecoration)
	p.AddBool(info.CustomColorHands)
	p.AddBool(info.CustomColorFeet)
	p.AddBool(info.CustomColorEyes)
	p.AddInt(info.ColorBody)
	p.AddInt(info.ColorMarking)
	p.AddInt(info.ColorDecoration)
	p.AddInt(info.ColorHands)
	p.AddInt(info.ColorFeet)
	p.AddInt(info.ColorEyes)
	return p.Bytes()
}

func (info *ClStartInfo) Unpack(u *packer.Unpacker) (err error) {
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
	info.ColorDecoration, err = u.NextInt()
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
	return err
}

func (msg *ClStartInfo) Header() *chunk7.ChunkHeader {
	return &msg.ChunkHeader
}

func (msg *ClStartInfo) SetHeader(header chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}
