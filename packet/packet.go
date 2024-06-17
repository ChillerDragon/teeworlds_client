package packet

import "slices"

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

func (header *PacketHeader) Unpack(packet []byte) {
	header.Flags.Unpack(packet)
	header.Ack = (int(packet[0]&0x3) << 8) | int(packet[1])
	header.NumChunks = int(packet[2])
	copy(header.Token[:], packet[3:7])
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
