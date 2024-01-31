package models

import "time"

type UserType uint8

const (
	USER_ADMIN  UserType = iota
	USER_MEMBER UserType = iota
)

type User struct {
	ID        int      `gorm:"primaryKey"`
	FirstName string   `gorm:"size:30"`
	LastName  string   `gorm:"size:30"`
	Username  string   `gorm:"size:50;uniqueIndex"`
	Email     string   `gorm:"size:50;uniqueIndex"`
	Password  string   `gorm:"size:128"`
	Type      UserType `gorm:"size:1;default:1"`
	UpdatedAt time.Time
	CreatedAt time.Time
	Bio       UserBio `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE,OnUpdate:CASCADE"`
}

type UserGender string

const (
	FEMALE UserGender = "FEMALE"
	MALE   UserGender = "MALE"
)

type UserBio struct {
	ID        int `gorm:"primaryKey"`
	UserID    int
	Avatar    *string    `gorm:"size:255;default:media/user_avatar.png"`
	Gender    UserGender `gorm:"size:6;default:MALE"`
	Birth     *string    `gorm:"type:date"`
	StudentID uint
	UpdatedAt time.Time
	CreatedAt time.Time
}
