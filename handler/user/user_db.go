package user

import (
	"github.com/tgo-team/tgo-talkapi/handler/db"
	"github.com/tgo-team/tgo-talkapi/utils/qreflect"
)

type Dao interface {
	InsertUser(user *User) error
	// QueryUserWithUsername 根据用户名查询用户
	QueryUserWithUsername(username string) (*User, error)
	// QueryUser 查询单个用户 根据可选参数
	QueryUserWithOption(talkId uint64, openId string) (*User, error)
}

type User struct {
	Nickname string `json:"nickname"` // 昵称
	Username string `json:"username"` // 用户名
	Password string `json:"password"`
	Sex      int    `json:"sex"`     // 0. 未设置 1.男 2.女
	OpenId   string `json:"open_id"` // 用户名为一ID
	TalkId   uint64 `json:"talk_id"` // 用户在talk服务上的唯一ID
	Zone     string `json:"zone"`    // 区号
	Mobile   string `json:"mobile"`  // 手机号
	db.BaseModel
}

type DefaultDao struct {
}

func NewDefaultDao() *DefaultDao {
	return &DefaultDao{}
}

func (d *DefaultDao) InsertUser(user *User) error {
	_, err := db.NewSession().InsertInto("user").Columns(qreflect.TagNameJsonNames(*user)...).Record(user).Exec()
	return err
}

func (d *DefaultDao) QueryUserWithUsername(username string) (*User, error) {
	var user *User
	err := db.NewSession().Select("*").From("user").Where("(username=? or mobile=?) and is_deleted=0", username, username).LoadOne(&user)
	return user, err
}

// QueryUser 查询单个用户 根据可选参数
func (d *DefaultDao) QueryUserWithOption(talkId uint64, openId string) (*User, error) {
	var user *User
	builder := db.NewSession().Select("*").From("user")
	hasOption := false
	if talkId > 0 {
		hasOption = true
		builder = builder.Where("talk_id=?", talkId)
	}
	if openId != "" {
		hasOption = true
		builder = builder.Where("open_id=?", talkId)
	}

	if !hasOption {
		return nil, nil
	}
	err := builder.Limit(1).LoadOne(&user)
	return user, err
}
