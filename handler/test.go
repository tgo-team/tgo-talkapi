package handler

import (
	"github.com/tgo-team/tgo-core/tgo"
	"github.com/tgo-team/tgo-core/tgo/packets"
	"github.com/tgo-team/tgo-talkapi/ctrl"
	"github.com/tgo-team/tgo-talkapi/test"
	_ "github.com/tgo-team/tgo-talkapi/tgo/protocol/mqtt"
	"github.com/tgo-team/tgo-talkapi/tgo/server/tcp"
	_ "github.com/tgo-team/tgo-talkapi/tgo/storage/memory"
	"net"
	"testing"
	"time"
)

func SendCmdPacket(t *testing.T, conn net.Conn, tg *tgo.TGO, CmdPacket *packets.CmdPacket) {
	WritePacket(t, conn, CmdPacket, tg)
}

func WritePacket(t *testing.T, conn net.Conn, packet packets.Packet, tg *tgo.TGO) {
	pingData, err := tg.GetOpts().Pro.EncodePacket(packet)
	test.Nil(t, err)
	_, err = conn.Write(pingData)
}

func StartTGO(t *testing.T) (*tgo.TGO, *ctrl.Controller,  net.Conn) {
	opts := tgo.NewOptions()
	opts.TCPAddress = "0.0.0.0:0"
	opts.Log = test.NewLog(t)
	tg := tgo.New(opts)
	// 开启控制器
	controller := ctrl.New(tg)
	controller.Start()
	err := tg.Start()
	test.Nil(t, err)

	var tcpServer *tcp.Server
	for _, server := range tg.Servers {
		s, ok := server.(*tcp.Server)
		if ok {
			tcpServer = s
		}
	}
	conn, err := MustConnectServer(tcpServer.RealTCPAddr())
	test.Nil(t, err)
	return tg, controller, conn
}

func MustConnectServer(tcpAddr *net.TCPAddr) (net.Conn, error) {
	conn, err := net.DialTimeout("tcp", tcpAddr.String(), time.Second)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
