package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Quest struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"user_id,omitempty"`
	Name string             `json:"name"`
}

type QuestDTO struct {
	Name string `json:"name"`
}
