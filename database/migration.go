package database

import (
	"BelajarKafka/config"
	"BelajarKafka/helper"
	"BelajarKafka/models"
	"fmt"
)

func MigrateDatabase() {
	fmt.Println("running migration...")
	err := config.CreateDBConnection().AutoMigrate(
		&models.User{},
	)
	if err != nil {
		helper.Exception(err)
		panic(err)
	}
}
