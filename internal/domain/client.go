package domain

type Client struct {
	ID     int         `json:"id,identity" gorm:"primary_key"`
	Name   string      `json:"name"`
	Image  string      `json:"image"`
	Quest  []QuestInfo `json:"quests" gorm:"foreignKey:ClientID;references:ID"`
	Coupon []Coupon    `json:"coupons" gorm:"foreignKey:ClientID;references:ID"`
}

type ClientDTO struct {
	ID    int
	Name  string `json:"name"`
	Image string `json:"image"`
}
