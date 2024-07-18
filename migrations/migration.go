package migrations

import (
	"log"

	"github.com/juniorteye/devCamp/database"
	"github.com/juniorteye/devCamp/model"
)

func Migrate() {
	db := database.DB.Db
	err := db.AutoMigrate(&model.Bootcamp{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}
