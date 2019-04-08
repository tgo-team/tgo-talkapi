package user

import (
	"github.com/tgo-team/tgo-talkapi/handler/db"
	"github.com/tgo-team/tgo-talkapi/utils/qreflect"
)

type User struct {
	Nickname string `json:"nickname"`                        // 昵称
	Username string `json:"username"` // 用户名
	Password string `json:"password"`
	Sex      int `json:"sex"`                    // 0. 未设置 1.男 2.女
	OpenId   string `json:"open_id"` // 用户名为一ID
	TalkId   uint64 `json:"talk_id"`                       // 用户在talk服务上的唯一ID
	Zone     string `json:"zone"`         // 区号
	Mobile   string `json:"mobile"`         // 手机号
	db.BaseModel
}

func InsertUser(user *User) error {
	_,err := db.NewSession().InsertInto("user").Columns(qreflect.TagNameJsonNames(*user)...).Record(user).Exec()
	return err
}

func QueryUser(username string) (*User,error)  {
	var user *User
	err := db.NewSession().Select("*").From("user").Where("username=? or mobile=?",username,username).LoadOne(&user)
	return user,err
}
