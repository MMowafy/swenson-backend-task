package routes

import "swenson-he-task/application"

func GetRoutes() []application.Route {
	var appRoutes []application.Route
	appRoutes = append(appRoutes, coffeeProductRoutes...)
	return appRoutes
}
