package models

import (
	"GoRepo/source/usedtobehere/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Audience struct {
	Id           string `json:"id" gorm:"primaryKey"`
	Asset_id     string
	Country      string                `json:"description"`
	Age_min      int                   `json:"agemin"`
	Age_max      int                   `json:"agemax"`
	CreationDate timestamppb.Timestamp `json:"creation_date"`
	UpdateDate   timestamppb.Timestamp `json:"update_date"`
	Gender       string                `json:"gender"`
	Asset        models.Asset          `gorm:"foreignKey:Id"`
}
