package model

import "time"

type User struct {
	ID          uint      `gorm:"primaryKey"`
	Username    string    `gorm:"type:varchar(255);not null;unique"`
	Email       string    `gorm:"type:varchar(255);not null;unique"`
	Password    string    `gorm:"type:varchar(255);not null"`
	Phonenumber string    `gorm:"type:varchar(255);not null"`
	Address     string    `gorm:"type:varchar(255);not null"`
	Hobby       string    `gorm:"type:varchar(255);not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}
