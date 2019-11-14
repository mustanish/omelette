package entity

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name        string    `gorm:"type:varchar(100);DEFAULT:NULL"`
	UserName    *string   `gorm:"type:varchar(100);unique_index;DEFAULT:NULL"`
	Email       *string   `gorm:"type:varchar(100);unique_index;DEFAULT:NULL"`
	EmailVerify int64     `gorm:"type:tinyint(1);DEFAULT:0"`
	PhoneNumber *string   `gorm:"type:varchar(10);unique_index;DEFAULT:NULL"`
	PhoneVerify int64     `gorm:"type:tinyint(1);DEFAULT:0"`
	Age         int64     `gorm:"type:int(5);DEFAULT:0"`
	OTP         int64     `gorm:"type:int(6);DEFAULT:0"`
	OtpType     string    `gorm:"type:varchar(20);DEFAULT:NULL"`
	OtpValidity int64     `gorm:"type:int(10);DEFAULT:0"`
	LastLogedIn time.Time `gorm:"type:timestamp;DEFAULT:CURRENT_TIMESTAMP"`
}
