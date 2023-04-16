package domain

type Videos struct {
	Id          uint   `gorm:"primarykey;autoIncrement;not null"`
	Name        string `gorm:"type:varchar(255);"`
	Description string `gorm:"type:varchar(255);"`
	Icon        string `gorm:"type:varchar(255);"`
	VideoURL    string `gorm:"type:varchar(255);"`
	Views       int    `gorm:"type:integer default 0"`
	Size        int64  `gorm:"type:integer"`
	ChannelId   uint   `gorm:"foreignKey:id"`
	Channel     Channel
	CreatedAt   string `gorm:"type:time with time zone"`
	IsBlock     bool   `gorm:"type:boolean;default:false"`
	// IsHide      bool   `gorm:"type:boolean;default:false"`
}

func (channel *Channel) Get() (*Channel, error) {
	err := Db.Where("id = ?", channel.Id).First(channel).Error

	if err != nil {
		return nil, err
	}

	return channel, nil
}

func (videos *Videos) GetAllVideosFromChannel() []Videos {
	var video []Videos
	err := Db.Where("channel_id = ? and is_block = false", videos.ChannelId).Find(&video).Error
	if err != nil {
		return nil
	}

	return video

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
	err := Db.Where("is_block = false").Find(&results).Error
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
	Db.Where("id = ?", video.Id).Updates(video)
}
