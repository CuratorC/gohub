package user

import "gohub/app/models"

type User struct {
	models.BaseModel

	Name     string `gorm:"type:varchar(255);not null" json:"name,omitempty"`
	Email    string `gorm:"type:varchar(255);not null" json:"-"`
	Phone    string `gorm:"type:varchar(255);not null" json:"-"`
	Password string `gorm:"type:varchar(255);not null" json:"-"`

	models.CommonTimestampsField
}
