package main

import (
	"encoding/json"
	"errors"
	_ "github.com/lib/pq"
	"io/ioutil"
	"os"
	"swenson-he-task/application"
	"swenson-he-task/infrastructure"
	"swenson-he-task/models"
	"swenson-he-task/routes"
)

func main() {
	app := application.NewApplication()
	migrateAndSeedDB()
	app.SetRoutes(routes.GetRoutes())
	app.StartServer()
}

// This function for testing purpose only migration should be done in other ways
func migrateAndSeedDB() {
	repository := infrastructure.NewCoffeeProductRepository()
	db := repository.DB.DropTableIfExists(&models.CoffeeProduct{})
	if db.Error != nil {
		application.GetLogger().Fatal("Error while migrating CoffeeProduct table => ", db.Error.Error())
	}

	db = db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	if db.Error != nil {
		application.GetLogger().Fatal(errors.New("Error while creating uuid-ossp"))
		application.GetLogger().Fatal(db.Error)
	}

	db = repository.DB.AutoMigrate(&models.CoffeeProduct{})
	if db.Error != nil {
		application.GetLogger().Fatal("Error while migrating CoffeeProduct table => ", db.Error.Error())
	}

	data, err := os.Open("data.json")
	if err != nil {
		application.GetLogger().Fatal("Error reading data.json ", err.Error())
	}

	application.GetLogger().Info("Successfully Opened data.json")
	defer data.Close()

	byteValue, _ := ioutil.ReadAll(data)

	var coffeeProducts []models.CoffeeProduct
	json.Unmarshal(byteValue, &coffeeProducts)

	for _, product := range coffeeProducts {
		err = repository.DB.Create(&product).Error
		if err != nil {
			application.GetLogger().Fatal("Error while seeding CoffeeProduct table => ", err.Error())
		}
	}
}
