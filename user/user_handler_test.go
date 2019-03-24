package user

import (
	"github.com/golang/protobuf/proto"
	"github.com/tgo-team/tgo-core/tgo"
	"github.com/tgo-team/tgo-core/tgo/packets"
	"github.com/tgo-team/tgo-talkapi/test"
	_ "github.com/tgo-team/tgo-talkapi/tgo/protocol/mqtt"
	"github.com/tgo-team/tgo-talkapi/tgo/server/tcp"
	_ "github.com/tgo-team/tgo-talkapi/tgo/storage/memory"
	"net"
	"testing"
	"time"
)

func TestUserLogin(t *testing.T) {
	tg := startTGO(t)
	StartUser(tg)
	var tcpServer *tcp.Server
	for _, server := range tg.Servers {
		s, ok := server.(*tcp.Server)
		if ok {
			tcpServer = s
		}
	}
	conn, err := MustConnectServer(tcpServer.RealTCPAddr())
	test.Nil(t, err)

	loginReq := &UserLoginReq{}
	loginReq.Username = "1"
	loginReq.Password = "1"
	data,err := proto.Marshal(loginReq)
	test.Nil(t,err)
	cp := packets.NewCMDPacket(100, data)
	cp.TokenFlag = true
	cp.Token = "2334"
	sendCMDPacket(t, conn, tg, cp)
	time.Sleep(time.Millisecond * 500)
}

func sendCMDPacket(t *testing.T, conn net.Conn, tg *tgo.TGO, cmdPacket *packets.CMDPacket) {
	WritePacket(t, conn, cmdPacket, tg)
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
