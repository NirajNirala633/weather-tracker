package router

import (
	"net/http"
	"weather-tracker/controller"
)

var WeatherRoutes = Routes{
	Route{"Get temp of particular city", http.MethodGet, "/getWeather", controller.GetCityWeather},
}
