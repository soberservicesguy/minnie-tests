package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/soberservicesguy/minnie-test-backend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(dotEnvPath string) {
	var err error
	p := config.Config(dotEnvPath, "DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("please choose proper port")
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=%s", config.Config(dotEnvPath, "DB_HOST"), port, config.Config(dotEnvPath, "DB_USER"), config.Config(dotEnvPath, "DB_PASSWORD"), config.Config(dotEnvPath, "DB_NAME"), config.Config(dotEnvPath, "DB_TIMEZONE"))
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")

}
