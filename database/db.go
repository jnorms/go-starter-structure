package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"jnorms.dev/common"
	"log"
)

type Database struct {
	*common.Database
}

func (database *Database) Connect() *gorm.DB {
	db, err := gorm.Open(database.driver(), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connect database")
	}

	return db
}

func (database *Database) driver() gorm.Dialector {
	host := database.Host
	user := database.User
	password := database.Password
	dbname := database.Name
	port := database.Port
	sslmode := database.Ssl
	timezone := database.Timezone

	switch database.Driver {
	case common.DB_DRIVER_MYSQL:
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
			host,
			user,
			password,
			dbname,
			port,
			sslmode,
			timezone,
		)

		return mysql.Open(dsn)
	case common.DB_DRIVER_POTSGRES:
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
			host,
			user,
			password,
			dbname,
			port,
			sslmode,
			timezone,
		)

		return postgres.Open(dsn)
	}

	return nil
}
