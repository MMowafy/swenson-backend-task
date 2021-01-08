package infrastructure

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"swenson-he-task/application"
	"swenson-he-task/models"
	"swenson-he-task/utils"
)

type CoffeeProductRepository struct {
	DB *gorm.DB
}

func NewCoffeeProductRepository() *CoffeeProductRepository {
	coffeeProductRepository := &CoffeeProductRepository{
		application.GetDBConnectionByName("swenson_db"),
	}
	coffeeProductRepository.DB.LogMode(application.GetConfig().GetBool("app.db_logs"))
	return coffeeProductRepository
}

func (coffeeProductRepository *CoffeeProductRepository) List(listRequest *utils.ListRequest) []models.CoffeeProduct {
	var coffeeProductList []models.CoffeeProduct

	sortField := fmt.Sprintf(" \"%s\" %s ", listRequest.OrderBy, listRequest.Order)
	db := utils.GenerateSqlCondition(coffeeProductRepository.DB, listRequest.Filters).
		Order(sortField).
		Offset(listRequest.Skip).
		Limit(listRequest.PageSize).
		Find(&coffeeProductList)

	if db.Error != nil {
		application.GetLogger().Error("Failed to get beneficiary with err ==> ", db.Error.Error())
		return nil
	}

	return coffeeProductList
}
