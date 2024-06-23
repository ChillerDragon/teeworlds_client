package messages7

import (
	"fmt"

	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type CtrlConnect struct {
	Token [4]byte
}

func (msg *CtrlConnect) MsgId() int {
	return network7.MsgCtrlConnect
}

func (msg *CtrlConnect) MsgType() network7.MsgType {
	return network7.TypeControl
}

func (msg *CtrlConnect) System() bool {
	return false
}

func (msg *CtrlConnect) Vital() bool {
	return false
}

func (msg *CtrlConnect) Pack() []byte {
	result := make([]byte, 0, 1+len(msg.Token)+len(TokenPadding))
	result = append(result, network7.MsgCtrlConnect)
	result = append(result, msg.Token[:]...)
	result = append(result, TokenPadding...)
	return result
}

func (msg *CtrlConnect) Unpack(u *packer.Unpacker) error {
	data := u.Bytes()
	if len(data) != 4 {
		return fmt.Errorf("invalid number of token bytes: %d, expected 4", len(data))
	}

	copy(msg.Token[:], data)
	return nil
}

func (msg *CtrlConnect) Header() *chunk7.ChunkHeader {
	return nil
}

func (msg *CtrlConnect) SetHeader(header chunk7.ChunkHeader) {
}
