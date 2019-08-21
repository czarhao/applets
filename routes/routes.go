package routes

import (
	"applets/routes/api"
	"github.com/gin-gonic/gin"
)

func CreateRoute() *gin.Engine {
	route := initRouter()
	setRouter(route)
	return route
}

func initRouter() *gin.Engine {
	route := gin.New()
	route.Use(gin.Logger())
	route.Use(gin.Recovery())
	return route
}

func setRouter(route *gin.Engine) {
	route.GET("/registered/:sno/:spw", api.Registered)
	route.GET("/info/:sno/:spw", api.Info)
	route.GET("/schedule/:sno/:spw", api.Schedule)
	route.GET("/grade/:sno/:spw", api.Grade)
	route.GET("/refresh/:types/:sno/:spw", api.Refresh)
	route.GET("/article/:start", api.Article)
}
