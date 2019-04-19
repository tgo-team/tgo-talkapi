package contacts

import (
	"github.com/golang/protobuf/proto"
	"github.com/tgo-team/tgo-core/tgo/packets"
	"github.com/tgo-team/tgo-talkapi/cmd"
	"github.com/tgo-team/tgo-talkapi/handler"
	"github.com/tgo-team/tgo-talkapi/test"
	"testing"
	"time"
)

func TestSyncContacts(t *testing.T) {
	cfg := handler.GetTestConfig()
	tg,controller,conn := handler.StartTGO(t,cfg)

	controller.Cache.Set("token:2334","123456",0)

	u := Handler{
		dao:NewTestDao(),
	}
	u.RegisterHandler(controller)

	syncContactsReq := &SyncContactsReq{
		SyncKey: "2019-11-12@12",
		Limit: 100,
	}
	data,_ :=proto.Marshal(syncContactsReq)
	cp := packets.NewCmdPacket("sync_contacts", data)
	cp.TokenFlag = true
	cp.Token = "2334"
	cmdackPacket := handler.SendCmdPacket(t, conn, tg, cp)

	test.Equal(t,cmd.SUCCESS,cmdackPacket.Status)

	resp := SyncContactsResp{}
	err := proto.Unmarshal(cmdackPacket.Payload,&resp)
	test.Nil(t,err)

	println(resp.SyncKey)

	time.Sleep(time.Millisecond*50)
}

type TestDao struct {
	cacheMap map[string]interface{}
}

func NewTestDao() *TestDao  {
	return &TestDao{
		cacheMap: map[string]interface{}{},
	}
}

func (d *TestDao) SyncContacts(syncKey string,openId string,limit uint64) ([]*ContactsDetail,error)  {

	return []*ContactsDetail{
		{
			Contacts:Contacts{
				OpenId: "23344",
				RelationOpenId: "1223",
			},
		},
	},nil
}
