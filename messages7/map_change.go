package messages7

import (
	"fmt"

	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
	"github.com/teeworlds-go/go-teeworlds-protocol/varint"
)

type MapChange struct {
	ChunkHeader chunk7.ChunkHeader

	Name                        string
	Crc                         int
	Size                        int
	NumResponseChunksPerRequest int
	ChunkSize                   int
	Sha256                      [32]byte
}

func (msg *MapChange) MsgId() int {
	return network7.MsgSysMapChange
}

func (msg *MapChange) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *MapChange) System() bool {
	return true
}

func (msg *MapChange) Vital() bool {
	return true
}

func (msg *MapChange) Pack() []byte {
	p := packer.NewPacker(make([]byte,
		0,
		len(msg.Name)+1+
			4*varint.MaxVarintLen32+
			len(msg.Sha256),
	))
	p.AddString(msg.Name)
	p.AddInt(msg.Crc)
	p.AddInt(msg.Size)
	p.AddInt(msg.NumResponseChunksPerRequest)
	p.AddInt(msg.ChunkSize)
	p.AddBytes(msg.Sha256[:])
	return p.Bytes()
}

func (msg *MapChange) Unpack(u *packer.Unpacker) (err error) {
	msg.Name, err = u.NextString()
	if err != nil {
		return err
	}
	msg.Crc, err = u.NextInt()
	if err != nil {
		return err
	}
	msg.Size, err = u.NextInt()
	if err != nil {
		return err
	}
	msg.NumResponseChunksPerRequest, err = u.NextInt()
	if err != nil {
		return err
	}
	msg.ChunkSize, err = u.NextInt()
	if err != nil {
		return err
	}

	sha256Slice := u.Bytes()
	if len(sha256Slice) != 32 {
		return fmt.Errorf("sha256 checksum size is not 32: %d", len(sha256Slice))
	}
	copy(msg.Sha256[:], sha256Slice)
	return nil
}

func (msg *MapChange) Header() *chunk7.ChunkHeader {
	return &msg.ChunkHeader
}

func (msg *MapChange) SetHeader(header chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}
