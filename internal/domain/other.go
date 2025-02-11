package domain

import (
	"gorm.io/gorm"
	"time"
)

var Db *gorm.DB

type Channel struct {
	Id          uint        `gorm:"primarykey;autoIncrement;not null"`
	OwnerId     uint        `gorm:"not null; foreignKey:id onUpdate:CASCADE; onDelete:CASCADE"`
	Owner       UserModel   `json:"-"`
	Name        string      `gorm:"type:varchar(255);"`
	Description string      `gorm:"type:varchar(255);"`
	SocialLink  string      `gorm:"type:varchar(255);"`
	Banner      string      `gorm:"type:varchar(255);"`
	Icon        string      `gorm:"type:varchar(255);"`
	Subscribers []UserModel `gorm:"many2many:channel_subscription;"`
}

type Role struct {
	Id          uint `gorm:"primarykey;autoIncrement;not null"`
	ChannelId   int
	Channel     Channel     `gorm:"foreignKey:channel_id; onUpdate:CASCADE; onDelete:CASCADE"`
	Users       []UserModel `gorm:"many2many:user_roles; onUpdate:CASCADE; onDelete:CASCADE"`
	Weight      int         `gorm:"integer; default 0"`
	Permission  int64       `gorm:"type:bigint"`
	Name        string      `gorm:"type:varchar(255);"`
	Description string      `gorm:"type:varchar(255);"`
}

type UserModel struct {
	Id            uint      			`gorm:"primarykey;autoIncrement;not null"`
	Icon          string    			`gorm:"type:varchar(255);"`
	Username      string    			`gorm:"type:varchar(255);not null"`
	Email         string    			`gorm:"type:varchar(255);"`
	Password      string    			`gorm:"type:varchar(255);"`
	Permission    int64     			`gorm:"type:bigint;default:1380863"`
	Incredentials string    			`gorm:"column:credentials type:text"`
	ValideAccount bool      			`gorm:"type:bool; default false"`
	Disable       bool      			`gorm:"type:bool; default false"`
	Online        bool      			`gorm:"type:bool; default false"`
	Subscribtion  []Channel 			`gorm:"many2many:channel_subscription;  onUpdate:CASCADE; onDelete:CASCADE"`
	Roles         []Role    			`gorm:"many2many:user_roles; onUpdate:CASCADE; onDelete:CASCADE"`
	CreatedAt     time.Time             `gorm:"default:CURRENT_TIMESTAMP"`
}

type Message struct {
	Id      uint       `gorm:"primarykey;autoIncrement;not null"`
	Content string     `json:"Content"`
	VideoId uint       `gorm:"foreignKey:id"`
	UserId  uint       `gorm:"foreignKey:id"`
	User    UserModel
	Created string     `gorm:"type:time without time zone"`
}

func (user *UserModel) TableName() string {
	return "users"
}

func (r *Role) TableName() string {
	return "roles"
}

func (channel *Channel) TableName() string {
	return "channels"
}

func (message *Message) TableName() string {
	return "messages"
}