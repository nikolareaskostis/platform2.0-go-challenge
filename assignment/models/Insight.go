package models

type Insight struct {
	Id          string `json:"id" gorm:"primaryKey"`
	Description string `json:"description"`
	Asset       Asset  `gorm:"foreignKey:Id"`
}
