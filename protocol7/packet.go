package protocol7

import (
	"fmt"
	"slices"

	"github.com/teeworlds-go/go-teeworlds-protocol/chunk7"
	"github.com/teeworlds-go/go-teeworlds-protocol/messages7"
	"github.com/teeworlds-go/go-teeworlds-protocol/network7"
	"github.com/teeworlds-go/go-teeworlds-protocol/packer"
	"github.com/teeworlds-go/huffman"
)

const (
	packetFlagControl     = 1
	packetFlagResend      = 2
	packetFlagCompression = 4
	packetFlagConnless    = 8
)

type PacketFlags struct {
	Connless    bool
	Compression bool
	Resend      bool
	Control     bool
}

type PacketHeader struct {
	Flags     PacketFlags
	Ack       int
	NumChunks int
	Token     [4]byte

	// connless
	ResponseToken [4]byte
}

type Packet struct {
	Header   PacketHeader
	Messages []messages7.NetMessage
}

func PackChunk(msg messages7.NetMessage, connection *Session) []byte {
	if _, ok := msg.(*messages7.Unknown); ok {
		return msg.Pack()
	}
	if msg.MsgType() == network7.TypeControl {
		return msg.Pack()
	}

	msgId := msg.MsgId() << 1
	if msg.System() {
		msgId |= 1
	}

	if msg.Vital() {
		connection.Sequence++
	}

	msgAndSys := packer.PackInt(msgId)
	payload := msg.Pack()

	if msg.Header() == nil {
		header := chunk7.ChunkHeader{
			Flags: chunk7.ChunkFlags{
				Vital: msg.Vital(),
			},
			Seq: connection.Sequence,
		}
		msg.SetHeader(header)
	}

	msg.Header().Size = len(msgAndSys) + len(payload)

	data := slices.Concat(
		msg.Header().Pack(),
		msgAndSys,
		payload,
	)

	return data
}

func (packet *Packet) unpackSystem(msgId int, chunk chunk7.Chunk, u *packer.Unpacker) (bool, error) {
	var msg messages7.NetMessage
	switch msgId {
	case network7.MsgSysInfo:
		msg = &messages7.Info{}
	case network7.MsgSysMapChange:
		msg = &messages7.MapChange{}
	case network7.MsgSysMapData:
		msg = &messages7.MapData{}
	case network7.MsgSysServerInfo:
		msg = &messages7.ServerInfo{}
	case network7.MsgSysConReady:
		msg = &messages7.ConReady{}
	case network7.MsgSysSnap:
		msg = &messages7.Snap{}
	case network7.MsgSysSnapEmpty:
		msg = &messages7.SnapEmpty{}
	case network7.MsgSysSnapSingle:
		msg = &messages7.SnapSingle{}
	case network7.MsgSysSnapSmall:
		msg = &messages7.SnapSmall{}
	case network7.MsgSysInputTiming:
		msg = &messages7.InputTiming{}
	case network7.MsgSysRconAuthOn:
		msg = &messages7.RconAuthOn{}
	case network7.MsgSysRconAuthOff:
		msg = &messages7.RconAuthOff{}
	case network7.MsgSysRconLine:
		msg = &messages7.RconLine{}
	case network7.MsgSysRconCmdAdd:
		msg = &messages7.RconCmdAdd{}
	case network7.MsgSysRconCmdRem:
		msg = &messages7.RconCmdRem{}
	case network7.MsgSysAuthChallenge:
		msg = &messages7.AuthChallenge{}
	case network7.MsgSysAuthResult:
		msg = &messages7.AuthResult{}
	case network7.MsgSysReady:
		msg = &messages7.Ready{}
	case network7.MsgSysEnterGame:
		msg = &messages7.EnterGame{}
	case network7.MsgSysInput:
		msg = &messages7.Input{}
	case network7.MsgSysRconCmd:
		msg = &messages7.RconCmd{}
	case network7.MsgSysRconAuth:
		msg = &messages7.RconAuth{}
	case network7.MsgSysRequestMapData:
		msg = &messages7.RequestMapData{}
	case network7.MsgSysAuthStart:
		msg = &messages7.AuthStart{}
	case network7.MsgSysAuthResponse:
		msg = &messages7.AuthResponse{}
	case network7.MsgSysPing:
		msg = &messages7.Ping{}
	case network7.MsgSysPingReply:
		msg = &messages7.PingReply{}
	case network7.MsgSysError:
		msg = &messages7.Error{}
	case network7.MsgSysMaplistEntryAdd:
		msg = &messages7.MaplistEntryAdd{}
	case network7.MsgSysMaplistEntryRem:
		msg = &messages7.MaplistEntryRem{}
	default:
		return false, nil
	}

	msg.SetHeader(chunk.Header)
	err := msg.Unpack(u)
	if err != nil {
		return false, err
	}

	packet.Messages = append(packet.Messages, msg)
	return true, nil
}

