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

// Get 通过 ID 获取用户
func Get(idStr string) (userModel User) {
	database.DB.Where("id", idStr).First(&userModel)
	return
}

// Save 保存
func (userModel *User) Save() (rowsAffected int64) {
	result := database.DB.Save(&userModel)
	return result.RowsAffected
}

// All 获取所有用户数据
func All() (users []User) {
	database.DB.Find(&users)
	return
}
