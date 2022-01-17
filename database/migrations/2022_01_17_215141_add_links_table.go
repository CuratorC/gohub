package migrations

import (
	"database/sql"
	"gohub/app/models"
	"gohub/pkg/logger"
	"gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type Link struct {
		models.BaseModel

		Name string `gorm:"type:varchar(255);not null"`
		URL  string `gorm:"type:varchar(255);default:null"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		err := migrator.AutoMigrate(&Link{})
		if err != nil {
			logger.ErrorString("migration", "links", "AutoMigrate error")
			return
		}
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		err := migrator.DropTable(&Link{})
		if err != nil {
			logger.ErrorString("migration", "links", "DropTable error")
			return
		}
	}

	migrate.Add("2022_01_17_215141_add_links_table", up, down)
}
