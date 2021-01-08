package controllers

import (
	"net/http"
	"swenson-he-task/application"
	"swenson-he-task/services"
)

type CoffeeProductController struct {
	*application.BaseController
	coffeeProdService *services.CoffeeProductService
}

func NewCoffeeProductController() *CoffeeProductController {
	return &CoffeeProductController{
		application.NewBaseController(),
		services.NewCoffeeProductService(),
	}
}

func (coffeeProductController *CoffeeProductController) List(w http.ResponseWriter, r *http.Request) {

	list, err := coffeeProductController.coffeeProdService.List(r)
	if err != nil {
		coffeeProductController.JsonError(w, "Failed to get list of coffee products", http.StatusBadRequest)
	}

	coffeeProductController.Json(w, list, http.StatusOK)
}
