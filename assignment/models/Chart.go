package models

import "google.golang.org/protobuf/types/known/timestamppb"

type Chart_point struct {
	Id       string
	Chart_id string
	X_v      string
	Y_v      string
}

type Chart struct {
	Id            string                `json:"id" gorm:"primaryKey"`
	X_description string                `json:"x_description"`
	Y_description string                `json:"y_description"`
	CreationDate  timestamppb.Timestamp `json:"creation_date"`
	UpdateDate    timestamppb.Timestamp `json:"update_date"`
	Charints      []Chart_point
	Asset         Asset `gorm:"foreignKey:Id"`
}
