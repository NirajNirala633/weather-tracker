package router

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Name    string
	Method  string
	Path    string
	Handler func(*gin.Context)
}

type routes struct {
	router *gin.Engine
}

type Routes []Route

func (r routes) Weather(g *gin.RouterGroup) {
	weatherRoute := g.Group("/" + os.Getenv("Weather"))
	for _, route := range WeatherRoutes {
		switch route.Method {
		case http.MethodPost:
			weatherRoute.POST(route.Path, route.Handler)
		case http.MethodGet:
			weatherRoute.GET(route.Path, route.Handler)
		case http.MethodPut:
			weatherRoute.PUT(route.Path, route.Handler)
		case http.MethodDelete:
			weatherRoute.DELETE(route.Path, route.Handler)
		}
	}
}

func Routing() {
	r := routes{
		router: gin.Default(),
	}
	grouping := r.router.Group(os.Getenv("Version"))
	r.Weather(grouping)

	r.router.Run(":" + os.Getenv("Port"))
}
