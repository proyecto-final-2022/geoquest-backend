package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
)

type Quest struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"user_id,omitempty"`
	Name string             `json:"name"`
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
	Tags          []string `json:"tags"`
}
