package model

type MessageModel struct {
	Model
	Content      string
	FromID       uint      `gorm:"default:NULL"`
	ReplyingToID uint      `gorm:"default:NULL"`
	From         UserModel `gorm:"default:NULL"`
	ReplyingTo   UserModel `gorm:"default:NULL"`
}
