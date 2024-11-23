package user

import (
	"gohub/app/models"
	"gohub/pkg/database"
	"gohub/pkg/hash"
)

type User struct {
	models.BaseModel

	Name     string `gorm:"type:varchar(255);not null" json:"name,omitempty"`
	Email    string `gorm:"type:varchar(255);not null" json:"-"`
	Phone    string `gorm:"type:varchar(255);not null" json:"-"`
	Password string `gorm:"type:varchar(255);not null" json:"-"`

	models.CommonTimestampsField
}

// Create 创建用户，通过 User.ID 来判断是否创建成功
func (userModel *User) Create() {
	database.DB.Create(&userModel)
}

// ComparePassword 密码是否正确
func (userModel *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, userModel.Password)
}
