package messages7

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

func TestPackStartInfo(t *testing.T) {
	want := []byte{
		0x67, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x00,
		0x00, 0x40, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x73, 0x77, 0x61, 0x72,
		0x64, 0x00, 0x64, 0x75, 0x6f, 0x64, 0x6f, 0x6e, 0x6e, 0x79, 0x00,
		0x00, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x00, 0x73,
		0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x00, 0x73, 0x74, 0x61,
		0x6e, 0x64, 0x61, 0x72, 0x64, 0x00, 0x01, 0x01, 0x00, 0x00, 0x00,
		0x00, 0x80, 0xfc, 0xaf, 0x05, 0xeb, 0x83, 0xd0, 0x0a, 0x80, 0xfe,
		0x07, 0x80, 0xfe, 0x07, 0x80, 0xfe, 0x07, 0x80, 0xfe, 0x07,
	}

	info := ClStartInfo{
		Name:                  "gopher",
		Clan:                  "",
		Country:               -1,
		Body:                  "greensward",
		Marking:               "duodonny",
		Decoration:            "",
		Hands:                 "standard",
		Feet:                  "standard",
		Eyes:                  "standard",
		CustomColorBody:       true,
		CustomColorMarking:    true,
		CustomColorDecoration: false,
		CustomColorHands:      false,
		CustomColorFeet:       false,
		CustomColorEyes:       false,
		ColorBody:             5635840,
		ColorMarking:          -11141356,
		ColorDecoration:       65408,
		ColorHands:            65408,
		ColorFeet:             65408,
		ColorEyes:             65408,
	}

	got := info.Pack()
	require.Equal(t, want, got)
}

func TestUnpackStartInfo(t *testing.T) {
	u := packer.Unpacker{}
	u.Reset([]byte{
		0x67, 0x6f, 0x70, 0x68, 0x65, 0x72, 0x00,
		0x00, 0x40, 0x67, 0x72, 0x65, 0x65, 0x6e, 0x73, 0x77, 0x61, 0x72,
		0x64, 0x00, 0x64, 0x75, 0x6f, 0x64, 0x6f, 0x6e, 0x6e, 0x79, 0x00,
		0x00, 0x73, 0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x00, 0x73,
		0x74, 0x61, 0x6e, 0x64, 0x61, 0x72, 0x64, 0x00, 0x73, 0x74, 0x61,
		0x6e, 0x64, 0x61, 0x72, 0x64, 0x00, 0x01, 0x01, 0x00, 0x00, 0x00,
		0x00, 0x80, 0xfc, 0xaf, 0x05, 0xeb, 0x83, 0xd0, 0x0a, 0x80, 0xfe,
		0x07, 0x80, 0xfe, 0x07, 0x80, 0xfe, 0x07, 0x80, 0xfe, 0x07,
	})

	info := ClStartInfo{}
	info.Unpack(&u)

	{
		got := info.Eyes
		want := "standard"
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}

		got = info.Decoration
		want = ""
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}

		got = info.Marking
		want = "duodonny"
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

	{
		got := info.ColorDecoration
		want := 65408
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

	wantedInfo := ClStartInfo{
		Name:                  "gopher",
		Clan:                  "",
		Country:               -1,
		Body:                  "greensward",
		Marking:               "duodonny",
		Decoration:            "",
		Hands:                 "standard",
		Feet:                  "standard",
		Eyes:                  "standard",
		CustomColorBody:       true,
		CustomColorMarking:    true,
		CustomColorDecoration: false,
		CustomColorHands:      false,
		CustomColorFeet:       false,
		CustomColorEyes:       false,
		ColorBody:             5635840,
		ColorMarking:          -11141356,
		ColorDecoration:       65408,
		ColorHands:            65408,
		ColorFeet:             65408,
		ColorEyes:             65408,
	}

	if !reflect.DeepEqual(info, wantedInfo) {
		t.Errorf("got %v, wanted %v", info, wantedInfo)
	}
}
