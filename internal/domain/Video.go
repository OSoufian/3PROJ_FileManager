package domain

import (
	"strings"
	// "time"
)

type Videos struct {
	Id            uint   `gorm:"primarykey;autoIncrement;not null"`
	Name          string `gorm:"type:varchar(255);"`
	Description   string `gorm:"type:varchar(1500);"`
	Icon          string `gorm:"type:varchar(255);"`
	VideoURL      string `gorm:"type:varchar(255);"`
	Views         int    `gorm:"type:integer"`
	Size          int64  `gorm:"type:integer"`
	ChannelId     uint   `gorm:"foreignKey:id"`
	Channel       Channel
	CreatedAt     string `gorm:"column:created_at"`
	CreationDate  string `gorm:"column:creation_date"`
	IsBlock       bool   `gorm:"type:boolean;default:false"`
	IsHide        bool   `gorm:"type:boolean;default:false"`
}

func (channel *Channel) Get() (*Channel, error) {
	err := Db.Where("id = ?", channel.Id).First(channel).Error

	if err != nil {
		return nil, err
	}

	return channel, nil
}

func (videos *Videos) GetAllVideosFromChannel(orderBy ...string) []Videos {
	var video []Videos
	db := Db.Where("channel_id = ?", videos.ChannelId)

	if len(orderBy) > 0 {
		orderColumns := strings.Join(orderBy, ", ")
		db = db.Order(orderColumns)
	}

	err := db.Find(&video).Error
	if err != nil {
		return nil
	}

	return video
}

func (videos *Videos) DeleteAllVideosFromChannel(channelID uint) error {
	var videoList []Videos

	db := Db.Where("channel_id = ?", channelID)

	err := db.Find(&videoList).Error
	if err != nil {
		return err
	}

	// Delete each video and its associated messages
	for _, video := range videoList {
		// Delete messages associated with the video
		if err := Db.Where("video_id = ?", video.Id).Delete(&Message{}).Error; err != nil {
			return err
		}

		// Delete the video
		if err := Db.Delete(&video).Error; err != nil {
			return err
		}
	}

	return nil
}

func (video *Videos) TableName() string {
	return "video_info"
}

func (video *Videos) Create() error {
	tx := Db.Create(video)

	return tx.Error
}

func (video *Videos) Get() *Videos {

	tx := Db.Where("name = ?", video.Name).Find(video)
	if tx.RowsAffected == 0 {
		return nil
	}
	return video

}

func (video *Videos) GetById() *Videos {

	tx := Db.Where("id = ?", video.Id).Find(video)
	if tx.RowsAffected == 0 {
		return nil
	}
	return video

}

func (video *Videos) GetAll() ([]Videos, error) {
	var results []Videos
	err := Db.Where("is_block = false and is_hide = false").Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (video *Videos) GetSearch(search string) ([]Videos, error) {
	var results []Videos
	err := Db.Where("LOWER(name) LIKE ? OR LOWER(description) LIKE ? and is_block = false and is_hide = false order by created_at asc", "%"+strings.ToLower(search)+"%", "%"+strings.ToLower(search)+"%").Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (video *Videos) Find() bool {
	tx := Db.Where("video_url = ?", video.VideoURL).Find(video)
	return tx.RowsAffected != 0
}

func (video *Videos) Delete() {
	Db.Delete(video)
}

func (video *Videos) Update() {
	Db.Model(&Videos{}).Where("id = ?", video.Id).Updates(map[string]interface{}{
		"Name":         video.Name,
		"Description":  video.Description,
		"Icon":         video.Icon,
		"IsHide":       video.IsHide,
		"IsBlock":      video.IsBlock,
		"Views":        video.Views,
	})
}

// package domain

// import "strings"

// type Videos struct {
// 	Id            uint   `gorm:"primarykey;autoIncrement;not null"`
// 	Name          string `gorm:"type:varchar(255);"`
// 	Description   string `gorm:"type:varchar(1500);"`
// 	Icon          string `gorm:"type:varchar(255);"`
// 	VideoURL      string `gorm:"type:varchar(255);"`
// 	Views         int    `gorm:"type:integer"`
// 	Size          int64  `gorm:"type:integer"`
// 	ChannelId     uint   `gorm:"foreignKey:id"`
// 	Channel       Channel
// 	CreatedAt     string `gorm:"type:time with time zone"`
// 	CreationDate  string `gorm:"type:time with time zone"`
// 	IsBlock       bool   `gorm:"type:boolean;default:false"`
// 	IsHide        bool   `gorm:"type:boolean;default:false"`
// }

// func (channel *Channel) Get() (*Channel, error) {
// 	err := Db.Where("id = ?", channel.Id).First(channel).Error

// 	if err != nil {
// 		return nil, err
// 	}

// 	return channel, nil
// }

// func (videos *Videos) GetAllVideosFromChannel(orderBy ...string) []Videos {
// 	var video []Videos
// 	db := Db.Where("channel_id = ?", videos.ChannelId)

// 	if len(orderBy) > 0 {
// 		orderColumns := strings.Join(orderBy, ", ")
// 		db = db.Order(orderColumns)
// 	}

// 	err := db.Find(&video).Error
// 	if err != nil {
// 		return nil
// 	}

// 	return video
// }

// func (video *Videos) TableName() string {
// 	return "video_info"
// }

// func (video *Videos) Create() error {
// 	tx := Db.Create(video)

// 	return tx.Error
// }

// func (video *Videos) Get() *Videos {

// 	tx := Db.Where("name = ?", video.Name).Find(video)
// 	if tx.RowsAffected == 0 {
// 		return nil
// 	}
// 	return video

// }

// func (video *Videos) GetById() *Videos {

// 	tx := Db.Where("id = ?", video.Id).Find(video)
// 	if tx.RowsAffected == 0 {
// 		return nil
// 	}
// 	return video

// }

// func (video *Videos) GetAll(orderBy ...string) ([]Videos, error) {
// 	var results []Videos
// 	db := Db.Where("is_block = false and is_hide = false")

// 	if len(orderBy) > 0 {
// 		orderColumns := strings.Join(orderBy, ", ")
// 		db = db.Order(orderColumns)
// 	}

// 	err := db.Find(&video).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return results, nil
// }

// func (video *Videos) GetSearch(search string, orderBy ...string) ([]Videos, error) {
// 	var results []Videos
// 	db := Db.Where("LOWER(name) LIKE ? OR LOWER(description) LIKE ? and is_block = false and is_hide = false", "%"+strings.ToLower(search)+"%", "%"+strings.ToLower(search)+"%")

// 	if len(orderBy) > 0 {
// 		orderColumns := strings.Join(orderBy, ", ")
// 		db = db.Order(orderColumns)
// 	}

// 	err := db.Find(&video).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return results, nil
// }

// func (video *Videos) Find() bool {
// 	tx := Db.Where("video_url = ?", video.VideoURL).Find(video)
// 	return tx.RowsAffected != 0
// }

// func (video *Videos) Delete() {
// 	Db.Delete(video)
// }

// func (video *Videos) Update() {
// 	Db.Where("id = ?", video.Id).Updates(video)
// }
