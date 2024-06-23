package messages7

// this message is shared between client and server
// but this implementation is assuming we are sending from a client

import (
	"fmt"

	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
)

type CtrlToken struct {
	Token [4]byte
}

func (msg *CtrlToken) MsgId() int {
	return network7.MsgCtrlToken
}

func (msg *CtrlToken) MsgType() network7.MsgType {
	return network7.TypeControl
}

func (msg *CtrlToken) System() bool {
	return false
}

func (msg *CtrlToken) Vital() bool {
	return false
}

func (msg *CtrlToken) Pack() []byte {
	result := make([]byte, 0, 1+len(msg.Token)+len(TokenPadding))
	result = append(result, network7.MsgCtrlToken)
	result = append(result, msg.Token[:]...)
	result = append(result, TokenPadding...)
	return result
}

func (msg *CtrlToken) Unpack(u *packer.Unpacker) (err error) {
	data := u.Bytes()
	if len(data) != 4 {
		return fmt.Errorf("invalid token size: %d, expected 4", len(data))
	}

	copy(msg.Token[:], data)
	return nil
}

func (msg *CtrlToken) Header() *chunk7.ChunkHeader {
	return nil
}

func (msg *CtrlToken) SetHeader(header chunk7.ChunkHeader) {}
