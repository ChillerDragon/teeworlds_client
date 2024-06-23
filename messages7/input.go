package messages7

import (
	"fmt"

	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
	"github.com/teeworlds-go/go-teeworlds-protocol/varint"
)

type Input struct {
	ChunkHeader chunk7.ChunkHeader

	AckGameTick    int
	PredictionTick int
	Size           int

	Direction    int
	TargetX      int
	TargetY      int
	Jump         int
	Fire         int
	Hook         int
	PlayerFlags  int
	WantedWeapon network7.Weapon
	NextWeapon   network7.Weapon
	PrevWeapon   network7.Weapon
}

func (msg *Input) MsgId() int {
	return network7.MsgSysInput
}

func (msg *Input) MsgType() network7.MsgType {
	return network7.TypeNet
}

func (msg *Input) System() bool {
	return true
}

func (msg *Input) Vital() bool {
	return false
}

func (msg *Input) Pack() []byte {
	p := packer.NewPacker(make([]byte, 0, 10*varint.MaxVarintLen32))

	p.AddInt(msg.Direction)
	p.AddInt(msg.TargetX)
	p.AddInt(msg.TargetY)
	p.AddInt(msg.Jump)
	p.AddInt(msg.Fire)
	p.AddInt(msg.Hook)
	p.AddInt(msg.PlayerFlags)
	p.AddInt(int(msg.WantedWeapon))
	p.AddInt(int(msg.NextWeapon))
	p.AddInt(int(msg.PrevWeapon))
	return p.Bytes()
}

func (msg *Input) Unpack(u *packer.Unpacker) (err error) {
	msg.Direction, err = u.NextInt()
	if err != nil {
		return err
	}
	msg.TargetX, err = u.NextInt()
	if err != nil {
		return err
	}
	msg.TargetY, err = u.NextInt()
	if err != nil {
		return err
	}
	msg.Jump, err = u.NextInt()
	if err != nil {
		return err
	}
	msg.Fire, err = u.NextInt()
	if err != nil {
		return err
	}
	msg.Hook, err = u.NextInt()
	if err != nil {
		return err
	}
	msg.PlayerFlags, err = u.NextInt()
	if err != nil {
		return err
	}

	wantWeapon, err := u.NextInt()
	if err != nil {
		return err
	}
	msg.WantedWeapon, err = newWeapon(wantWeapon)
	if err != nil {
		return err
	}

	nextWeapon, err := u.NextInt()
	if err != nil {
		return err
	}
	msg.NextWeapon, err = newWeapon(nextWeapon)
	if err != nil {
		return err
	}

	prevWeapon, err := u.NextInt()
	if err != nil {
		return err
	}

	msg.PrevWeapon, err = newWeapon(prevWeapon)
	if err != nil {
		return err
	}
	return nil
}

func (msg *Input) Header() *chunk7.ChunkHeader {
	return &msg.ChunkHeader
}

func (msg *Input) SetHeader(header chunk7.ChunkHeader) {
	msg.ChunkHeader = header
}

// validate weapon value
func newWeapon(w int) (network7.Weapon, error) {
	if 0 <= w || w < int(network7.NumWeapons) {
		return network7.Weapon(w), nil
	}
	return 0, fmt.Errorf("invalid weapon: %d", w)
}
