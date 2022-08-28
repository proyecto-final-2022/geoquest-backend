package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
)

type Quest struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"user_id,omitempty"`
	Name string             `json:"name"`
}

type QuestCompletion struct {
	gorm.Model
	ID        int       `json:"id,identity" gorm:"primary_key"`
	QuestID   int       `json:"quest_id"`
	UserID    int       `json:"user_id"`
	StartTime time.Time `json:"completion_time"`
	EndTime   time.Time `json:"end_time"`
}

type QuestCompletionDTO struct {
	UserID    int       `json:"user_id"`
	StartTime time.Time `json:"completion_time"`
	EndTime   time.Time `json:"end_time"`
	Hours     float64   `json:"hours"`
	Minutes   float64   `json:"minutes"`
	Seconds   float64   `json:"seconds"`
}

type QuestTeamCompletion struct {
	gorm.Model
	ID        int       `json:"id,identity" gorm:"primary_key"`
	TeamID    int       `json:"team_id"`
	QuestID   int       `json:"quest_id"`
	StartTime time.Time `json:"completion_time"`
	EndTime   time.Time `json:"end_time"`
}

type QuestTeamCompletionDTO struct {
	Users     []UserDTO `json:"users"`
	StartTime time.Time `json:"completion_time"`
	EndTime   time.Time `json:"end_time"`
	Hours     float64   `json:"hours"`
	Minutes   float64   `json:"minutes"`
	Seconds   float64   `json:"seconds"`
}

type QuestInfo struct {
	gorm.Model
	ID            int     `json:"id,identity" gorm:"primary_key"`
	ClientID      int     `json:"client_id"`
	Name          string  `json:"name"`
	Qualification float32 `json:"qualification"`
	Description   string  `json:"description"`
	Difficulty    string  `json:"difficulty"`
	Duration      string  `json:"duration"`
	Image         string  `json:"image_url"`
	Completions   int     `json:"completions"`
	Tags          []Tag   `json:"tags" gorm:"foreignKey:QuestID;references:ID"`
}

type Tag struct {
	gorm.Model
	ID          int    `json:"id,identity" gorm:"primary_key"`
	QuestID     int    `json:"quest_id"`
	Description string `json:"description"`
}

type QuestDTO struct {
	Name string `json:"name"`
}

type QuestInfoDTO struct {
	ID            int      `json:"id"`
	Name          string   `json:"name"`
	Qualification float32  `json:"qualification"`
	Description   string   `json:"description"`
	Difficulty    string   `json:"difficulty"`
	Duration      string   `json:"duration"`
	Completions   int      `json:"completions"`
	Image         string   `json:"image_url"`
	Tags          []string `json:"tags"`
}
