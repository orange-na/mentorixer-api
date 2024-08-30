package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     			string `gorm:"not null"`
	Email    			string `gorm:"uniqueIndex;not null;type:varchar(255)"`
	EncryptedPassword 	string `gorm:"not null"`
	Friends             []Friend
	Rooms               []Room
	Messages            []Message
}

type Friend struct {
	gorm.Model
	UserID      		uint 	`gorm:"not null"`
	User 	  			User 	`gorm:"foreignKey:UserID"`
	Name    			string 	`gorm:"not null"`
	Mbti   				string 	`gorm:"not null"`
	Age    				int 	`gorm:"not null"`
	Gender   		    string	`gorm:"not null"`
	ProfilePictureUrl 	*string	
	Description 		*string
	Rooms               []Room
	Messages            []Message
}

type Room struct {
	gorm.Model
    UserID     uint    `gorm:"not null"`
    User       User    `gorm:"foreignKey:UserID"`
    FriendID   uint    `gorm:"not null"`
    Friend     Friend  `gorm:"foreignKey:FriendID"`
	Messages   []Message
}

type Message struct {
	gorm.Model
	RoomID     uint
	Room       Room    `gorm:"foreignKey:RoomID"`
    UserID     *uint    
    User       *User    `gorm:"foreignKey:UserID"`
    FriendID   *uint    
    Friend     *Friend  `gorm:"foreignKey:FriendID"`
	Content    string   `gorm:"not null"`
}
