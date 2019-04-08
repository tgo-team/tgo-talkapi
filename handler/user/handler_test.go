package user

import (
	"github.com/golang/protobuf/proto"
	"github.com/tgo-team/tgo-core/tgo/packets"
	"github.com/tgo-team/tgo-talkapi/cmd"
	"github.com/tgo-team/tgo-talkapi/handler"
	"github.com/tgo-team/tgo-talkapi/test"
	"testing"
	"time"
)

func TestUserLogin(t *testing.T) {
	tg,controller,conn := handler.StartTGO(t)
	u := NewHandler()
	u.RegisterHandler(controller)

	loginReq := &UserLoginReq{
		Username: "test",
		Password: "123456",
	}
	data,_ :=proto.Marshal(loginReq)
	cp := packets.NewCmdPacket("login", data)
	cp.TokenFlag = true
	cp.Token = "2334"
	handler.SendCmdPacket(t, conn, tg, cp)

	respPacket,err := tg.GetOpts().Pro.DecodePacket(conn)
	test.Nil(t,err)

	cmdackPacket := respPacket.(*packets.CmdackPacket)

	test.Equal(t,cmd.SUCCESS,cmdackPacket.Status)



	time.Sleep(time.Millisecond*50)
}
