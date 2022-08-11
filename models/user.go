package models

import (
	"context"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"reflect"
	"time"
)

type User struct {
	ID       	int64  `json:"id" gorm:"AUTO_INCREMENT;primaryKey;NOT NULL"`
	Username  	string  `json:"username" gorm:"column:username;NOT NULL"`
	Password    string `json:"password" gorm:"column:password;NOT NULL"`
	CreatedAt     string          	`json:"createdAt" gorm:"<-:false;column:created_at;NULL"`
	Token    string `json:"token,omitempty" gorm:"-"`
}

func (u User) IsEmpty() bool {
	return reflect.DeepEqual(User{}, u)
}

func (User) TableName() string {
	return "users"
}

func (u *User) Register() (err error) {
	if err = sqlDB.Ping(); db == nil || sqlDB == nil || err != nil {
		err = connect()
	}
	if err == nil {
		hashedPassword, e := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if e != nil {
			return e
		}
		u.Password = string(hashedPassword)
		err = db.Create(u).Error
	}
	return
}

func (u *User) Login() (err error) {
	if err = sqlDB.Ping(); db == nil || sqlDB == nil || err != nil {
		err = connect()
	}
	if err != nil {
		return
	}
	password := u.Password
	err = db.Where(&User{Username: u.Username}).First(u).Error
	if err != nil {
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err == nil {
		u.Token = uuid.New().String()
		err = bm.Put(context.Background(), u.Token, u.ID,  1*time.Hour)
	}
	return
}

func (u *User) Get() (err error) {
	if err = sqlDB.Ping(); db == nil || sqlDB == nil || err != nil {
		err = connect()
	}
	if err == nil {
		err = db.Where(u).First(u).Error
	}
	return
}
func CheckToken(token string) (userId int64, err error) {
	id, err := bm.Get(context.Background(), token)
	if err == nil && id.(int64) > 0 {
		return id.(int64), err
	}
	return
}
