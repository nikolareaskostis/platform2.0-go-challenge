package models

import "GoRepo/source/usedtobehere/models"

type Asset struct {
	Id           string        `json:"id" gorm:"primaryKey"`
	Users        []models.User `gorm:"many2many:users_assets"`
	Favourite    bool          `json:"favourite"`
	Description  string        `json:"description"`
	CreationDate string        `json:"creationDate"`
	UpdateDate   string        `json:"updateDate"`
}
