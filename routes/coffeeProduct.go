package routes

import (
	"net/http"
	"swenson-he-task/application"
	"swenson-he-task/controllers"
)

const CoffeResourceName = "/coffee-products"

var coffeeProductRoutes = []application.Route{
	{http.MethodGet, CoffeResourceName,
		func(writer http.ResponseWriter, request *http.Request) {
			controllers.NewCoffeeProductController().List(writer, request)
		},
	},
}