func (packet *Packet) unpackGame(msgId int, chunk chunk7.Chunk, u *packer.Unpacker) (bool, error) {

	var msg messages7.NetMessage
	switch msgId {
	case network7.MsgGameReadyToEnter:
		msg = &messages7.ReadyToEnter{}
	case network7.MsgGameSvMotd:
		msg = &messages7.SvMotd{}
	case network7.MsgGameSvChat:
		msg = &messages7.SvChat{}
	case network7.MsgGameSvClientInfo:
		msg = &messages7.SvClientInfo{}
	default:
		return false, nil
	}
	msg.SetHeader(chunk.Header)

	err := msg.Unpack(u)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (packet *Packet) unpackChunk(chunk chunk7.Chunk) (bool, error) {
	u := packer.NewUnpacker(chunk.Data)

	msg, err := u.NextInt()
	if err != nil {
		return false, err
	}

	sys := msg&1 != 0
	msg >>= 1

	if sys {
		return packet.unpackSystem(msg, chunk, u)
	}
	return packet.unpackGame(msg, chunk, u)
}

func (packet *Packet) unpackPayload(payload []byte) {
	chunks := chunk7.UnpackChunks(payload)

	var (
		ok  bool
		err error
	)
	for _, c := range chunks {
		ok, err = packet.unpackChunk(c)
		if err != nil || !ok {
			unknown := &messages7.Unknown{
				ChunkHeader: c.Header,
				Data:        append(c.Header.Pack(), c.Data...),
				Type:        network7.TypeNet,
			}
			packet.Messages = append(packet.Messages, unknown)
		}
	}
}

// Only returns an error if there is invalid huffman compression applied
// Unknown messages will be unpacked as messages7.Unknown
// There is no validation no messages will be dropped
func (packet *Packet) Unpack(data []byte) error {
	header, payload := data[:7], data[7:]
	err := packet.Header.Unpack(header)
	if err != nil {
		return err
	}

	if packet.Header.Flags.Control {
		unpacker := packer.NewUnpacker(payload)

		ctrlMsg, err := unpacker.NextInt()
		if err != nil {
			return err
		}

		var msg messages7.NetMessage
		switch ctrlMsg {
		case network7.MsgCtrlToken:
			msg = &messages7.CtrlToken{}
		case network7.MsgCtrlAccept:
			msg = &messages7.CtrlAccept{}
		case network7.MsgCtrlClose:
			msg = &messages7.CtrlClose{}
		default:
			msg = &messages7.Unknown{
				Type: network7.TypeControl,
			}
		}

		err = msg.Unpack(unpacker)
		if err != nil {
			return err
		}
		return nil
	}

	if packet.Header.Flags.Compression {
		huff := huffman.Huffman{}
		payload, err = huff.Decompress(payload)
		if err != nil {
			return err
		}
	}

	packet.unpackPayload(payload)
	return nil
}

func (packet *Packet) Pack(connection *Session) []byte {
	payload := []byte{}
	control := false

	for _, msg := range packet.Messages {
		payload = append(payload, PackChunk(msg, connection)...)
		if msg.MsgType() == network7.TypeControl {
			control = true
		}
	}

	packet.Header.NumChunks = len(packet.Messages)

	if control {
		packet.Header.Flags.Connless = false
		packet.Header.Flags.Compression = false
		packet.Header.Flags.Resend = false
		packet.Header.Flags.Control = true
	}

	if packet.Header.Flags.Compression {
		// TODO: store huffman object in connection to avoid reallocating memory
		huff := huffman.Huffman{}
		var err error
		payload, err = huff.Compress(payload)
		if err != nil {
			panic(err)
		}
	}

	return slices.Concat(
		packet.Header.Pack(),
		payload,
	)
}

func (header *PacketHeader) Pack() []byte {
	flags := 0
	if header.Flags.Control {
		flags |= packetFlagControl
	}
	if header.Flags.Resend {
		flags |= packetFlagResend
	}
	if header.Flags.Compression {
		flags |= packetFlagCompression
	}
	if header.Flags.Connless {
		flags |= packetFlagConnless
	}

	if header.Flags.Connless {
		version := 1
		return slices.Concat(
			[]byte{byte(((packetFlagConnless << 2) & 0x0fc) | (version & 0x03))},
			header.Token[:],
			header.ResponseToken[:],
		)
	}

	return slices.Concat(
		[]byte{
			byte(((flags << 2) & 0xfc) | ((header.Ack >> 8) & 0x03)),
			byte(header.Ack & 0x0ff),
			byte(header.NumChunks),
		},
		header.Token[:],
	)
}

func (header *PacketHeader) Unpack(packet []byte) error {
	if len(packet) < 7 {
		return fmt.Errorf("failed to unpack packet header: size %d", len(packet))
	}

	header.Flags.Unpack(packet)
	header.Ack = (int(packet[0]&0x3) << 8) | int(packet[1])
	header.NumChunks = int(packet[2])
	copy(header.Token[:], packet[3:7])
	return nil
}

func (flags *PacketFlags) Unpack(packetHeaderRaw []byte) {
	flagBits := packetHeaderRaw[0] >> 2
	flags.Control = (flagBits & packetFlagControl) != 0
	flags.Resend = (flagBits & packetFlagResend) != 0
	flags.Compression = (flagBits & packetFlagCompression) != 0
	flags.Connless = (flagBits & packetFlagConnless) != 0
}

func (flags *PacketFlags) Pack() []byte {
	data := 0

	if flags.Control {
		data |= packetFlagControl
	}
	if flags.Resend {
		data |= packetFlagResend
	}
	if flags.Compression {
		data |= packetFlagCompression
	}
	if flags.Connless {
		data |= packetFlagConnless
	}

	return []byte{byte(data)}
}
