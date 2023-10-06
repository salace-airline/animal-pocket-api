package database

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/salace-airline/animalpocketresources/models"

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
		"host=db user=%s password=%s dbname=%s port=5432 sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
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
	DB.Db.Migrator().DropTable(&models.Fish{}, &models.Bug{}, &models.SeaCreature{})
	log.Println("running migrations")
	DB.Db.AutoMigrate(&models.Fish{}, &models.Bug{}, &models.SeaCreature{})
}

func IntegrateResources() {
	log.Println("integrate resources: fishes, bugs and sea creatures")
	// Fish table
	fileFish, err := os.Open("./database/fishes.json")
	if err != nil {
		log.Fatal(err)
	}
	defer fileFish.Close()

	var fishes []models.Fish
	decoderFish := json.NewDecoder(fileFish)
	err = decoderFish.Decode(&fishes)
	if err != nil {
		log.Fatal(err)
	}

	for _, fish := range fishes {
		err := DB.Db.Create(&fish).Error
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Fishes were successfully insert into database!")

	// Bug table
	fileBug, err := os.Open("./database/bugs.json")
	if err != nil {
		log.Fatal(err)
	}
	defer fileBug.Close()

	var bugs []models.Bug
	decoderBug := json.NewDecoder(fileBug)
	err = decoderBug.Decode(&bugs)
	if err != nil {
		log.Fatal(err)
	}

	for _, bug := range bugs {
		err := DB.Db.Create(&bug).Error
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Bugs were successfully insert into database!")

	// Sea creature table
	fileSeaCreature, err := os.Open("./database/bugs.json")
	if err != nil {
		log.Fatal(err)
	}
	defer fileSeaCreature.Close()

	var seaCreatures []models.SeaCreature
	decoderSea := json.NewDecoder(fileSeaCreature)
	err = decoderSea.Decode(&seaCreatures)
	if err != nil {
		log.Fatal(err)
	}

	for _, SeaCreature := range seaCreatures {
		err := DB.Db.Create(&SeaCreature).Error
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Sea Creatures were successfully insert into database!")
}
