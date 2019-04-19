package contacts

import (
	"github.com/tgo-team/tgo-talkapi/handler/db"
	"strconv"
	"strings"
)

type Dao interface {
	// 同步联系人
	SyncContacts(syncKey string, openId string, limit uint64) ([]*ContactsDetail, error)
}

type ContactsDetail struct {
	Contacts
	RelationNickname string `json:"relation_nickname"`
	TalkId           uint64 `json:"talk_id"`
}

type Contacts struct {
	OpenId         string `json:"open_id"`          // 用户open_id
	RelationOpenId string `json:"relation_open_id"` // 关联的openId
	Type           int    `json:"type"`
	Remark         string `json:"remark"` // 备注
	db.BaseModel
}

type DefaultDao struct {
}

func NewDefaultDao() *DefaultDao {
	return &DefaultDao{}
}

/// 通过同步key 查询用户的联系人
func (d *DefaultDao) SyncContacts(syncKey string, openId string, limit uint64) ([]*ContactsDetail, error) {
	var updatedAt string
	var id int64
	var err error
	if syncKey != "" {
		syncKeyArray := strings.Split(syncKey, "@")
		updatedAt = syncKeyArray[0]
		id, err = strconv.ParseInt(syncKeyArray[1], 10, 64)
		if err != nil {
			return nil, err
		}
	}
	/**
	防止相同时间的数据量大于limit的时候 出现重复问题，所以这里做了id查询
	*/
	var contacts []*ContactsDetail
	builder := db.NewSession().Select("contacts.*,user.nickname relation_nickname,user.talk_id").From("contacts").LeftJoin("user", "contacts.relation_open_id=user.open_id").Where("contacts.open_id=?", openId).OrderDir("contacts.updated_at", true).OrderDir("contacts.id", true)
	if updatedAt == "" {
		_, err = builder.Limit(limit).Load(&contacts)
		return contacts, err
	}
	// 查询重复更新时间的数据
	_, err = builder.Where("contacts.updated_at=? and contacts.id > ?", updatedAt, id).Load(&contacts)
	if err != nil {
		return nil, err
	}
	// 如果重复更新时间的数据等于或大于限制数量 就直接返回查询的数据
	if contacts != nil && uint64(len(contacts)) >= limit {
		return contacts, nil
	}
	var newlimit uint64
	newlimit = limit
	var contacts2 []*ContactsDetail
	builder2 := db.NewSession().Select("contacts.*,user.nickname relation_nickname,user.talk_id").From("contacts").LeftJoin("user", "contacts.relation_open_id=user.open_id").Where("contacts.open_id=?", openId).OrderDir("contacts.updated_at", true).OrderDir("contacts.id", true)
	if contacts != nil && len(contacts) > 0 {
		newlimit = limit - uint64(len(contacts))
	}
	_, err = builder2.Where("contacts.updated_at>?", updatedAt).Limit(newlimit).Load(&contacts2)
	if err != nil {
		return nil, err
	}
	newContacts := make([]*ContactsDetail, 0)
	newContacts = append(newContacts, contacts...)
	newContacts = append(newContacts, contacts2...)

	return newContacts, nil
}
