package user_models

import (
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"server/common/models"
)

// FriendsModel 好友表
type FriendsModel struct {
	models.Model

	SendUserID     uint      `json:"sendUserID"` // 发起验证方
	SendUser       UserModel `json:"-" gorm:"foreignKey:SendUserID"`
	RecvUserID     uint      `json:"recvUserID"` // 接受验证方
	RecvUser       UserModel `json:"-" gorm:"foreignKey:RecvUserID"`
	SendUserNotice string    `gorm:"size:128" json:"sendUserNotice"` // 发送方备注
	RecvUserNotice string    `gorm:"size:128" json:"recvUserNotice"` // 接收方备注

	// Notice string `gorm:"size:128" json:"notice"` // 备注
}

func (*FriendsModel) TableName() string {
	return "friends"
}

func (f *FriendsModel) IsFriend(db *gorm.DB, A, B uint) bool {
	err := db.Where("((send_user_id = ? and recv_user_id = ?) or (send_user_id = ? and recv_user_id = ? ))", A, B, B, A).
		Take(f).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	} else if err != nil {
		logx.Error(err)
		return false

	}
	return true
}

func (f *FriendsModel) Friends(db *gorm.DB, userID uint) (list []FriendsModel) {
	db.Find(&list, "send_user_id = ? or recv_user_id = ?", userID, userID)
	return
}

func (f *FriendsModel) GetUserNotice(userID uint) string {
	if userID == f.SendUserID {
		// 如果我是发起方
		return f.SendUserNotice
	}
	if userID == f.RecvUserID {
		// 如果我是接收方
		return f.RecvUserNotice
	}
	return ""
}
