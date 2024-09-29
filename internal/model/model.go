package model

import (
	"time"
)

type User struct {
	ID     			    uint   `gorm:"primaryKey" json:"id"`
	Name     			string `gorm:"not null" json:"name"`
	Email    			string `gorm:"uniqueIndex;not null;type:varchar(255)" json:"email"`
	EncryptedPassword 	string `gorm:"not null" json:"-"`
	CreatedAt 			time.Time
	UpdatedAt 			time.Time
	DeletedAt 			*time.Time `gorm:"index"`
	Friends             []Friend
	Rooms               []Room
	Messages            []Message
}

type Friend struct {
	ID                  uint 	`gorm:"primaryKey" json:"id"`
	UserID      		uint 	`gorm:"not null" json:"userId"`
	User 	  			User 	`gorm:"foreignKey:UserID"`
	Name    			string 	`gorm:"not null" json:"name"`
	Mbti   				string 	`gorm:"not null" json:"mbti"`
	Age    				int 	`gorm:"not null" json:"age"`
	Gender   		    string	`gorm:"not null" json:"gender"`
	ProfilePictureUrl 	*string	`json:"profilePictureUrl"`
	Description 		*string `json:"description"`
	CreatedAt 			time.Time
	UpdatedAt 			time.Time
	DeletedAt 			*time.Time `gorm:"index"`
	Rooms               []Room     `gorm:"constraint:OnDelete:CASCADE"`
	Messages            []Message  `gorm:"constraint:OnDelete:CASCADE"`
}

type Room struct {
	ID         uint    `gorm:"primaryKey" json:"id"`
    UserID     uint    `gorm:"not null" json:"userId"`
    User       User    `gorm:"foreignKey:UserID" json:"-"`
    FriendID   uint    `gorm:"not null;constraint:OnDelete:CASCADE" json:"friendId"`
    Friend     Friend  `gorm:"foreignKey:FriendID" json:"-"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `gorm:"index"`
	Messages   []Message `gorm:"constraint:OnDelete:CASCADE"`
}

type Message struct {
	ID         uint     `gorm:"primaryKey" json:"id"`
	RoomID     uint     `json:"roomId"`
	Room       Room     `gorm:"foreignKey:RoomID;constraint:OnDelete:CASCADE" json:"-"`
    UserID     *uint    `json:"userId"`
    User       *User    `gorm:"foreignKey:UserID" json:"-"`
    FriendID   *uint    `json:"friendId"`
    Friend     *Friend  `gorm:"foreignKey:FriendID;constraint:OnDelete:CASCADE" json:"-"`
	Content    string   `gorm:"not null" json:"content"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time `gorm:"index"`
}