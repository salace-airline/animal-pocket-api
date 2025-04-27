package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Dbinstance struct {
	Db *gorm.DB
}

var DB Dbinstance

func ConnectDb() {
	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	DB = Dbinstance{
		Db: db,
	}
}

func Migrate() {
	DB.Db.Migrator().DropTable(&models.User{}, &models.Fish{}, &models.Bug{}, &models.SeaCreature{})
	log.Println("running migrations")
	DB.Db.AutoMigrate(&models.User{}, &models.Fish{}, &models.Bug{}, &models.SeaCreature{})
}

// Remove the IntegrateResources function as data integration will now be handled by SQL scripts.
