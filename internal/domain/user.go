package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int      `json:"id,identity" gorm:"primary_key"`
	Name     string   `json:"name"`
	Username string   `json:"username" gorm:"unique"`
	Email    string   `json:"email" gorm:"unique"`
	Password string   `json:"password"`
	Image    int      `json:"image"`
	Coupons  []Coupon `json:"coupons" gorm:"foreignKey:UserID;references:ID"`
}

type UserDTO struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Image    int    `json:"image"`
}

type UserFriends struct {
	gorm.Model
	ID       int `json:"id,identity" gorm:"primary_key"`
	UserID   int `json:"user_id"`
	FriendID int `json:"friend_id"`
}

type Team struct {
	gorm.Model
	ID int `json:"id,identity" gorm:"primary_key"`
}

type UserXTeam struct {
	gorm.Model
	ID      int  `json:"id,identity" gorm:"primary_key"`
	TeamID  int  `json:"team_id"`
	UserID  int  `json:"user_id"`
	QuestID int  `json:"quest_id"`
	Accept  bool `json:"accepted"`
}

type WaitRoomDTO struct {
	UsersAccepted []UserDTO `json:"users_accepted"`
}

type Notification struct {
	gorm.Model
	ID         int       `json:"id,identity" gorm:"primary_key"`
	SenderID   int       `json:"sender_id"`
	ReceiverID int       `json:"receiver_id"`
	Type       string    `json:"type"`
	QuestName  string    `json:"quest_name"`
	TeamID     int       `json:"team_id"`
	QuestID    int       `json:"quest_id"`
	SentTime   time.Time `json:"sent_time"`
}

type NotificationDTO struct {
	ID         int    `json:"id,identity" gorm:"primary_key"`
	SenderID   int    `json:"sender_id"`
	TeamID     int    `json:"team_id"`
	QuestID    int    `json:"quest_id"`
	SenderName string `json:"sender_name"`
	Type       string `json:"type"`
	QuestName  string `json:"quest_name"`
}
