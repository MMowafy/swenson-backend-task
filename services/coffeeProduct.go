package services

import (
	"errors"
	"net/http"
	"reflect"
	"swenson-he-task/infrastructure"
	"swenson-he-task/models"
	"swenson-he-task/utils"
)

type CoffeeProductService struct {
	repository    *infrastructure.CoffeeProductRepository
	coffeeProduct *models.CoffeeProduct
}

func NewCoffeeProductService() *CoffeeProductService {
	return &CoffeeProductService{
		infrastructure.NewCoffeeProductRepository(),
		models.NewCoffeeProduct(),
	}
}

func (coffeeProductService *CoffeeProductService) List(r *http.Request) ([]models.CoffeeProduct, error) {

	listRequest := utils.NewListRequest(r, reflect.ValueOf(coffeeProductService.coffeeProduct).Elem())
	coffeeList := coffeeProductService.repository.List(listRequest)
	if coffeeList == nil {
		return nil, errors.New("something went wrong please try again later")
	}

	return coffeeList, nil
}
