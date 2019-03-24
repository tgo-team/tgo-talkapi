package mqtt

import (
	"bytes"
	"fmt"
	"github.com/tgo-team/tgo-core/tgo/packets"
	"io"
)

func (m *MQTTCodec) decodeCMD(fh *packets.FixedHeader,reader io.Reader) (*packets.CMDPacket, error) {
	c :=packets.NewCMDPacketWithHeader(*fh)
	c.CMD = packets.DecodeUint16(reader)
	c.TokenFlag = 1 & (packets.DecodeByte(reader)>>7) > 0
	var payloadLength = c.RemainingLength - 2 -1  // payloadLength = 剩余长度 - CMD长度 - TokenFlag长度
	if c.TokenFlag {
		c.Token = packets.DecodeString(reader)
		payloadLength = payloadLength - 2 - len(c.Token) //  （字符串长度标识(2 byte) +  Token字符串长度）
	}
	if payloadLength < 0 {
		return nil,fmt.Errorf("error upacking cmd, payload length < 0 for %d",payloadLength)
	}
	c.Payload = make([]byte, payloadLength)
	_, err := reader.Read(c.Payload)
	return c,err
}

func (m *MQTTCodec) encodeCMD(packet packets.Packet) ([]byte, error) {
	c := packet.(*packets.CMDPacket)
	var body bytes.Buffer
	body.Write(packets.EncodeUint16(c.CMD))
	body.WriteByte(packets.BoolToByte(c.TokenFlag)<<7)
	if c.TokenFlag {
		body.Write(packets.EncodeString(c.Token))
	}
	body.Write(c.Payload)
	c.RemainingLength = body.Len()
	return body.Bytes(),nil
}
