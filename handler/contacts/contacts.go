package contacts

import (
	"github.com/jinzhu/gorm"
	"github.com/tgo-team/tgo-talkapi/handler/user"
)

type Contacts struct {
	User       user.User `gorm:"foreignkey:UId"`
	UId        uint64
	TargetUser user.User `gorm:"foreignkey:TargetUId"`
	TargetUId uint64
	gorm.Model
}
