package domain

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID            int      `json:"id,identity" gorm:"primary_key"`
	Name          string   `json:"name"`
	Username      string   `json:"username" gorm:"unique"`
	Email         string   `json:"email" gorm:"unique"`
	Password      string   `json:"password"`
	FirebaseToken string   `json:"firebaseToken"`
	Image         int      `json:"image" gorm:"default:1"`
	Manual        bool     `json:"manual" gorm:"default:false"`
	Google        bool     `json:"google" gorm:"default:false"`
	Facebook      bool     `json:"facebook" gorm:"default:false"`
	Coupons       []Coupon `json:"coupons" gorm:"foreignKey:UserID;references:ID"`

	//Achivements
	MadeFriend_ac         bool `json:"madeFriend_ac"`
	StartedQuest_ac       bool `json:"startedQuest_ac"`
	FinishedQuest_ac      bool `json:"finishedQuest_ac"`
	FinishedTeamQuest_ac  bool `json:"finishedTeamQuest_ac"`
	RatedQuest_ac         bool `json:"ratedQuest_ac"`
	UsedCoupon_ac         bool `json:"usedCoupon_ac"`
	FinishedFiveQuests_ac bool `json:"finishedFiveQuests_ac"`
	TopThreeRanking_ac    bool `json:"topThreeRanking_ac"`
	FiftyMinutes_ac       bool `json:"fiftyMinutes_ac"`
}

type UserDTO struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	FirebaseToken string `json:"firebaseToken"`
	Image         int    `json:"image"`
	Manual        bool   `json:"manual"`
	Google        bool   `json:"google"`
	Facebook      bool   `json:"facebook"`

	//Achivements
	MadeFriend_ac         bool `json:"madeFriend_ac"`
	StartedQuest_ac       bool `json:"startedQuest_ac"`
	FinishedQuest_ac      bool `json:"finishedQuest_ac"`
	FinishedTeamQuest_ac  bool `json:"finishedTeamQuest_ac"`
	RatedQuest_ac         bool `json:"ratedQuest_ac"`
	UsedCoupon_ac         bool `json:"usedCoupon_ac"`
	FinishedFiveQuests_ac bool `json:"finishedFiveQuests_ac"`
	TopThreeRanking_ac    bool `json:"topThreeRanking_ac"`
	FiftyMinutes_ac       bool `json:"fiftyMinutes_ac"`
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
	ID          int       `json:"id,identity" gorm:"primary_key"`
	SenderID    int       `json:"sender_id"`
	SenderImage int       `json:"sender_image"`
	ReceiverID  int       `json:"receiver_id"`
	Type        string    `json:"type"`
	QuestName   string    `json:"quest_name"`
	TeamID      int       `json:"team_id"`
	QuestID     int       `json:"quest_id"`
	SentTime    time.Time `json:"sent_time"`
}

type NotificationDTO struct {
	ID          int    `json:"id,identity" gorm:"primary_key"`
	SenderID    int    `json:"sender_id"`
	TeamID      int    `json:"team_id"`
	QuestID     int    `json:"quest_id"`
	SenderName  string `json:"sender_name"`
	SenderImage int    `json:"sender_image"`
	Type        string `json:"type"`
	QuestName   string `json:"quest_name"`
}
