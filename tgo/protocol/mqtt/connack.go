package mqtt

import (
	"bytes"
	"github.com/tgo-team/tgo-core/tgo/packets"
	"io"
)

func (m *MQTTCodec) decodeConnack(fh *packets.FixedHeader,reader io.Reader) (*packets.ConnackPacket, error) {
	c :=packets.NewConnackPacketWithHeader(*fh)
	packets.DecodeByte(reader)
	c.ReturnCode = packets.ConnReturnCode(packets.DecodeByte(reader))
	c.FixedHeader = *fh
	return c,nil
}

func (m *MQTTCodec) encodeConnack(packet packets.Packet) ([]byte, error) {
	c := packet.(*packets.ConnackPacket)
	var body bytes.Buffer
	body.WriteByte(0x00)
	body.WriteByte(byte(c.ReturnCode))
	c.RemainingLength = body.Len()
	return body.Bytes(),nil
}