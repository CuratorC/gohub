package migrations

import (
	"database/sql"
	"gohub/pkg/logger"
	"gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type User struct {
		City         string `gorm:"type:varchar(10)"`
		Introduction string `gorm:"type:varchar(255);"`
		Avatar       string `gorm:"type:varchar(255);default:null"`
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		err := migrator.AutoMigrate(&User{})
		if err != nil {
			logger.ErrorString("migration", "users", "AutoMigrate error")
			return
		}
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		err := migrator.DropColumn(&User{}, "City")
		logger.LogIf(err)
		err = migrator.DropColumn(&User{}, "Introduction")
		logger.LogIf(err)
		err = migrator.DropColumn(&User{}, "Avatar")
		logger.LogIf(err)
	}

	migrate.Add("2022_01_18_092416_add_fields_to_user", up, down)
}
