package domain

import "time"

type Videos struct {
	Id           uint      `gorm:"primarykey;autoIncrement;not null"`
	Name         string    `gorm:"type:varchar(255);"`
	Description  string    `gorm:"type:varchar(255);"`
	Icon         string    `gorm:"type:varchar(255);"`
	VideoURL     string    `gorm:"type:varchar(255);"`
	Views        int       `gorm:"type:integer default:0"`
	ChannelId    uint      `gorm:"foreignKey:id"`
	CreationDate time.Time `gorm:"type:datetime"`
	IsBlock      bool      `gorm:"type:boolean;default:false"`
}

func (channel *Channel) Get() *Channel {
	tx := Db.Where("id = ?", channel.Id).First(channel)

	if tx.RowsAffected == 0 {
		return nil
	}

	return channel
}

func (video *Videos) GetChannel() *Channel {
	var channel *Channel
	err := Db.Joins("JOIN channels c ON c.id = video_info.channel_id").
		Where("video_info.id = ?", video.Id).
		First(&channel).Error
	if err != nil {
		return nil
	}

	return channel
}

func (video *Videos) GetAllVideosFromChannel() []Videos {
	var videos []Videos 
	err := Db.Where("channel_id = ?", video.ChannelId)
	if err != nil {
		return nil
	}

	return videos

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
	err := Db.Find(&results).Error
	if err != nil {
		return nil, err
	}
	return results, nil
}

func (video *Videos) Find() bool {
	tx := Db.Where("videourl = ?", video.VideoURL).Find(video)
	return tx.RowsAffected != 0
}

func (video *Videos) Delete() {
	Db.Delete(video)
}

func (video *Videos) Update() {
	Db.Save(&video)
}
