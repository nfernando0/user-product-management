package migration

import (
	"fmt"
	"products/db"
	"products/models"
)

func RunMigration() {
	err := db.DB.AutoMigrate(
		&models.User{},
		&models.Product{},
	)

	if err != nil {
		panic(err)
	}

	fmt.Println("Database migrated")
}
