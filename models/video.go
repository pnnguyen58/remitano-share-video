package models

import (
	"errors"
	"reflect"
)

type Video struct {
	Id       		int64  `json:"id"  gorm:"AUTO_INCREMENT;primaryKey;NOT NULL"`
	UserId       		int64  `json:"userId"  gorm:"column:user_id;NOT NULL"`
	Title  			string  `json:"title" gorm:"column:title;NOT NULL"`
	Description  	string  `json:"description" gorm:"column:description;NULL"`
	Link  			string  `json:"link" gorm:"column:link;NOT NULL"`
	CreatedAt     string          	`json:"createdAt" gorm:"<-:false;column:created_at;NULL"`
}

func (v Video) IsEmpty() bool {
	return reflect.DeepEqual(Video{}, v)
}

func (Video) TableName() string {
	return "videos"
}

func (v *Video) Share() (err error) {
	if err = sqlDB.Ping(); db == nil || sqlDB == nil || err != nil {
		err = connect()
	}
	if err == nil {
		u := User{ID: v.UserId}
		err = u.Get()
		if err != nil {
			return
		}
		if u.Username != "" {
			err = db.Create(v).Error
		} else {
			return errors.New("user not found")
		}
	}
	return
}

func GetVideos(id int64, limit int, offset int) (videos []Video, err error) {
	if err = sqlDB.Ping(); db == nil || sqlDB == nil || err != nil {
		err = connect()
	}
	if err == nil {
		if limit <= 0 {
			limit = -1
		}
		if offset <= 0 {
			offset = -1
		}
		err = db.Where(&Video{UserId: id}).Limit(limit).Offset(offset).Find(&videos).Error
	}
	return
}