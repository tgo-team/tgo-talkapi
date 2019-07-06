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
	cfg := handler.GetTestConfig()
	tg, controller, conn := handler.StartTGO(t, cfg)
	u := &Handler{
		cfg: handler.GetTestConfig(),
		dao: NewTestDao(),
	}
	u.RegisterHandler(controller)

	loginReq := &UserLoginReq{
		Username: "test",
		Password: "123456",
	}
	data, _ := proto.Marshal(loginReq)
	cp := packets.NewCmdPacket("login", data)
	cp.TokenFlag = true
	cp.Token = "2334"
	cmdackPacket := handler.SendCmdPacket(t, conn, tg, cp)

	test.Equal(t, cmd.SUCCESS, cmdackPacket.Status)

	time.Sleep(time.Millisecond * 50)
}

type TestDao struct {
	cacheMap map[string]interface{}
}

func NewTestDao() *TestDao {
	return &TestDao{
		cacheMap: map[string]interface{}{},
	}
}

func (t *TestDao) InsertUser(user *User) error {

	t.cacheMap[user.Username] = user
	return nil

}
func (t *TestDao) QueryUser(username string) (*User, error) {

	return &User{Username: "test", Password: "123456"}, nil
}

func (t *TestDao) QueryUserWithUsername(username string) (*User, error) {
	return &User{Username: "test", Password: "123456"}, nil
}
func (t *TestDao) QueryUserWithOption(talkId uint64, openId string) (*User, error) {
	return &User{Username: "test", Password: "123456"}, nil
}
