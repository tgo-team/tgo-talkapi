package ctrl

import (
	"github.com/tgo-team/tgo-core/tgo"
	"github.com/tgo-team/tgo-core/tgo/packets"
	"github.com/tgo-team/tgo-talkapi/cmd"
	"github.com/tgo-team/tgo-talkapi/test"
	_ "github.com/tgo-team/tgo-talkapi/tgo/protocol/mqtt"
	"github.com/tgo-team/tgo-talkapi/tgo/server/tcp"
	_ "github.com/tgo-team/tgo-talkapi/tgo/storage/memory"
	"net"
	"testing"
	"time"
)

func TestControllerServer(t *testing.T) {
	tg := startTGO(t)
	var tcpServer *tcp.Server
	for _, server := range tg.Servers {
		s, ok := server.(*tcp.Server)
		if ok {
			tcpServer = s
		}
	}
	resultChan := make(chan int,0)
	// 开启控制器
	controller := New(tg)
	controller.RegisterCMDHandler("login", func(ctx *cmd.Context) {
		resultChan<- 1
	})
	controller.Start()
	conn, err := MustConnectServer(tcpServer.RealTCPAddr())
	test.Nil(t, err)

	data := []byte{0x1}
	test.Nil(t, err)
	cp := packets.NewCmdPacket("login", data)
	cp.TokenFlag = true
	cp.Token = "2334"
	sendCmdPacket(t, conn, tg, cp)
	<- resultChan
}

func sendCmdPacket(t *testing.T, conn net.Conn, tg *tgo.TGO, CmdPacket *packets.CmdPacket) {
	WritePacket(t, conn, CmdPacket, tg)
}

func WritePacket(t *testing.T, conn net.Conn, packet packets.Packet, tg *tgo.TGO) {
	pingData, err := tg.GetOpts().Pro.EncodePacket(packet)
	test.Nil(t, err)
	_, err = conn.Write(pingData)
}

func startTGO(t *testing.T) *tgo.TGO {
	opts := tgo.NewOptions()
	opts.TCPAddress = "0.0.0.0:0"
	opts.Log = test.NewLog(t)
	tg := tgo.New(opts)
	err := tg.Start()
	test.Nil(t, err)
	return tg
}

func MustConnectServer(tcpAddr *net.TCPAddr) (net.Conn, error) {
	conn, err := net.DialTimeout("tcp", tcpAddr.String(), time.Second)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
