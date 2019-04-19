package contacts

import (
	"fmt"
	"github.com/tgo-team/tgo-talkapi/cmd"
	"github.com/tgo-team/tgo-talkapi/ctrl"
	"github.com/tgo-team/tgo-talkapi/middleware"
)

type Handler struct {
	dao Dao
}

func NewHandler() *Handler {
	return &Handler{
		dao: NewDefaultDao(),
	}
}

// 同步联系人
func (h *Handler) SyncContacts(c *cmd.Context) {
	openId := c.CurrentOpenId()
	syncContactsReq := &SyncContactsReq{}
	err := c.UnmarshalProto(c.Param(), syncContactsReq)
	if err != nil {
		c.Error("数据格式有误！-> %v", err)
		c.ReplyErrorMsg("数据格式有误！")
		return
	}
	contacts, err := h.dao.SyncContacts(syncContactsReq.SyncKey, openId, syncContactsReq.Limit)
	if err != nil {
		c.Error("数据格式有误！-> %v", err)
		c.ReplyErrorMsg("数据格式有误！")
		return
	}

	newSyncKey := syncContactsReq.SyncKey
	if contacts != nil && len(contacts) > 0 {
		lastContacts := contacts[len(contacts)-1]
		newSyncKey = fmt.Sprintf("%s@%d",lastContacts.UpdatedAt.String(),lastContacts.Id)
	}

	contactsVos := make([]*ContactsVo,0,len(contacts))
	for _,contact :=range contacts {
		contactsVos = append(contactsVos,h.contactsDetailToVo(contact))
	}

	syncContactsResp := &SyncContactsResp{
		SyncKey: newSyncKey,
		Contacts:contactsVos,
	}
	respData := c.MarshalProto(syncContactsResp)
	c.ReplySuccess(respData)

}

func (h *Handler) RegisterHandler(c *ctrl.Controller) {
	// 同步联系人
	c.RegisterHandlerFuncs("sync_contacts", middleware.Auth, h.SyncContacts)
}

func (h *Handler) contactsDetailToVo(contacts *ContactsDetail) *ContactsVo {
	contactsVo := &ContactsVo{}
	contactsVo.Id = contacts.Id
	contactsVo.OpenId = contacts.OpenId
	contactsVo.RelationOpenId = contacts.RelationOpenId
	contactsVo.Remark = contacts.Remark
	contactsVo.CreatedAt = contacts.CreatedAt.String()
	contactsVo.UpdatedAt = contacts.UpdatedAt.String()
	contactsVo.Type = int32(contacts.Type)
	contactsVo.RelationNickname = contacts.RelationNickname
	contactsVo.TalkId =  contacts.TalkId
	return contactsVo
}