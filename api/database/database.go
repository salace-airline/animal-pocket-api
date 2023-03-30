package database

import (
	"encoding/json"
	"fmt"
	"io"
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

	// Verify if tables exists
	if db.Migrator().HasTable("Fish") && db.Migrator().HasTable("Bug") && db.Migrator().HasTable("SeaCreature") {
		fmt.Println("Fish, Bug and SeaCreature exist in database.")
	} else {
		db.AutoMigrate(&models.Fish{}, &models.Bug{}, &models.SeaCreature{})
		fmt.Println("The table doesn't exist in database. \n Running migrations.")

		//db.Migrator().DropTable(&models.Fish{}, &models.Bug{}, &models.SeaCreature{})

		// Fish table
		fileFish, err := os.Open("./database/fishes.json")
		if err != nil {
			fmt.Println("Error during file opening:", err)
			return
		}
		defer fileFish.Close()

		byteValueFish, _ := io.ReadAll(fileFish)
		jsonStringFish := string(byteValueFish)

		var dataFish map[string][]models.Fish
		json.Unmarshal([]byte(jsonStringFish), &dataFish)
		ressourcesFish := dataFish["fishes"]
		for _, ressourceFish := range ressourcesFish {
			err := db.Create(&ressourceFish).Error
			if err != nil {
				fmt.Println(err)
			}
		}
		fmt.Println("Fishes were succesfully insert into database!")

		// Bug table
		fileBug, err := os.Open("./database/bugs.json")
		if err != nil {
			fmt.Println("Error during file opening:", err)
			return
		}
		defer fileBug.Close()

		byteValueBug, _ := io.ReadAll(fileBug)
		jsonStringBug := string(byteValueBug)

		var dataBug map[string][]models.Bug
		json.Unmarshal([]byte(jsonStringBug), &dataBug)
		ressourcesBug := dataBug["bugs"]
		for _, ressourceBug := range ressourcesBug {
			err := db.Create(&ressourceBug).Error
			if err != nil {
				fmt.Println(err)
			}
		}
		fmt.Println("Bugs were succesfully insert into database!")

		// Sea creature table
		fileSeaCreature, err := os.Open("./database/seaCreatures.json")
		if err != nil {
			fmt.Println("Error during file opening:", err)
			return
		}
		defer fileSeaCreature.Close()

		byteValueSeaCreature, _ := io.ReadAll(fileSeaCreature)
		jsonStringSeaCreature := string(byteValueSeaCreature)

		var dataSeaCreature map[string][]models.SeaCreature
		json.Unmarshal([]byte(jsonStringSeaCreature), &dataSeaCreature)
		ressourcesSeaCreature := dataSeaCreature["sea_creatures"]
		for _, ressourceSeaCreature := range ressourcesSeaCreature {
			err := db.Create(&ressourceSeaCreature).Error
			if err != nil {
				fmt.Println(err)
			}
		}
		fmt.Println("Sea Creatures were succesfully insert into database!")
	}

	DB = Dbinstance{
		Db: db,
	}
}
