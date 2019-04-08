package mqtt

import (
	"bytes"
	"fmt"
	"github.com/tgo-team/tgo-core/tgo/packets"
	"io"
)

/**
  Header | Remaining Length  |  status  |   payload |
  1 byte |       xxx         |   2 byte |    xxx    |

 */

func (m *MQTTCodec) decodeCmdack(fh *packets.FixedHeader,reader io.Reader) (*packets.CmdackPacket, error) {
	c :=packets.NewCmdackPacketWithHeader(*fh)
	c.CMD = packets.DecodeString(reader)
	c.Status = packets.DecodeUint16(reader)
	var payloadLength = c.RemainingLength - (len(c.CMD) + 2) - 2 // payloadLength = 剩余长度 - CMD长度 - 状态长度
	if payloadLength < 0 {
		return nil,fmt.Errorf("Error upacking cmd, payload length < 0")
	}
	c.Payload = make([]byte, payloadLength)
	_, err := reader.Read(c.Payload)
	return c,err
}

func (m *MQTTCodec) encodeCmdack(packet packets.Packet) ([]byte, error) {
	c := packet.(*packets.CmdackPacket)
	var body bytes.Buffer
	body.Write(packets.EncodeString(c.CMD))
	body.Write(packets.EncodeUint16(c.Status))
	body.Write(c.Payload)
	c.RemainingLength = body.Len()
	return body.Bytes(),nil
}