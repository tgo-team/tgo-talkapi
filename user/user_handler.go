package user

import (
	"github.com/golang/protobuf/proto"
	"github.com/tgo-team/tgo-core/tgo"
)

type User struct {
	t *tgo.TGO
}

func StartUser(t *tgo.TGO) *User {
	u := &User{}
	u.t = t
	u.route()
	return u
}

func (u *User) route() {
	u.t.Match("cmd:100", u.Login)

}

func (u *User) Login(m *tgo.MContext) {
	packet :=  m.CMDPacket()
	m.Info("用户登录！-%v",packet.Payload)
	loginReq := &UserLoginReq{}
	err := proto.Unmarshal(packet.Payload,loginReq)
	if err!=nil {
		m.Error("解析请求数据出错！-> %v",err)
		return
	}
	m.Info("loginReq--%v",loginReq.String())
}

func (u *User) Register(m *tgo.MContext)  {

}